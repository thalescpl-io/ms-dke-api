package testjson

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/jonboulle/clockwork"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"gotest.tools/gotestsum/log"
)

// Action of TestEvent
type Action string

// nolint: unused
const (
	ActionRun    Action = "run"
	ActionPause  Action = "pause"
	ActionCont   Action = "cont"
	ActionPass   Action = "pass"
	ActionBench  Action = "bench"
	ActionFail   Action = "fail"
	ActionOutput Action = "output"
	ActionSkip   Action = "skip"
)

// TestEvent is a structure output by go tool test2json and go test -json.
type TestEvent struct {
	// Time encoded as an RFC3339-format string
	Time    time.Time
	Action  Action
	Package string
	Test    string
	// Elapsed time in seconds
	Elapsed float64
	// Output of test or benchmark
	Output string
	// raw is the raw JSON bytes of the event
	raw []byte
	// RunID from the ScanConfig which produced this test event.
	RunID int
}

// PackageEvent returns true if the event is a package start or end event
func (e TestEvent) PackageEvent() bool {
	return e.Test == ""
}

// ElapsedFormatted returns Elapsed formatted in the go test format, ex (0.00s).
func (e TestEvent) ElapsedFormatted() string {
	return fmt.Sprintf("(%.2fs)", e.Elapsed)
}

// Bytes returns the serialized JSON bytes that were parsed to create the event.
func (e TestEvent) Bytes() []byte {
	return e.raw
}

// Package is the set of TestEvents for a single go package
type Package struct {
	Total   int
	running map[string]TestCase
	Failed  []TestCase
	Skipped []TestCase
	Passed  []TestCase

	// mapping of root TestCase ID to all sub test IDs. Used to mitigate
	// github.com/golang/go/issues/29755.
	// In the future when that bug is fixed this mapping can likely be removed.
	subTests map[int][]int

	// output printed by test cases, indexed by TestCase.ID. Package output is
	// saved with key 0.
	output map[int][]string
	// coverage stores the code coverage output for the package without the
	// trailing newline (ex: coverage: 91.1% of statements).
	coverage string
	// action identifies if the package passed or failed. A package may fail
	// with no test failures if an init() or TestMain exits non-zero.
	// skip indicates there were no tests.
	action Action
	// cached is true if the package was marked as (cached)
	cached bool
}

// Result returns if the package passed, failed, or was skipped because there
// were no tests.
func (p *Package) Result() Action {
	return p.action
}

// Elapsed returns the sum of the elapsed time for all tests in the package.
func (p *Package) Elapsed() time.Duration {
	elapsed := time.Duration(0)
	for _, testcase := range p.TestCases() {
		elapsed += testcase.Elapsed
	}
	return elapsed
}

// TestCases returns all the test cases.
func (p *Package) TestCases() []TestCase {
	tc := append([]TestCase{}, p.Passed...)
	tc = append(tc, p.Failed...)
	tc = append(tc, p.Skipped...)
	return tc
}

// LastFailedByName returns the most recent test with name in the list of Failed
// tests. If no TestCase is found with that name, an empty TestCase is returned.
//
// LastFailedByName may be used by formatters to find the TestCase.ID for the current
// failing TestEvent. It is very likely the last TestCase in Failed, but this method
// provides a little more safety if that ever changes.
func (p *Package) LastFailedByName(name string) TestCase {
	for i := len(p.Failed) - 1; i >= 0; i-- {
		if p.Failed[i].Test == name {
			return p.Failed[i]
		}
	}
	return TestCase{}
}

// Output returns the full test output for a test.
//
// Unlike OutputLines() it does not return lines from subtests in some cases.
func (p *Package) Output(id int) string {
	return strings.Join(p.output[id], "")
}

// OutputLines returns the full test output for a test as a slice of strings.
//
// As a workaround for test output being attributed to the wrong subtest, if:
//   - the TestCase is a root TestCase (not a subtest), and
//   - the TestCase has no subtest failures;
// then all output for every subtest under the root test is returned.
// See https://github.com/golang/go/issues/29755.
func (p *Package) OutputLines(tc TestCase) []string {
	_, sub := SplitTestName(tc.Test)
	lines := p.output[tc.ID]

	// If this is a subtest, or a root test case with subtest failures the
	// subtest failure output should contain the relevant lines, so we don't need
	// extra lines.
	if sub != "" || tc.hasSubTestFailed {
		return lines
	}

	result := make([]string, 0, len(lines)+1)
	result = append(result, lines...)
	for _, sub := range p.subTests[tc.ID] {
		result = append(result, p.output[sub]...)
	}
	return result
}

func (p *Package) addOutput(id int, output string) {
	// TODO: limit size of buffered test output
	p.output[id] = append(p.output[id], output)
}

// SplitTestName into root test name and any subtest names.
func SplitTestName(name string) (root, sub string) {
	parts := strings.SplitN(name, "/", 2)
	if len(parts) < 2 {
		return name, ""
	}
	return parts[0], parts[1]
}

func (p *Package) removeOutput(id int) {
	delete(p.output, id)

	skipped := tcIDSet(p.Skipped)
	for _, sub := range p.subTests[id] {
		if _, isSkipped := skipped[sub]; !isSkipped {
			delete(p.output, sub)
		}
	}
}

func tcIDSet(skipped []TestCase) map[int]struct{} {
	result := make(map[int]struct{})
	for _, tc := range skipped {
		result[tc.ID] = struct{}{}
	}
	return result
}

// TestMainFailed returns true if the package failed, but there were no tests.
// This may occur if the package init() or TestMain exited non-zero.
func (p *Package) TestMainFailed() bool {
	return p.action == ActionFail && len(p.Failed) == 0
}

const neverFinished time.Duration = -1

// end adds any tests that were missing an ActionFail TestEvent to the list of
// Failed, and returns a slice of artificial TestEvent for the missing ones.
//
// This is done to work around 'go test' not sending the ActionFail TestEvents
// in some cases, when a test panics.
func (p *Package) end() []TestEvent {
	result := make([]TestEvent, 0, len(p.running))
	for k, tc := range p.running {
		tc.Elapsed = neverFinished
		p.Failed = append(p.Failed, tc)

		result = append(result, TestEvent{
			Action:  ActionFail,
			Package: tc.Package,
			Test:    tc.Test,
			Elapsed: float64(neverFinished),
		})
		delete(p.running, k)
	}
	return result
}

// TestCase stores the name and elapsed time for a test case.
type TestCase struct {
	// ID is unique ID for each test case. A test run may include multiple instances
	// of the same Package and Name if -count is used, or if the input comes from
	// multiple runs. The ID can be used to uniquely reference an instance of a
	// test case.
	ID      int
	Package string
	Test    string
	Elapsed time.Duration
	// RunID from the ScanConfig which produced this test case.
	RunID int
	// hasSubTestFailed is true when a subtest of this TestCase has failed. It is
	// used to find root TestCases which have no failing subtests.
	hasSubTestFailed bool
}

func newPackage() *Package {
	return &Package{
		output:   make(map[int][]string),
		running:  make(map[string]TestCase),
		subTests: make(map[int][]int),
	}
}

// Execution of one or more test packages
type Execution struct {
	started   time.Time
	packages  map[string]*Package
	errors    []string
	done      bool
	lastRunID int
}

func (e *Execution) add(event TestEvent) {
	pkg, ok := e.packages[event.Package]
	if !ok {
		pkg = newPackage()
		e.packages[event.Package] = pkg
	}
	if event.PackageEvent() {
		e.addPackageEvent(pkg, event)
		return
	}
	pkg.addTestEvent(event)
}

func (e *Execution) addPackageEvent(pkg *Package, event TestEvent) {
	switch event.Action {
	case ActionPass, ActionFail:
		pkg.action = event.Action
	case ActionOutput:
		if isCoverageOutput(event.Output) {
			pkg.coverage = strings.TrimRight(event.Output, "\n")
		}
		if isCachedOutput(event.Output) {
			pkg.cached = true
		}
		pkg.addOutput(0, event.Output)
	}
}

func (p *Package) addTestEvent(event TestEvent) {
	tc := p.running[event.Test]
	root, subTest := SplitTestName(event.Test)

	switch event.Action {
	case ActionRun:
		// Incremental total before using it as the ID, because ID 0 is used for
		// the package output
		p.Total++
		tc := TestCase{
			Package: event.Package,
			Test:    event.Test,
			ID:      p.Total,
			RunID:   event.RunID,
		}
		p.running[event.Test] = tc

		if subTest != "" {
			rootID := p.running[root].ID
			p.subTests[rootID] = append(p.subTests[rootID], tc.ID)
		}
		return
	case ActionOutput, ActionBench:
		p.addOutput(tc.ID, event.Output)
		return
	case ActionPause, ActionCont:
		return
	}

	delete(p.running, event.Test)
	tc.Elapsed = elapsedDuration(event.Elapsed)

	switch event.Action {
	case ActionFail:
		p.Failed = append(p.Failed, tc)

		// If this is a subtest, mark the root test as having a failed subtest
		if subTest != "" {
			rootTestCase := p.running[root]
			rootTestCase.hasSubTestFailed = true
			p.running[root] = rootTestCase
		}
	case ActionSkip:
		p.Skipped = append(p.Skipped, tc)

	case ActionPass:
		p.Passed = append(p.Passed, tc)

		// Do not immediately remove output for subtests, to work around a bug
		// in 'go test' where output is attributed to the wrong sub test.
		// github.com/golang/go/issues/29755.
		if subTest != "" {
			return
		}

		// Remove test output once a test passes, it wont be used.
		p.removeOutput(tc.ID)
		// Remove subtest mapping, it is only used when a test fails.
		delete(p.subTests, tc.ID)
	}
}

func elapsedDuration(elapsed float64) time.Duration {
	return time.Duration(elapsed*1000) * time.Millisecond
}

func isCoverageOutput(output string) bool {
	return all(
		strings.HasPrefix(output, "coverage:"),
		strings.HasSuffix(output, "% of statements\n"))
}

func isCachedOutput(output string) bool {
	return strings.Contains(output, "\t(cached)")
}

// OutputLines returns the full test output for a test as an slice of lines.
// This function is a convenient wrapper around Package.OutputLines() to
// support the hiding of output in the summary.
//
// See Package.OutLines() for more details.
func (e *Execution) OutputLines(tc TestCase) []string {
	return e.packages[tc.Package].OutputLines(tc)
}

// Package returns the Package by name.
func (e *Execution) Package(name string) *Package {
	return e.packages[name]
}

// Packages returns a sorted list of all package names.
func (e *Execution) Packages() []string {
	return sortedKeys(e.packages)
}

var clock = clockwork.NewRealClock()

// Elapsed returns the time elapsed since the execution started.
func (e *Execution) Elapsed() time.Duration {
	return clock.Now().Sub(e.started)
}

// Failed returns a list of all the failed test cases.
func (e *Execution) Failed() []TestCase {
	if e == nil {
		return nil
	}
	var failed []TestCase //nolint:prealloc
	for _, name := range sortedKeys(e.packages) {
		pkg := e.packages[name]

		// Add package-level failure output if there were no failed tests.
		if pkg.TestMainFailed() {
			failed = append(failed, TestCase{Package: name})
		}
		failed = append(failed, pkg.Failed...)
	}
	return failed
}

// FilterFailedUnique filters a slice of failed TestCases by removing root test
// case that have failed subtests.
func FilterFailedUnique(tcs []TestCase) []TestCase {
	var result []TestCase
	for _, tc := range tcs {
		if !tc.hasSubTestFailed {
			result = append(result, tc)
		}
	}
	return result
}

func sortedKeys(pkgs map[string]*Package) []string {
	keys := make([]string, 0, len(pkgs))
	for key := range pkgs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// Skipped returns a list of all the skipped test cases.
func (e *Execution) Skipped() []TestCase {
	skipped := make([]TestCase, 0, len(e.packages))
	for _, pkg := range sortedKeys(e.packages) {
		skipped = append(skipped, e.packages[pkg].Skipped...)
	}
	return skipped
}

// Total returns a count of all test cases.
func (e *Execution) Total() int {
	total := 0
	for _, pkg := range e.packages {
		total += pkg.Total
	}
	return total
}

func (e *Execution) addError(err string) {
	// Build errors start with a header
	if strings.HasPrefix(err, "# ") {
		return
	}
	e.errors = append(e.errors, err)
}

// Errors returns a list of all the errors.
func (e *Execution) Errors() []string {
	return e.errors
}

func (e *Execution) end() []TestEvent {
	e.done = true
	var result []TestEvent // nolint: prealloc
	for _, pkg := range e.packages {
		result = append(result, pkg.end()...)
	}
	return result
}

// newExecution returns a new Execution and records the current time as the
// time the test execution started.
func newExecution() *Execution {
	return &Execution{
		started:  clock.Now(),
		packages: make(map[string]*Package),
	}
}

// ScanConfig used by ScanTestOutput.
type ScanConfig struct {
	// RunID is a unique identifier for the run. It may be set to the pid of the
	// process, or some other identifier. It will stored as the TestCase.RunID.
	RunID int
	// Stdout is a reader that yields the test2json output stream.
	Stdout io.Reader
	// Stderr is a reader that yields stderr from the 'go test' process. Often
	// it contains build errors, or panics. Stderr may be nil.
	Stderr io.Reader
	// Handler is a set of callbacks for receiving TestEvents and stderr text.
	Handler EventHandler
	// Execution to populate while scanning. If nil a new one will be created
	// and returned from ScanTestOutput.
	Execution *Execution
}

// EventHandler is called by ScanTestOutput for each event and write to stderr.
type EventHandler interface {
	// Event is called for every TestEvent, with the current value of Execution.
	// It may return an error to stop scanning.
	Event(event TestEvent, execution *Execution) error
	// Err is called for every line from the Stderr reader and may return an
	// error to stop scanning.
	Err(text string) error
}

// ScanTestOutput reads lines from config.Stdout and config.Stderr, populates an
// Execution, calls the Handler for each event, and returns the Execution.
//
// If config.Handler is nil, a default no-op handler will be used.
func ScanTestOutput(config ScanConfig) (*Execution, error) {
	if config.Stdout == nil {
		return nil, fmt.Errorf("stdout reader must be non-nil")
	}
	if config.Handler == nil {
		config.Handler = noopHandler{}
	}
	if config.Stderr == nil {
		config.Stderr = new(bytes.Reader)
	}
	execution := config.Execution
	if execution == nil {
		execution = newExecution()
	}
	execution.done = false
	execution.lastRunID = config.RunID

	var group errgroup.Group
	group.Go(func() error {
		return readStdout(config, execution)
	})
	group.Go(func() error {
		return readStderr(config, execution)
	})
	if err := group.Wait(); err != nil {
		return execution, err
	}
	for _, event := range execution.end() {
		if err := config.Handler.Event(event, execution); err != nil {
			return execution, err
		}
	}

	return execution, nil
}

func readStdout(config ScanConfig, execution *Execution) error {
	scanner := bufio.NewScanner(config.Stdout)
	for scanner.Scan() {
		raw := scanner.Bytes()
		event, err := parseEvent(raw)
		switch {
		case err == errBadEvent:
			// nolint: errcheck
			config.Handler.Err(errBadEvent.Error() + ": " + scanner.Text())
			continue
		case err != nil:
			return errors.Wrapf(err, "failed to parse test output: %s", string(raw))
		}

		event.RunID = config.RunID
		execution.add(event)
		if err := config.Handler.Event(event, execution); err != nil {
			return err
		}
	}
	return errors.Wrap(scanner.Err(), "failed to scan test output")
}

func readStderr(config ScanConfig, execution *Execution) error {
	scanner := bufio.NewScanner(config.Stderr)
	for scanner.Scan() {
		line := scanner.Text()
		if err := config.Handler.Err(line); err != nil {
			return fmt.Errorf("failed to handle stderr: %v", err)
		}
		if isGoModuleOutput(line) {
			continue
		}
		execution.addError(line)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan stderr: %v", err)
	}
	return nil
}

func isGoModuleOutput(scannerText string) bool {
	prefixes := []string{
		"go: copying",
		"go: creating",
		"go: downloading",
		"go: extracting",
		"go: finding",
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(scannerText, prefix) {
			return true
		}
	}
	return false
}

func parseEvent(raw []byte) (TestEvent, error) {
	// TODO: this seems to be a bug in the `go test -json` output
	if bytes.HasPrefix(raw, []byte("FAIL")) {
		log.Warnf("invalid TestEvent: %v", string(raw))
		return TestEvent{}, errBadEvent
	}

	event := TestEvent{}
	err := json.Unmarshal(raw, &event)
	event.raw = raw
	return event, err
}

var errBadEvent = errors.New("bad output from test2json")

type noopHandler struct{}

func (s noopHandler) Event(TestEvent, *Execution) error {
	return nil
}

func (s noopHandler) Err(string) error {
	return nil
}

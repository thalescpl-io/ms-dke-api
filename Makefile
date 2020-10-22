.PHONY: gen gen-proto
PROTOTOOL_IMAGE:=uber/prototool:latest

gen: gen-proto
gen-proto:
		@mkdir -p generated
		@docker run -ti -v $(shell pwd):/work $(PROTOTOOL_IMAGE) prototool generate --walk-timeout 60s || true
		@cp -r generated/github.com/thalscpl-io/ms-dke-api/dke/* ./dke/
		@rm -rf generated

gen-prototol:

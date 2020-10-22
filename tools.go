// +build tools

package ms_dke_api

// tools.go is not meant to be compiled with the project.  It's solely
// here to pin some dependencies on tools we use in the build.
// go modules will include version management and vendoring of
// these tools, even though this source file has a build tag that
// isn't include by the real build.

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "github.com/mitchellh/protoc-gen-go-json"
	_ "golang.org/x/tools/cmd/cover"
	_ "gotest.tools/gotestsum"

)

// These are projects we're vendoring for the sake of their proto files
// Trying to use go's vendoring as much as possible for fetching and tracking
// the proto file sources.

// Unfortunately, go's toolchain won't work with projects that don't look like
// go projects, so these are projects we just need to vendor manually, which
// sucks.  proto dependency management is pretty messy right now.
// Still keeping a list of them here just to remember what they are.
//import _ "github.com/googleapis/api-common-protos"
//import _ "github.com/googleapis/googleapis"

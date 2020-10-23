all: gen app

SHELL := /bin/bash -o pipefail

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

TMP_BASE := .tmp
TMP := $(TMP_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
TMP_BIN = $(TMP)/bin
TMP_VERSIONS := $(TMP)/versions

export GO111MODULE := on
export GOBIN := $(abspath $(TMP_BIN))
export PATH := $(GOBIN):$(PATH)

# These are the only variables that ever should change.
GOLANG_PROTOBUF_VERSION := v1.3.1
GOGO_PROTOBUF_VERSION := v1.2.1
GRPC_GATEWAY_VERSION := v1.14.0
GRPC_WEB_VERSION := v1.2.1
PROTOTOOL_VERSION := v1.10.0

GOLANG_PROTOBUF := $(TMP_VERSIONS)/golang_protobuf/$(GOLANG_PROTOBUF_VERSION)
$(GOLANG_PROTOBUF):
	$(eval GOLANG_PROTOBUF_TMP := $(shell mktemp -d))
	cd $(GOLANG_PROTOBUF_TMP); go get github.com/golang/protobuf/protoc-gen-go@${GOLANG_PROTOBUF_VERSION}
	@rm -rf $(GOLANG_PROTOBUF_TMP)
	@rm -rf $(dir $(GOLANG_PROTOBUF))
	@mkdir -p $(dir $(GOLANG_PROTOBUF))
	@touch $(GOLANG_PROTOBUF)

GOGO_PROTOBUF := $(TMP_VERSIONS)/gogo_protobuf/$(GOGO_PROTOBUF_VERSION)
$(GOGO_PROTOBUF):
	$(eval GOGO_PROTOBUF_TMP := $(shell mktemp -d))
	cd $(GOGO_PROTOBUF_TMP); go get github.com/gogo/protobuf/protoc-gen-gofast@${GOGO_PROTOBUF_VERSION} \
	                                github.com/gogo/protobuf/protoc-gen-gogo@${GOGO_PROTOBUF_VERSION} \
	                                github.com/gogo/protobuf/protoc-gen-gogofast@${GOGO_PROTOBUF_VERSION} \
	                                github.com/gogo/protobuf/protoc-gen-gogofaster@${GOGO_PROTOBUF_VERSION} \
	                                github.com/gogo/protobuf/protoc-gen-gogoslick@${GOGO_PROTOBUF_VERSION}
	@rm -rf $(GOGO_PROTOBUF_TMP)
	@rm -rf $(dir $(GOGO_PROTOBUF))
	@mkdir -p $(dir $(GOGO_PROTOBUF))
	@touch $(GOGO_PROTOBUF)

GRPC_GATEWAY := $(TMP_VERSIONS)/grpc-gateway/$(GRPC_GATEWAY_VERSION)
$(GRPC_GATEWAY):
	$(eval GRPC_GATEWAY_TMP := $(shell mktemp -d))
	cd $(GRPC_GATEWAY_TMP); go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@$(GRPC_GATEWAY_VERSION)
	cd $(GRPC_GATEWAY_TMP); go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@$(GRPC_GATEWAY_VERSION)
	@rm -rf $(GRPC_GATEWAY_TMP)
	@rm -rf $(dir $(GRPC_GATEWAY))
	@mkdir -p $(dir $(GRPC_GATEWAY))
	@touch $(GRPC_GATEWAY)

PROTOTOOL := $(TMP_VERSIONS)/prototool/$(PROTOTOOL_VERSION)
$(PROTOTOOL):
	$(eval PROTOTOOL_TMP := $(shell mktemp -d))
	cd $(PROTOTOOL_TMP); go get github.com/uber/prototool/cmd/prototool@$(PROTOTOOL_VERSION)
	@rm -rf $(PROTOTOOL_TMP)
	@rm -rf $(dir $(PROTOTOOL))
	@mkdir -p $(dir $(PROTOTOOL))
	@touch $(PROTOTOOL)

.PHONY: deps gen gen-proto

gen: gen-proto
gen-proto: deps
		mkdir -p generated
		./prototool generate --walk-timeout 60s
		cp -r generated/github.com/thalscpl-io/ms-dke-api/dke/* ./dke/
		rm -rf generated

prototool:
	curl -L https://github.com/uber/prototool/releases/download/v$(PROTOTOOL_VERSION)/prototool-$(UNAME_S)-$(UNAME_M) -o prototool
	chmod a+x prototool
	go get -u github.com/golang/protobuf/protoc-gen-go

gen-prototol:

deps: $(GOLANG_PROTOBUF) $(GOGO_PROTOBUF) $(GRPC_GATEWAY) $(PROTOTOOL)
	@echo > /dev/null

app: gen main.go keystore/main.go
	go build

key.pem:
	openssl genrsa -out key.pem

key.pub: key.pem
	openssl rsa -in key.pem -pubout > key.pub

test: key.pub
	curl -v -L -H"Content-Type: application/json" -d "{\"key_name\": \"TestKey1\", \"key_id\": \"D798B899-3350-4F5C-A608-2EDA37CB0EBD\", \"alg\": \"RSA-OAEP-256\", \"value\": \"$(shell openssl pkeyutl -encrypt -inkey key.pub -pubin -pkeyopt rsa_padding_mode:oaep -pkeyopt rsa_oaep_md:sha256 -in input.txt | base64 -w0)\"}" http://localhost:5000/TestKey1/D798B899-3350-4F5C-A608-2EDA37CB0EBD/decrypt

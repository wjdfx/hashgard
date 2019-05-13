PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_MODULES=$(shell go list ./... | grep 'x')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
CAT := $(if $(filter $(OS),Windows_NT),type,cat)

GOBIN ?= $(GOPATH)/bin
GOSUM := $(shell which gosum)

export GO111MODULE = on


BUILD_FLAGS = -ldflags "-X github.com/hashgard/hashgard/version.Version=$(VERSION) \
    -X github.com/hashgard/hashgard/version.Commit=$(COMMIT)"

ifneq ($(GOSUM),)
ldflags += -X github.com/hashgard/hashgard/version.VendorDirHash=$(shell $(GOSUM) go.sum)
endif

all: get_tools install lint test


########################################
### Tools

###
# Find OS and Go environment
# GO contains the Go binary
# FS contains the OS file separator
###

ifeq ($(OS),Windows_NT)
  GO := $(shell where go.exe 2> NUL)
  FS := \\
else
  GO := $(shell command -v go 2> /dev/null)
  FS := /
endif

ifeq ($(GO),)
  $(error could not find go. Is it in PATH? $(GO))
endif

GOPATH ?= $(shell $(GO) env GOPATH)
GITHUBDIR := $(GOPATH)$(FS)src$(FS)github.com
GOLANGCI_LINT_VERSION := v1.15.0
GOLANGCI_LINT_HASHSUM := ac897cadc180bf0c1a4bf27776c410debad27205b22856b861d41d39d06509cf

###
# Functions
###

go_get = $(if $(findstring Windows_NT,$(OS)),\
IF NOT EXIST $(GITHUBDIR)$(FS)$(1)$(FS) ( mkdir $(GITHUBDIR)$(FS)$(1) ) else (cd .) &\
IF NOT EXIST $(GITHUBDIR)$(FS)$(1)$(FS)$(2)$(FS) ( cd $(GITHUBDIR)$(FS)$(1) && git clone https://github.com/$(1)/$(2) ) else (cd .) &\
,\
mkdir -p $(GITHUBDIR)$(FS)$(1) &&\
(test ! -d $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && cd $(GITHUBDIR)$(FS)$(1) && git clone https://github.com/$(1)/$(2)) || true &&\
)\
cd $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && git fetch origin && git checkout -q $(3)

go_install = $(call go_get,$(1),$(2),$(3)) && cd $(GITHUBDIR)$(FS)$(1)$(FS)$(2) && $(GO) install

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dir := $(shell cd $(shell dirname $(mkfile_path)); pwd)

###
# tools
###

get_tools: tools-stamp
tools-stamp: $(GOBIN)/golangci-lint $(GOBIN)/statik $(GOBIN)/goimports $(GOBIN)/gosum
	touch $@

$(GOBIN)/golangci-lint: contrib/install-golangci-lint.sh $(GOBIN)/gosum
	bash contrib/install-golangci-lint.sh $(GOBIN) $(GOLANGCI_LINT_VERSION) $(GOLANGCI_LINT_HASHSUM)

$(GOBIN)/statik:
	$(call go_install,rakyll,statik,v0.1.5)

$(GOBIN)/goimports:
	go get golang.org/x/tools/cmd/goimports@v0.0.0-20190114222345-bf090417da8b

$(GOBIN)/gosum:
	go install -mod=readonly ./cmd/gosum/

tools-clean:
	cd $(GOBIN) && rm -f golangci-lint statik goimports gosum
	rm -f tools-stamp


########################################
### CI

ci: get_tools install test_cover lint test


########################################
### Dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: get_tools go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

distclean: clean
	rm -rf vendor/


########################################
### Build/Install

update_gaia_lite_docs:
	@statik -src=client/lcd/swagger-ui -dest=client/lcd -f

build-linux: go.sum
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(MAKE) build

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/hashgard.exe ./cmd/hashgard
	go build -mod=readonly $(BUILD_FLAGS) -o build/hashgardcli.exe ./cmd/hashgardcli
	go build -mod=readonly $(BUILD_FLAGS) -o build/hashgardlcd.exe ./cmd/hashgardlcd
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/hashgard ./cmd/hashgard
	go build -mod=readonly $(BUILD_FLAGS) -o build/hashgardcli ./cmd/hashgardcli
	go build -mod=readonly $(BUILD_FLAGS) -o build/hashgardlcd ./cmd/hashgardlcd
	go build -mod=readonly $(BUILD_FLAGS) -o build/hashgardkeyutil ./cmd/hashgardkeyutil
	go build -mod=readonly $(BUILD_FLAGS) -o build/hashgardreplay ./cmd/hashgardreplay
endif


install: go.sum update_gaia_lite_docs
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hashgard
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hashgardcli
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hashgardlcd
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hashgardkeyutil
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/hashgardreplay


########################################
### Documentation

godocs:
	@echo "--> Wait a few seconds and visit http://localhost:6060/pkg/github.com/hashgard/hashgard"
	godoc -http=:6060


########################################
### Testing

test: test_unit

test_unit:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_NOSIMULATION)

test_cover:
	@export VERSION=$(VERSION); bash tests/test_cover.sh

test_cli:
	@go test -mod=readonly -p 4 `go list ./cli_test/...` -tags=cli_test

ci-lint:
	golangci-lint run
	go vet -composites=false -tests=false ./...
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

lint: get_tools ci-lint


# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: build install  \
clean distclean \
test test_unit test_cover lint\
build-linux go-mod-cache\
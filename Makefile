PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell git describe --tags --long | sed 's/v\(.*\)/\1/')
BUILD_TAGS = netgo ledger
BUILD_FLAGS = -tags "${BUILD_TAGS}" -ldflags "-X github.com/hashgard/hashgard/version.Version=${VERSION}"
GCC := $(shell command -v gcc 2> /dev/null)
LEDGER_ENABLED ?= true
UNAME_S := $(shell uname -s)

GLIDE_CHECK := $(shell command -v glide 2> /dev/null)

all: get_tools get_vendor_deps install test_lint test

########################################
### CI

ci: get_tools get_vendor_deps install test_cover test_lint test

########################################
### Build/Install

check-ledger:
ifeq ($(LEDGER_ENABLED),true)
   	ifeq ($(UNAME_S),OpenBSD)
   		$(info "OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988)")
TMP_BUILD_TAGS := $(BUILD_TAGS)
BUILD_TAGS = $(filter-out ledger, $(TMP_BUILD_TAGS))
   	else
   	   	ifndef GCC
   	   	   $(error "gcc not installed for ledger support, please install or set LEDGER_ENABLED to false in the Makefile")
   	   	endif
   	endif
else
TMP_BUILD_TAGS := $(BUILD_TAGS)
BUILD_TAGS = $(filter-out ledger, $(TMP_BUILD_TAGS))
endif

update_gaia_lite_docs:
	@statik -src=vendor/github.com/cosmos/cosmos-sdk/client/lcd/swagger-ui -dest=vendor/github.com/cosmos/cosmos-sdk/client/lcd -f

build: check-ledger update_gaia_lite_docs
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/hashgard.exe ./cmd/hashgard
	go build $(BUILD_FLAGS) -o build/hashgardcli.exe ./cmd/hashgardcli
else
	go build $(BUILD_FLAGS) -o build/hashgard ./cmd/hashgard
	go build $(BUILD_FLAGS) -o build/hashgardcli ./cmd/hashgardcli
endif

build-linux:
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install: check-ledger update_gaia_lite_docs
	go install $(BUILD_FLAGS) ./cmd/hashgard
	go install $(BUILD_FLAGS) ./cmd/hashgardcli


########################################
### Tools & Dependencies

get_tools:
ifdef GLIDE_CHECK
	@echo "Glide is already installed."
else
	@echo "Installing glide"
	curl https://glide.sh/get | sh
endif
	go get github.com/rakyll/statik
	go get github.com/alecthomas/gometalinter

get_vendor_deps:
	@echo "--> Generating vendor directory via glide install"
	@rm -rf ./vendor
	@glide install
	
update_vendor_deps:
	@echo "--> Running glide update"
	@glide update


########################################
### Testing

test_unit:
	@VERSION=$(VERSION) go test $(PACKAGES_NOSIMULATION)

test: test_unit

test_lint:
	gometalinter --config=tests/gometalinter.json ./...
	!(gometalinter --exclude /usr/lib/go/src/ --exclude 'vendor/*' --disable-all --enable='errcheck' --vendor ./... )
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s

test_cover:
	@export VERSION=$(VERSION); bash tests/test_cover.sh



# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: build install install_debug dist \
get_tools get_dev_tools get_vendor_deps test test_cli test_unit \
test_cover test_lint \
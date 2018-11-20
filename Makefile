PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell git describe --tags --long | sed 's/v\(.*\)/\1/')
BUILD_TAGS = netgo ledger
BUILD_FLAGS = -tags "${BUILD_TAGS}" -ldflags "-X github.com/cosmos/cosmos-sdk/version.Version=${VERSION}"
GCC := $(shell command -v gcc 2> /dev/null)
LEDGER_ENABLED ?= true
UNAME_S := $(shell uname -s)

all: get_tools get_vendor_deps install

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

build: check-ledger
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/hashgard.exe ./cmd/hashgard
	go build $(BUILD_FLAGS) -o build/hashgardcli.exe ./cmd/hashgardcli
else
	go build $(BUILD_FLAGS) -o build/hashgard ./cmd/hashgard
	go build $(BUILD_FLAGS) -o build/hashgardcli ./cmd/hashgardcli
endif

build-linux:
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install: check-ledger
	go install $(BUILD_FLAGS) ./cmd/hashgard
	go install $(BUILD_FLAGS) ./cmd/hashgardcli


########################################
### Tools & dependencies

GLIDE_CHECK := $(shell command -v glide 2> /dev/null)

get_tools:
	@echo "--> Installing tools"
	ifdef GLIDE_CHECK
		@echo "Glide is already installed."
	else
		@echo "Installing glide"
		curl https://glide.sh/get | sh
	endif

get_vendor_deps:
	@echo "--> Generating vendor directory via glide install"
	@rm -rf ./vendor
	@glide install
	
update_vendor_deps:
	@echo "--> Running glide update"
	@glide update

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: build build_cosmos-sdk-cli build_examples install install_examples install_cosmos-sdk-cli install_debug dist \
check_tools check_dev_tools get_tools get_dev_tools get_vendor_deps draw_deps test test_cli test_unit \
test_cover test_lint benchmark devdoc_init devdoc devdoc_save devdoc_update \
build-linux build-docker-gaiadnode localnet-start localnet-stop \
format check-ledger test_sim_gaia_nondeterminism test_sim_modules test_sim_gaia_fast \
test_sim_gaia_multi_seed test_sim_gaia_import_export update_tools update_dev_tools
PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
PACKAGES_MODULES=$(shell go list ./... | grep 'x')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')

VERSION := $(subst v,,$(shell git describe --tags --long))
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_FLAGS = -ldflags "-X github.com/hashgard/hashgard/version.Version=${VERSION}"
GLIDE_CHECK := $(shell command -v glide 2> /dev/null)

all: get_tools get_vendor_deps install


########################################
### Tools

GLIDE = github.com/Masterminds/glide
GOLINT = github.com/tendermint/lint/golint
GOMETALINTER = gopkg.in/alecthomas/gometalinter.v2
UNCONVERT = github.com/mdempsky/unconvert
INEFFASSIGN = github.com/gordonklaus/ineffassign
MISSPELL = github.com/client9/misspell/cmd/misspell
ERRCHECK = github.com/kisielk/errcheck
UNPARAM = mvdan.cc/unparam
STATIK = github.com/rakyll/statik

GLIDE_CHECK := $(shell command -v glide 2> /dev/null)
GOLINT_CHECK := $(shell command -v golint 2> /dev/null)
GOMETALINTER_CHECK := $(shell command -v gometalinter.v2 2> /dev/null)
UNCONVERT_CHECK := $(shell command -v unconvert 2> /dev/null)
INEFFASSIGN_CHECK := $(shell command -v ineffassign 2> /dev/null)
MISSPELL_CHECK := $(shell command -v misspell 2> /dev/null)
ERRCHECK_CHECK := $(shell command -v errcheck 2> /dev/null)
UNPARAM_CHECK := $(shell command -v unparam 2> /dev/null)
STATIK_CHECK := $(shell command -v statik 2> /dev/null)


check_tools:
ifndef GLIDE_CHECK
	@echo "No glide in path.  Install with 'make get_tools'."
else
	@echo "Found glide in path."
endif
ifndef STATIK_CHECK
	@echo "No statik in path.  Install with 'make get_tools'."
else
	@echo "Found statik in path."
endif


check_dev_tools:
	$(MAKE) check_tools
ifndef GOLINT_CHECK
	@echo "No golint in path.  Install with 'make get_dev_tools'."
else
	@echo "Found golint in path."
endif
ifndef GOMETALINTER_CHECK
	@echo "No gometalinter in path.  Install with 'make get_dev_tools'."
else
	@echo "Found gometalinter in path."
endif
ifndef UNCONVERT_CHECK
	@echo "No unconvert in path.  Install with 'make get_dev_tools'."
else
	@echo "Found unconvert in path."
endif
ifndef INEFFASSIGN_CHECK
	@echo "No ineffassign in path.  Install with 'make get_dev_tools'."
else
	@echo "Found ineffassign in path."
endif
ifndef MISSPELL_CHECK
	@echo "No misspell in path.  Install with 'make get_dev_tools'."
else
	@echo "Found misspell in path."
endif
ifndef ERRCHECK_CHECK
	@echo "No errcheck in path.  Install with 'make get_dev_tools'."
else
	@echo "Found errcheck in path."
endif
ifndef UNPARAM_CHECK
	@echo "No unparam in path.  Install with 'make get_dev_tools'."
else
	@echo "Found unparam in path."
endif


get_tools:
ifdef GLIDE_CHECK_CHECK
	@echo "Glide is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing glide"
	go get -v $(GLIDE)
endif
ifdef STATIK_CHECK
	@echo "Statik is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing statik"
	go version
	go get -v $(STATIK)
endif


get_dev_tools:
	$(MAKE) get_tools
ifdef GOLINT_CHECK
	@echo "Golint is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing golint"
	go get -v $(GOLINT)
endif
ifdef GOMETALINTER_CHECK
	@echo "Gometalinter.v2 is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing gometalinter.v2"
	go get -v $(GOMETALINTER)
endif
ifdef UNCONVERT_CHECK
	@echo "Unconvert is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing unconvert"
	go get -v $(UNCONVERT)
endif
ifdef INEFFASSIGN_CHECK
	@echo "Ineffassign is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing ineffassign"
	go get -v $(INEFFASSIGN)
endif
ifdef MISSPELL_CHECK
	@echo "misspell is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing misspell"
	go get -v $(MISSPELL)
endif
ifdef ERRCHECK_CHECK
	@echo "errcheck is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing errcheck"
	go get -v $(ERRCHECK)
endif
ifdef UNPARAM_CHECK
	@echo "unparam is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing unparam"
	go get -v $(UNPARAM)
endif
ifdef STATIK_CHECK
	@echo "statik is already installed.  Run 'make update_tools' to update."
else
	@echo "Installing statik"
	go get -v $(STATIK)
endif


update_tools:
	@echo "Updating dep"
	go get -u -v $(DEP)


########################################
### Dependencies

get_vendor_deps:
	@echo "--> Generating vendor directory via glide install"
	@rm -rf ./vendor
	@glide install

update_vendor_deps:
	@echo "--> Running glide update"
	@glide update



########################################
### Build/Install

update_gaia_lite_docs:
	@statik -src=vendor/github.com/cosmos/cosmos-sdk/client/lcd/swagger-ui -dest=vendor/github.com/cosmos/cosmos-sdk/client/lcd -f


build:
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/hashgard.exe ./cmd/hashgard
	go build $(BUILD_FLAGS) -o build/hashgardcli.exe ./cmd/hashgardcli
	go build $(BUILD_FLAGS) -o build/hashgardkeyutil.exe ./cmd/hashgardkeyutil
	go build $(BUILD_FLAGS) -o build/hashgardreplay.exe ./cmd/hashgardreplay
else
	go build $(BUILD_FLAGS) -o build/hashgard ./cmd/hashgard
	go build $(BUILD_FLAGS) -o build/hashgardcli ./cmd/hashgardcli
	go build $(BUILD_FLAGS) -o build/hashgardkeyutil ./cmd/hashgardkeyutil
	go build $(BUILD_FLAGS) -o build/hashgardreplay ./cmd/hashgardreplay
endif


install: update_gaia_lite_docs
	go install $(BUILD_FLAGS) ./cmd/hashgard
	go install $(BUILD_FLAGS) ./cmd/hashgardcli
	go install $(BUILD_FLAGS) ./cmd/hashgardkeyutil
	go install $(BUILD_FLAGS) ./cmd/hashgardreplay


########################################
### Testing

test_unit:
	@VERSION=$(VERSION) go test $(PACKAGES_NOSIMULATION)

test: test_unit test_cover

test_lint:
	gometalinter --config=tests/gometalinter.json ./...
	!(gometalinter --exclude /usr/lib/go/src/ --exclude 'vendor/*' --disable-all --enable='errcheck' --vendor ./... | grep -v "vendor/")
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s

test_cover:
	@export VERSION=$(VERSION); bash tests/test_cover.sh


# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: build install  \
get_tools get_dev_tools get_vendor_deps test test_cli test_unit \
test_cover test_lint \
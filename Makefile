PACKAGE  = bitbucket.org/mexisme/get-secrets
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
GOPATH   = $(CURDIR)/.gopath~
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))
TESTPKGS = $(shell env GOPATH=$(GOPATH) $(GO) list -f '{{ if .TestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))

GO      = go
# GO_TEST = go test
GO_TEST = $(GINKGO) -r -p -v
GODOC   = godoc
GOFMT   = gofmt
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m>>\033[0m")

.PHONY: all
all: fmt lint vendor | $(BASE) ; $(info $(M) building executable...) @ ## Build program binary
	$Q cd $(BASE) && $(GO) build \
		-tags release \
		-ldflags '-X $(PACKAGE)/cmd.Version=$(VERSION) -X $(PACKAGE)/cmd.BuildDate=$(DATE)' \
		-o bin/$(PACKAGE) main.go

$(BASE): ; $(info $(M) setting GOPATH...)
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

# Tools

GOLINT = $(BIN)/golint
$(BIN)/golint: | $(BASE) ; $(info $(M) building golint...)
	$Q go get github.com/golang/lint/golint

# GOCOVMERGE = $(BIN)/gocovmerge
# $(BIN)/gocovmerge: | $(BASE) ; $(info $(M) building gocovmerge...)
# 	$Q go get github.com/wadey/gocovmerge

# GOCOV = $(BIN)/gocov
# $(BIN)/gocov: | $(BASE) ; $(info $(M) building gocov...)
# 	$Q go get github.com/axw/gocov/...

# GOCOVXML = $(BIN)/gocov-xml
# $(BIN)/gocov-xml: | $(BASE) ; $(info $(M) building gocov-xml...)
# 	$Q go get github.com/AlekSi/gocov-xml

# GO2XUNIT = $(BIN)/go2xunit
# $(BIN)/go2xunit: | $(BASE) ; $(info $(M) building go2xunit...)
# 	$Q go get github.com/tebeka/go2xunit

.PHONY: go-dep
GODEP = $(BIN)/dep
go-dep: $(GODEP)
$(GODEP): | $(BASE) ; $(info $(M) building go-dep...)
	$Q go get -u github.com/golang/dep/cmd/dep

GINKGO = $(BIN)/ginkgo
$(GINKGO): | $(BASE) ; $(info $(M) building ginkgo...)
	$Q go get -u github.com/onsi/ginkgo/ginkgo

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
INTEGRATION_TEST_TARGETS := test-integration

.PHONY: $(TEST_TARGETS) $(INTEGRATION_TEST_TARGETS) \
	test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): SKIP_ARGS=-skip=Integration
$(INTEGRATION_TEST_TARGETS): SKIP_ARGS=-focus=Integration
$(TEST_TARGETS) $(INTEGRATION_TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS) $(INTEGRATION_TEST_TARGETS): test
check test tests: fmt lint quick-test
quick-check quick-test: vendor | $(BASE) $(GINKGO) ; $(info $(M) running $(NAME:%=% )tests...) @ ## Run tests
	$Q cd $(BASE) && $(GO_TEST) $(ARGS) $(SKIP_ARGS) $(TESTPKGS)

# COVERAGE_MODE = atomic
# COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
# COVERAGE_XML = $(COVERAGE_DIR)/coverage.xml
# COVERAGE_HTML = $(COVERAGE_DIR)/index.html
# .PHONY: test-coverage test-coverage-tools
# test-coverage-tools: | $(GOCOVMERGE) $(GOCOV) $(GOCOVXML)
# test-coverage: fmt lint quick-test-coverage
# quick-test-coverage: COVERAGE_DIR := $(CURDIR)/test/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
# quick-test-coverage: vendor test-coverage-tools | $(BASE) ; $(info $(M) running coverage tests...) @ ## Run coverage tests
# 	$Q mkdir -p $(COVERAGE_DIR)/coverage
# 	$Q cd $(BASE) && for pkg in $(TESTPKGS); do \
# 		$(GO_TEST) \
# 			-covermode=$(COVERAGE_MODE) \
# 			-coverprofile="$(COVERAGE_DIR)/coverage/`echo $$pkg | tr "/" "-"`.cover" $$pkg ;\
# 	 done
# 	$Q $(GOCOVMERGE) $(COVERAGE_DIR)/coverage/*.cover > $(COVERAGE_PROFILE)
# 	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
# 	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

# Code format

.PHONY: lint
lint: vendor | $(BASE) $(GOLINT) ; $(info $(M) running golint...) @ ## Run golint
	$Q cd $(BASE) && ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=1 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info $(M) running gofmt...) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

# Dependency management

Gopkg.lock: Gopkg.toml | $(BASE) $(GODEP) ; $(info $(M) updating dependencies...)
	$Q cd $(BASE) && $(GODEP) ensure -update && $(GODEP) prune
	@touch $@
vendor: Gopkg.lock | $(BASE) $(GODEP) ; $(info $(M) retrieving dependencies...)
	$Q cd $(BASE) && $(GODEP) ensure
	@ln -sf . vendor/src
	@touch $@
.PHONY: dep-init
go-dep-init: Gopkg.lock | $(BASE) $(GODEP) ; $(info $(M) retrieving dependencies...)
	$Q cd $(BASE) && $(GODEP) init
	@touch $@

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning...)	@ ## Cleanup everything
	@rm -rf $(GOPATH)
	@rm -rf bin
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)

SWEEP?=repositories,teams
PKG_NAME=github
TEST?=./$(PKG_NAME)/...
WEBSITE_REPO=github.com/hashicorp/terraform-website

COVERAGEARGS?=-race -coverprofile=coverage.txt -covermode=atomic

# VARIABLE REFERENCE:
#
# Test-specific variables:
#   T=<pattern>       - Test name pattern (e.g., TestAccGithubRepository)
#   COV=true          - Enable coverage
#
#
# Examples:
#   make test T=TestMigrate                               # Run only schema migration unit tests
#   make test COV=true                                    # Run all unit tests with coverage
#   make testacc T=TestAccGithubRepositories\$$ COV=true  # Run only acceptance tests for a specific Test name with coverage

ifneq ($(origin T), undefined)
	RUNARGS = -run='$(T)'
endif

ifneq ($(origin COV), undefined)
	RUNARGS += $(COVERAGEARGS)
endif

default: build

tools:
	go install github.com/client9/misspell/cmd/misspell@v0.3.4
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6.0
	go install github.com/bflad/tfproviderlint/cmd/tfproviderlintx@latest

build: lintcheck
	CGO_ENABLED=0 go build -ldflags="-s -w" ./...

fmt:
	@echo "==> Fixing source code formatting..."
	golangci-lint fmt ./...

lint:
	@echo "==> Checking source code against linters and fixing..."
	golangci-lint run --fix ./...

lintcheck:
	@echo "==> Checking source code against linters..."
	golangci-lint run ./...

tf-provider-lint:
	@branch=$$(git rev-parse --abbrev-ref HEAD); \
	printf "==> Running TF provider lint on branch: \033[1m%s\033[0m...\n" "🌿 $$branch 🌿";
	# Disabled linter rules:
	# AT001: TestCase missing CheckDestroy - not yet adopted across all tests
	# AT003: should use underscores in acc test names - not yet adopted across all tests
	# AT004: provider declaration should be omitted - intentionally kept for provider configuration tests
	# AT006: acc tests should not contain multiple resource.Test() invocations - not yet adopted across all tests
	# XAT001: acceptance test should use ErrorCheck - not all resources support destroy verification
	# XR003: resource should configure Timeouts - not yet adopted across all resources
	# XR007: avoid os/exec.Command - intentionally used for GitHub CLI integration and ssh-keygen in tests
	# XS002: schema should use keys in alphabetical order - not sure we want to enforce this
	tfproviderlintx \
		-AT001=false \
		-AT003=false \
		-AT004=false \
		-AT006=false \
		-XAT001=false \
		-XR003=false \
		-XR007=false \
		-XS002=false \
		$(TEST)

test:
	@branch=$$(git rev-parse --abbrev-ref HEAD); \
	printf "==> Running unit tests on branch: \033[1m%s\033[0m...\n" "🌿 $$branch 🌿"
	CGO_ENABLED=0 go test $(TEST) \
		-timeout=30s \
		-parallel=4 \
		-v \
		-skip '^TestAcc' \
		$(RUNARGS) $(TESTARGS) \
		-count 1;

testacc:
	@branch=$$(git rev-parse --abbrev-ref HEAD); \
	printf "==> Running acceptance tests on branch: \033[1m%s\033[0m...\n" "🌿 $$branch 🌿"
	TF_ACC=1 CGO_ENABLED=0 go test $(TEST) -v -run '^TestAcc' $(RUNARGS) $(TESTARGS) -timeout 120m -count=1

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test $(TEST) -v -sweep=$(SWEEP) $(SWEEPARGS)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-lint:
	@echo "==> Checking website against linters..."
	@misspell -error -source=text website/

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc fmt lint lintcheck tools website website-lint website-test sweep tf-provider-lint

SWEEP?=repositories,teams
PKG_NAME=github
TEST?=./$(PKG_NAME)/...
WEBSITE_REPO=github.com/hashicorp/terraform-website

COVERAGEARGS?=-race -coverprofile=coverage.txt -covermode=atomic
BIN="$$(pwd -P)"/bin

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

bin/golangci-lint:
	mkdir -p $(BIN)
	GOBIN=$(BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.10.1

tools_go_files = $(shell find tools \( -name '*.go' -or -name '*.mod' -or -name '*.sum' \) -and -not -name '*_test.go' -maxdepth 4)
bin/custom-gcl: bin/golangci-lint $(tools_go_files)
	$(BIN)/golangci-lint custom --name custom-gcl --destination $(BIN)

tools: bin/custom-gcl go.sum

go.sum: go.mod $(shell find github -name '*.go')
	go mod tidy

build: lintcheck
	CGO_ENABLED=0 go build -ldflags="-s -w" ./...

fmt: tools
	@echo "==> Fixing source code formatting..."
	$(BIN)/custom-gcl fmt ./... ./tools/...

lint: tools
	@echo "==> Checking source code against linters and fixing..."
	$(BIN)/custom-gcl run --fix ./...

lintcheck: tools
	@branch=$$(git rev-parse --abbrev-ref HEAD); \
	printf "==> Checking source code against linters on branch: \033[1m%s\033[0m...\n" "🌿 $$branch 🌿"
	$(BIN)/custom-gcl run ./...

.golangci.new.yml: .golangci.yml .golangci.strict.yml
	yq eval-all 'select(fileIndex == 0) *+ select(fileIndex == 1)' .golangci.yml .golangci.strict.yml > .golangci.new.yml 

lintcheck-new: tools .golangci.new.yml
	@branch=$$(git rev-parse --abbrev-ref HEAD); \
	printf "==> Checking source code against linters on branch: \033[1m%s\033[0m...\n" "🌿 $$branch 🌿"
	$(BIN)/custom-gcl run ./... --new-from-merge-base main --config .golangci.new.yml

lintcheck-strict: tools .golangci.new.yml
	@branch=$$(git rev-parse --abbrev-ref HEAD); \
	printf "==> Checking source code against linters on branch: \033[1m%s\033[0m...\n" "🌿 $$branch 🌿"
	$(BIN)/custom-gcl run ./... --config .golangci.new.yml

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

test-tools:
	@echo "==> Running tools tests..."
	$(MAKE) test -C tools/tfproviderlint

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

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc fmt lint lintcheck lintcheck-new lintcheck-strict tools test-tools website website-test sweep

SWEEP?=repositories,teams
PKG_NAME=github
TEST?=./$(PKG_NAME)/...

COVERAGEARGS?=-race -coverprofile=coverage.txt -covermode=atomic

RUMDL_ARGS?=--output-format text

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
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

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

generatedocs:
	@cd tools; go generate ./...

validatedocs:
	@cd tools; go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs validate --provider-dir ..

fmtdocs:
	@rumdl fmt --fix ./docs

lintdocs: validatedocs
	@rumdl check $(RUMDL_ARGS) ./docs

checkdocs: generatedocs
	@git diff --quiet ||\
		{ echo "New file modification detected in the Git working tree. Please check in before commit."; git --no-pager diff --name-only | uniq | awk '{print "  - " $$0}'; \
		if [ "${CI}" = true ]; then\
			exit 1;\
		fi;}

mdfmt:
	@rumdl fmt --fix .

mdlint:
	@rumdl check $(RUMDL_ARGS) .

.PHONY: build test testacc fmt lint lintcheck tools sweep generatedocs validatedocs fmtdocs lintdocs checkdocs mdfmt mdlint

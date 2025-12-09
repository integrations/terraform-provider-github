TEST?=$$(go list ./... |grep -v 'vendor')
PKG_NAME=github

default: build

tools:
	go install github.com/golangci/misspell/cmd/misspell@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6
	go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.24

build: fmtcheck
	CGO_ENABLED=0 go build -ldflags="-s -w" ./...

fmt:
	@echo "==> Fixing source code with golangci-lint..."
	golangci-lint fmt ./...

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run ./...

test:
	CGO_ENABLED=0 go test ./...
	# commenting this out for release tooling, please run testacc instead

testacc: fmtcheck
	TF_ACC=1 CGO_ENABLED=0 go test $(TEST) -v $(TESTARGS) -timeout 120m

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	CGO_ENABLED=0 go test -c $(TEST) $(TESTARGS)

docs-generate:
	@echo "==> Generating docs..."
	tfplugindocs generate

docs-lint:
	@echo "==> Checking docs against linters..."
	@misspell -error -source=text docs/
	@tfplugindocs validate


.PHONY: build test testacc fmt fmtcheck lint tools test-compile docs-generate docs-lint

BIN := bin/terraform-provider-tester
PKG := github.com/github/terraform-provider-tester

# Version stamped into the binary. Prefers a `vX.Y.Z` tag, else `v0.1.0-dev`.
# `--dirty` marks an uncommitted working tree.
VERSION := $(shell git describe --tags --match "v*" --always --dirty 2>/dev/null || echo v0.1.0-dev)
LDFLAGS := -s -w -X $(PKG)/cli.version=$(VERSION)

.PHONY: build test vet run version checkdocs
build:
	CGO_ENABLED=0 go build -trimpath -ldflags="$(LDFLAGS)" -o $(BIN) .
test:
	go test ./...
vet:
	go vet ./...
# checkdocs verifies the docs reference only real env vars and subcommands.
# It is fork-safe and independent of the provider's tfplugindocs pipeline.
checkdocs:
	go test ./cli/ -run 'TestDocs|TestCLIReferenceCoversEverySubcommand|TestEnvScannerCatchesDrift'
run:
	go run .
version:
	@echo $(VERSION)

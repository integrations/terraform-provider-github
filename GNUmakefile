TEST?=$$(go list ./... |grep -v 'vendor')
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=github

default: build

tools:
	go install github.com/client9/misspell/cmd/misspell@v0.3.4
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6.0

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

.PHONY: build test testacc fmt fmtcheck lint tools test-compile website website-lint website-test

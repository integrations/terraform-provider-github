test_go_sums = $(patsubst %/go.mod,%/go.sum,$(wildcard checks/*/testdata/go.mod)) 

.SECONDEXPANSION:                                                                                                                                                                                                                                                                                                        
checks/%/testdata/go.sum: checks/%/testdata/go.mod $$(wildcard checks/$$*/*.go)
	@cd $(dir $@) && go mod tidy

go.sum: go.mod $(shell find checks -name '*.go' -maxdepth 2)
	@go mod tidy

test: go.sum $(test_go_sums)
	CGO_ENABLED=0 go test ./... -v -count=1 -timeout=30s -parallel=4

.PHONY: test
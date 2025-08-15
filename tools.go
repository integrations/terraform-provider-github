//go:build tools
// +build tools

package main

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"
)

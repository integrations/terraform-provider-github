//go:build tools

package main

import (
	_ "github.com/golangci/golangci-lint/v2/cmd/golangci-lint" //nolint:typecheck
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)

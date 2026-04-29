//go:build generate

package tools

import (
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)

// Format Terraform code for use in documentation.
//go:generate terraform fmt -recursive ../examples/

// Generate documentation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-dir .. --rendered-provider-name GitHub

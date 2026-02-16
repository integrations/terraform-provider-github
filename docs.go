package main

// Format Terraform code for use in documentation.
//go:generate terraform fmt -recursive examples/

// Generate documentation.
//go:generate go tool github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --rendered-provider-name=GitHub

// Check for misspellings in documentation.
//go:generate go tool github.com/client9/misspell/cmd/misspell -error -i "docs/**/*.md"

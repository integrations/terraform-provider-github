package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/integrations/terraform-provider-github/v6/github"
)

var (
	// These will be set by GoReleaser.
	version string = "dev"
	commit  string = "none"
)

func main() {
	opts := &plugin.ServeOpts{
		ProviderAddr: "registry.terraform.io/integrations/github",
		ProviderFunc: github.NewProvider(version, commit),
	}

	plugin.Serve(opts)
}

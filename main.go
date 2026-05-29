package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/integrations/terraform-provider-github/v6/github"
)

func main() {
	opts := &plugin.ServeOpts{
		ProviderAddr: "registry.terraform.io/integrations/github",
		ProviderFunc: github.NewProvider(),
	}

	plugin.Serve(opts)
}

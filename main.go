//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name github

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/integrations/terraform-provider-github/v6/github"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: github.Provider,
	})
}

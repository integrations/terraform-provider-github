package github

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubAppToken() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubAppTokenRead,

		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub App's identifier.",
			},
			"installation_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub App installation's identifier.",
			},
			"pem_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub App's PEM file content; `\\n` can be used for newlines.",
			},
			"token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The generated token from the credentials.",
			},
		},
	}
}

func dataSourceGithubAppTokenRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	appID := d.Get("app_id").(string)
	installationID := d.Get("installation_id").(string)
	pemFile := d.Get("pem_file").(string)

	// The Go encoding/pem package only decodes PEM formatted blocks
	// that contain new lines. Some platforms, like Terraform Cloud,
	// do not support new lines within Environment Variables.
	// Any occurrence of \n in the `pem_file` argument's value
	// (explicit value, or default value taken from
	// GITHUB_APP_PEM_FILE Environment Variable) is replaced with an
	// actual new line character before decoding.
	pemFile = strings.ReplaceAll(pemFile, `\n`, "\n")

	token, err := GenerateOAuthTokenFromApp(meta.(*Owner).v3client.BaseURL, appID, installationID, pemFile)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("token", token)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("id")

	return nil
}

package github

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubAppToken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubAppTokenRead,

		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["app_auth.id"],
			},
			"installation_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["app_auth.installation_id"],
			},
			"pem_file": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["app_auth.pem_file"],
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "https://api.github.com/",
				Description: descriptions["base_url"],
			},
			"token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func dataSourceGithubAppTokenRead(d *schema.ResourceData, meta interface{}) error {
	appID := d.Get("app_id").(string)
	installationID := d.Get("installation_id").(string)
	pemFile := d.Get("pem_file").(string)
	baseURL := d.Get("base_url").(string)

	// The Go encoding/pem package only decodes PEM formatted blocks
	// that contain new lines. Some platforms, like Terraform Cloud,
	// do not support new lines within Environment Variables.
	// Any occurrence of \n in the `pem_file` argument's value
	// (explicit value, or default value taken from
	// GITHUB_APP_PEM_FILE Environment Variable) is replaced with an
	// actual new line character before decoding.
	pemFile = strings.Replace(pemFile, `\n`, "\n", -1)

	token, err := GenerateOAuthTokenFromApp(baseURL, appID, installationID, pemFile)
	if err != nil {
		return err
	}
	d.Set("token", token)
	d.SetId("id")

	return nil
}

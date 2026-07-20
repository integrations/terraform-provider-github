package github

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsEnvironmentPublicKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsEnvironmentPublicKeyRead,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubActionsEnvironmentPublicKeyRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repository, _ := d.Get("repository").(string)
	envName, _ := d.Get("environment").(string)

	publicKey, _, err := client.Actions.GetEnvPublicKey(ctx, owner, repository, url.PathEscape(envName))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(publicKey.GetKeyID())
	if err := d.Set("key_id", publicKey.GetKeyID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("key", publicKey.GetKey()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

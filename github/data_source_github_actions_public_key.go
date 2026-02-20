package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsPublicKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsPublicKeyRead,

		Schema: map[string]*schema.Schema{
			"repository": {
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

func dataSourceGithubActionsPublicKeyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name

	client := meta.(*Owner).v3client

	publicKey, _, err := client.Actions.GetRepoPublicKey(ctx, owner, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(publicKey.GetKeyID())
	err = d.Set("key_id", publicKey.GetKeyID())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("key", publicKey.GetKey())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

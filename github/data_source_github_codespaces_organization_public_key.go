package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubCodespacesOrganizationPublicKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubCodespacesOrganizationPublicKeyRead,

		Schema: map[string]*schema.Schema{
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

func dataSourceGithubCodespacesOrganizationPublicKeyRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	publicKey, _, err := client.Codespaces.GetOrgPublicKey(ctx, owner)
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

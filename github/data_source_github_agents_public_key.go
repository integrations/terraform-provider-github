package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubAgentsPublicKey() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubAgentsPublicKeyRead,
		Description: "Retrieves the public key used to encrypt a repository's agent secrets.",
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the agent public key.",
			},
			"key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Base64-encoded agent public key.",
			},
		},
	}
}

func dataSourceGithubAgentsPublicKeyRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	repository, _ := d.Get("repository").(string)
	key, _, err := meta.v3client.Agents.GetRepoPublicKey(ctx, meta.name, repository)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(key.GetKeyID())
	if err := d.Set("key_id", key.GetKeyID()); err != nil {
		return diag.FromErr(err)
	}
	return diag.FromErr(d.Set("key", key.GetKey()))
}

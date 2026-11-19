package github

import (
	"context"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubAgentsOrganizationSecrets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubAgentsOrganizationSecretsRead,
		Description: "Lists organization agent secret metadata without revealing secret values.",
		Schema: map[string]*schema.Schema{
			"secrets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Organization agent secret metadata.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the secret.",
						},
						"visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Repository access to the secret: all, private, or selected.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of when the secret was created.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of when the secret was last updated.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubAgentsOrganizationSecretsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	if err := checkOrganization(m); err != nil {
		return diag.FromErr(err)
	}
	meta, _ := m.(*Owner)
	secrets := []map[string]string{}
	for secret, err := range meta.v3client.Agents.ListOrgSecretsIter(ctx, meta.name, &github.ListOptions{PerPage: meta.maxPerPage}) {
		if err != nil {
			return diag.FromErr(err)
		}
		secrets = append(secrets, map[string]string{"name": secret.Name, "visibility": secret.Visibility, "created_at": secret.CreatedAt.String(), "updated_at": secret.UpdatedAt.String()})
	}
	d.SetId(meta.name)
	return diag.FromErr(d.Set("secrets", secrets))
}

package github

import (
	"context"

	"github.com/google/go-github/v84/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationSecrets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsOrganizationSecretsRead,

		Schema: map[string]*schema.Schema{
			"secrets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"visibility": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubActionsOrganizationSecretsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	options := github.ListOptions{
		PerPage: 100,
	}

	var allSecrets []map[string]string
	for {
		secrets, resp, err := client.Actions.ListOrgSecrets(ctx, owner, &options)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, secret := range secrets.Secrets {
			new_secret := map[string]string{
				"name":       secret.Name,
				"created_at": secret.CreatedAt.String(),
				"updated_at": secret.UpdatedAt.String(),
				"visibility": secret.Visibility,
			}
			allSecrets = append(allSecrets, new_secret)

		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	d.SetId(owner)
	err := d.Set("secrets", allSecrets)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

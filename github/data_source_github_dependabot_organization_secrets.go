package github

import (
	"context"

	"github.com/google/go-github/v89/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubDependabotOrganizationSecrets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubDependabotOrganizationSecretsRead,

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

func dataSourceGithubDependabotOrganizationSecretsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	options := github.ListOptions{
		PerPage: meta.maxPerPage,
	}

	var all_secrets []map[string]string
	for {
		secrets, resp, err := client.Dependabot.ListOrgSecrets(ctx, owner, &options)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, secret := range secrets.Secrets {
			new_secret := map[string]string{
				"name":       secret.Name,
				"visibility": secret.Visibility,
				"created_at": secret.CreatedAt.String(),
				"updated_at": secret.UpdatedAt.String(),
			}
			all_secrets = append(all_secrets, new_secret)

		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	d.SetId(owner)
	err := d.Set("secrets", all_secrets)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

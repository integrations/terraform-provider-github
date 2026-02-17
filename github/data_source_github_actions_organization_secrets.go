package github

import (
	"context"

	"github.com/google/go-github/v82/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationSecrets() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the list of secrets for a GitHub organization.",
		Read:        dataSourceGithubActionsOrganizationSecretsRead,

		Schema: map[string]*schema.Schema{
			"secrets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of secrets for the organization.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the secret.",
						},
						"visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The visibility of the secret (all, private, or selected).",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of the secret creation.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of the secret last update.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubActionsOrganizationSecretsRead(d *schema.ResourceData, meta any) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	options := github.ListOptions{
		PerPage: 100,
	}

	var all_secrets []map[string]string
	for {
		secrets, resp, err := client.Actions.ListOrgSecrets(ctx, owner, &options)
		if err != nil {
			return err
		}
		for _, secret := range secrets.Secrets {
			new_secret := map[string]string{
				"name":       secret.Name,
				"created_at": secret.CreatedAt.String(),
				"updated_at": secret.UpdatedAt.String(),
				"visibility": secret.Visibility,
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
		return err
	}

	return nil
}

package github

import (
	"context"
	"net/url"

	"github.com/google/go-github/v82/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsEnvironmentSecrets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsEnvironmentSecretsRead,

		Schema: map[string]*schema.Schema{
			"full_name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name"},
				Description:   "Full name of the repository (in org/name format).",
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"full_name"},
				Description:   "The name of the repository.",
			},
			"environment": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the environment.",
			},
			"secrets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of secrets for the environment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the secret.",
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

func dataSourceGithubActionsEnvironmentSecretsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	var repoName string

	envName := d.Get("environment").(string)

	if fullName, ok := d.GetOk("full_name"); ok {
		var err error
		owner, repoName, err = splitRepoFullName(fullName.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if name, ok := d.GetOk("name"); ok {
		repoName = name.(string)
	}

	if repoName == "" {
		return diag.Errorf("one of %q or %q has to be provided", "full_name", "name")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	options := github.ListOptions{
		PerPage: maxPerPage,
	}

	var all_secrets []map[string]string
	for {
		secrets, resp, err := client.Actions.ListEnvSecrets(ctx, int(repo.GetID()), url.PathEscape(envName), &options)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, secret := range secrets.Secrets {
			new_secret := map[string]string{
				"name":       secret.Name,
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

	id, err := buildID(repoName, escapeIDPart(envName))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("secrets", all_secrets); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

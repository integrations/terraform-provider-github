package github

import (
	"context"
	"net/url"

	"github.com/google/go-github/v89/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsEnvironmentSecrets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsEnvironmentSecretsRead,

		Schema: map[string]*schema.Schema{
			"full_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"full_name", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"full_name", "name"},
			},
			"environment": {
				Type:     schema.TypeString,
				Required: true,
			},
			"secrets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
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

func dataSourceGithubActionsEnvironmentSecretsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	var repoName string
	envName, _ := d.Get("environment").(string)

	if fullName, ok := d.GetOk("full_name"); ok {
		fn, _ := fullName.(string)
		o, r, err := splitRepoFullName(fn)
		if err != nil {
			return diag.FromErr(err)
		}
		owner = o
		repoName = r
	} else {
		if name, ok := d.GetOk("name"); ok {
			repoName, _ = name.(string)
		}
	}

	var all_secrets []map[string]string
	for secret, err := range client.Actions.ListEnvSecretsIter(ctx, owner, repoName, url.PathEscape(envName), &github.ListOptions{PerPage: maxPerPage}) {
		if err != nil {
			return diag.FromErr(err)
		}

		new_secret := map[string]string{
			"name":       secret.Name,
			"created_at": secret.CreatedAt.String(),
			"updated_at": secret.UpdatedAt.String(),
		}
		all_secrets = append(all_secrets, new_secret)
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

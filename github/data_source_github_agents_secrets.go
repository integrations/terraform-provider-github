package github

import (
	"context"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubAgentsSecrets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubAgentsSecretsRead,
		Description: "Lists repository agent secret metadata without revealing secret values.",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "full_name"},
				Description:  "Repository name within the provider's owner.",
			},
			"full_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"name", "full_name"},
				Description:  "Full repository name in owner/repository format.",
			},
			"secrets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Repository agent secret metadata.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the secret.",
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

func dataSourceGithubAgentsSecretsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	owner := meta.name
	repository, _ := d.Get("name").(string)
	if fullName, ok := d.GetOk("full_name"); ok {
		var err error
		fullNameString, _ := fullName.(string)
		owner, repository, err = splitRepoFullName(fullNameString)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	secrets := []map[string]string{}
	for secret, err := range meta.v3client.Agents.ListRepoSecretsIter(ctx, owner, repository, &github.ListOptions{PerPage: meta.maxPerPage}) {
		if err != nil {
			return diag.FromErr(err)
		}
		secrets = append(secrets, map[string]string{"name": secret.Name, "created_at": secret.CreatedAt.String(), "updated_at": secret.UpdatedAt.String()})
	}
	d.SetId(owner + "/" + repository)
	return diag.FromErr(d.Set("secrets", secrets))
}

package github

import (
	"context"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRepositories() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationRepositoriesRead,

		Description: "Data source to list all organization repositories.",

		Schema: map[string]*schema.Schema{
			"repositories": {
				Description: "Organization repositories.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the repository.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"node_id": {
							Description: "Node ID of the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"visibility": {
							Description: "Visibility of the repository; one of `public`, `private`, or `internal`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"archived": {
							Description: "Whether the repository is archived.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	repos := make([]map[string]any, 0)
	for repo, err := range meta.v3client.Repositories.ListByOrgIter(ctx, meta.name, &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: maxPerPage}}) {
		if err != nil {
			return diag.FromErr(err)
		}

		r := map[string]any{
			"id":         repo.GetID(),
			"node_id":    repo.GetNodeID(),
			"name":       repo.GetName(),
			"visibility": repo.GetVisibility(),
			"archived":   repo.GetArchived(),
		}
		repos = append(repos, r)
	}

	d.SetId(meta.name)

	if err := d.Set("repositories", repos); err != nil {
		return diag.Errorf("error setting repositories: %v", err)
	}

	return nil
}

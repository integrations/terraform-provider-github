package github

import (
	"context"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRepositories() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationRepositoriesRead,
		Schema: map[string]*schema.Schema{
			"ignore_archived_repositories": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"repositories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"archived": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"visibility": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	client := meta.v3client
	org := meta.name

	options := github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	ignoreArchived := d.Get("ignore_archived_repositories").(bool)
	var allRepositories []map[string]any
	iter := client.Repositories.ListByOrgIter(ctx, org, &options)
	for repository, err := range iter {
		if err != nil {
			return diag.FromErr(err)
		}
		archived := repository.GetArchived()
		if ignoreArchived && archived {
			continue
		}
		repo := map[string]any{
			"id":         repository.GetID(),
			"node_id":    repository.GetNodeID(),
			"name":       repository.GetName(),
			"archived":   archived,
			"visibility": repository.GetVisibility(),
		}
		allRepositories = append(allRepositories, repo)
	}

	d.SetId(org)
	d.Set("repositories", allRepositories)

	return nil
}

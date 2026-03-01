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
						"repo_id": {
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

func dataSourceGithubOrganizationRepositoriesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	org := meta.(*Owner).name
	client3 := meta.(*Owner).v3client

	options := github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	ignoreArchived := d.Get("ignore_archived_repositories").(bool)
	var allRepositories []map[string]interface{}
	for {
		repositories, resp, err := client3.Repositories.ListByOrg(ctx, org, &options)
		if err != nil {
			return diag.FromErr(err)
		}
		for _, repository := range repositories {
			if ignoreArchived && repository.GetArchived() {
				continue
			}
			repo := make(map[string]interface{})
			repo["repo_id"] = repository.GetID()
			repo["node_id"] = repository.GetNodeID()
			repo["name"] = repository.GetName()
			repo["archived"] = repository.GetArchived()
			repo["visibility"] = repository.GetVisibility()
			allRepositories = append(allRepositories, repo)
		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	d.SetId(org)
	d.Set("repositories", allRepositories)

	return nil
}

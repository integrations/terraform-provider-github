package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryBranches() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the branches for a repository.",
		Read:        dataSourceGithubRepositoryBranchesRead,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository to retrieve the branches from.",
			},
			"only_protected_branches": {
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"only_non_protected_branches"},
				Description:   "If true, the branches attributes will be populated only with protected branches.",
			},
			"only_non_protected_branches": {
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"only_protected_branches"},
				Description:   "If true, the branches attributes will be populated only with non protected branches.",
			},
			"branches": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of this repository's branches.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the branch.",
						},
						"protected": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the branch is protected.",
						},
					},
				},
			},
		},
	}
}

func flattenBranches(branches []*github.Branch) []map[string]any {
	results := make([]map[string]any, 0)
	if branches == nil {
		return results
	}

	for _, branch := range branches {
		branchMap := make(map[string]any)
		branchMap["name"] = branch.GetName()
		branchMap["protected"] = branch.GetProtected()
		results = append(results, branchMap)
	}

	return results
}

func dataSourceGithubRepositoryBranchesRead(d *schema.ResourceData, meta any) error {
	ctx := context.Background()
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	onlyProtectedBranches := d.Get("only_protected_branches").(bool)
	onlyNonProtectedBranches := d.Get("only_non_protected_branches").(bool)
	var listBranchOptions *github.BranchListOptions
	if onlyProtectedBranches {
		listBranchOptions = &github.BranchListOptions{
			Protected: &onlyProtectedBranches,
		}
	} else if onlyNonProtectedBranches {
		listBranchOptions = &github.BranchListOptions{
			Protected: &onlyProtectedBranches,
		}
	} else {
		listBranchOptions = &github.BranchListOptions{}
	}

	results := make([]map[string]any, 0)
	for {
		branches, resp, err := client.Repositories.ListBranches(ctx, orgName, repoName, listBranchOptions)
		if err != nil {
			return err
		}
		results = append(results, flattenBranches(branches)...)

		if resp.NextPage == 0 {
			break
		}

		listBranchOptions.Page = resp.NextPage
	}

	d.SetId(fmt.Sprintf("%s/%s", orgName, repoName))
	err := d.Set("repository", repoName)
	if err != nil {
		return err
	}
	err = d.Set("branches", results)
	if err != nil {
		return err
	}

	return nil
}

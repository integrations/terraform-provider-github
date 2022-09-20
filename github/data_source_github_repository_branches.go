package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubRepositoryBranches() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoryBranchesRead,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"only_protected_branches": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"only_non_protected_branches": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"branches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protected": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func flattenBranches(branches []*github.Branch) []interface{} {
	if branches == nil {
		return []interface{}{}
	}

	branchList := make([]interface{}, 0, len(branches))
	for _, branch := range branches {
		branchMap := make(map[string]interface{})
		branchMap["name"] = branch.GetName()
		branchMap["protected"] = branch.GetProtected()
		branchList = append(branchList, branchMap)
	}

	return branchList
}

func dataSourceGithubRepositoryBranchesRead(d *schema.ResourceData, meta interface{}) error {
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

	branches, _, err := client.Repositories.ListBranches(context.TODO(), orgName, repoName, listBranchOptions)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", orgName, repoName))
	d.Set("repository", repoName)
	d.Set("branches", flattenBranches(branches))

	return nil
}

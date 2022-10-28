package github

import (
	"fmt"
	"sort"

	"github.com/google/go-github/v48/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubRepositories() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoriesRead,

		Schema: map[string]*schema.Schema{
			"query": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sort": {
				Type:         schema.TypeString,
				Default:      "updated",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"stars", "fork", "updated"}, false),
			},
			"include_repo_id": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"results_per_page": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      100,
				ValidateFunc: validation.IntBetween(0, 100),
			},
			"full_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"repo_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceGithubRepositoriesRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v4client
	searchQuery := d.Get("query").(string)
	sortBy := d.Get("sort").(string)
	includeRepoId := d.Get("include_repo_id").(bool)
	resultsPerPage := d.Get("results_per_page").(int)

	var query RepositoriesQuery

	variables := map[string]interface{}{
		"first":       githubv4.Int(resultsPerPage),
		"cursor":      (*githubv4.String)(nil),
		"searchQuery": githubv4.String(searchQuery),
	}

	var repos []RepositoryStruct

	for {
		err = client.Query(meta.(*Owner).StopContext, &query, variables)
		if err != nil {
			return err
		}

		additionalRepos := flattenGitHubRepos(query)
		repos = append(repos, additionalRepos...)

		if !query.Search.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Search.PageInfo.EndCursor)
	}

	// Sort repos by stars/fork/updated (default)
	if sortBy == "stars" {
		sort.Slice(repos, func(i, j int) bool {
			return repos[i].StargazerCount > repos[j].StargazerCount
		})
	} else if sortBy == "fork" {
		sort.Slice(repos, func(i, j int) bool {
			return repos[i].ForkCount > repos[j].ForkCount
		})
	} else {
		sort.Slice(repos, func(i, j int) bool {
			return repos[i].UpdatedAt.Unix() > repos[j].UpdatedAt.Unix()
		})
	}

	var names []string
	var fullNames []string
	var ids []int64

	for _, v := range repos {
		names = append(names, string(v.Name))
		fullNames = append(fullNames, string(v.NameWithOwner))
		ids = append(ids, int64(v.DatabaseID))
	}

	d.SetId(fmt.Sprintf("github-repositories/%s", searchQuery))
	d.Set("names", names)
	d.Set("full_names", fullNames)

	if includeRepoId {
		d.Set("repo_ids", ids)
	}

	return nil
}

func flattenGitHubRepos(rq RepositoriesQuery) []RepositoryStruct {
	nodes := rq.Search.Nodes
	if len(nodes) == 0 {
		return make([]RepositoryStruct, 0)
	}

	flatRepos := make([]RepositoryStruct, len(nodes))

	for i, node := range nodes {
		flatRepos[i] = node.RepositoryStruct
	}
	return flatRepos
}

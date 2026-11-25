package github

import (
	"context"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubRepositories() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoriesRead,

		Schema: map[string]*schema.Schema{
			"query": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sort": {
				Type:             schema.TypeString,
				Default:          "updated",
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"stars", "fork", "updated"}, false)),
			},
			"include_repo_id": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"results_per_page": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          100,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 1000)),
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

func dataSourceGithubRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client

	includeRepoId, _ := d.Get("include_repo_id").(bool)
	resultsPerPage, _ := d.Get("results_per_page").(int)

	query, _ := d.Get("query").(string)
	sort, _ := d.Get("sort").(string)
	opt := &github.SearchOptions{
		Sort:    sort,
		PerPage: resultsPerPage,
	}

	fullNames, names, repoIDs, err := searchGithubRepositories(ctx, client, query, opt)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(query)
	err = d.Set("full_names", fullNames)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("names", names)
	if err != nil {
		return diag.FromErr(err)
	}
	if includeRepoId {
		err = d.Set("repo_ids", repoIDs)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func searchGithubRepositories(ctx context.Context, client *github.Client, query string, opt *github.SearchOptions) ([]string, []string, []int64, error) {
	fullNames := make([]string, 0)

	names := make([]string, 0)

	repoIDs := make([]int64, 0)

	for {
		results, resp, err := client.Search.Repositories(ctx, query, opt)
		if err != nil {
			return fullNames, names, repoIDs, err
		}

		for _, repo := range results.Repositories {
			fullNames = append(fullNames, repo.GetFullName())
			names = append(names, repo.GetName())
			repoIDs = append(repoIDs, repo.GetID())
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return fullNames, names, repoIDs, nil
}

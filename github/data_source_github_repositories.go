package github

import (
	"context"
	"log"

	"github.com/google/go-github/v21/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGithubRepositories() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRepositoriesRead,

		Schema: map[string]*schema.Schema{
			"query": {
				Type:     schema.TypeString,
				Required: true,
			},
			"full_names": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"names": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceGithubRepositoriesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Organization).client

	query := d.Get("query").(string)

	log.Printf("[DEBUG] Searching for GitHub repositories: %q", query)
	fullNames, names, err := searchGithubRepositories(client, query)
	if err != nil {
		return err
	}

	d.SetId(query)
	d.Set("full_names", fullNames)
	d.Set("names", names)

	return nil
}

func searchGithubRepositories(client *github.Client, query string) ([]string, []string, error) {
	fullNames := make([]string, 0, 0)
	names := make([]string, 0, 0)

	opt := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	for {
		results, resp, err := client.Search.Repositories(context.TODO(), query, opt)
		if err != nil {
			return fullNames, names, err
		}

		for _, repo := range results.Repositories {
			fullNames = append(fullNames, repo.GetFullName())
			names = append(names, repo.GetName())
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return fullNames, names, nil
}

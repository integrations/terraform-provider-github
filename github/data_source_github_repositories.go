package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/shurcooL/githubv4"
)

const (
	ORGANIZATION_REPOSITORIES = "repositories"
)

const (
	REPOSITORY_ID   = "repository_id"
	REPOSITORY_NAME = "name"
)

type PageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}

func dataSourceGithubRepositories() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			// Computed
			ORGANIZATION_REPOSITORIES: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						REPOSITORY_ID: {
							Type:     schema.TypeString,
							Computed: true,
						},
						REPOSITORY_NAME: {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},

		Read: dataSourceGithubRepositoriesRead,
	}
}

func dataSourceGithubRepositoriesRead(d *schema.ResourceData, meta interface{}) error {
	var query struct {
		Organization struct {
			Repositories struct {
				Nodes []struct {
					ID   githubv4.ID
					Name githubv4.String
				}
				PageInfo PageInfo
			} `graphql:"repositories(first: $first, after: $cursor)"`
		} `graphql:"organization(login: $login)"`
	}
	variables := map[string]interface{}{
		"login":  githubv4.String(meta.(*Organization).name),
		"first":  githubv4.Int(100),
		"cursor": (*githubv4.String)(nil),
	}

	var allRepositories []struct {
		ID   githubv4.ID
		Name githubv4.String
	}

	ctx := context.Background()
	client := meta.(*Organization).v4client
	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return err
		}

		allRepositories = append(allRepositories, query.Organization.Repositories.Nodes...)

		if !query.Organization.Repositories.PageInfo.HasNextPage {
			break
		}

		variables["cursor"] = githubv4.NewString(query.Organization.Repositories.PageInfo.EndCursor)
	}

	var repositories []map[string]interface{}
	for _, r := range allRepositories {
		repository := make(map[string]interface{})
		repository[REPOSITORY_ID] = fmt.Sprintf("%s", r.ID)
		repository[REPOSITORY_NAME] = string(r.Name)
		repositories = append(repositories, repository)
	}
	err := d.Set(ORGANIZATION_REPOSITORIES, repositories)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/repositories", meta.(*Organization).name))

	return nil
}

package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/shurcooL/githubv4"
)

const (
	REPOSITORY_NAME = "name"
)

func dataSourceGithubRepository() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			// Input
			REPOSITORY_NAME: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "",
			},
		},

		Read: dataSourceGithubRepositoryRead,
	}
}

func dataSourceGithubRepositoryRead(d *schema.ResourceData, meta interface{}) error {
	var query struct {
		Repository struct {
			ID githubv4.ID
		} `graphql:"repository(owner:$owner, name:$name)"`
	}
	variables := map[string]interface{}{
		"owner": githubv4.String(meta.(*Organization).name),
		"name":  githubv4.String(d.Get("name").(string)),
	}

	ctx := context.Background()
	client := meta.(*Organization).v4client
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s", query.Repository.ID))

	return nil
}

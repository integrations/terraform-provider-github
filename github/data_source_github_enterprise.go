package github

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubEnterprise() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubEnterpriseRead,
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubEnterpriseRead(data *schema.ResourceData, meta interface{}) error {
	var query struct {
		Enterprise struct {
			ID          githubv4.String
			Name        githubv4.String
			Description githubv4.String
			CreatedAt   githubv4.String
			Url         githubv4.String
		} `graphql:"enterprise(slug: $slug)"`
	}

	slug := data.Get("slug").(string)
	client := meta.(*Owner).v4client
	variables := map[string]interface{}{
		"slug": githubv4.String(slug),
	}
	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return err
	}
	if query.Enterprise.ID == "" {
		return fmt.Errorf("could not find enterprise %v", slug)
	}
	data.SetId(string(query.Enterprise.ID))
	data.Set("name", query.Enterprise.Name)
	data.Set("description", query.Enterprise.Description)
	data.Set("created_at", query.Enterprise.CreatedAt)
	data.Set("url", query.Enterprise.Url)

	return nil
}

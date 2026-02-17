package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubEnterprise() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about a GitHub Enterprise.",
		Read:        dataSourceGithubEnterpriseRead,
		Schema: map[string]*schema.Schema{
			"database_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The database ID of the enterprise.",
			},
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL slug identifying the enterprise.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the enterprise.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the enterprise.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time the enterprise was created.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the enterprise.",
			},
		},
	}
}

func dataSourceGithubEnterpriseRead(data *schema.ResourceData, meta any) error {
	var query struct {
		Enterprise struct {
			ID          githubv4.String
			DatabaseId  githubv4.Int
			Name        githubv4.String
			Description githubv4.String
			CreatedAt   githubv4.String
			Url         githubv4.String
		} `graphql:"enterprise(slug: $slug)"`
	}

	slug := data.Get("slug").(string)
	client := meta.(*Owner).v4client
	variables := map[string]any{
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
	err = data.Set("name", query.Enterprise.Name)
	if err != nil {
		return err
	}
	err = data.Set("description", query.Enterprise.Description)
	if err != nil {
		return err
	}
	err = data.Set("created_at", query.Enterprise.CreatedAt)
	if err != nil {
		return err
	}
	err = data.Set("url", query.Enterprise.Url)
	if err != nil {
		return err
	}
	err = data.Set("database_id", query.Enterprise.DatabaseId)
	if err != nil {
		return err
	}

	return nil
}

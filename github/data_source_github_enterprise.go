package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubEnterprise() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubEnterpriseRead,
		Schema: map[string]*schema.Schema{
			"database_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
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

func dataSourceGithubEnterpriseRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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

	slug := d.Get("slug").(string)
	client := meta.(*Owner).v4client
	variables := map[string]any{
		"slug": githubv4.String(slug),
	}
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return diag.FromErr(err)
	}
	if query.Enterprise.ID == "" {
		return diag.Errorf("could not find enterprise %v", slug)
	}
	d.SetId(string(query.Enterprise.ID))
	err = d.Set("name", query.Enterprise.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", query.Enterprise.Description)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("created_at", query.Enterprise.CreatedAt)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("url", query.Enterprise.Url)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("database_id", query.Enterprise.DatabaseId)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

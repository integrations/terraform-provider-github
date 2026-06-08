package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubApp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubAppRead,
		Description: "Get information about an app.",

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the app.",
			},
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL-friendly name of your GitHub App.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The app's description.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The app's full name.",
			},
			"node_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Node ID of the app.",
			},
		},
	}
}

func dataSourceGithubAppRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	slug := d.Get("slug").(string)

	client := meta.(*Owner).v3client

	app, _, err := client.Apps.Get(ctx, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(app.GetID(), 10))
	err = d.Set("description", app.GetDescription())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("name", app.GetName())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("node_id", app.GetNodeID())
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

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
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The OAuth client ID of the app.",
			},
			"external_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The app's homepage URL.",
			},
			"permissions": {
				Type:        schema.TypeMap,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The permissions the app requests, keyed by permission name.",
			},
			"events": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The webhook events the app subscribes to.",
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

	permissions, err := flattenInstallationPermissions(app.Permissions)
	if err != nil {
		return diag.FromErr(err)
	}

	events := app.Events
	if events == nil {
		events = []string{}
	}

	d.SetId(strconv.FormatInt(app.GetID(), 10))
	values := map[string]any{
		"description":  app.GetDescription(),
		"name":         app.GetName(),
		"node_id":      app.GetNodeID(),
		"client_id":    app.GetClientID(),
		"external_url": app.GetExternalURL(),
		"permissions":  permissions,
		"events":       events,
	}
	for k, v := range values {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubApp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubAppRead,

		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubAppRead(d *schema.ResourceData, meta any) error {
	slug := d.Get("slug").(string)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	app, _, err := client.Apps.Get(ctx, slug)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(app.GetID(), 10))
	err = d.Set("description", app.GetDescription())
	if err != nil {
		return err
	}
	err = d.Set("name", app.GetName())
	if err != nil {
		return err
	}
	err = d.Set("node_id", app.GetNodeID())
	if err != nil {
		return err
	}

	return nil
}

package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func dataSourceGithubAppRead(d *schema.ResourceData, meta interface{}) error {
	slug := d.Get("slug").(string)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	app, _, err := client.Apps.Get(ctx, slug)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(app.GetID(), 10))
	d.Set("description", app.GetDescription())
	d.Set("name", app.GetName())
	d.Set("node_id", app.GetNodeID())

	return nil
}

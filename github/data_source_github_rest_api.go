package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubRestApi() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubRestApiRead,

		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"code": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"headers": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"body": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubRestApiRead(d *schema.ResourceData, meta interface{}) error {
	u := d.Get("endpoint").(string)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	var body map[string]interface{}

	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	resp, _ := client.Do(ctx, req, &body)

	d.SetId(resp.Header.Get("x-github-request-id"))
	d.Set("code", resp.StatusCode)
	d.Set("status", resp.Status)
	d.Set("headers", resp.Header)
	d.Set("body", body)

	return nil
}

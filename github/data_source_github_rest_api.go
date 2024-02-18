package github

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"body": {
				Type:     schema.TypeString,
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

	h, err := json.Marshal(resp.Header)
	if err != nil {
		return err
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	d.SetId(resp.Header.Get("x-github-request-id"))
	if err = d.Set("code", resp.StatusCode); err != nil {
		return err
	}
	if err = d.Set("status", resp.Status); err != nil {
		return err
	}
	if err = d.Set("headers", string(h)); err != nil {
		return err
	}
	if err = d.Set("body", string(b)); err != nil {
		return err
	}
	return nil
}

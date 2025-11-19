package github

import (
	"context"
	"encoding/json"
	"io"

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

func dataSourceGithubRestApiRead(d *schema.ResourceData, meta any) error {
	u := d.Get("endpoint").(string)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(ctx, req, nil)
	if err != nil && resp.StatusCode != 404 {
		return err
	}

	h, err := json.Marshal(resp.Header)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)
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

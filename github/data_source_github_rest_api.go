package github

import (
	"context"
	"encoding/json"
	"io"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRestApi() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRestApiRead,

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

func dataSourceGithubRestApiRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	u := d.Get("endpoint").(string)

	client := meta.(*Owner).v3client

	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := client.Do(ctx, req, nil)
	if err != nil && resp.StatusCode != 404 {
		return diag.FromErr(err)
	}

	h, err := json.Marshal(resp.Header)
	if err != nil {
		return diag.FromErr(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Header.Get("x-github-request-id"))
	if err = d.Set("code", resp.StatusCode); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("status", resp.Status); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("headers", string(h)); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("body", string(b)); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

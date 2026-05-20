package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubClientConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubClientConfigRead,

		Schema: map[string]*schema.Schema{
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_organization": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"base_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubClientConfigRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner := meta.(*Owner)
	client := owner.v3client

	if err := d.Set("owner", owner.name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("is_organization", owner.IsOrganization); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("base_url", client.BaseURL.String()); err != nil {
		return diag.FromErr(err)
	}

	var username string
	if user, _, err := client.Users.Get(ctx, ""); err == nil {
		username = user.GetLogin()
	}
	if err := d.Set("username", username); err != nil {
		return diag.FromErr(err)
	}

	id := owner.name
	if id == "" {
		id = username
	}
	if id == "" {
		id = client.BaseURL.String()
	}
	d.SetId(id)

	return nil
}

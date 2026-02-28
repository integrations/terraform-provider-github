package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubMembership() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubMembershipRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"organization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceGithubMembershipRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	username := d.Get("username").(string)

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	if configuredOrg := d.Get("organization").(string); configuredOrg != "" {
		orgName = configuredOrg
	}

	membership, resp, err := client.Organizations.GetOrgMembership(ctx,
		username, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(membership.GetOrganization().GetLogin(), membership.GetUser().GetLogin()))

	err = d.Set("username", membership.GetUser().GetLogin())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("role", membership.GetRole())
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("etag", resp.Header.Get("ETag"))
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("state", membership.GetState())
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

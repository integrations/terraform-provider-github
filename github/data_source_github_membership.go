package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubMembership() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubMembershipRead,

		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The username to lookup in the organization.",
			},
			"organization": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The organization to check for the above username.",
			},
			"role": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The role the user has within the organization. Can be 'admin' or 'member'.",
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An etag representing the membership object.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the user's membership in the organization. Can be 'active' or 'pending'.",
			},
		},
	}
}

func dataSourceGithubMembershipRead(d *schema.ResourceData, meta any) error {
	username := d.Get("username").(string)

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	if configuredOrg := d.Get("organization").(string); configuredOrg != "" {
		orgName = configuredOrg
	}

	ctx := context.Background()

	membership, resp, err := client.Organizations.GetOrgMembership(ctx,
		username, orgName)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(membership.GetOrganization().GetLogin(), membership.GetUser().GetLogin()))

	err = d.Set("username", membership.GetUser().GetLogin())
	if err != nil {
		return err
	}
	err = d.Set("role", membership.GetRole())
	if err != nil {
		return err
	}
	err = d.Set("etag", resp.Header.Get("ETag"))
	if err != nil {
		return err
	}
	err = d.Set("state", membership.GetState())
	if err != nil {
		return err
	}
	return nil
}

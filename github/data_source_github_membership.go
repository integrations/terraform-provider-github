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
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"username", "user_id"},
				Description:  "The username (login) to lookup in the organization. Exactly one of `username` or `user_id` must be set.",
			},
			"user_id": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"username", "user_id"},
				Description:  "The GitHub numeric user ID to lookup in the organization. Stable across username changes. Exactly one of `username` or `user_id` must be set.",
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
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	if configuredOrg := d.Get("organization").(string); configuredOrg != "" {
		orgName = configuredOrg
	}

	// Resolve to username (login). If user_id is provided, resolve it via
	// GET /user/{id} since GitHub's membership endpoints only accept the
	// username. This makes the data source robust against username changes.
	var username string
	if v, ok := d.GetOk("user_id"); ok {
		userID := int64(v.(int))
		user, _, err := client.Users.GetByID(ctx, userID)
		if err != nil {
			return diag.FromErr(err)
		}
		username = user.GetLogin()
	} else {
		username = d.Get("username").(string)
	}

	membership, resp, err := client.Organizations.GetOrgMembership(ctx, username, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(membership.GetOrganization().GetLogin(), membership.GetUser().GetLogin()))

	if err = d.Set("username", membership.GetUser().GetLogin()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("user_id", membership.GetUser().GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("role", membership.GetRole()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("etag", resp.Header.Get("ETag")); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("state", membership.GetState()); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

package github

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseTeamMembership() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieves information about a user's membership in a GitHub enterprise team.",
		ReadContext: dataSourceGithubEnterpriseTeamMembershipRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug of the enterprise.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"team_slug": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug of the enterprise team.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"username": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The username of the user.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"user_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The ID of the user.",
			},
		},
	}
}

func dataSourceGithubEnterpriseTeamMembershipRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))
	teamSlug := strings.TrimSpace(d.Get("team_slug").(string))
	username := strings.TrimSpace(d.Get("username").(string))

	// Get the membership using the SDK
	user, _, err := client.Enterprise.GetTeamMembership(ctx, enterpriseSlug, teamSlug, username)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildEnterpriseTeamMembershipID(enterpriseSlug, teamSlug, username))
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("team_slug", teamSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("username", username); err != nil {
		return diag.FromErr(err)
	}
	if user != nil && user.ID != nil {
		if err := d.Set("user_id", int(*user.ID)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

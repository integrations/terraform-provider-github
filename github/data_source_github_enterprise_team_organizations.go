package github

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseTeamOrganizations() *schema.Resource {
	return &schema.Resource{
		Description: "Lists organizations assigned to a GitHub enterprise team.",
		ReadContext: dataSourceGithubEnterpriseTeamOrganizationsRead,

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
			"organization_slugs": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Set of organization slugs that the enterprise team is assigned to.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},
		},
	}
}

func dataSourceGithubEnterpriseTeamOrganizationsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))
	teamSlug := strings.TrimSpace(d.Get("team_slug").(string))
	orgs, err := listAllEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, teamSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	slugs := make([]string, 0, len(orgs))
	for _, org := range orgs {
		if org.Login != nil && *org.Login != "" {
			slugs = append(slugs, *org.Login)
		}
	}

	d.SetId(buildEnterpriseTeamOrganizationsID(enterpriseSlug, teamSlug))
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("team_slug", teamSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("organization_slugs", slugs); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

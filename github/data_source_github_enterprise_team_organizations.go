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
			"enterprise_team": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug or ID of the enterprise team.",
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
	enterpriseTeam := strings.TrimSpace(d.Get("enterprise_team").(string))
	orgs, err := listEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, enterpriseTeam)
	if err != nil {
		return diag.FromErr(err)
	}

	slugs := make([]string, 0, len(orgs))
	for _, org := range orgs {
		if org.Login != "" {
			slugs = append(slugs, org.Login)
		}
	}

	d.SetId(buildSlashTwoPartID(enterpriseSlug, enterpriseTeam))
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enterprise_team", enterpriseTeam); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("organization_slugs", slugs); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

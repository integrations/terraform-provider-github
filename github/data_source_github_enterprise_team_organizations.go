package github

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseTeamOrganizations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubEnterpriseTeamOrganizationsRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"enterprise_team": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug or ID of the enterprise team.",
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

func dataSourceGithubEnterpriseTeamOrganizationsRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))
	enterpriseTeam := strings.TrimSpace(d.Get("enterprise_team").(string))
	if enterpriseSlug == "" {
		return fmt.Errorf("enterprise_slug must not be empty")
	}
	if enterpriseTeam == "" {
		return fmt.Errorf("enterprise_team must not be empty")
	}

	ctx := context.Background()
	orgs, err := listEnterpriseTeamOrganizations(ctx, client, enterpriseSlug, enterpriseTeam)
	if err != nil {
		return err
	}

	slugs := make([]string, 0, len(orgs))
	for _, org := range orgs {
		if org.Login != "" {
			slugs = append(slugs, org.Login)
		}
	}

	d.SetId(buildSlashTwoPartID(enterpriseSlug, enterpriseTeam))
	_ = d.Set("enterprise_slug", enterpriseSlug)
	_ = d.Set("enterprise_team", enterpriseTeam)
	_ = d.Set("organization_slugs", slugs)
	return nil
}

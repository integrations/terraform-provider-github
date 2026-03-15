package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationSecurityManagers() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This data source is deprecated in favour of using the github_organization_role_teams data source.",

		ReadContext: dataSourceGithubOrganizationSecurityManagersRead,

		Schema: map[string]*schema.Schema{
			"teams": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "Unique identifier of the team.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"slug": {
							Description: "Name based identifier of the team.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name of the team.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"permission": {
							Description: "Permission that the team will have for its repositories.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationSecurityManagersRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	allTeams := make([]any, 0)

	//nolint:staticcheck // SA1019: ListSecurityManagerTeams is deprecated but still needed for legacy compatibility
	teams, _, err := client.Organizations.ListSecurityManagerTeams(ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, team := range teams {
		t := map[string]any{
			"id":         int(team.GetID()),
			"slug":       team.GetSlug(),
			"name":       team.GetName(),
			"permission": team.GetPermission(),
		}
		allTeams = append(allTeams, t)
	}

	d.SetId(fmt.Sprintf("%s/github-org-security-managers", orgName))
	if err := d.Set("teams", allTeams); err != nil {
		return diag.Errorf("error setting teams: %v", err)
	}

	return nil
}

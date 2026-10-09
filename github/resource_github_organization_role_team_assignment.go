package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationRoleTeamAssignment() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This resource is deprecated in favor of the github_organization_role_team resource.",

		CreateContext: resourceGithubOrganizationRoleTeamAssignmentCreate,
		ReadContext:   resourceGithubOrganizationRoleTeamAssignmentRead,
		DeleteContext: resourceGithubOrganizationRoleTeamAssignmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub team slug.",
				ForceNew:    true,
			},
			"role_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GitHub organization role id",
				ForceNew:    true,
			},
		},
	}
}

func resourceGithubOrganizationRoleTeamAssignmentCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.v3client
	orgName := meta.name

	teamSlug := d.Get("team_slug").(string)
	roleIDString := d.Get("role_id").(string)

	roleID, err := strconv.ParseInt(roleIDString, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Organizations.AssignOrgRoleToTeam(ctx, orgName, teamSlug, roleID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(teamSlug, roleIDString))
	return resourceGithubOrganizationRoleTeamAssignmentRead(ctx, d, meta)
}

func resourceGithubOrganizationRoleTeamAssignmentRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	teamSlug, roleIDString, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	roleID, err := strconv.ParseInt(roleIDString, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	// There is no api for checking a specific team role assignment, so instead we iterate over all teams assigned to the role
	// go-github pagination (https://github.com/google/go-github?tab=readme-ov-file#pagination)
	options := &github.ListOptions{
		PerPage: meta.maxPerPage,
	}
	var foundTeam *github.Team
	for {
		teams, resp, err := client.Organizations.ListTeamsAssignedToOrgRole(ctx, orgName, roleID, options)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, team := range teams {
			if team.GetSlug() == teamSlug {
				foundTeam = team
				break
			}
		}

		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	if foundTeam == nil {
		tflog.Warn(ctx, "Removing team organization role association from state because it no longer exists in GitHub", map[string]any{"resource_id": d.Id()})
		d.SetId("")
		return nil
	}

	if err = d.Set("team_slug", teamSlug); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("role_id", roleIDString); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRoleTeamAssignmentDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.v3client
	orgName := meta.name

	teamSlug, roleIDString, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	roleID, err := strconv.ParseInt(roleIDString, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Organizations.RemoveOrgRoleFromTeam(ctx, orgName, teamSlug, roleID)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

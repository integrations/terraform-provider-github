package github

import (
	"context"
	"errors"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationSecurityManager() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This resource is deprecated in favor of the github_organization_role_team resource.",

		CreateContext: resourceGithubOrganizationSecurityManagerCreate,
		ReadContext:   resourceGithubOrganizationSecurityManagerRead,
		UpdateContext: resourceGithubOrganizationSecurityManagerUpdate,
		DeleteContext: resourceGithubOrganizationSecurityManagerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the team to manage.",
			},
		},
	}
}

func getSecurityManagerRole(client *github.Client, ctx context.Context, orgName string) (*github.CustomOrgRole, error) {
	roles, _, err := client.Organizations.ListRoles(ctx, orgName)
	if err != nil {
		return nil, err
	}

	for _, role := range roles.CustomRepoRoles {
		if *role.Name == "security_manager" {
			return role, nil
		}
	}

	return nil, errors.New("security manager role not found")
}

func resourceGithubOrganizationSecurityManagerCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	orgName := meta.name
	teamSlug := d.Get("team_slug").(string)

	client := meta.v3client

	team, _, err := client.Teams.GetTeamBySlug(ctx, orgName, teamSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	smRole, err := getSecurityManagerRole(client, ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Organizations.AssignOrgRoleToTeam(ctx, orgName, teamSlug, smRole.GetID())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(team.GetID(), 10))

	return resourceGithubOrganizationSecurityManagerRead(ctx, d, meta)
}

func resourceGithubOrganizationSecurityManagerRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	orgName := meta.name
	client := meta.v3client

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	teamId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	ctx = context.WithValue(ctx, ctxId, d.Id())

	smRole, err := getSecurityManagerRole(client, ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	// There is no endpoint for getting a single security manager team, so get the list and filter.
	options := &github.ListOptions{PerPage: meta.maxPerPage}
	var smTeam *github.Team = nil
	for {
		smTeams, resp, err := client.Organizations.ListTeamsAssignedToOrgRole(ctx, orgName, smRole.GetID(), options)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, t := range smTeams {
			if t.GetID() == teamId {
				smTeam = t
				break
			}
		}

		// Break when we've found the team or there are no more pages.
		if smTeam != nil || resp.NextPage == 0 {
			break
		}

		options.Page = resp.NextPage
	}

	if smTeam == nil {
		tflog.Warn(ctx, "Removing organization security manager team from state because it no longer exists in GitHub", map[string]any{"resource_id": d.Id()})
		d.SetId("")
		return nil
	}

	if err = d.Set("team_slug", smTeam.GetSlug()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationSecurityManagerUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	orgId := meta.id
	orgName := meta.name
	teamId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.v3client
	ctx = context.WithValue(ctx, ctxId, d.Id())

	team, _, err := client.Teams.GetTeamByID(ctx, orgId, teamId)
	if err != nil {
		return diag.FromErr(err)
	}

	smRole, err := getSecurityManagerRole(client, ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	// Adding the same team is a no-op.
	_, err = client.Organizations.AssignOrgRoleToTeam(ctx, orgName, team.GetSlug(), smRole.GetID())
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubOrganizationSecurityManagerRead(ctx, d, meta)
}

func resourceGithubOrganizationSecurityManagerDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	orgName := meta.name
	teamSlug := d.Get("team_slug").(string)

	client := meta.v3client
	ctx = context.WithValue(ctx, ctxId, d.Id())

	smRole, err := getSecurityManagerRole(client, ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Organizations.RemoveOrgRoleFromTeam(ctx, orgName, teamSlug, smRole.GetID())
	return diag.FromErr(err)
}

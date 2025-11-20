package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationRoleTeamAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationRoleTeamAssignmentCreate,
		Read:   resourceGithubOrganizationRoleTeamAssignmentRead,
		Delete: resourceGithubOrganizationRoleTeamAssignmentDelete,
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

func resourceGithubOrganizationRoleTeamAssignmentCreate(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	teamSlug := d.Get("team_slug").(string)
	roleIDString := d.Get("role_id").(string)

	roleID, err := strconv.ParseInt(roleIDString, 10, 64)
	if err != nil {
		return err
	}

	_, err = client.Organizations.AssignOrgRoleToTeam(ctx, orgName, teamSlug, roleID)
	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(teamSlug, roleIDString))
	return resourceGithubOrganizationRoleTeamAssignmentRead(d, meta)
}

func resourceGithubOrganizationRoleTeamAssignmentRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	ctx := context.Background()
	orgName := meta.(*Owner).name

	teamSlug, roleIDString, err := parseTwoPartID(d.Id(), "team_slug", "role_id")
	if err != nil {
		return err
	}

	roleID, err := strconv.ParseInt(roleIDString, 10, 64)
	if err != nil {
		return err
	}

	// There is no api for checking a specific team role assignment, so instead we iterate over all teams assigned to the role
	// go-github pagination (https://github.com/google/go-github?tab=readme-ov-file#pagination)
	options := &github.ListOptions{
		PerPage: 100,
	}
	var foundTeam *github.Team
	for {
		teams, resp, err := client.Organizations.ListTeamsAssignedToOrgRole(ctx, orgName, roleID, options)
		if err != nil {
			return err
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
		log.Printf("[WARN] Removing team organization role association %s from state because it no longer exists in GitHub", d.Id())
		d.SetId("")
		return nil
	}

	if err = d.Set("team_slug", teamSlug); err != nil {
		return err
	}
	if err = d.Set("role_id", roleIDString); err != nil {
		return err
	}

	return nil
}

func resourceGithubOrganizationRoleTeamAssignmentDelete(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := context.Background()

	teamSlug, roleIDString, err := parseTwoPartID(d.Id(), "team_slug", "role_id")
	if err != nil {
		return err
	}

	roleID, err := strconv.ParseInt(roleIDString, 10, 64)
	if err != nil {
		return err
	}

	_, err = client.Organizations.RemoveOrgRoleFromTeam(ctx, orgName, teamSlug, roleID)
	if err != nil {
		return err
	}

	return nil
}

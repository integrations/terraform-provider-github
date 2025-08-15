package github

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/google/go-github/v74/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationSecurityManager() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubOrganizationSecurityManagerCreate,
		Read:   resourceGithubOrganizationSecurityManagerRead,
		Update: resourceGithubOrganizationSecurityManagerUpdate,
		Delete: resourceGithubOrganizationSecurityManagerDelete,
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

func getSecurityManagerRole(client *github.Client, ctx context.Context, orgName string) (*github.CustomOrgRoles, error) {
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

func resourceGithubOrganizationSecurityManagerCreate(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	orgName := meta.(*Owner).name
	teamSlug := d.Get("team_slug").(string)

	client := meta.(*Owner).v3client
	ctx := context.Background()

	team, _, err := client.Teams.GetTeamBySlug(ctx, orgName, teamSlug)
	if err != nil {
		return err
	}

	smRole, err := getSecurityManagerRole(client, ctx, orgName)
	if err != nil {
		return err
	}

	_, err = client.Organizations.AssignOrgRoleToTeam(ctx, orgName, teamSlug, smRole.GetID())
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(team.GetID(), 10))

	return resourceGithubOrganizationSecurityManagerRead(d, meta)
}

func resourceGithubOrganizationSecurityManagerRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	orgName := meta.(*Owner).name
	teamId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	smRole, err := getSecurityManagerRole(client, ctx, orgName)
	if err != nil {
		return err
	}

	// There is no endpoint for getting a single security manager team, so get the list and filter.
	options := &github.ListOptions{PerPage: 100}
	var smTeam *github.Team = nil
	for {
		smTeams, resp, err := client.Organizations.ListTeamsAssignedToOrgRole(ctx, orgName, smRole.GetID(), options)
		if err != nil {
			return err
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
		log.Printf("[WARN] Removing organization security manager team %s from state because it no longer exists in GitHub", d.Id())
		d.SetId("")
		return nil
	}

	if err = d.Set("team_slug", smTeam.GetSlug()); err != nil {
		return err
	}

	return nil
}

func resourceGithubOrganizationSecurityManagerUpdate(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	orgId := meta.(*Owner).id
	orgName := meta.(*Owner).name
	teamId, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	team, _, err := client.Teams.GetTeamByID(ctx, orgId, teamId)
	if err != nil {
		return err
	}

	smRole, err := getSecurityManagerRole(client, ctx, orgName)
	if err != nil {
		return err
	}

	// Adding the same team is a no-op.
	_, err = client.Organizations.AssignOrgRoleToTeam(ctx, orgName, team.GetSlug(), smRole.GetID())
	if err != nil {
		return err
	}

	return resourceGithubOrganizationSecurityManagerRead(d, meta)
}

func resourceGithubOrganizationSecurityManagerDelete(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	orgName := meta.(*Owner).name
	teamSlug := d.Get("team_slug").(string)

	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	smRole, err := getSecurityManagerRole(client, ctx, orgName)
	if err != nil {
		return err
	}

	_, err = client.Organizations.RemoveOrgRoleFromTeam(ctx, orgName, teamSlug, smRole.GetID())
	return err
}

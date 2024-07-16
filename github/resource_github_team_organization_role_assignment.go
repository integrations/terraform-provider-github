package github

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	abs "github.com/microsoft/kiota-abstractions-go"
	"github.com/octokit/go-sdk/pkg/github/models"
	"github.com/octokit/go-sdk/pkg/github/orgs"
)

func resourceGithubTeamOrganizationRoleAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubTeamOrganizationRoleAssignmentCreate,
		Read:   resourceGithubTeamOrganizationRoleAssignmentRead,
		Delete: resourceGithubTeamOrganizationRoleAssignmentDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"team_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The organization custom repository role to create.",
				ForceNew:    true,
			},
			"role_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The base role for the custom repository role.",
				ForceNew:    true,
			},
		},
	}
}

func newOctokitClientDefaultRequestConfig() *abs.RequestConfiguration[abs.DefaultQueryParameters] {
	headers := abs.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")
	_ = headers.TryAdd("X-GitHub-Api-Version", "2022-11-28")

	return &abs.RequestConfiguration[abs.DefaultQueryParameters]{
		QueryParameters: &abs.DefaultQueryParameters{},
		Headers:         headers,
	}
}

func resourceGithubTeamOrganizationRoleAssignmentCreate(d *schema.ResourceData, meta interface{}) error {
	octokitClient := meta.(*Owner).octokitClient

	orgName := meta.(*Owner).name
	ctx := context.Background()

	teamSlug := d.Get("team_slug").(string)
	roleIDString := d.Get("role_id").(string)
	roleID, err := strconv.ParseInt(roleIDString, 10, 32)

	if err != nil {
		return err
	}

	defaultRequestConfig := newOctokitClientDefaultRequestConfig()
	err = octokitClient.Orgs().ByOrg(orgName).OrganizationRoles().Teams().ByTeam_slug(teamSlug).ByRole_id(int32(roleID)).Put(ctx, defaultRequestConfig)
	if err != nil {
		return err
	}

	d.SetId(buildThreePartID(orgName, teamSlug, roleIDString))
	return resourceGithubTeamOrganizationRoleAssignmentRead(d, meta)
}

func resourceGithubTeamOrganizationRoleAssignmentRead(d *schema.ResourceData, meta interface{}) error {
	octokitClient := meta.(*Owner).octokitClient

	ctx := context.Background()
	orgName := meta.(*Owner).name

	teamSlug := d.Get("team_slug").(string)
	roleIDString := d.Get("role_id").(string)
	roleID, err := strconv.ParseInt(roleIDString, 10, 32)

	headers := abs.NewRequestHeaders()
	_ = headers.TryAdd("Accept", "application/vnd.github.v3+json")
	_ = headers.TryAdd("X-GitHub-Api-Version", "2022-11-28")
	requestConfig := &abs.RequestConfiguration[orgs.ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters]{
		QueryParameters: &orgs.ItemOrganizationRolesItemTeamsRequestBuilderGetQueryParameters{},
		Headers:         headers,
	}

	// TODO: Unsure if this handles pagination
	// If it doesn't the go-github sdk does
	// client := meta.(*Owner).v3client
	// client.Organizations.ListTeamsAssignedToOrgRole()
	roleTeamAssignments, err := octokitClient.Orgs().ByOrg(orgName).OrganizationRoles().ByRole_id(int32(roleID)).Teams().Get(ctx, requestConfig)
	if err != nil {
		return err
	}

	var roleAssignment models.TeamRoleAssignmentable
	for _, roleTeamAssignment := range roleTeamAssignments {
		if *roleTeamAssignment.GetSlug() == teamSlug {
			roleAssignment = roleTeamAssignment
		}
	}

	if roleAssignment == nil {
		log.Printf("[WARN] Removing team organization role association %s from state because it no longer exists in GitHub", d.Id())
		d.SetId("")
		return nil
	}

	if err = d.Set("team_slug", roleAssignment.GetSlug()); err != nil {
		return err
	}
	if err = d.Set("role_id", roleIDString); err != nil {
		return err
	}

	return nil
}

func resourceGithubTeamOrganizationRoleAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
	octokitClient := meta.(*Owner).octokitClient

	orgName := meta.(*Owner).name
	ctx := context.Background()

	teamSlug := d.Get("team_slug").(string)
	roleIDString := d.Get("role_id").(string)
	roleID, err := strconv.ParseInt(roleIDString, 10, 32)

	if err != nil {
		return err
	}

	defaultRequestConfig := newOctokitClientDefaultRequestConfig()
	err = octokitClient.Orgs().ByOrg(orgName).OrganizationRoles().Teams().ByTeam_slug(teamSlug).ByRole_id(int32(roleID)).Delete(ctx, defaultRequestConfig)
	if err != nil {
		return err
	}

	return nil
}

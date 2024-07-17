package github

import (
	"context"
	"log"
	"strconv"

	"github.com/google/go-github/v62/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	abs "github.com/microsoft/kiota-abstractions-go"
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
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

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
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	restClient := meta.(*Owner).v3client

	ctx := context.Background()
	orgName := meta.(*Owner).name

	teamSlug := d.Get("team_slug").(string)
	roleIDString := d.Get("role_id").(string)
	roleID, err := strconv.ParseInt(roleIDString, 10, 32)
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
		teams, resp, err := restClient.Organizations.ListTeamsAssignedToOrgRole(ctx, orgName, roleID, options)
		if err != nil {
			return err
		}

		for _, team := range teams {
			if team.GetSlug() == teamSlug {
				foundTeam = team
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

	if err = d.Set("team_slug", foundTeam.GetSlug()); err != nil {
		return err
	}
	if err = d.Set("role_id", roleIDString); err != nil {
		return err
	}

	return nil
}

func resourceGithubTeamOrganizationRoleAssignmentDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

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

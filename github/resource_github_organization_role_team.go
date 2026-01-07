package github

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubOrganizationRoleTeam() *schema.Resource {
	return &schema.Resource{
		Description: "Manage an association between an organization role and a team.",

		CreateContext: resourceGithubOrganizationRoleTeamCreate,
		ReadContext:   resourceGithubOrganizationRoleTeamRead,
		DeleteContext: resourceGithubOrganizationRoleTeamDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "The ID of the organization role.",
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
			},
			"team_slug": {
				Description: "The slug of the team name.",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceGithubOrganizationRoleTeamCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId := int64(d.Get("role_id").(int))
	teamSlug := d.Get("team_slug").(string)

	_, err = client.Organizations.AssignOrgRoleToTeam(ctx, orgName, teamSlug, roleId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(strconv.FormatInt(roleId, 10), teamSlug))

	return nil
}

func resourceGithubOrganizationRoleTeamRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleIdString, teamSlug, err := parseTwoPartID(d.Id(), "role_id", "team_slug")
	if err != nil {
		return diag.FromErr(err)
	}
	roleId, err := strconv.ParseInt(roleIdString, 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	opts := &github.ListOptions{
		PerPage: maxPerPage,
	}

	var team *github.Team
	for {
		teams, resp, err := client.Organizations.ListTeamsAssignedToOrgRole(ctx, orgName, roleId, opts)
		if err != nil {
			return diag.FromErr(err)
		}

		for _, t := range teams {
			if t.GetSlug() == teamSlug {
				team = t
				break
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	if team == nil {
		log.Printf("[INFO] Removing organization role team (%d:%s) from state because it no longer exists in GitHub", roleId, teamSlug)
		d.SetId("")
		return nil
	}

	if err = d.Set("role_id", roleId); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("team_slug", teamSlug); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubOrganizationRoleTeamDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId := int64(d.Get("role_id").(int))
	teamSlug := d.Get("team_slug").(string)

	_, err = client.Organizations.RemoveOrgRoleFromTeam(ctx, orgName, teamSlug, roleId)
	if err != nil {
		return diag.FromErr(fmt.Errorf("Error deleting organization role team %d %s: %w", roleId, teamSlug, err))
	}

	return nil
}

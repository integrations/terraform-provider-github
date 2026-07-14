package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

type projectV2TeamNode struct {
	ID           githubv4.String
	Slug         githubv4.String
	Organization struct {
		Login githubv4.String
	}
}

func resourceGithubTeamProject() *schema.Resource {
	return &schema.Resource{
		Description:   "Links a team to a GitHub Projects V2 project.",
		CreateContext: resourceGithubTeamProjectCreate,
		ReadContext:   resourceGithubTeamProjectRead,
		DeleteContext: resourceGithubTeamProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubTeamProjectImport,
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Node ID of the Projects V2 project.",
			},
			"organization": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Organization login. Defaults to the provider owner.",
			},
			"team_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Team slug.",
			},
		},
	}
}

func resourceGithubTeamProjectCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	team, err := queryProjectV2TeamBySlug(ctx, meta.(*Owner).v4client, projectV2TeamOrganization(d, meta), d.Get("team_slug").(string))
	if err != nil {
		return diag.FromErr(err)
	}
	var mutation struct {
		LinkProjectV2ToTeam struct {
			Team projectV2TeamNode
		} `graphql:"linkProjectV2ToTeam(input: $input)"`
	}
	input := githubv4.LinkProjectV2ToTeamInput{ProjectID: githubv4.ID(d.Get("project_id").(string)), TeamID: githubv4.ID(team.ID)}
	if err := meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil); err != nil {
		return diag.FromErr(err)
	}
	id, err := buildID(d.Get("project_id").(string), string(team.ID))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	return resourceGithubTeamProjectRead(ctx, d, meta)
}

func resourceGithubTeamProjectRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID, teamID, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	client := meta.(*Owner).v4client
	linked, err := projectV2HasTeam(ctx, client, projectID, teamID)
	if isProjectV2NotFound(err) || (err == nil && !linked) {
		tflog.Info(ctx, "Removing team project link from state because it no longer exists in GitHub", map[string]any{"id": d.Id()})
		d.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}
	team, err := queryProjectV2TeamByID(ctx, client, teamID)
	if err != nil {
		return diag.FromErr(err)
	}
	for key, value := range map[string]any{"project_id": projectID, "organization": team.Organization.Login, "team_slug": team.Slug} {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(fmt.Errorf("setting %s: %w", key, err))
		}
	}
	return nil
}

func resourceGithubTeamProjectDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	projectID, teamID, err := parseID2(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	var mutation struct {
		UnlinkProjectV2FromTeam struct {
			Team projectV2TeamNode
		} `graphql:"unlinkProjectV2FromTeam(input: $input)"`
	}
	input := githubv4.UnlinkProjectV2FromTeamInput{ProjectID: githubv4.ID(projectID), TeamID: githubv4.ID(teamID)}
	err = meta.(*Owner).v4client.Mutate(ctx, &mutation, input, nil)
	if err != nil && !isProjectV2NotFound(err) {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubTeamProjectImport(_ context.Context, d *schema.ResourceData, _ any) ([]*schema.ResourceData, error) {
	if _, _, err := parseID2(d.Id()); err != nil {
		return nil, err
	}
	return []*schema.ResourceData{d}, nil
}

func projectV2TeamOrganization(d *schema.ResourceData, meta any) string {
	if organization := d.Get("organization").(string); organization != "" {
		return organization
	}
	return meta.(*Owner).name
}

func queryProjectV2TeamBySlug(ctx context.Context, client *githubv4.Client, organization, slug string) (projectV2TeamNode, error) {
	var query struct {
		Organization struct {
			Team projectV2TeamNode `graphql:"team(slug: $slug)"`
		} `graphql:"organization(login: $organization)"`
	}
	err := client.Query(ctx, &query, map[string]any{"organization": githubv4.String(organization), "slug": githubv4.String(slug)})
	return query.Organization.Team, err
}

func queryProjectV2TeamByID(ctx context.Context, client *githubv4.Client, id string) (projectV2TeamNode, error) {
	var query struct {
		Node struct {
			Team projectV2TeamNode `graphql:"... on Team"`
		} `graphql:"node(id: $id)"`
	}
	err := client.Query(ctx, &query, map[string]any{"id": githubv4.ID(id)})
	return query.Node.Team, err
}

func projectV2HasTeam(ctx context.Context, client *githubv4.Client, projectID, teamID string) (bool, error) {
	var after *githubv4.String
	for {
		var query struct {
			Node struct {
				Project struct {
					Teams struct {
						Nodes    []struct{ ID githubv4.String }
						PageInfo PageInfo
					} `graphql:"teams(first: 100, after: $after)"`
				} `graphql:"... on ProjectV2"`
			} `graphql:"node(id: $id)"`
		}
		err := client.Query(ctx, &query, map[string]any{"id": githubv4.ID(projectID), "after": after})
		if err != nil {
			return false, err
		}
		for _, team := range query.Node.Project.Teams.Nodes {
			if string(team.ID) == teamID {
				return true, nil
			}
		}
		if !bool(query.Node.Project.Teams.PageInfo.HasNextPage) {
			return false, nil
		}
		cursor := query.Node.Project.Teams.PageInfo.EndCursor
		after = &cursor
	}
}

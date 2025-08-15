package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubOrganizationTeams() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationTeamsRead,

		Schema: map[string]*schema.Schema{
			"root_teams_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"summary_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"results_per_page": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          100,
				ValidateDiagFunc: toDiagFunc(validation.IntBetween(0, 100), "results_per_page"),
			},
			"teams": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"slug": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"privacy": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"members": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"repositories": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"parent": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationTeamsRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name
	rootTeamsOnly := d.Get("root_teams_only").(bool)
	summaryOnly := d.Get("summary_only").(bool)
	resultsPerPage := d.Get("results_per_page").(int)

	var query TeamsQuery

	variables := map[string]any{
		"first":         githubv4.Int(resultsPerPage),
		"login":         githubv4.String(orgName),
		"cursor":        (*githubv4.String)(nil),
		"rootTeamsOnly": githubv4.Boolean(rootTeamsOnly),
		"summaryOnly":   githubv4.Boolean(summaryOnly),
	}

	var teams []any
	for {
		err = client.Query(meta.(*Owner).StopContext, &query, variables)
		if err != nil {
			return err
		}

		additionalTeams := flattenGitHubTeams(query)
		teams = append(teams, additionalTeams...)

		if !query.Organization.Teams.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.NewString(query.Organization.Teams.PageInfo.EndCursor)
	}

	d.SetId(string(query.Organization.ID))
	err = d.Set("teams", teams)
	if err != nil {
		return err
	}

	return nil
}

func flattenGitHubTeams(tq TeamsQuery) []any {
	teams := tq.Organization.Teams.Nodes

	if len(teams) == 0 {
		return make([]any, 0)
	}

	flatTeams := make([]any, len(teams))

	for i, team := range teams {
		t := make(map[string]any)

		t["id"] = team.DatabaseID
		t["node_id"] = team.ID
		t["slug"] = team.Slug
		t["name"] = team.Name
		t["description"] = team.Description
		t["privacy"] = team.Privacy
		members := team.Members.Nodes
		flatMembers := make([]string, len(members))

		for i, member := range members {
			flatMembers[i] = string(member.Login)
		}

		t["members"] = flatMembers

		parentTeam := make(map[string]any)
		parentTeam["id"] = team.Parent.ID
		parentTeam["slug"] = team.Parent.Slug
		parentTeam["name"] = team.Parent.Name
		t["parent"] = parentTeam

		repositories := team.Repositories.Nodes

		flatRepositories := make([]string, len(repositories))

		for i, repository := range repositories {
			flatRepositories[i] = string(repository.Name)
		}

		t["repositories"] = flatRepositories

		flatTeams[i] = t
	}

	return flatTeams
}

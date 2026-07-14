package github

import (
	"context"

	"github.com/google/go-github/v89/github"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubRepositoryTeams() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubRepositoryTeamsRead,

		Description: "Data source to list all teams with access to a repository.",

		Schema: map[string]*schema.Schema{
			"full_name": {
				Description:  "The full name of the repository (e.g. `owner/repo`).",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"full_name", "name"},
				Deprecated:   "The `full_name` attribute is deprecated and will be removed in a future version of the provider. Use `name` instead.",
			},
			"name": {
				Description:  "The name of the repository.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"full_name", "name"},
			},
			"teams": {
				Description: "Teams with access to the repository.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the team.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"node_id": {
							Description: "Node ID of the team.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"slug": {
							Description: "Slug of the team name.",
							Type:        schema.TypeString,
							Required:    true,
						},
						"name": {
							Description: "Name of the team.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "Description of the team.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": {
							Description: "Ownership type of the team; one of `enterprise` or `organization`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"privacy": {
							Description: "Privacy level of the team; one of `secret` or `closed`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"permission": {
							Description: "Legacy default repository permission for the team (typically pull, push, or admin), used when adding a repository without specifying an explicit permission. This does not represent effective access for all repositories or custom repository roles.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"access_source": {
							Description: "Source of the team's access to the repository; one of `direct`, `organization`, or `enterprise`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubRepositoryTeamsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	// TODO: Reinstate this check for v7 when `full_name` is removed from the schema.
	// if ok, diags := checkOrganizationOK(meta); !ok {
	// 	return diags
	// }

	var owner, repoName string
	if v, ok := d.GetOk("full_name"); ok {
		fullName, _ := v.(string)
		o, r, err := splitRepoFullName(fullName)
		if err != nil {
			return diag.FromErr(err)
		}
		owner = o
		repoName = r
	} else {
		owner = meta.name
		repoName, _ = d.Get("name").(string)
	}

	var teams []map[string]any
	for team, err := range meta.v3client.Repositories.ListTeamsIter(ctx, owner, repoName, &github.ListOptions{PerPage: maxPerPage}) {
		if err != nil {
			return diag.FromErr(err)
		}

		t := map[string]any{
			"id":            int(team.GetID()),
			"node_id":       team.GetNodeID(),
			"slug":          team.GetSlug(),
			"name":          team.GetName(),
			"description":   team.GetDescription(),
			"type":          team.GetType(),
			"privacy":       team.GetPrivacy(),
			"permission":    team.GetPermission(),
			"access_source": team.GetAccessSource(),
		}

		teams = append(teams, t)
	}

	d.SetId(repoName)

	if err := d.Set("teams", teams); err != nil {
		return diag.Errorf("error setting teams: %v", err)
	}

	return nil
}

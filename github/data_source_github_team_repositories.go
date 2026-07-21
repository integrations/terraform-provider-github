package github

import (
	"context"
	"iter"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubTeamRepositories() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubTeamRepositoriesRead,

		Description: "Data source to list all team repositories.",

		Schema: map[string]*schema.Schema{
			"team_id": {
				Description:      "ID of the team. One of `team_id` or `slug` must be specified.",
				Type:             schema.TypeInt,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"team_id", "slug"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
			},
			"slug": {
				Description:      "Slug of the team name. One of `team_id` or `slug` must be specified.",
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ExactlyOneOf:     []string{"team_id", "slug"},
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
			"repositories": {
				Description: "Team repositories.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the repository.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"node_id": {
							Description: "Node ID of the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "Name of the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"visibility": {
							Description: "Visibility of the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"archived": {
							Description: "Whether the repository is archived.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
						"role_name": {
							Description: "Role the team has for the repository.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubTeamRepositoriesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	opts := &github.ListOptions{
		PerPage: meta.maxPerPage,
	}

	var iter iter.Seq2[*github.Repository, error]
	if v, ok := d.GetOk("team_id"); ok {
		teamIDInt, _ := v.(int)
		teamID := int64(teamIDInt)
		iter = meta.v3client.Teams.ListTeamReposByIDIter(ctx, meta.id, teamID, opts)
	} else {
		slug, _ := d.Get("slug").(string)
		iter = meta.v3client.Teams.ListTeamReposBySlugIter(ctx, meta.name, slug, opts)
	}

	repos := make([]map[string]any, 0)
	for repo, err := range iter {
		if err != nil {
			return diag.FromErr(err)
		}

		r := map[string]any{
			"id":         repo.GetID(),
			"node_id":    repo.GetNodeID(),
			"name":       repo.GetName(),
			"visibility": repo.GetVisibility(),
			"archived":   repo.GetArchived(),
			"role_name":  repo.GetRoleName(),
		}
		repos = append(repos, r)
	}

	d.SetId(meta.name)

	if err := d.Set("repositories", repos); err != nil {
		return diag.Errorf("error setting repositories: %v", err)
	}

	return nil
}

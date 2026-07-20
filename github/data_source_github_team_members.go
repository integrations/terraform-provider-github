package github

import (
	"context"
	"iter"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubTeamMembers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubTeamMembersRead,

		Description: "Data source to list all team members.",

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
			"members": {
				Description: "Team members.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the member.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"node_id": {
							Description: "Node ID of the member.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"login": {
							Description: "Login of the member.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"email": {
							Description: "Email of the member.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"role": {
							Description: "Role of the member in the team; can be one of `member` or `maintainer`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"inherited": {
							Description: "Whether the member is inherited from a parent team.",
							Type:        schema.TypeBool,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubTeamMembersRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	opts := &github.TeamListTeamMembersOptions{
		ListOptions: github.ListOptions{
			PerPage: maxPerPage,
		},
	}

	var iter iter.Seq2[*github.User, error]
	if v, ok := d.GetOk("team_id"); ok {
		teamIDInt, _ := v.(int)
		teamID := int64(teamIDInt)
		iter = meta.v3client.Teams.ListTeamMembersByIDIter(ctx, meta.id, teamID, opts)
	} else {
		slug, _ := d.Get("slug").(string)
		iter = meta.v3client.Teams.ListTeamMembersBySlugIter(ctx, meta.name, slug, opts)
	}

	members := make([]map[string]any, 0)
	for user, err := range iter {
		if err != nil {
			return diag.FromErr(err)
		}

		u := map[string]any{
			"id":        user.GetID(),
			"node_id":   user.GetNodeID(),
			"login":     user.GetLogin(),
			"email":     user.GetEmail(),
			"role":      user.GetRole(),
			"inherited": user.GetInherited(),
		}
		members = append(members, u)
	}

	d.SetId(meta.name)

	if err := d.Set("members", members); err != nil {
		return diag.Errorf("error setting members: %v", err)
	}

	return nil
}

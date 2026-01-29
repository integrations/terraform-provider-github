package github

import (
	"context"
	"strings"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	teamIDKey                    = "team_id"
	teamSlugKey                  = "slug"
	teamNameKey                  = "name"
	teamDescriptionKey           = "description"
	teamOrganizationSelectionKey = "organization_selection_type"
	teamGroupIDKey               = "group_id"
)

func dataSourceGithubEnterpriseTeams() *schema.Resource {
	return &schema.Resource{
		Description: "Lists all GitHub enterprise teams in an enterprise.",
		ReadContext: dataSourceGithubEnterpriseTeamsRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug of the enterprise.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.All(validation.StringIsNotWhiteSpace, validation.StringIsNotEmpty)),
			},
			"teams": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "All teams in the enterprise.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						teamIDKey: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The numeric ID of the enterprise team.",
						},
						teamSlugKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The slug of the enterprise team.",
						},
						teamNameKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the enterprise team.",
						},
						teamDescriptionKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A description of the enterprise team.",
						},
						teamOrganizationSelectionKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies which organizations in the enterprise should have access to this team.",
						},
						teamGroupIDKey: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the IdP group to assign team membership with.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseTeamsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := strings.TrimSpace(d.Get("enterprise_slug").(string))
	teams, err := listAllEnterpriseTeams(ctx, client, enterpriseSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	flat := make([]any, 0, len(teams))
	for _, t := range teams {
		m := map[string]any{
			teamIDKey:   int(t.ID),
			teamSlugKey: t.Slug,
			teamNameKey: t.Name,
		}
		if t.Description != nil {
			m[teamDescriptionKey] = *t.Description
		} else {
			m[teamDescriptionKey] = ""
		}
		orgSel := ""
		if t.OrganizationSelectionType != nil {
			orgSel = *t.OrganizationSelectionType
		}
		if orgSel == "" {
			orgSel = "disabled"
		}
		m[teamOrganizationSelectionKey] = orgSel
		if t.GroupID != "" {
			m[teamGroupIDKey] = t.GroupID
		} else {
			m[teamGroupIDKey] = ""
		}
		flat = append(flat, m)
	}

	d.SetId(enterpriseSlug)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("teams", flat); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// listAllEnterpriseTeams returns all enterprise teams with pagination handled.
func listAllEnterpriseTeams(ctx context.Context, client *github.Client, enterpriseSlug string) ([]*github.EnterpriseTeam, error) {
	var all []*github.EnterpriseTeam
	opt := &github.ListOptions{PerPage: maxPerPage}

	for {
		teams, resp, err := client.Enterprise.ListTeams(ctx, enterpriseSlug, opt)
		if err != nil {
			return nil, err
		}
		all = append(all, teams...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return all, nil
}

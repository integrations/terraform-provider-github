package github

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
						"team_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The numeric ID of the enterprise team.",
						},
						"slug": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The slug of the enterprise team.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the enterprise team.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A description of the enterprise team.",
						},
						"organization_selection_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies which organizations in the enterprise should have access to this team.",
						},
						"group_id": {
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
	teams, err := listEnterpriseTeams(ctx, client, enterpriseSlug)
	if err != nil {
		return diag.FromErr(err)
	}

	flat := make([]any, 0, len(teams))
	for _, t := range teams {
		m := map[string]any{
			"team_id": int(t.ID),
			"slug":    t.Slug,
			"name":    t.Name,
		}
		if t.Description != nil {
			m["description"] = *t.Description
		} else {
			m["description"] = ""
		}
		orgSel := t.OrganizationSelectionType
		if orgSel == "" {
			orgSel = "disabled"
		}
		m["organization_selection_type"] = orgSel
		if t.GroupID != nil {
			m["group_id"] = *t.GroupID
		} else {
			m["group_id"] = ""
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

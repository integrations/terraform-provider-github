package github

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubTeamExternalGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Data source to retrieve external groups for a specific GitHub team.",
		ReadContext: dataSourceGithubTeamExternalGroupsRead,
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the GitHub team.",
			},
			"external_groups": {
				Description: "List of external groups connected to the team.",
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Description: "ID of the external group.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"group_name": {
							Description: "Name of the external group.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"updated_at": {
							Description: "Timestamp of the last update to the external group.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubTeamExternalGroupsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	orgName := meta.name
	slug, ok := d.Get("slug").(string)
	if !ok {
		return diag.Errorf("expected type of %s to be string", d.Get("slug"))
	}

	externalGroups, _, err := meta.v3client.Teams.ListExternalGroupsForTeamBySlug(ctx, orgName, slug)
	if err != nil {
		return diag.FromErr(err)
	}

	groups := make([]map[string]any, 0)
	for _, group := range externalGroups.Groups {
		g := map[string]any{
			"group_id":   group.GetGroupID(),
			"group_name": group.GetGroupName(),
			"updated_at": group.GetUpdatedAt().Format(time.RFC3339),
		}

		groups = append(groups, g)
	}

	id, err := buildID(orgName, slug)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("external_groups", groups); err != nil {
		return diag.Errorf("error setting external_groups: %v", err)
	}

	return nil
}

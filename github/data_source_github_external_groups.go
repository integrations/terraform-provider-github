package github

import (
	"context"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubExternalGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubExternalGroupsRead,
		Description: "Data source to list all external groups in an organization.",
		Schema: map[string]*schema.Schema{
			"display_name_filter": {
				Description: "Filter external groups by display name.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"external_groups": {
				Description: "List of external groups in the organization.",
				Type:        schema.TypeList,
				Computed:    true,
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

func dataSourceGithubExternalGroupsRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if err := checkOrganization(meta); err != nil {
		return diag.FromErr(err)
	}

	var filter *string
	if v, ok := d.GetOk("display_name_filter"); ok {
		s, _ := v.(string)
		filter = &s
	}

	groups := make([]map[string]any, 0)
	for group, err := range meta.v3client.Teams.ListExternalGroupsIter(ctx, meta.name, &github.ListExternalGroupsOptions{DisplayName: filter, ListOptions: github.ListOptions{PerPage: meta.maxPerPage}}) {
		if err != nil {
			return diag.FromErr(err)
		}

		g := map[string]any{
			"group_id":   group.GetGroupID(),
			"group_name": group.GetGroupName(),
			"updated_at": group.GetUpdatedAt().Format(time.RFC3339),
		}

		groups = append(groups, g)
	}

	var id string
	if filter != nil {
		s, err := buildID(meta.name, *filter)
		if err != nil {
			return diag.FromErr(err)
		}
		id = s
	} else {
		id = meta.name
	}

	d.SetId(id)

	if err := d.Set("external_groups", groups); err != nil {
		return diag.Errorf("error setting external_groups: %v", err)
	}

	return nil
}

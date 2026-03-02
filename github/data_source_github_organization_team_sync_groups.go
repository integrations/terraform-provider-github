package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationTeamSyncGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationTeamSyncGroupsRead,

		Schema: map[string]*schema.Schema{
			"prefix_filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Filters the results to return only those that begin with the specified value.",
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"group_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationTeamSyncGroupsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	orgName := meta.(*Owner).name
	query := d.Get("q").(string)
	options := &github.ListIDPGroupsOptions{
		Query: query,
		ListCursorOptions: github.ListCursorOptions{
			PerPage: maxPerPage,
		},
	}

	groups := make([]any, 0)
	for {
		idpGroupList, resp, err := client.Teams.ListIDPGroupsInOrganization(ctx, orgName, options)
		if err != nil {
			return diag.FromErr(err)
		}

		result := flattenGithubIDPGroupList(idpGroupList)

		groups = append(groups, result...)

		if resp.NextPageToken == "" {
			break
		}
		options.Page = resp.NextPageToken
	}

	d.SetId(fmt.Sprintf("%s/github-org-team-sync-groups/%s", orgName, query))
	if err := d.Set("groups", groups); err != nil {
		return diag.Errorf("error setting groups: %v", err)
	}

	return nil
}

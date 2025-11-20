package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/shurcooL/githubv4"
)

func dataSourceGithubOrganizationIpAllowList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubOrganizationIpAllowListRead,

		Schema: map[string]*schema.Schema{
			"ip_allow_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"allow_list_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"created_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_at": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationIpAllowListRead(d *schema.ResourceData, meta any) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	ctx := context.Background()
	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	type PageInfo struct {
		StartCursor     githubv4.String
		EndCursor       githubv4.String
		HasNextPage     githubv4.Boolean
		HasPreviousPage githubv4.Boolean
	}

	type IpAllowListEntry struct {
		ID             githubv4.String
		Name           githubv4.String
		AllowListValue githubv4.String
		IsActive       githubv4.Boolean
		CreatedAt      githubv4.String
		UpdatedAt      githubv4.String
	}

	type IpAllowListEntries struct {
		Nodes      []IpAllowListEntry
		PageInfo   PageInfo
		TotalCount githubv4.Int
	}

	var query struct {
		Organization struct {
			ID                 githubv4.String
			IpAllowListEntries IpAllowListEntries `graphql:"ipAllowListEntries(first: 100, after: $entriesCursor)"`
		} `graphql:"organization(login: $login)"`
	}

	variables := map[string]any{
		"login":         githubv4.String(orgName),
		"entriesCursor": (*githubv4.String)(nil),
	}

	var ipAllowList []any
	var ipAllowListEntries []IpAllowListEntry

	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return err
		}

		ipAllowListEntries = append(ipAllowListEntries, query.Organization.IpAllowListEntries.Nodes...)
		if !query.Organization.IpAllowListEntries.PageInfo.HasNextPage {
			break
		}
		variables["entriesCursor"] = githubv4.NewString(query.Organization.IpAllowListEntries.PageInfo.EndCursor)
	}
	for index := range ipAllowListEntries {
		ipAllowList = append(ipAllowList, map[string]any{
			"id":               ipAllowListEntries[index].ID,
			"name":             ipAllowListEntries[index].Name,
			"allow_list_value": ipAllowListEntries[index].AllowListValue,
			"is_active":        ipAllowListEntries[index].IsActive,
			"created_at":       ipAllowListEntries[index].CreatedAt,
			"updated_at":       ipAllowListEntries[index].UpdatedAt,
		})
	}

	d.SetId(string(query.Organization.ID))
	err = d.Set("ip_allow_list", ipAllowList)
	if err != nil {
		return err
	}

	return nil
}

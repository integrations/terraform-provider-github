package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

type IpAllowListEntry struct {
	id               githubv4.String
	name             githubv4.String
	allow_list_value githubv4.String
	is_active        githubv4.Boolean
	created_at       githubv4.String
	updated_at       githubv4.String
}

/**
 * Returns all IP allow list entries for an organization.
 * This util function is used by both data_source and resource elements.
 */
func getOrganizationIpAllowListEntries(meta interface{}) ([]IpAllowListEntry, error) {
	err := checkOrganization(meta)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	client := meta.(*Owner).v4client
	orgName := meta.(*Owner).name

	type IpAllowListEntryGql struct {
		ID             githubv4.String
		Name           githubv4.String
		AllowListValue githubv4.String
		IsActive       githubv4.Boolean
		CreatedAt      githubv4.String
		UpdatedAt      githubv4.String
	}

	type IpAllowListEntriesGql struct {
		Nodes      []IpAllowListEntryGql
		PageInfo   PageInfo
		TotalCount githubv4.Int
	}

	var query struct {
		Organization struct {
			ID                 githubv4.String
			IpAllowListEntries IpAllowListEntriesGql `graphql:"ipAllowListEntries(first: 100, after: $entriesCursor)"`
		} `graphql:"organization(login: $login)"`
	}

	variables := map[string]interface{}{
		"login":         githubv4.String(orgName),
		"entriesCursor": (*githubv4.String)(nil),
	}

	var ipAllowListEntriesGql []IpAllowListEntryGql

	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}

		ipAllowListEntriesGql = append(ipAllowListEntriesGql, query.Organization.IpAllowListEntries.Nodes...)
		if !query.Organization.IpAllowListEntries.PageInfo.HasNextPage {
			break
		}
		variables["entriesCursor"] = githubv4.NewString(query.Organization.IpAllowListEntries.PageInfo.EndCursor)
	}

	// Translate the graphql response to terraform state.
	var ipAllowListEntries []IpAllowListEntry
	for index := range ipAllowListEntriesGql {
		ipAllowListEntries = append(ipAllowListEntries, IpAllowListEntry{
			id:               ipAllowListEntriesGql[index].ID,
			name:             ipAllowListEntriesGql[index].Name,
			allow_list_value: ipAllowListEntriesGql[index].AllowListValue,
			is_active:        ipAllowListEntriesGql[index].IsActive,
			created_at:       ipAllowListEntriesGql[index].CreatedAt,
			updated_at:       ipAllowListEntriesGql[index].UpdatedAt,
		})
	}

	return ipAllowListEntries, nil
}

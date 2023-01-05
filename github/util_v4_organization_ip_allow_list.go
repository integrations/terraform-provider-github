package github

import (
	"context"

	"github.com/shurcooL/githubv4"
)

type IpAllowListEntry struct {
	ID             githubv4.String  `mapstructure:"id"`
	Name           githubv4.String  `mapstructure:"name"`
	AllowListValue githubv4.String  `mapstructure:"allow_list_value"`
	IsActive       githubv4.Boolean `mapstructure:"is_active"`
	CreatedAt      githubv4.String  `mapstructure:"created_at"`
	UpdatedAt      githubv4.String  `mapstructure:"updated_at"`
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

	variables := map[string]interface{}{
		"login":         githubv4.String(orgName),
		"entriesCursor": (*githubv4.String)(nil),
	}

	var ipAllowListEntries []IpAllowListEntry

	for {
		err := client.Query(ctx, &query, variables)
		if err != nil {
			return nil, err
		}

		ipAllowListEntries = append(ipAllowListEntries, query.Organization.IpAllowListEntries.Nodes...)
		if !query.Organization.IpAllowListEntries.PageInfo.HasNextPage {
			break
		}
		variables["entriesCursor"] = githubv4.NewString(query.Organization.IpAllowListEntries.PageInfo.EndCursor)
	}

	return ipAllowListEntries, nil
}

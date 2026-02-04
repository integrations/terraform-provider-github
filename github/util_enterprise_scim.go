package github

import (
	"context"

	gh "github.com/google/go-github/v81/github"
)

// enterpriseSCIMListAllGroups fetches all SCIM groups for an enterprise with automatic pagination.
func enterpriseSCIMListAllGroups(ctx context.Context, client *gh.Client, enterprise, filter string, count int) ([]*gh.SCIMEnterpriseGroupAttributes, *gh.SCIMEnterpriseGroups, error) {
	startIndex := 1
	all := make([]*gh.SCIMEnterpriseGroupAttributes, 0)
	var firstResp *gh.SCIMEnterpriseGroups

	for {
		opts := &gh.ListProvisionedSCIMGroupsEnterpriseOptions{
			StartIndex: gh.Ptr(startIndex),
			Count:      gh.Ptr(count),
		}
		if filter != "" {
			opts.Filter = gh.Ptr(filter)
		}

		page, _, err := client.Enterprise.ListProvisionedSCIMGroups(ctx, enterprise, opts)
		if err != nil {
			return nil, nil, err
		}

		if firstResp == nil {
			firstResp = page
		}

		all = append(all, page.Resources...)

		if len(page.Resources) == 0 {
			break
		}
		if page.TotalResults != nil && len(all) >= *page.TotalResults {
			break
		}

		startIndex += len(page.Resources)
	}

	if firstResp == nil {
		firstResp = &gh.SCIMEnterpriseGroups{
			Schemas:      []string{gh.SCIMSchemasURINamespacesListResponse},
			TotalResults: gh.Ptr(len(all)),
			StartIndex:   gh.Ptr(1),
			ItemsPerPage: gh.Ptr(count),
		}
	}

	return all, firstResp, nil
}

// enterpriseSCIMListAllUsers fetches all SCIM users for an enterprise with automatic pagination.
func enterpriseSCIMListAllUsers(ctx context.Context, client *gh.Client, enterprise, filter string, count int) ([]*gh.SCIMEnterpriseUserAttributes, *gh.SCIMEnterpriseUsers, error) {
	startIndex := 1
	all := make([]*gh.SCIMEnterpriseUserAttributes, 0)
	var firstResp *gh.SCIMEnterpriseUsers

	for {
		opts := &gh.ListProvisionedSCIMUsersEnterpriseOptions{
			StartIndex: gh.Ptr(startIndex),
			Count:      gh.Ptr(count),
		}
		if filter != "" {
			opts.Filter = gh.Ptr(filter)
		}

		page, _, err := client.Enterprise.ListProvisionedSCIMUsers(ctx, enterprise, opts)
		if err != nil {
			return nil, nil, err
		}

		if firstResp == nil {
			firstResp = page
		}

		all = append(all, page.Resources...)

		if len(page.Resources) == 0 {
			break
		}
		if page.TotalResults != nil && len(all) >= *page.TotalResults {
			break
		}

		startIndex += len(page.Resources)
	}

	if firstResp == nil {
		firstResp = &gh.SCIMEnterpriseUsers{
			Schemas:      []string{gh.SCIMSchemasURINamespacesListResponse},
			TotalResults: gh.Ptr(len(all)),
			StartIndex:   gh.Ptr(1),
			ItemsPerPage: gh.Ptr(count),
		}
	}

	return all, firstResp, nil
}

func flattenEnterpriseSCIMMeta(meta *gh.SCIMEnterpriseMeta) []any {
	if meta == nil {
		return nil
	}
	m := map[string]any{
		"resource_type": meta.ResourceType,
	}
	if meta.Created != nil {
		m["created"] = meta.Created.String()
	}
	if meta.LastModified != nil {
		m["last_modified"] = meta.LastModified.String()
	}
	if meta.Location != nil {
		m["location"] = *meta.Location
	}
	return []any{m}
}

func flattenEnterpriseSCIMGroupMembers(members []*gh.SCIMEnterpriseDisplayReference) []any {
	out := make([]any, 0, len(members))
	for _, m := range members {
		item := map[string]any{
			"value": m.Value,
		}
		if m.Ref != nil {
			item["ref"] = *m.Ref
		}
		if m.Display != nil {
			item["display_name"] = *m.Display
		}
		out = append(out, item)
	}
	return out
}

func flattenEnterpriseSCIMGroup(group *gh.SCIMEnterpriseGroupAttributes) map[string]any {
	m := map[string]any{
		"schemas": group.Schemas,
		"members": flattenEnterpriseSCIMGroupMembers(group.Members),
		"meta":    flattenEnterpriseSCIMMeta(group.Meta),
	}
	if group.ID != nil {
		m["id"] = *group.ID
	}
	if group.ExternalID != nil {
		m["external_id"] = *group.ExternalID
	}
	if group.DisplayName != nil {
		m["display_name"] = *group.DisplayName
	}
	return m
}

func flattenEnterpriseSCIMUserName(name *gh.SCIMEnterpriseUserName) []any {
	if name == nil {
		return nil
	}
	m := map[string]any{
		"family_name": name.FamilyName,
		"given_name":  name.GivenName,
	}
	if name.Formatted != nil {
		m["formatted"] = *name.Formatted
	}
	if name.MiddleName != nil {
		m["middle_name"] = *name.MiddleName
	}
	return []any{m}
}

func flattenEnterpriseSCIMUserEmails(emails []*gh.SCIMEnterpriseUserEmail) []any {
	out := make([]any, 0, len(emails))
	for _, e := range emails {
		out = append(out, map[string]any{
			"value":   e.Value,
			"type":    e.Type,
			"primary": e.Primary,
		})
	}
	return out
}

func flattenEnterpriseSCIMUserRoles(roles []*gh.SCIMEnterpriseUserRole) []any {
	out := make([]any, 0, len(roles))
	for _, r := range roles {
		item := map[string]any{
			"value": r.Value,
		}
		if r.Display != nil {
			item["display"] = *r.Display
		}
		if r.Type != nil {
			item["type"] = *r.Type
		}
		if r.Primary != nil {
			item["primary"] = *r.Primary
		}
		out = append(out, item)
	}
	return out
}

func flattenEnterpriseSCIMUser(user *gh.SCIMEnterpriseUserAttributes) map[string]any {
	m := map[string]any{
		"schemas":      user.Schemas,
		"user_name":    user.UserName,
		"display_name": user.DisplayName,
		"external_id":  user.ExternalID,
		"active":       user.Active,
		"name":         flattenEnterpriseSCIMUserName(user.Name),
		"emails":       flattenEnterpriseSCIMUserEmails(user.Emails),
		"roles":        flattenEnterpriseSCIMUserRoles(user.Roles),
		"meta":         flattenEnterpriseSCIMMeta(user.Meta),
	}
	if user.ID != nil {
		m["id"] = *user.ID
	}
	return m
}

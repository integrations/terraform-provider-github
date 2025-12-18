package github

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	gh "github.com/google/go-github/v67/github"
)

const enterpriseSCIMAcceptHeader = "application/scim+json"

type enterpriseSCIMListOptions struct {
	Filter             string
	ExcludedAttributes string
	StartIndex         int
	Count              int
}

type enterpriseSCIMListResponse[T any] struct {
	Schemas      []string `json:"schemas,omitempty"`
	TotalResults int      `json:"totalResults,omitempty"`
	StartIndex   int      `json:"startIndex,omitempty"`
	ItemsPerPage int      `json:"itemsPerPage,omitempty"`
	Resources    []T      `json:"Resources,omitempty"`
}

type enterpriseSCIMMeta struct {
	ResourceType  string `json:"resourceType,omitempty"`
	Created       string `json:"created,omitempty"`
	LastModified  string `json:"lastModified,omitempty"`
	Location      string `json:"location,omitempty"`
	Version       string `json:"version,omitempty"`
	ETag          string `json:"eTag,omitempty"`
	PasswordChgAt string `json:"passwordChangedAt,omitempty"`
}

type enterpriseSCIMGroupMember struct {
	Value   string `json:"value,omitempty"`
	Ref     string `json:"$ref,omitempty"`
	Display string `json:"display,omitempty"`
}

type enterpriseSCIMGroup struct {
	Schemas     []string                    `json:"schemas,omitempty"`
	ID          string                      `json:"id,omitempty"`
	ExternalID  string                      `json:"externalId,omitempty"`
	DisplayName string                      `json:"displayName,omitempty"`
	Members     []enterpriseSCIMGroupMember `json:"members,omitempty"`
	Meta        *enterpriseSCIMMeta         `json:"meta,omitempty"`
}

type enterpriseSCIMUserName struct {
	Formatted  string `json:"formatted,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
	GivenName  string `json:"givenName,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
}

type enterpriseSCIMUserEmail struct {
	Value   string `json:"value,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type enterpriseSCIMUserRole struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Type    string `json:"type,omitempty"`
	Primary bool   `json:"primary,omitempty"`
}

type enterpriseSCIMUser struct {
	Schemas     []string                  `json:"schemas,omitempty"`
	ID          string                    `json:"id,omitempty"`
	ExternalID  string                    `json:"externalId,omitempty"`
	UserName    string                    `json:"userName,omitempty"`
	DisplayName string                    `json:"displayName,omitempty"`
	Active      bool                      `json:"active,omitempty"`
	Name        *enterpriseSCIMUserName   `json:"name,omitempty"`
	Emails      []enterpriseSCIMUserEmail `json:"emails,omitempty"`
	Roles       []enterpriseSCIMUserRole  `json:"roles,omitempty"`
	Meta        *enterpriseSCIMMeta       `json:"meta,omitempty"`
}

func enterpriseSCIMListURL(path string, opts enterpriseSCIMListOptions) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	q := u.Query()
	if opts.Filter != "" {
		q.Set("filter", opts.Filter)
	}
	if opts.ExcludedAttributes != "" {
		q.Set("excludedAttributes", opts.ExcludedAttributes)
	}
	if opts.StartIndex > 0 {
		q.Set("startIndex", strconv.Itoa(opts.StartIndex))
	}
	if opts.Count > 0 {
		q.Set("count", strconv.Itoa(opts.Count))
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func enterpriseSCIMGet[T any](ctx context.Context, client *gh.Client, urlStr string, out *T) (*gh.Response, error) {
	req, err := client.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", enterpriseSCIMAcceptHeader)
	return client.Do(ctx, req, out)
}

func enterpriseSCIMListAllGroups(ctx context.Context, client *gh.Client, enterprise string, filter string, excludedAttributes string, count int) ([]enterpriseSCIMGroup, *enterpriseSCIMListResponse[enterpriseSCIMGroup], error) {
	startIndex := 1
	all := make([]enterpriseSCIMGroup, 0)
	var firstResp *enterpriseSCIMListResponse[enterpriseSCIMGroup]

	for {
		path := fmt.Sprintf("scim/v2/enterprises/%s/Groups", enterprise)
		urlStr, err := enterpriseSCIMListURL(path, enterpriseSCIMListOptions{
			Filter:             filter,
			ExcludedAttributes: excludedAttributes,
			StartIndex:         startIndex,
			Count:              count,
		})
		if err != nil {
			return nil, nil, err
		}

		page := enterpriseSCIMListResponse[enterpriseSCIMGroup]{}
		_, err = enterpriseSCIMGet(ctx, client, urlStr, &page)
		if err != nil {
			return nil, nil, err
		}

		if firstResp == nil {
			snap := page
			firstResp = &snap
		}

		all = append(all, page.Resources...)

		if len(page.Resources) == 0 {
			break
		}
		if page.TotalResults > 0 && len(all) >= page.TotalResults {
			break
		}

		startIndex += len(page.Resources)
	}

	if firstResp == nil {
		firstResp = &enterpriseSCIMListResponse[enterpriseSCIMGroup]{
			Schemas:      []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"},
			TotalResults: len(all),
			StartIndex:   1,
			ItemsPerPage: count,
			Resources:    nil,
		}
	}

	return all, firstResp, nil
}

func enterpriseSCIMListAllUsers(ctx context.Context, client *gh.Client, enterprise string, filter string, excludedAttributes string, count int) ([]enterpriseSCIMUser, *enterpriseSCIMListResponse[enterpriseSCIMUser], error) {
	startIndex := 1
	all := make([]enterpriseSCIMUser, 0)
	var firstResp *enterpriseSCIMListResponse[enterpriseSCIMUser]

	for {
		path := fmt.Sprintf("scim/v2/enterprises/%s/Users", enterprise)
		urlStr, err := enterpriseSCIMListURL(path, enterpriseSCIMListOptions{
			Filter:             filter,
			ExcludedAttributes: excludedAttributes,
			StartIndex:         startIndex,
			Count:              count,
		})
		if err != nil {
			return nil, nil, err
		}

		page := enterpriseSCIMListResponse[enterpriseSCIMUser]{}
		_, err = enterpriseSCIMGet(ctx, client, urlStr, &page)
		if err != nil {
			return nil, nil, err
		}

		if firstResp == nil {
			snap := page
			firstResp = &snap
		}

		all = append(all, page.Resources...)

		if len(page.Resources) == 0 {
			break
		}
		if page.TotalResults > 0 && len(all) >= page.TotalResults {
			break
		}

		startIndex += len(page.Resources)
	}

	if firstResp == nil {
		firstResp = &enterpriseSCIMListResponse[enterpriseSCIMUser]{
			Schemas:      []string{"urn:ietf:params:scim:api:messages:2.0:ListResponse"},
			TotalResults: len(all),
			StartIndex:   1,
			ItemsPerPage: count,
			Resources:    nil,
		}
	}

	return all, firstResp, nil
}

func enterpriseSCIMGetGroup(ctx context.Context, client *gh.Client, enterprise, scimGroupID string) (*enterpriseSCIMGroup, error) {
	path := fmt.Sprintf("scim/v2/enterprises/%s/Groups/%s", enterprise, scimGroupID)
	group := enterpriseSCIMGroup{}
	_, err := enterpriseSCIMGet(ctx, client, path, &group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func enterpriseSCIMGetUser(ctx context.Context, client *gh.Client, enterprise, scimUserID string) (*enterpriseSCIMUser, error) {
	path := fmt.Sprintf("scim/v2/enterprises/%s/Users/%s", enterprise, scimUserID)
	user := enterpriseSCIMUser{}
	_, err := enterpriseSCIMGet(ctx, client, path, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func flattenEnterpriseSCIMMeta(meta *enterpriseSCIMMeta) []any {
	if meta == nil {
		return nil
	}
	return []any{map[string]any{
		"resource_type":       meta.ResourceType,
		"created":             meta.Created,
		"last_modified":       meta.LastModified,
		"location":            meta.Location,
		"version":             meta.Version,
		"etag":                meta.ETag,
		"password_changed_at": meta.PasswordChgAt,
	}}
}

func flattenEnterpriseSCIMGroupMembers(members []enterpriseSCIMGroupMember) []any {
	out := make([]any, 0, len(members))
	for _, m := range members {
		out = append(out, map[string]any{
			"value":        m.Value,
			"ref":          m.Ref,
			"display_name": m.Display,
		})
	}
	return out
}

func flattenEnterpriseSCIMGroup(group enterpriseSCIMGroup) map[string]any {
	return map[string]any{
		"schemas":      group.Schemas,
		"id":           group.ID,
		"external_id":  group.ExternalID,
		"display_name": group.DisplayName,
		"members":      flattenEnterpriseSCIMGroupMembers(group.Members),
		"meta":         flattenEnterpriseSCIMMeta(group.Meta),
	}
}

func flattenEnterpriseSCIMUserName(name *enterpriseSCIMUserName) []any {
	if name == nil {
		return nil
	}
	return []any{map[string]any{
		"formatted":   name.Formatted,
		"family_name": name.FamilyName,
		"given_name":  name.GivenName,
		"middle_name": name.MiddleName,
	}}
}

func flattenEnterpriseSCIMUserEmails(emails []enterpriseSCIMUserEmail) []any {
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

func flattenEnterpriseSCIMUserRoles(roles []enterpriseSCIMUserRole) []any {
	out := make([]any, 0, len(roles))
	for _, r := range roles {
		out = append(out, map[string]any{
			"value":   r.Value,
			"display": r.Display,
			"type":    r.Type,
			"primary": r.Primary,
		})
	}
	return out
}

func flattenEnterpriseSCIMUser(user enterpriseSCIMUser) map[string]any {
	return map[string]any{
		"schemas":      user.Schemas,
		"id":           user.ID,
		"external_id":  user.ExternalID,
		"user_name":    user.UserName,
		"display_name": user.DisplayName,
		"active":       user.Active,
		"name":         flattenEnterpriseSCIMUserName(user.Name),
		"emails":       flattenEnterpriseSCIMUserEmails(user.Emails),
		"roles":        flattenEnterpriseSCIMUserRoles(user.Roles),
		"meta":         flattenEnterpriseSCIMMeta(user.Meta),
	}
}

package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	githubv3 "github.com/google/go-github/v67/github"
)

const enterpriseTeamsAPIVersion = "2022-11-28"

func parseSlashTwoPartID(id, left, right string) (string, string, error) {
	parts := strings.SplitN(id, "/", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("unexpected ID format (%q); expected %s/%s", id, left, right)
	}
	return parts[0], parts[1], nil
}

func buildSlashTwoPartID(a, b string) string {
	return fmt.Sprintf("%s/%s", a, b)
}

func parseSlashThreePartID(id, left, center, right string) (string, string, string, error) {
	parts := strings.SplitN(id, "/", 3)
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("unexpected ID format (%q); expected %s/%s/%s", id, left, center, right)
	}
	return parts[0], parts[1], parts[2], nil
}

func buildSlashThreePartID(a, b, c string) string {
	return fmt.Sprintf("%s/%s/%s", a, b, c)
}

func enterpriseTeamsAddListOptions(u string, opt *githubv3.ListOptions) string {
	if opt == nil {
		return u
	}
	vals := url.Values{}
	if opt.Page != 0 {
		vals.Set("page", strconv.Itoa(opt.Page))
	}
	if opt.PerPage != 0 {
		vals.Set("per_page", strconv.Itoa(opt.PerPage))
	}
	enc := vals.Encode()
	if enc == "" {
		return u
	}
	if strings.Contains(u, "?") {
		return u + "&" + enc
	}
	return u + "?" + enc
}

func enterpriseTeamsNewRequest(client *githubv3.Client, method, urlStr string, body any) (*http.Request, error) {
	req, err := client.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	// These endpoints are versioned and currently in public preview.
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", enterpriseTeamsAPIVersion)
	return req, nil
}

type enterpriseTeam struct {
	ID                        int64   `json:"id"`
	Name                      string  `json:"name"`
	Description               *string `json:"description"`
	Slug                      string  `json:"slug"`
	GroupID                   *string `json:"group_id"`
	OrganizationSelectionType string  `json:"organization_selection_type"`
}

type enterpriseTeamCreateRequest struct {
	Name                      string  `json:"name"`
	Description               *string `json:"description,omitempty"`
	OrganizationSelectionType *string `json:"organization_selection_type,omitempty"`
	GroupID                   *string `json:"group_id,omitempty"`
}

type enterpriseTeamUpdateRequest struct {
	Name                      *string `json:"name,omitempty"`
	Description               *string `json:"description,omitempty"`
	OrganizationSelectionType *string `json:"organization_selection_type,omitempty"`
	GroupID                   *string `json:"group_id,omitempty"`
}

func parseEnterpriseTeam(raw json.RawMessage) (*enterpriseTeam, error) {
	// The API docs are inconsistent about whether this returns an object or an
	// array with one element, so we try both.
	var t enterpriseTeam
	if err := json.Unmarshal(raw, &t); err == nil {
		if t.ID != 0 || t.Slug != "" || t.Name != "" {
			return &t, nil
		}
	}

	var ts []enterpriseTeam
	if err := json.Unmarshal(raw, &ts); err == nil {
		if len(ts) > 0 {
			return &ts[0], nil
		}
	}

	return nil, fmt.Errorf("unexpected enterprise team response")
}

func listEnterpriseTeams(ctx context.Context, client *githubv3.Client, enterpriseSlug string) ([]enterpriseTeam, error) {
	all := []enterpriseTeam{}
	opt := &githubv3.ListOptions{PerPage: maxPerPage}

	for {
		u := enterpriseTeamsAddListOptions(fmt.Sprintf("enterprises/%s/teams", enterpriseSlug), opt)
		req, err := enterpriseTeamsNewRequest(client, "GET", u, nil)
		if err != nil {
			return nil, err
		}

		var pageTeams []enterpriseTeam
		resp, err := client.Do(ctx, req, &pageTeams)
		if err != nil {
			return nil, err
		}
		all = append(all, pageTeams...)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return all, nil
}

func getEnterpriseTeamBySlug(ctx context.Context, client *githubv3.Client, enterpriseSlug, teamSlug string) (*enterpriseTeam, *githubv3.Response, error) {
	u := fmt.Sprintf("enterprises/%s/teams/%s", enterpriseSlug, teamSlug)
	req, err := enterpriseTeamsNewRequest(client, "GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var raw json.RawMessage
	resp, err := client.Do(ctx, req, &raw)
	if err != nil {
		return nil, resp, err
	}

	te, err := parseEnterpriseTeam(raw)
	return te, resp, err
}

func createEnterpriseTeam(ctx context.Context, client *githubv3.Client, enterpriseSlug string, reqBody enterpriseTeamCreateRequest) (*enterpriseTeam, *githubv3.Response, error) {
	u := fmt.Sprintf("enterprises/%s/teams", enterpriseSlug)
	req, err := enterpriseTeamsNewRequest(client, "POST", u, reqBody)
	if err != nil {
		return nil, nil, err
	}

	var raw json.RawMessage
	resp, err := client.Do(ctx, req, &raw)
	if err != nil {
		return nil, resp, err
	}

	te, err := parseEnterpriseTeam(raw)
	return te, resp, err
}

func updateEnterpriseTeam(ctx context.Context, client *githubv3.Client, enterpriseSlug, teamSlug string, reqBody enterpriseTeamUpdateRequest) (*enterpriseTeam, *githubv3.Response, error) {
	u := fmt.Sprintf("enterprises/%s/teams/%s", enterpriseSlug, teamSlug)
	req, err := enterpriseTeamsNewRequest(client, "PATCH", u, reqBody)
	if err != nil {
		return nil, nil, err
	}

	var raw json.RawMessage
	resp, err := client.Do(ctx, req, &raw)
	if err != nil {
		return nil, resp, err
	}

	te, err := parseEnterpriseTeam(raw)
	return te, resp, err
}

func deleteEnterpriseTeam(ctx context.Context, client *githubv3.Client, enterpriseSlug, teamSlug string) (*githubv3.Response, error) {
	u := fmt.Sprintf("enterprises/%s/teams/%s", enterpriseSlug, teamSlug)
	req, err := enterpriseTeamsNewRequest(client, "DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func findEnterpriseTeamByID(ctx context.Context, client *githubv3.Client, enterpriseSlug string, id int64) (*enterpriseTeam, error) {
	teams, err := listEnterpriseTeams(ctx, client, enterpriseSlug)
	if err != nil {
		return nil, err
	}
	for _, t := range teams {
		if t.ID == id {
			copy := t
			return &copy, nil
		}
	}
	return nil, nil
}

type enterpriseOrg struct {
	Login string `json:"login"`
	ID    int64  `json:"id"`
}

func listEnterpriseTeamOrganizations(ctx context.Context, client *githubv3.Client, enterpriseSlug, enterpriseTeam string) ([]enterpriseOrg, error) {
	all := []enterpriseOrg{}
	opt := &githubv3.ListOptions{PerPage: maxPerPage}

	for {
		u := enterpriseTeamsAddListOptions(fmt.Sprintf("enterprises/%s/teams/%s/organizations", enterpriseSlug, enterpriseTeam), opt)
		req, err := enterpriseTeamsNewRequest(client, "GET", u, nil)
		if err != nil {
			return nil, err
		}

		var pageOrgs []enterpriseOrg
		resp, err := client.Do(ctx, req, &pageOrgs)
		if err != nil {
			// Some docs show a single object; tolerate that.
			var ghErr *githubv3.ErrorResponse
			if errors.As(err, &ghErr) {
				return nil, err
			}
			return nil, err
		}
		all = append(all, pageOrgs...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return all, nil
}

type enterpriseTeamOrgSlugsRequest struct {
	OrganizationSlugs []string `json:"organization_slugs"`
}

func addEnterpriseTeamOrganizations(ctx context.Context, client *githubv3.Client, enterpriseSlug, enterpriseTeam string, orgSlugs []string) error {
	if len(orgSlugs) == 0 {
		return nil
	}
	u := fmt.Sprintf("enterprises/%s/teams/%s/organizations/add", enterpriseSlug, enterpriseTeam)
	req, err := enterpriseTeamsNewRequest(client, "POST", u, enterpriseTeamOrgSlugsRequest{OrganizationSlugs: orgSlugs})
	if err != nil {
		return err
	}
	_, err = client.Do(ctx, req, nil)
	return err
}

func removeEnterpriseTeamOrganizations(ctx context.Context, client *githubv3.Client, enterpriseSlug, enterpriseTeam string, orgSlugs []string) (*githubv3.Response, error) {
	if len(orgSlugs) == 0 {
		return nil, nil
	}
	u := fmt.Sprintf("enterprises/%s/teams/%s/organizations/remove", enterpriseSlug, enterpriseTeam)
	req, err := enterpriseTeamsNewRequest(client, "POST", u, enterpriseTeamOrgSlugsRequest{OrganizationSlugs: orgSlugs})
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req, nil)
	return resp, err
}

func getEnterpriseTeamMembership(ctx context.Context, client *githubv3.Client, enterpriseSlug, enterpriseTeam, username string) (*githubv3.Response, error) {
	u := fmt.Sprintf("enterprises/%s/teams/%s/memberships/%s", enterpriseSlug, enterpriseTeam, username)
	req, err := enterpriseTeamsNewRequest(client, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req, nil)
	return resp, err
}

type enterpriseTeamMembership struct {
	State string `json:"state"`
	Role  string `json:"role"`
}

func parseEnterpriseTeamMembership(raw json.RawMessage) (*enterpriseTeamMembership, error) {
	var m enterpriseTeamMembership
	if err := json.Unmarshal(raw, &m); err == nil {
		if m.State != "" || m.Role != "" {
			return &m, nil
		}
	}

	var ms []enterpriseTeamMembership
	if err := json.Unmarshal(raw, &ms); err == nil {
		if len(ms) > 0 {
			return &ms[0], nil
		}
	}

	// If the API ever returns an empty object, keep a non-nil struct for callers.
	return &enterpriseTeamMembership{}, nil
}

func getEnterpriseTeamMembershipDetails(ctx context.Context, client *githubv3.Client, enterpriseSlug, enterpriseTeam, username string) (*enterpriseTeamMembership, *githubv3.Response, error) {
	u := fmt.Sprintf("enterprises/%s/teams/%s/memberships/%s", enterpriseSlug, enterpriseTeam, username)
	req, err := enterpriseTeamsNewRequest(client, "GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var raw json.RawMessage
	resp, err := client.Do(ctx, req, &raw)
	if err != nil {
		return nil, resp, err
	}

	m, err := parseEnterpriseTeamMembership(raw)
	return m, resp, err
}

func addEnterpriseTeamMember(ctx context.Context, client *githubv3.Client, enterpriseSlug, enterpriseTeam, username string) (*githubv3.Response, error) {
	u := fmt.Sprintf("enterprises/%s/teams/%s/memberships/%s", enterpriseSlug, enterpriseTeam, username)
	req, err := enterpriseTeamsNewRequest(client, "PUT", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req, nil)
	return resp, err
}

func removeEnterpriseTeamMember(ctx context.Context, client *githubv3.Client, enterpriseSlug, enterpriseTeam, username string) (*githubv3.Response, error) {
	u := fmt.Sprintf("enterprises/%s/teams/%s/memberships/%s", enterpriseSlug, enterpriseTeam, username)
	req, err := enterpriseTeamsNewRequest(client, "DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req, nil)
	return resp, err
}

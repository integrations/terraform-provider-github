package github

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-github/v67/github"
)

type enterpriseCostCenter struct {
	ID                string                         `json:"id"`
	Name              string                         `json:"name"`
	State             string                         `json:"state,omitempty"`
	AzureSubscription string                         `json:"azure_subscription,omitempty"`
	Resources         []enterpriseCostCenterResource `json:"resources,omitempty"`
}

type enterpriseCostCenterResource struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type enterpriseCostCenterListResponse struct {
	CostCenters []enterpriseCostCenter `json:"costCenters"`
}

type enterpriseCostCenterCreateRequest struct {
	Name string `json:"name"`
}

type enterpriseCostCenterUpdateRequest struct {
	Name string `json:"name"`
}

type enterpriseCostCenterArchiveResponse struct {
	Message         string `json:"message"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	CostCenterState string `json:"costCenterState"`
}

type enterpriseCostCenterResourcesRequest struct {
	Users         []string `json:"users,omitempty"`
	Organizations []string `json:"organizations,omitempty"`
	Repositories  []string `json:"repositories,omitempty"`
}

type enterpriseCostCenterAssignResponse struct {
	Message             string `json:"message"`
	ReassignedResources []struct {
		ResourceType       string `json:"resource_type"`
		Name               string `json:"name"`
		PreviousCostCenter string `json:"previous_cost_center"`
	} `json:"reassigned_resources"`
}

type enterpriseCostCenterRemoveResponse struct {
	Message string `json:"message"`
}

func enterpriseCostCentersList(ctx context.Context, client *github.Client, enterpriseSlug, state string) ([]enterpriseCostCenter, error) {
	u, err := url.Parse(fmt.Sprintf("enterprises/%s/settings/billing/cost-centers", enterpriseSlug))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	if state != "" {
		q.Set("state", state)
	}
	u.RawQuery = q.Encode()

	req, err := client.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var result enterpriseCostCenterListResponse
	_, err = client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return result.CostCenters, nil
}

func enterpriseCostCenterGet(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string) (*enterpriseCostCenter, error) {
	req, err := client.NewRequest("GET", fmt.Sprintf("enterprises/%s/settings/billing/cost-centers/%s", enterpriseSlug, costCenterID), nil)
	if err != nil {
		return nil, err
	}

	var result enterpriseCostCenter
	_, err = client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func enterpriseCostCenterCreate(ctx context.Context, client *github.Client, enterpriseSlug, name string) (*enterpriseCostCenter, error) {
	req, err := client.NewRequest("POST", fmt.Sprintf("enterprises/%s/settings/billing/cost-centers", enterpriseSlug), &enterpriseCostCenterCreateRequest{Name: name})
	if err != nil {
		return nil, err
	}

	var result enterpriseCostCenter
	_, err = client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func enterpriseCostCenterUpdate(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID, name string) (*enterpriseCostCenter, error) {
	req, err := client.NewRequest("PATCH", fmt.Sprintf("enterprises/%s/settings/billing/cost-centers/%s", enterpriseSlug, costCenterID), &enterpriseCostCenterUpdateRequest{Name: name})
	if err != nil {
		return nil, err
	}

	var result enterpriseCostCenter
	_, err = client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func enterpriseCostCenterArchive(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string) (*enterpriseCostCenterArchiveResponse, error) {
	req, err := client.NewRequest("DELETE", fmt.Sprintf("enterprises/%s/settings/billing/cost-centers/%s", enterpriseSlug, costCenterID), nil)
	if err != nil {
		return nil, err
	}

	var result enterpriseCostCenterArchiveResponse
	_, err = client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func enterpriseCostCenterAssignResources(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string, reqBody enterpriseCostCenterResourcesRequest) (*enterpriseCostCenterAssignResponse, error) {
	req, err := client.NewRequest("POST", fmt.Sprintf("enterprises/%s/settings/billing/cost-centers/%s/resource", enterpriseSlug, costCenterID), &reqBody)
	if err != nil {
		return nil, err
	}

	var result enterpriseCostCenterAssignResponse
	_, err = client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

//nolint:unparam
func enterpriseCostCenterRemoveResources(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string, reqBody enterpriseCostCenterResourcesRequest) (*enterpriseCostCenterRemoveResponse, error) {
	req, err := client.NewRequest("DELETE", fmt.Sprintf("enterprises/%s/settings/billing/cost-centers/%s/resource", enterpriseSlug, costCenterID), &reqBody)
	if err != nil {
		return nil, err
	}

	var result enterpriseCostCenterRemoveResponse
	_, err = client.Do(ctx, req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func enterpriseCostCenterSplitResources(resources []enterpriseCostCenterResource) (users, organizations, repositories []string) {
	for _, r := range resources {
		switch strings.ToLower(r.Type) {
		case "user":
			users = append(users, r.Name)
		case "org", "organization":
			organizations = append(organizations, r.Name)
		case "repo", "repository":
			repositories = append(repositories, r.Name)
		}
	}
	return users, organizations, repositories
}

func stringSliceToAnySlice(v []string) []any {
	out := make([]any, 0, len(v))
	for _, s := range v {
		out = append(out, s)
	}
	return out
}

func is404(err error) bool {
	var ghErr *github.ErrorResponse
	if errors.As(err, &ghErr) && ghErr.Response != nil {
		return ghErr.Response.StatusCode == 404
	}
	return false
}

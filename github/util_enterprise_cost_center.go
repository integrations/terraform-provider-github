package github

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-github/v92/github"
)

const (
	maxCostCenterResourcesPerRequest = 50

	// CostCenterResourceType constants match the API response values.
	CostCenterResourceTypeUser = "User"
	CostCenterResourceTypeOrg  = "Org"
	CostCenterResourceTypeRepo = "Repo"
)

type queryParameterTransport struct {
	base   http.RoundTripper
	values url.Values
}

func (t queryParameterTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	requestURL := *req.URL
	query := requestURL.Query()
	for key, values := range t.values {
		query.Del(key)
		for _, value := range values {
			query.Add(key, value)
		}
	}
	requestURL.RawQuery = query.Encode()
	req.URL = &requestURL

	return t.base.RoundTrip(req)
}

func githubClientWithListOptions(client *github.Client, opts github.ListOptions) (*github.Client, error) {
	query := url.Values{}
	if opts.Page > 0 {
		query.Set("page", strconv.Itoa(opts.Page))
	}
	if opts.PerPage > 0 {
		query.Set("per_page", strconv.Itoa(opts.PerPage))
	}

	transport := client.Client().Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	return client.Clone(github.WithTransport(queryParameterTransport{
		base:   transport,
		values: query,
	}))
}

func getEnterpriseCostCenter(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string, perPage int) (*github.CostCenter, error) {
	if perPage < 1 {
		return nil, fmt.Errorf("per-page limit must be at least 1")
	}

	var costCenter *github.CostCenter
	for page := 1; ; page++ {
		pagedClient, err := githubClientWithListOptions(client, github.ListOptions{Page: page, PerPage: perPage})
		if err != nil {
			return nil, err
		}

		result, resp, err := pagedClient.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
		if err != nil {
			return nil, err
		}
		if result == nil {
			return nil, fmt.Errorf("GitHub returned an empty response when reading cost center %q in enterprise %q", costCenterID, enterpriseSlug)
		}

		if costCenter == nil {
			costCenter = &github.CostCenter{
				ID:                result.ID,
				Name:              result.Name,
				State:             result.State,
				AzureSubscription: result.AzureSubscription,
			}
		}
		costCenter.Resources = append(costCenter.Resources, result.Resources...)

		if resp == nil || resp.NextPage == 0 {
			return costCenter, nil
		}
	}
}

func listEnterpriseCostCenters(ctx context.Context, client *github.Client, enterpriseSlug string, opts *github.ListCostCenterOptions, perPage int) ([]*github.CostCenter, error) {
	if perPage < 1 {
		return nil, fmt.Errorf("per-page limit must be at least 1")
	}

	var costCenters []*github.CostCenter
	for page := 1; ; page++ {
		pagedClient, err := githubClientWithListOptions(client, github.ListOptions{Page: page, PerPage: perPage})
		if err != nil {
			return nil, err
		}

		result, resp, err := pagedClient.Enterprise.ListCostCenters(ctx, enterpriseSlug, opts)
		if err != nil {
			return nil, err
		}
		if result == nil {
			return nil, fmt.Errorf("GitHub returned an empty response when listing cost centers for enterprise %q", enterpriseSlug)
		}

		costCenters = append(costCenters, result.CostCenters...)
		if resp == nil || resp.NextPage == 0 {
			return costCenters, nil
		}
	}
}

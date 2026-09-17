package github

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Cost center resource management constants and retry functions.
const (
	maxCostCenterResourcesPerRequest = 50
	costCenterResourcesRetryTimeout  = 5 * time.Minute

	// CostCenterResourceType constants match the API response values.
	CostCenterResourceTypeUser = "User"
	CostCenterResourceTypeOrg  = "Org"
	CostCenterResourceTypeRepo = "Repo"
)

type costCenterPage struct {
	ID                string                       `json:"id"`
	Name              string                       `json:"name"`
	Resources         []*github.CostCenterResource `json:"resources"`
	State             *string                      `json:"state,omitempty"`
	AzureSubscription *string                      `json:"azure_subscription,omitempty"`
	HasNextPage       bool                         `json:"has_next_page"`
}

func costCenterOwner(meta any) (*Owner, error) {
	owner, ok := meta.(*Owner)
	if !ok {
		return nil, fmt.Errorf("unexpected provider metadata type %T", meta)
	}
	return owner, nil
}

func costCenterString(d *schema.ResourceData, key string) (string, error) {
	value, ok := resourceKeysGetOk[string](d, key)
	if !ok {
		return "", fmt.Errorf("expected %q to be a non-empty string", key)
	}
	return value, nil
}

func costCenterStringSet(d *schema.ResourceData, key string) ([]string, error) {
	values, ok := resourceKeysGetOk[*schema.Set](d, key)
	if !ok {
		return nil, fmt.Errorf("expected %q to be a non-empty set of strings", key)
	}

	result := expandStringList(values.List())
	if len(result) != values.Len() {
		return nil, fmt.Errorf("expected %q to contain only non-empty strings", key)
	}
	return result, nil
}

func getEnterpriseCostCenter(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string) (*github.CostCenter, error) {
	const resourcesPerPage = 100

	var costCenter *github.CostCenter
	for page := 1; ; page++ {
		query := url.Values{}
		query.Set("page", strconv.Itoa(page))
		query.Set("per_page", strconv.Itoa(resourcesPerPage))

		endpoint := fmt.Sprintf(
			"enterprises/%s/settings/billing/cost-centers/%s?%s",
			url.PathEscape(enterpriseSlug),
			url.PathEscape(costCenterID),
			query.Encode(),
		)
		req, err := client.NewRequest(ctx, http.MethodGet, endpoint, nil)
		if err != nil {
			return nil, err
		}

		result := new(costCenterPage)
		_, err = client.Do(req, result)
		if err != nil {
			return nil, err
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

		if !result.HasNextPage {
			return costCenter, nil
		}
	}
}

// retryCostCenterRemoveResources removes resources from a cost center with retry logic.
// Uses retry.RetryContext for exponential backoff on transient errors.
func retryCostCenterRemoveResources(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string, req github.CostCenterResourceRequest) diag.Diagnostics {
	err := retry.RetryContext(ctx, costCenterResourcesRetryTimeout, func() *retry.RetryError {
		_, _, err := client.Enterprise.RemoveResourcesFromCostCenter(ctx, enterpriseSlug, costCenterID, req)
		if err == nil {
			return nil
		}
		if errIsRetryable(err) {
			return retry.RetryableError(err)
		}
		return retry.NonRetryableError(err)
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

// retryCostCenterAddResources adds resources to a cost center with retry logic.
// Uses retry.RetryContext for exponential backoff on transient errors.
func retryCostCenterAddResources(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string, req github.CostCenterResourceRequest) diag.Diagnostics {
	err := retry.RetryContext(ctx, costCenterResourcesRetryTimeout, func() *retry.RetryError {
		_, _, err := client.Enterprise.AddResourcesToCostCenter(ctx, enterpriseSlug, costCenterID, req)
		if err == nil {
			return nil
		}
		if errIsRetryable(err) {
			return retry.RetryableError(err)
		}
		return retry.NonRetryableError(err)
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

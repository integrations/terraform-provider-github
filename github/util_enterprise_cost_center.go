package github

import (
	"context"
	"fmt"
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

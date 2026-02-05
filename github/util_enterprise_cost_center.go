package github

import (
	"context"
	"time"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

// Cost center resource management constants and retry functions.
const (
	maxResourcesPerRequest          = 50
	costCenterResourcesRetryTimeout = 5 * time.Minute
)

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

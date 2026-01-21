package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseCostCenter() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages an enterprise cost center in GitHub.",
		CreateContext: resourceGithubEnterpriseCostCenterCreate,
		ReadContext:   resourceGithubEnterpriseCostCenterRead,
		UpdateContext: resourceGithubEnterpriseCostCenterUpdate,
		DeleteContext: resourceGithubEnterpriseCostCenterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseCostCenterImport,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the cost center.",
			},
			"users": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The usernames assigned to this cost center.",
			},
			"organizations": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The organization logins assigned to this cost center.",
			},
			"repositories": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The repositories (full name) assigned to this cost center.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the cost center.",
			},
			"azure_subscription": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Azure subscription associated with the cost center.",
			},
		},
	}
}

func resourceGithubEnterpriseCostCenterCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	name := d.Get("name").(string)

	tflog.Info(ctx, "Creating enterprise cost center", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"name":            name,
	})

	cc, _, err := client.Enterprise.CreateCostCenter(ctx, enterpriseSlug, github.CostCenterRequest{Name: name})
	if err != nil {
		return diag.FromErr(err)
	}

	if cc == nil || cc.ID == "" {
		return diag.Errorf("failed to create cost center: missing id in response (unexpected API response; please retry or contact support)")
	}

	d.SetId(cc.ID)

	if hasCostCenterAssignmentsConfigured(d) {
		if diags := syncEnterpriseCostCenterAssignments(ctx, d, client, enterpriseSlug, cc.ID); diags.HasError() {
			return diags
		}
	}

	// Set computed fields from the API response
	state := strings.ToLower(cc.GetState())
	if state == "" {
		state = "active"
	}
	if err := d.Set("state", state); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("azure_subscription", cc.GetAzureSubscription()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Id()

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		if is404(err) {
			tflog.Warn(ctx, "Cost center not found, removing from state", map[string]any{
				"enterprise_slug": enterpriseSlug,
				"cost_center_id":  costCenterID,
			})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("name", cc.Name); err != nil {
		return diag.FromErr(err)
	}

	state := strings.ToLower(cc.GetState())
	if state == "" {
		state = "active"
	}
	if err := d.Set("state", state); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("azure_subscription", cc.GetAzureSubscription()); err != nil {
		return diag.FromErr(err)
	}

	if err := setCostCenterResourceFields(d, cc); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Id()

	// Check current state to prevent updates on archived cost centers
	currentState := d.Get("state").(string)
	if strings.EqualFold(currentState, "deleted") {
		return diag.Errorf("cannot update cost center %q because it is archived", costCenterID)
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		tflog.Info(ctx, "Updating enterprise cost center name", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"name":            name,
		})
		_, _, err := client.Enterprise.UpdateCostCenter(ctx, enterpriseSlug, costCenterID, github.CostCenterRequest{Name: name})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	if d.HasChange("users") || d.HasChange("organizations") || d.HasChange("repositories") {
		if diags := syncCostCenterAssignmentsFromState(ctx, d, client, enterpriseSlug, costCenterID); diags.HasError() {
			return diags
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Id()

	tflog.Info(ctx, "Archiving enterprise cost center", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"cost_center_id":  costCenterID,
	})

	_, _, err := client.Enterprise.DeleteCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		if is404(err) {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, costCenterID, err := parseTwoPartID(d.Id(), "enterprise_slug", "cost_center_id")
	if err != nil {
		return nil, fmt.Errorf("invalid import ID %q: expected format <enterprise_slug>:<cost_center_id>", d.Id())
	}

	d.SetId(costCenterID)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func syncEnterpriseCostCenterAssignments(ctx context.Context, d *schema.ResourceData, client *github.Client, enterpriseSlug, costCenterID string) diag.Diagnostics {
	desiredUsers := expandStringSet(getStringSetOrEmpty(d, "users"))
	desiredOrgs := expandStringSet(getStringSetOrEmpty(d, "organizations"))
	desiredRepos := expandStringSet(getStringSetOrEmpty(d, "repositories"))

	if len(desiredUsers)+len(desiredOrgs)+len(desiredRepos) > 0 {
		tflog.Info(ctx, "Assigning enterprise cost center resources", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
		})

		for _, batch := range chunkStringSlice(desiredUsers) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunkStringSlice(desiredOrgs) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Organizations: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunkStringSlice(desiredRepos) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

// syncCostCenterAssignmentsFromState syncs assignments using d.GetChange (Terraform as source of truth).
func syncCostCenterAssignmentsFromState(ctx context.Context, d *schema.ResourceData, client *github.Client, enterpriseSlug, costCenterID string) diag.Diagnostics {
	var toAddUsers, toRemoveUsers, toAddOrgs, toRemoveOrgs, toAddRepos, toRemoveRepos []string

	if d.HasChange("users") {
		oldSet, newSet := d.GetChange("users")
		toRemoveUsers, toAddUsers = diffSets(oldSet.(*schema.Set), newSet.(*schema.Set))
	}
	if d.HasChange("organizations") {
		oldSet, newSet := d.GetChange("organizations")
		toRemoveOrgs, toAddOrgs = diffSets(oldSet.(*schema.Set), newSet.(*schema.Set))
	}
	if d.HasChange("repositories") {
		oldSet, newSet := d.GetChange("repositories")
		toRemoveRepos, toAddRepos = diffSets(oldSet.(*schema.Set), newSet.(*schema.Set))
	}

	if len(toRemoveUsers)+len(toRemoveOrgs)+len(toRemoveRepos) > 0 {
		tflog.Info(ctx, "Removing enterprise cost center resources", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
		})

		for _, batch := range chunkStringSlice(toRemoveUsers) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunkStringSlice(toRemoveOrgs) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Organizations: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunkStringSlice(toRemoveRepos) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	if len(toAddUsers)+len(toAddOrgs)+len(toAddRepos) > 0 {
		tflog.Info(ctx, "Assigning enterprise cost center resources", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
		})

		for _, batch := range chunkStringSlice(toAddUsers) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunkStringSlice(toAddOrgs) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Organizations: batch}); diags.HasError() {
				return diags
			}
		}
		for _, batch := range chunkStringSlice(toAddRepos) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Repositories: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

func hasCostCenterAssignmentsConfigured(d *schema.ResourceData) bool {
	assignmentKeys := []string{"users", "organizations", "repositories"}
	for _, key := range assignmentKeys {
		if v, ok := d.GetOk(key); ok {
			if set, ok := v.(*schema.Set); ok && set != nil && set.Len() > 0 {
				return true
			}
		}
	}
	return false
}

func expandStringSet(set *schema.Set) []string {
	if set == nil {
		return nil
	}

	list := set.List()
	return expandStringList(list)
}

func getStringSetOrEmpty(d *schema.ResourceData, key string) *schema.Set {
	v, ok := d.GetOk(key)
	if !ok || v == nil {
		return schema.NewSet(schema.HashString, []any{})
	}

	set, ok := v.(*schema.Set)
	if !ok || set == nil {
		return schema.NewSet(schema.HashString, []any{})
	}

	return set
}

// diffSets returns elements to remove (in old but not new) and to add (in new but not old).
func diffSets(oldSet, newSet *schema.Set) (toRemove, toAdd []string) {
	for _, v := range oldSet.Difference(newSet).List() {
		toRemove = append(toRemove, v.(string))
	}
	for _, v := range newSet.Difference(oldSet).List() {
		toAdd = append(toAdd, v.(string))
	}
	return toRemove, toAdd
}

func isRetryableGithubResponseError(err error) bool {
	var ghErr *github.ErrorResponse
	if errors.As(err, &ghErr) && ghErr.Response != nil {
		switch ghErr.Response.StatusCode {
		case http.StatusConflict, http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable, http.StatusGatewayTimeout:
			return true
		default:
			return false
		}
	}
	return false
}

func costCenterSplitResources(resources []*github.CostCenterResource) (users, organizations, repositories []string) {
	for _, r := range resources {
		if r == nil {
			continue
		}
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

// setCostCenterResourceFields sets the resource-related fields on the schema.ResourceData.
func setCostCenterResourceFields(d *schema.ResourceData, cc *github.CostCenter) error {
	users, organizations, repositories := costCenterSplitResources(cc.Resources)
	if err := d.Set("users", flattenStringList(users)); err != nil {
		return err
	}
	if err := d.Set("organizations", flattenStringList(organizations)); err != nil {
		return err
	}
	if err := d.Set("repositories", flattenStringList(repositories)); err != nil {
		return err
	}
	return nil
}

// Cost center resource management constants and retry functions.
const (
	maxResourcesPerRequest          = 50
	costCenterResourcesRetryTimeout = 5 * time.Minute
)

// chunkStringSlice splits a slice into chunks of the max resources per request.
func chunkStringSlice(items []string) [][]string {
	if len(items) == 0 {
		return nil
	}
	chunks := make([][]string, 0, (len(items)+maxResourcesPerRequest-1)/maxResourcesPerRequest)
	for start := 0; start < len(items); start += maxResourcesPerRequest {
		end := min(start+maxResourcesPerRequest, len(items))
		chunks = append(chunks, items[start:end])
	}
	return chunks
}

// retryCostCenterRemoveResources removes resources from a cost center with retry logic.
// Uses retry.RetryContext for exponential backoff on transient errors.
func retryCostCenterRemoveResources(ctx context.Context, client *github.Client, enterpriseSlug, costCenterID string, req github.CostCenterResourceRequest) diag.Diagnostics {
	err := retry.RetryContext(ctx, costCenterResourcesRetryTimeout, func() *retry.RetryError {
		_, _, err := client.Enterprise.RemoveResourcesFromCostCenter(ctx, enterpriseSlug, costCenterID, req)
		if err == nil {
			return nil
		}
		if isRetryableGithubResponseError(err) {
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
		if isRetryableGithubResponseError(err) {
			return retry.RetryableError(err)
		}
		return retry.NonRetryableError(err)
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

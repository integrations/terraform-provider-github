package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseCostCenterUsers() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages user assignments for a GitHub enterprise cost center (authoritative).",
		CreateContext: resourceGithubEnterpriseCostCenterUsersCreateOrUpdate,
		ReadContext:   resourceGithubEnterpriseCostCenterUsersRead,
		UpdateContext: resourceGithubEnterpriseCostCenterUsersCreateOrUpdate,
		DeleteContext: resourceGithubEnterpriseCostCenterUsersDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseCostCenterUsersImport,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"cost_center_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the cost center.",
			},
			"usernames": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Usernames to assign to the cost center. This is authoritative - users not in this set will be removed.",
			},
		},
	}
}

func resourceGithubEnterpriseCostCenterUsersCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	// If this is Create, set the ID
	if d.Id() == "" {
		id, err := buildID(enterpriseSlug, costCenterID)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(id)
	}

	// Get current assignments from API
	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}

	// Extract current users
	currentUsers := make(map[string]bool)
	for _, r := range cc.Resources {
		if r != nil && r.Type == "User" {
			currentUsers[r.Name] = true
		}
	}

	// Get desired users from config
	desiredUsersSet := d.Get("usernames").(*schema.Set)
	desiredUsers := make(map[string]bool)
	for _, u := range desiredUsersSet.List() {
		desiredUsers[u.(string)] = true
	}

	// Calculate additions and removals
	var toAdd, toRemove []string
	for user := range desiredUsers {
		if !currentUsers[user] {
			toAdd = append(toAdd, user)
		}
	}
	for user := range currentUsers {
		if !desiredUsers[user] {
			toRemove = append(toRemove, user)
		}
	}

	// Remove users no longer desired
	if len(toRemove) > 0 {
		tflog.Info(ctx, "Removing users from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toRemove),
		})

		for _, batch := range chunkStringSlice(toRemove, maxResourcesPerRequest) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
	}

	// Add new users
	if len(toAdd) > 0 {
		tflog.Info(ctx, "Adding users to cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toAdd),
		})

		for _, batch := range chunkStringSlice(toAdd, maxResourcesPerRequest) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return resourceGithubEnterpriseCostCenterUsersRead(ctx, d, meta)
}

func resourceGithubEnterpriseCostCenterUsersRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		if errIs404(err) {
			tflog.Warn(ctx, "Cost center not found, removing from state", map[string]any{
				"enterprise_slug": enterpriseSlug,
				"cost_center_id":  costCenterID,
			})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	// Extract users from resources
	var users []string
	for _, r := range cc.Resources {
		if r != nil && r.Type == "User" {
			users = append(users, r.Name)
		}
	}

	if err := d.Set("usernames", flattenStringList(users)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterUsersDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	usernamesSet := d.Get("usernames").(*schema.Set)
	usernames := expandStringList(usernamesSet.List())

	if len(usernames) > 0 {
		tflog.Info(ctx, "Removing all users from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(usernames),
		})

		for _, batch := range chunkStringSlice(usernames, maxResourcesPerRequest) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterUsersImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, costCenterID, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid import ID %q: expected format <enterprise_slug>:<cost_center_id>", d.Id())
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	if err := d.Set("cost_center_id", costCenterID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

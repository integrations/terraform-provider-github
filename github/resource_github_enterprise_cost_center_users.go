package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseCostCenterUsers() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages user assignments for a GitHub enterprise cost center (authoritative).",
		CreateContext: resourceGithubEnterpriseCostCenterUsersCreate,
		ReadContext:   resourceGithubEnterpriseCostCenterUsersRead,
		UpdateContext: resourceGithubEnterpriseCostCenterUsersUpdate,
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

func resourceGithubEnterpriseCostCenterUsersCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeUser {
			return diag.Errorf("cost center %q already has users assigned; import the existing assignments first or remove them manually", costCenterID)
		}
	}

	desiredUsersSet := d.Get("usernames").(*schema.Set)
	toAdd := expandStringList(desiredUsersSet.List())

	tflog.Info(ctx, "Adding users to cost center", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"cost_center_id":  costCenterID,
		"count":           len(toAdd),
	})

	for _, batch := range chunkStringSlice(toAdd, maxCostCenterResourcesPerRequest) {
		if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
			return diags
		}
	}

	d.SetId(costCenterID)
	return nil
}

func resourceGithubEnterpriseCostCenterUsersUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}

	diff := make(map[string]bool)
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeUser {
			diff[ccResource.Name] = false
		}
	}

	var toAdd []string
	for _, user := range d.Get("usernames").(*schema.Set).List() {
		name := user.(string)
		if _, exists := diff[name]; exists {
			diff[name] = true
		} else {
			toAdd = append(toAdd, name)
		}
	}

	var toRemove []string
	for name, keep := range diff {
		if !keep {
			toRemove = append(toRemove, name)
		}
	}

	if len(toRemove) > 0 {
		tflog.Info(ctx, "Removing users from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toRemove),
		})

		for _, batch := range chunkStringSlice(toRemove, maxCostCenterResourcesPerRequest) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
	}

	if len(toAdd) > 0 {
		tflog.Info(ctx, "Adding users to cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toAdd),
		})

		for _, batch := range chunkStringSlice(toAdd, maxCostCenterResourcesPerRequest) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Users: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
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

	var users []string
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeUser {
			users = append(users, ccResource.Name)
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

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		if errIs404(err) {
			return nil
		}
		return diag.FromErr(err)
	}

	var usernames []string
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeUser {
			usernames = append(usernames, ccResource.Name)
		}
	}

	if len(usernames) > 0 {
		tflog.Info(ctx, "Removing all users from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(usernames),
		})

		for _, batch := range chunkStringSlice(usernames, maxCostCenterResourcesPerRequest) {
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

	d.SetId(costCenterID)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	if err := d.Set("cost_center_id", costCenterID); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseCostCenterOrganizations() *schema.Resource {
	return &schema.Resource{
		Description:   "Manages organization assignments for a GitHub enterprise cost center (authoritative).",
		CreateContext: resourceGithubEnterpriseCostCenterOrganizationsCreate,
		ReadContext:   resourceGithubEnterpriseCostCenterOrganizationsRead,
		UpdateContext: resourceGithubEnterpriseCostCenterOrganizationsUpdate,
		DeleteContext: resourceGithubEnterpriseCostCenterOrganizationsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubEnterpriseCostCenterOrganizationsImport,
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
			"organization_logins": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Organization logins to assign to the cost center. This is authoritative - organizations not in this set will be removed.",
			},
		},
	}
}

func resourceGithubEnterpriseCostCenterOrganizationsCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeOrg {
			return diag.Errorf("cost center %q already has organizations assigned; import the existing assignments first or remove them manually", costCenterID)
		}
	}

	desiredOrgsSet := d.Get("organization_logins").(*schema.Set)
	toAdd := expandStringList(desiredOrgsSet.List())

	tflog.Info(ctx, "Adding organizations to cost center", map[string]any{
		"enterprise_slug": enterpriseSlug,
		"cost_center_id":  costCenterID,
		"count":           len(toAdd),
	})

	for _, batch := range chunkStringSlice(toAdd, maxCostCenterResourcesPerRequest) {
		if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Organizations: batch}); diags.HasError() {
			return diags
		}
	}

	d.SetId(costCenterID)
	return nil
}

func resourceGithubEnterpriseCostCenterOrganizationsUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}

	diff := make(map[string]bool)
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeOrg {
			diff[ccResource.Name] = false
		}
	}

	var toAdd []string
	for _, org := range d.Get("organization_logins").(*schema.Set).List() {
		name := org.(string)
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
		tflog.Info(ctx, "Removing organizations from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toRemove),
		})

		for _, batch := range chunkStringSlice(toRemove, maxCostCenterResourcesPerRequest) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Organizations: batch}); diags.HasError() {
				return diags
			}
		}
	}

	if len(toAdd) > 0 {
		tflog.Info(ctx, "Adding organizations to cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(toAdd),
		})

		for _, batch := range chunkStringSlice(toAdd, maxCostCenterResourcesPerRequest) {
			if diags := retryCostCenterAddResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Organizations: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterOrganizationsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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

	var organizations []string
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeOrg {
			organizations = append(organizations, ccResource.Name)
		}
	}

	if err := d.Set("organization_logins", flattenStringList(organizations)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterOrganizationsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
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

	var organizations []string
	for _, ccResource := range cc.Resources {
		if ccResource != nil && ccResource.Type == CostCenterResourceTypeOrg {
			organizations = append(organizations, ccResource.Name)
		}
	}

	if len(organizations) > 0 {
		tflog.Info(ctx, "Removing all organizations from cost center", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
			"count":           len(organizations),
		})

		for _, batch := range chunkStringSlice(organizations, maxCostCenterResourcesPerRequest) {
			if diags := retryCostCenterRemoveResources(ctx, client, enterpriseSlug, costCenterID, github.CostCenterResourceRequest{Organizations: batch}); diags.HasError() {
				return diags
			}
		}
	}

	return nil
}

func resourceGithubEnterpriseCostCenterOrganizationsImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
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

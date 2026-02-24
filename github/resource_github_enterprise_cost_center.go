package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

	if err := d.Set("state", cc.GetState()); err != nil {
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

	// If the cost center is archived (deleted), remove from state
	if cc.GetState() == "deleted" {
		tflog.Warn(ctx, "Cost center is archived, removing from state", map[string]any{
			"enterprise_slug": enterpriseSlug,
			"cost_center_id":  costCenterID,
		})
		d.SetId("")
		return nil
	}

	if err := d.Set("name", cc.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", cc.GetState()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("azure_subscription", cc.GetAzureSubscription()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Id()

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
		if errIs404(err) {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCostCenterImport(ctx context.Context, d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, costCenterID, err := parseID2(d.Id())
	if err != nil {
		return nil, fmt.Errorf("invalid import ID %q: expected format <enterprise_slug>:<cost_center_id>", d.Id())
	}

	client := meta.(*Owner).v3client
	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		return nil, fmt.Errorf("error reading cost center %q in enterprise %q: %w", costCenterID, enterpriseSlug, err)
	}

	d.SetId(costCenterID)
	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	if err := d.Set("name", cc.Name); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

package github

import (
	"context"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseCostCenters() *schema.Resource {
	return &schema.Resource{
		Description: "Retrieves a list of GitHub enterprise cost centers.",
		ReadContext: dataSourceGithubEnterpriseCostCentersRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "all",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"all", "active", "deleted"}, false)),
				Description:      "Filter cost centers by state. Valid values are 'all', 'active', and 'deleted'.",
			},
			"cost_centers": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The list of cost centers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cost center ID.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
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
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseCostCentersRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner, ok := meta.(*Owner)
	if !ok {
		return diag.Errorf("unexpected provider metadata type %T", meta)
	}
	enterpriseSlug, ok := resourceKeysGetOk[string](d, "enterprise_slug")
	if !ok {
		return diag.Errorf("expected enterprise_slug to be a non-empty string")
	}
	stateFilter, ok := resourceKeysGetOk[string](d, "state")
	if !ok {
		return diag.Errorf("expected state to be a non-empty string")
	}

	var opts github.ListCostCenterOptions
	if stateFilter != "all" {
		opts.State = &stateFilter
	}

	costCenters, err := listEnterpriseCostCenters(ctx, owner.v3client, enterpriseSlug, &opts, owner.maxPerPage)
	if err != nil {
		return diag.FromErr(err)
	}

	items := make([]any, 0, len(costCenters))
	for _, cc := range costCenters {
		if cc == nil {
			continue
		}
		items = append(items, map[string]any{
			"id":                 cc.ID,
			"name":               cc.Name,
			"state":              cc.GetState(),
			"azure_subscription": cc.GetAzureSubscription(),
		})
	}

	id, err := buildID(enterpriseSlug, stateFilter)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	if err := d.Set("cost_centers", items); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

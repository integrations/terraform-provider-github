package github

import (
	"context"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseCostCenters() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve a list of enterprise cost centers.",
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
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"active", "deleted"}, false)),
				Description:      "Filter cost centers by state.",
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
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	var state *string
	if v, ok := d.GetOk("state"); ok {
		state = github.Ptr(v.(string))
	}

	result, _, err := client.Enterprise.ListCostCenters(ctx, enterpriseSlug, &github.ListCostCenterOptions{State: state})
	if err != nil {
		return diag.FromErr(err)
	}

	items := make([]any, 0, len(result.CostCenters))
	for _, cc := range result.CostCenters {
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

	stateStr := "all"
	if state != nil {
		stateStr = *state
	}
	id, err := buildID(enterpriseSlug, stateStr)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	if err := d.Set("cost_centers", items); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

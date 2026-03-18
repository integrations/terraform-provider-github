package github

import (
	"context"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseBillingUsage() *schema.Resource {
	return &schema.Resource{
		Description: "Gets a billing usage report for a GitHub enterprise.",
		ReadContext: dataSourceGithubEnterpriseBillingUsageRead,
		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The slug of the enterprise.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
			"year": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "If specified, only return results for a single year.",
				ValidateFunc: validation.IntAtLeast(2000),
			},
			"month": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "If specified, only return results for a single month. Value between 1 and 12.",
				ValidateFunc: validation.IntBetween(1, 12),
			},
			"day": {
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "If specified, only return results for a single day. Value between 1 and 31.",
				ValidateFunc: validation.IntBetween(1, 31),
			},
			"cost_center_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID corresponding to a cost center. Use `none` to target usage not associated to any cost center.",
			},
			"usage_items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of billing usage items.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date of the usage item.",
						},
						"product": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The product name.",
						},
						"sku": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SKU name.",
						},
						"quantity": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The quantity of usage.",
						},
						"unit_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of unit for the usage.",
						},
						"price_per_unit": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The price per unit of usage.",
						},
						"gross_amount": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The gross amount of usage.",
						},
						"discount_amount": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The discount amount applied.",
						},
						"net_amount": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The net amount after discounts.",
						},
						"organization_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The organization name associated with the usage.",
						},
						"repository_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repository name associated with the usage.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseBillingUsageRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)

	opts := &EnterpriseBillingUsageOptions{}
	if yearVal, ok := d.GetOk("year"); ok {
		opts.Year = github.Ptr(yearVal.(int))
	}
	if monthVal, ok := d.GetOk("month"); ok {
		opts.Month = github.Ptr(monthVal.(int))
	}
	if dayVal, ok := d.GetOk("day"); ok {
		opts.Day = github.Ptr(dayVal.(int))
	}
	if costCenterID, ok := d.GetOk("cost_center_id"); ok {
		opts.CostCenterID = github.Ptr(costCenterID.(string))
	}

	report, err := getEnterpriseBillingUsage(ctx, client, enterpriseSlug, opts)
	if err != nil {
		return diag.Errorf("error getting enterprise billing usage for %q: %s", enterpriseSlug, err)
	}

	id, err := buildID(enterpriseSlug, "billing-usage")
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("usage_items", flattenUsageItems(report.UsageItems)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

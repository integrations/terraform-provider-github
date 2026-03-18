package github

import (
	"context"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubEnterpriseBillingUsageSummary() *schema.Resource {
	return &schema.Resource{
		Description: "Gets a billing usage summary report for a GitHub enterprise. This API is in public preview and subject to change.",
		ReadContext: dataSourceGithubEnterpriseBillingUsageSummaryRead,
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
			"organization": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The organization name to query usage for.",
			},
			"repository": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The repository name to query usage for.",
			},
			"product": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The product name to query usage for.",
			},
			"sku": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SKU name to query usage for.",
			},
			"cost_center_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID corresponding to a cost center. Use `none` to target usage not associated to any cost center.",
			},
			"time_period": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The time period of the report.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"year": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The year of the time period.",
						},
						"month": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The month of the time period.",
						},
						"day": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The day of the time period.",
						},
					},
				},
			},
			"enterprise": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The enterprise name from the report.",
			},
			"usage_items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of usage summary items.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"gross_quantity": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The gross quantity of usage.",
						},
						"gross_amount": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The gross amount of usage.",
						},
						"discount_quantity": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The discount quantity applied.",
						},
						"discount_amount": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The discount amount applied.",
						},
						"net_quantity": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The net quantity after discounts.",
						},
						"net_amount": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "The net amount after discounts.",
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseBillingUsageSummaryRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)

	opts := &EnterpriseUsageSummaryOptions{}
	if yearVal, ok := d.GetOk("year"); ok {
		opts.Year = github.Ptr(yearVal.(int))
	}
	if monthVal, ok := d.GetOk("month"); ok {
		opts.Month = github.Ptr(monthVal.(int))
	}
	if dayVal, ok := d.GetOk("day"); ok {
		opts.Day = github.Ptr(dayVal.(int))
	}
	if orgVal, ok := d.GetOk("organization"); ok {
		opts.Organization = github.Ptr(orgVal.(string))
	}
	if repoVal, ok := d.GetOk("repository"); ok {
		opts.Repository = github.Ptr(repoVal.(string))
	}
	if productVal, ok := d.GetOk("product"); ok {
		opts.Product = github.Ptr(productVal.(string))
	}
	if skuVal, ok := d.GetOk("sku"); ok {
		opts.SKU = github.Ptr(skuVal.(string))
	}
	if costCenterID, ok := d.GetOk("cost_center_id"); ok {
		opts.CostCenterID = github.Ptr(costCenterID.(string))
	}

	report, err := getEnterpriseUsageSummary(ctx, client, enterpriseSlug, opts)
	if err != nil {
		return diag.Errorf("error getting enterprise billing usage summary for %q: %s", enterpriseSlug, err)
	}

	id, err := buildID(enterpriseSlug, "billing-usage-summary")
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("time_period", flattenTimePeriod(report.TimePeriod)); err != nil {
		return diag.FromErr(err)
	}
	if report.Enterprise != nil {
		if err := d.Set("enterprise", *report.Enterprise); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("usage_items", flattenUsageSummaryItems(report.UsageItems)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

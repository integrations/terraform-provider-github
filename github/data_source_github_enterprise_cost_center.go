package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseCostCenter() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve information about a specific enterprise cost center.",
		ReadContext: dataSourceGithubEnterpriseCostCenterRead,

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the enterprise.",
			},
			"cost_center_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the cost center.",
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
			"users": {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The usernames assigned to this cost center.",
			},
			"organizations": {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The organization logins assigned to this cost center.",
			},
			"repositories": {
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The repositories (full name) assigned to this cost center.",
			},
		},
	}
}

func dataSourceGithubEnterpriseCostCenterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	cc, _, err := client.Enterprise.GetCostCenter(ctx, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(costCenterID)
	if err := d.Set("name", cc.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("state", cc.GetState()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("azure_subscription", cc.GetAzureSubscription()); err != nil {
		return diag.FromErr(err)
	}

	users := make([]string, 0)
	organizations := make([]string, 0)
	repositories := make([]string, 0)
	for _, resource := range cc.Resources {
		if resource == nil {
			continue
		}
		switch resource.Type {
		case CostCenterResourceTypeUser:
			users = append(users, resource.Name)
		case CostCenterResourceTypeOrg:
			organizations = append(organizations, resource.Name)
		case CostCenterResourceTypeRepo:
			repositories = append(repositories, resource.Name)
		}
	}

	if err := d.Set("users", users); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("organizations", organizations); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("repositories", repositories); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

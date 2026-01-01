package github

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseCostCenter() *schema.Resource {
	return &schema.Resource{
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
			"resources": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubEnterpriseCostCenterRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	enterpriseSlug := d.Get("enterprise_slug").(string)
	costCenterID := d.Get("cost_center_id").(string)

	ctx = context.WithValue(ctx, ctxId, fmt.Sprintf("%s/%s", enterpriseSlug, costCenterID))

	cc, err := enterpriseCostCenterGet(ctx, client, enterpriseSlug, costCenterID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(costCenterID)
	_ = d.Set("name", cc.Name)

	state := strings.ToLower(cc.State)
	if state == "" {
		state = "active"
	}
	_ = d.Set("state", state)
	_ = d.Set("azure_subscription", cc.AzureSubscription)

	resources := make([]map[string]any, 0)
	for _, r := range cc.Resources {
		resources = append(resources, map[string]any{
			"type": r.Type,
			"name": r.Name,
		})
	}
	_ = d.Set("resources", resources)

	users, organizations, repositories := enterpriseCostCenterSplitResources(cc.Resources)
	sort.Strings(users)
	sort.Strings(organizations)
	sort.Strings(repositories)
	_ = d.Set("users", stringSliceToAnySlice(users))
	_ = d.Set("organizations", stringSliceToAnySlice(organizations))
	_ = d.Set("repositories", stringSliceToAnySlice(repositories))

	return nil
}

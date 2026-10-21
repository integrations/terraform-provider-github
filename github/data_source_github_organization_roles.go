package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRoles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationRolesRead,

		Description: "Data source to list all custom roles in an organization.",

		Schema: map[string]*schema.Schema{
			"roles": {
				Description: "Available organization roles.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the organization role.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"role_id": {
							Description: "ID of the organization role.",
							Type:        schema.TypeInt,
							Computed:    true,
							Deprecated:  "The `role_id` attribute is deprecated and will be removed in a future version of the provider. Use `id` instead.",
						},
						"name": {
							Description: "Name of the organization role.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "Description of the organization role.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"source": {
							Description: "Source of this role; one of `Predefined`, `Organization`, or `Enterprise`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"base_role": {
							Description: "System role from which this role inherits permissions.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"permissions": {
							Description: "Additional permissions included in this role.",
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationRolesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	ret, _, err := meta.v3client.Organizations.ListRoles(ctx, meta.name)
	if err != nil {
		return diag.FromErr(err)
	}

	allRoles := make([]any, ret.GetTotalCount())
	for i, role := range ret.CustomRepoRoles {
		r := map[string]any{
			"id":          int(role.GetID()),
			"role_id":     int(role.GetID()),
			"name":        role.GetName(),
			"description": role.GetDescription(),
			"source":      role.GetSource(),
			"base_role":   role.GetBaseRole(),
			"permissions": role.Permissions,
		}
		allRoles[i] = r
	}

	d.SetId(meta.name)

	if err := d.Set("roles", allRoles); err != nil {
		return diag.Errorf("error setting roles: %v", err)
	}

	return nil
}

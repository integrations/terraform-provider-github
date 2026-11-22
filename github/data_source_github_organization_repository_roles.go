package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRepositoryRoles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationRepositoryRolesRead,

		Description: "Data source to list all custom repository roles in an organization.",

		Schema: map[string]*schema.Schema{
			"roles": {
				Description: "Available organization repository roles.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "ID of the organization repository role.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"role_id": {
							Description: "ID of the organization repository role.",
							Type:        schema.TypeInt,
							Computed:    true,
							Deprecated:  "The `role_id` attribute is deprecated and will be removed in a future version of the provider. Use `id` instead.",
						},
						"name": {
							Description: "Name of the organization repository role.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"description": {
							Description: "Description of the organization repository role.",
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

func dataSourceGithubOrganizationRepositoryRolesRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	ret, _, err := meta.v3client.Organizations.ListCustomRepoRoles(ctx, meta.name)
	if err != nil {
		return diag.FromErr(err)
	}

	roles := make([]any, ret.GetTotalCount())
	for i, role := range ret.CustomRepoRoles {
		r := map[string]any{
			"id":          role.GetID(),
			"role_id":     role.GetID(),
			"name":        role.GetName(),
			"description": role.GetDescription(),
			"base_role":   role.GetBaseRole(),
			"permissions": role.Permissions,
		}
		roles[i] = r
	}

	d.SetId(meta.name)

	if err := d.Set("roles", roles); err != nil {
		return diag.Errorf("error setting roles: %v", err)
	}

	return nil
}

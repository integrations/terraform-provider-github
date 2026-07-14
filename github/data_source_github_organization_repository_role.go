package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRepositoryRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationRepositoryRoleRead,

		Description: "Data source to lookup a custom organization repository role.",

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "ID of the organization repository role.",
				Type:        schema.TypeInt,
				Required:    true,
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
	}
}

func dataSourceGithubOrganizationRepositoryRoleRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	roleIdInt, _ := d.Get("role_id").(int)
	roleID := int64(roleIdInt)

	role, _, err := meta.v3client.Organizations.GetCustomRepoRole(ctx, meta.name, roleID)
	if err != nil {
		return diag.FromErr(err)
	}

	r := map[string]any{
		"role_id":     role.GetID(),
		"name":        role.GetName(),
		"description": role.GetDescription(),
		"base_role":   role.GetBaseRole(),
		"permissions": role.Permissions,
	}

	d.SetId(strconv.FormatInt(role.GetID(), 10))

	for k, v := range r {
		if err := d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationRoleRead,

		Description: "Data source to lookup a custom organization role.",

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "ID of the organization role.",
				Type:        schema.TypeInt,
				Required:    true,
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
	}
}

func dataSourceGithubOrganizationRoleRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	roleIdInt, _ := d.Get("role_id").(int)
	roleID := int64(roleIdInt)

	role, _, err := meta.v3client.Organizations.GetOrgRole(ctx, meta.name, roleID)
	if err != nil {
		return diag.FromErr(err)
	}

	r := map[string]any{
		"role_id":     int(role.GetID()),
		"name":        role.GetName(),
		"description": role.GetDescription(),
		"source":      role.GetSource(),
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

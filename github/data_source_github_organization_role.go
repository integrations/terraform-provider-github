package github

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRole() *schema.Resource {
	return &schema.Resource{
		Description: "Lookup a custom organization role.",

		ReadContext: dataSourceGithubOrganizationRoleRead,

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "The ID of the organization role.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"name": {
				Description: "The name of the organization role.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "The description of the organization role.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"source": {
				Description: "The source of this role; one of `Predefined`, `Organization`, or `Enterprise`.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"base_role": {
				Description: "The system role from which this role inherits permissions.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"permissions": {
				Description: "A list of permissions included in this role.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
		},
	}
}

func dataSourceGithubOrganizationRoleRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId := int64(d.Get("role_id").(int))

	role, _, err := client.Organizations.GetOrgRole(ctx, orgName, roleId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(role.GetID(), 10))

	if err := d.Set("role_id", int(role.GetID())); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", role.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", role.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("source", role.GetSource()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("base_role", role.GetBaseRole()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("permissions", role.Permissions); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

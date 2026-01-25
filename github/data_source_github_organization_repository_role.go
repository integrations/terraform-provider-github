package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubOrganizationRepositoryRole() *schema.Resource {
	return &schema.Resource{
		Description: "Lookup a custom organization repository role.",

		ReadContext: dataSourceGithubOrganizationRepositoryRoleRead,

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description: "The ID of the organization repository role.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"name": {
				Description: "The name of the organization repository role.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"description": {
				Description: "The description of the organization repository role.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"base_role": {
				Description: "The system role from which this role inherits permissions.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"permissions": {
				Description: "The permissions included in this role.",
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
			},
		},
	}
}

func dataSourceGithubOrganizationRepositoryRoleRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	roleId := int64(d.Get("role_id").(int))

	// TODO: Use this code when go-github is at v68+
	// role, _, err := client.Organizations.GetCustomRepoRole(ctx, orgName, roleId)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	roles, _, err := client.Organizations.ListCustomRepoRoles(ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	var role *github.CustomRepoRoles
	for _, r := range roles.CustomRepoRoles {
		if r.GetID() == roleId {
			role = r
			break
		}
	}
	if role == nil {
		return diag.FromErr(fmt.Errorf("custom organization repo role with ID %d not found", roleId))
	}

	d.SetId(strconv.FormatInt(role.GetID(), 10))

	if err := d.Set("role_id", role.GetID()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", role.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", role.GetDescription()); err != nil {
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

package github

import (
	"context"
	"strconv"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGithubOrganizationRoleUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubOrganizationRoleUsersRead,

		Description: "Data source to list all users assigned to a custom organization role.",

		Schema: map[string]*schema.Schema{
			"role_id": {
				Description:      "ID of the organization role.",
				Type:             schema.TypeInt,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntAtLeast(1)),
			},
			"users": {
				Description: "Users assigned to the organization role.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Description: "ID of the user.",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"login": {
							Description: "Login of the user.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"assignment": {
							Description: "Relationship a user has with a role; one of `direct`, `indirect`, or `mixed`.",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGithubOrganizationRoleUsersRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)

	if ok, diags := checkOrganizationOK(meta); !ok {
		return diags
	}

	roleIdInt, _ := d.Get("role_id").(int)
	roleID := int64(roleIdInt)

	users := make([]any, 0)
	for user, err := range meta.v3client.Organizations.ListUsersAssignedToOrgRoleIter(ctx, meta.name, roleID, &github.ListOptions{PerPage: meta.maxPerPage}) {
		if err != nil {
			return diag.FromErr(err)
		}

		u := map[string]any{
			"user_id":    int(user.GetID()),
			"login":      user.GetLogin(),
			"assignment": user.GetAssignment(),
		}

		users = append(users, u)
	}

	d.SetId(strconv.FormatInt(roleID, 10))

	if err := d.Set("users", users); err != nil {
		return diag.Errorf("error setting users: %v", err)
	}

	return nil
}

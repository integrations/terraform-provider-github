package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseSCIMUser() *schema.Resource {
	return &schema.Resource{
		Description: "Lookup SCIM provisioning information for a single GitHub enterprise user.",
		ReadContext: dataSourceGithubEnterpriseSCIMUserRead,

		Schema: map[string]*schema.Schema{
			"enterprise": {
				Description: "The enterprise slug.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"scim_user_id": {
				Description: "The SCIM user ID.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"excluded_attributes": {
				Description: "Optional SCIM excludedAttributes query parameter.",
				Type:        schema.TypeString,
				Optional:    true,
			},

			"schemas": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "SCIM schemas for this user.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SCIM user ID.",
			},
			"external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The external ID for the user.",
			},
			"user_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SCIM userName.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SCIM displayName.",
			},
			"active": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the user is active.",
			},
			"name": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User name object.",
				Elem:        &schema.Resource{Schema: enterpriseSCIMUserNameSchema()},
			},
			"emails": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User emails.",
				Elem:        &schema.Resource{Schema: enterpriseSCIMUserEmailSchema()},
			},
			"roles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User roles.",
				Elem:        &schema.Resource{Schema: enterpriseSCIMUserRoleSchema()},
			},
			"meta": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Resource metadata.",
				Elem:        &schema.Resource{Schema: enterpriseSCIMMetaSchema()},
			},
		},
	}
}

func dataSourceGithubEnterpriseSCIMUserRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise := d.Get("enterprise").(string)
	scimUserID := d.Get("scim_user_id").(string)
	excluded := d.Get("excluded_attributes").(string)

	path := fmt.Sprintf("scim/v2/enterprises/%s/Users/%s", enterprise, scimUserID)
	urlStr, err := enterpriseSCIMListURL(path, enterpriseSCIMListOptions{ExcludedAttributes: excluded})
	if err != nil {
		return diag.FromErr(err)
	}

	user := enterpriseSCIMUser{}
	_, err = enterpriseSCIMGet(ctx, client, urlStr, &user)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", enterprise, scimUserID))

	_ = d.Set("schemas", user.Schemas)
	_ = d.Set("id", user.ID)
	_ = d.Set("external_id", user.ExternalID)
	_ = d.Set("user_name", user.UserName)
	_ = d.Set("display_name", user.DisplayName)
	_ = d.Set("active", user.Active)
	_ = d.Set("name", flattenEnterpriseSCIMUserName(user.Name))
	_ = d.Set("emails", flattenEnterpriseSCIMUserEmails(user.Emails))
	_ = d.Set("roles", flattenEnterpriseSCIMUserRoles(user.Roles))
	_ = d.Set("meta", flattenEnterpriseSCIMMeta(user.Meta))

	return nil
}

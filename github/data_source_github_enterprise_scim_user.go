package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseSCIMUser() *schema.Resource {
	s := enterpriseSCIMUserSchema()
	s["enterprise"] = &schema.Schema{
		Description: "The enterprise slug.",
		Type:        schema.TypeString,
		Required:    true,
	}
	s["scim_user_id"] = &schema.Schema{
		Description: "The SCIM user ID.",
		Type:        schema.TypeString,
		Required:    true,
	}

	return &schema.Resource{
		Description: "Lookup SCIM provisioning information for a single GitHub enterprise user.",
		ReadContext: dataSourceGithubEnterpriseSCIMUserRead,
		Schema:      s,
	}
}

func dataSourceGithubEnterpriseSCIMUserRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise := d.Get("enterprise").(string)
	scimUserID := d.Get("scim_user_id").(string)

	user, _, err := client.Enterprise.GetProvisionedSCIMUser(ctx, enterprise, scimUserID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", enterprise, scimUserID))

	if err := d.Set("schemas", user.Schemas); err != nil {
		return diag.FromErr(err)
	}
	if user.ID != nil {
		if err := d.Set("id", *user.ID); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("external_id", user.ExternalID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("user_name", user.UserName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("display_name", user.DisplayName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("active", user.Active); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", flattenEnterpriseSCIMUserName(user.Name)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("emails", flattenEnterpriseSCIMUserEmails(user.Emails)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("roles", flattenEnterpriseSCIMUserRoles(user.Roles)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("meta", flattenEnterpriseSCIMMeta(user.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseSCIMGroup() *schema.Resource {
	s := enterpriseSCIMGroupSchema()
	s["enterprise"] = &schema.Schema{
		Description: "The enterprise slug.",
		Type:        schema.TypeString,
		Required:    true,
	}
	s["scim_group_id"] = &schema.Schema{
		Description: "The SCIM group ID.",
		Type:        schema.TypeString,
		Required:    true,
	}

	return &schema.Resource{
		Description: "Retrieves SCIM provisioning information for a single GitHub enterprise group.",
		ReadContext: dataSourceGithubEnterpriseSCIMGroupRead,
		Schema:      s,
	}
}

func dataSourceGithubEnterpriseSCIMGroupRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise := d.Get("enterprise").(string)
	scimGroupID := d.Get("scim_group_id").(string)

	group, _, err := client.Enterprise.GetProvisionedSCIMGroup(ctx, enterprise, scimGroupID, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", enterprise, scimGroupID))

	if err := d.Set("schemas", group.Schemas); err != nil {
		return diag.FromErr(err)
	}
	if group.ID != nil {
		if err := d.Set("id", *group.ID); err != nil {
			return diag.FromErr(err)
		}
	}
	if group.ExternalID != nil {
		if err := d.Set("external_id", *group.ExternalID); err != nil {
			return diag.FromErr(err)
		}
	}
	if group.DisplayName != nil {
		if err := d.Set("display_name", *group.DisplayName); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("members", flattenEnterpriseSCIMGroupMembers(group.Members)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("meta", flattenEnterpriseSCIMMeta(group.Meta)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

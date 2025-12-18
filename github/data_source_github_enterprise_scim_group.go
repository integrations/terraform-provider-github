package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubEnterpriseSCIMGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Lookup SCIM provisioning information for a single GitHub enterprise group.",
		ReadContext: dataSourceGithubEnterpriseSCIMGroupRead,

		Schema: map[string]*schema.Schema{
			"enterprise": {
				Description: "The enterprise slug.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"scim_group_id": {
				Description: "The SCIM group ID.",
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
				Description: "SCIM schemas for this group.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SCIM group ID.",
			},
			"external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The external ID for the group.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SCIM group displayName.",
			},
			"members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Group members.",
				Elem:        &schema.Resource{Schema: enterpriseSCIMGroupMemberSchema()},
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

func dataSourceGithubEnterpriseSCIMGroupRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterprise := d.Get("enterprise").(string)
	scimGroupID := d.Get("scim_group_id").(string)
	excluded := d.Get("excluded_attributes").(string)

	path := fmt.Sprintf("scim/v2/enterprises/%s/Groups/%s", enterprise, scimGroupID)
	urlStr, err := enterpriseSCIMListURL(path, enterpriseSCIMListOptions{ExcludedAttributes: excluded})
	if err != nil {
		return diag.FromErr(err)
	}

	group := enterpriseSCIMGroup{}
	_, err = enterpriseSCIMGet(ctx, client, urlStr, &group)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", enterprise, scimGroupID))

	_ = d.Set("schemas", group.Schemas)
	_ = d.Set("id", group.ID)
	_ = d.Set("external_id", group.ExternalID)
	_ = d.Set("display_name", group.DisplayName)
	_ = d.Set("members", flattenEnterpriseSCIMGroupMembers(group.Members))
	_ = d.Set("meta", flattenEnterpriseSCIMMeta(group.Meta))

	return nil
}

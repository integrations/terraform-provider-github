package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationOIDCCustomPropertyInclusions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsOrganizationOIDCCustomPropertyInclusionsRead,

		Schema: map[string]*schema.Schema{
			"custom_property_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of custom property names included in the OIDC token.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceGithubActionsOrganizationOIDCCustomPropertyInclusionsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	inclusions, err := listOrgOIDCCustomPropertyInclusions(ctx, client, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	propertyNames := make([]string, len(inclusions))
	for i, inclusion := range inclusions {
		propertyNames[i] = inclusion.PropertyName
	}

	d.SetId(orgName)
	if err := d.Set("custom_property_names", propertyNames); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

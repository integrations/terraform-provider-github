package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateRead,

		Schema: map[string]*schema.Schema{
			"include_claim_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	orgName := meta.name

	err := checkOrganization(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	template, _, err := client.Actions.GetOrgOIDCSubjectClaimCustomTemplate(ctx, orgName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(orgName)
	if err = d.Set("include_claim_keys", template.IncludeClaimKeys); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

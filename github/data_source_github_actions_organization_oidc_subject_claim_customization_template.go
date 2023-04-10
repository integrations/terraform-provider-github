package github

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateRead,

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

func dataSourceGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*Owner).v3client
	orgName := meta.(*Owner).name
	ctx := meta.(*Owner).StopContext

	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	template, _, err := client.Actions.GetOrgOIDCSubjectClaimCustomTemplate(ctx, orgName)

	if err != nil {
		return err
	}

	d.SetId(orgName)
	d.Set("include_claim_keys", template.IncludeClaimKeys)

	return nil
}

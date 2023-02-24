package github

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"use_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
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

func dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	repository := d.Get("name").(string)
	owner := meta.(*Owner).name
	ctx := meta.(*Owner).StopContext

	template, _, err := client.Actions.GetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repository)

	if err != nil {
		return err
	}

	d.SetId(repository)
	d.Set("use_default", template.UseDefault)
	d.Set("include_claim_keys", template.IncludeClaimKeys)

	return nil
}

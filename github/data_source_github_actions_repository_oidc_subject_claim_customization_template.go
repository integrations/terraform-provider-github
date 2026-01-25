package github

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to retrieve the OpenID Connect subject claim customization template for a repository.",
		Read:        dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository.",
			},
			"use_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the repository uses the default template.",
			},
			"include_claim_keys": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of OpenID Connect claim keys.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client

	repository := d.Get("name").(string)
	owner := meta.(*Owner).name
	ctx := meta.(*Owner).StopContext

	template, _, err := client.Actions.GetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repository)
	if err != nil {
		return err
	}

	d.SetId(repository)
	err = d.Set("use_default", template.UseDefault)
	if err != nil {
		return err
	}
	err = d.Set("include_claim_keys", template.IncludeClaimKeys)
	if err != nil {
		return err
	}

	return nil
}

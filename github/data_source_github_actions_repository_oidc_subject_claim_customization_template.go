package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead,

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

func dataSourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	repository := d.Get("name").(string)
	owner := meta.(*Owner).name

	template, _, err := client.Actions.GetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repository)
	err = d.Set("use_default", template.UseDefault)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("include_claim_keys", template.IncludeClaimKeys)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

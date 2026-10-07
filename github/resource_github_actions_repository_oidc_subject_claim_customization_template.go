package github

import (
	"context"
	"errors"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateCreateOrUpdate,
		ReadContext:   resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead,
		UpdateContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateCreateOrUpdate,
		DeleteContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The name of the repository.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
			},
			"use_default": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether to use the default template or not. If 'true', 'include_claim_keys' must not be set.",
			},
			"include_claim_keys": {
				Type:        schema.TypeList,
				Optional:    true,
				MinItems:    1,
				Description: "A list of OpenID Connect claims.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name

	useDefault := d.Get("use_default").(bool)
	includeClaimKeys, hasClaimKeys := d.GetOk("include_claim_keys")

	if useDefault && hasClaimKeys {
		return diag.FromErr(errors.New("include_claim_keys cannot be set when use_default is true"))
	}

	customOIDCSubjectClaimTemplate := github.OIDCSubjectClaimCustomTemplate{
		UseDefault: &useDefault,
	}

	if includeClaimKeys != nil {

		includeClaimKeysVal := includeClaimKeys.([]any)

		claimsStr := make([]string, len(includeClaimKeysVal))

		for i, v := range includeClaimKeysVal {
			claimsStr[i] = v.(string)
		}

		customOIDCSubjectClaimTemplate.IncludeClaimKeys = claimsStr
	}

	_, err := client.Actions.SetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repository, customOIDCSubjectClaimTemplate)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repository)
	return resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead(ctx, d, meta)
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	repository := d.Id()
	owner := meta.(*Owner).name

	template, _, err := client.Actions.GetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repository)
	if err != nil {
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(ctx, err, d, "actions repository oidc subject claim customization template", map[string]any{"owner": owner, "repository": repository}))
	}

	if err = d.Set("repository", repository); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("use_default", template.UseDefault); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("include_claim_keys", template.IncludeClaimKeys); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	// Reset the repository to use the default claims
	// https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#using-the-default-subject-claims
	client := meta.(*Owner).v3client

	repository := d.Get("repository").(string)
	owner := meta.(*Owner).name

	customOIDCSubjectClaimTemplate := github.OIDCSubjectClaimCustomTemplate{
		UseDefault: new(true),
	}

	_, err := client.Actions.SetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repository, customOIDCSubjectClaimTemplate)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

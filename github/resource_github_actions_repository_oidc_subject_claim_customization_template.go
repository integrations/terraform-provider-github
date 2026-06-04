package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateStateUpgradeV0,
				Version: 0,
			},
		},

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The name of the repository.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
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

		CustomizeDiff: customdiff.All(
			diffRepository,
		),

		CreateContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateCreateOrUpdate,
		ReadContext:   resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead,
		UpdateContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateCreateOrUpdate,
		DeleteContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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
		return diag.Errorf("include_claim_keys cannot be set when use_default is true")
	}

	customOIDCSubjectClaimTemplate := &github.OIDCSubjectClaimCustomTemplate{
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

	repo, _, err := client.Repositories.Get(ctx, owner, repository)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repository)
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead(ctx, d, meta)
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	ctx = tflog.SetField(ctx, "id", d.Id())

	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	repoName := d.Get("repository").(string)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Repository not found, removing from state.", map[string]any{"repository": repoName})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	template, _, err := client.Actions.GetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(deleteResourceOn404AndSwallow304OtherwiseReturnError(err, d, "actions repository oidc subject claim customization template (%s, %s)", owner, repoName))
	}

	if err = d.Set("repository", repo.GetName()); err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("repository_id", int(repo.GetID())); err != nil {
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

	useDefault := true
	customOIDCSubjectClaimTemplate := &github.OIDCSubjectClaimCustomTemplate{
		UseDefault: &useDefault,
	}

	_, err := client.Actions.SetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repository, customOIDCSubjectClaimTemplate)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

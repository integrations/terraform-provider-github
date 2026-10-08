package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateCreate,
		ReadContext:   resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead,
		UpdateContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateUpdate,
		DeleteContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateImport,
		},

		CustomizeDiff: diffRepository,

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateResourceV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateInstanceStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Creates and manages an OpenID Connect subject claim customization template for a repository.",

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
				Description: "The ID of the repository.",
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

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)

	template, err := expandRepositoryOIDCSubjectClaimCustomTemplate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Creating repository OIDC subject claim customization template", map[string]any{"repository": repoName})

	if _, err := client.Actions.SetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repoName, template); err != nil {
		return diag.FromErr(err)
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repoName)

	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)

	template, _, err := client.Actions.GetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repoName)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing repository OIDC subject claim customization template from state because it no longer exists in GitHub", map[string]any{"repository": repoName})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("use_default", template.UseDefault); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("include_claim_keys", template.IncludeClaimKeys); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)

	template, err := expandRepositoryOIDCSubjectClaimCustomTemplate(d)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, "Updating repository OIDC subject claim customization template", map[string]any{"repository": repoName})

	if _, err := client.Actions.SetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repoName, template); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(repoName)

	return nil
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Get("repository").(string)

	tflog.Debug(ctx, "Resetting repository OIDC subject claim customization template to default", map[string]any{"repository": repoName})

	// Delete resets the repository to the default claims.
	// https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/about-security-hardening-with-openid-connect#using-the-default-subject-claims
	template := github.OIDCSubjectClaimCustomTemplate{
		UseDefault: new(true),
	}

	if _, err := client.Actions.SetRepoOIDCSubjectClaimCustomTemplate(ctx, owner, repoName, template); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Repository no longer exists, nothing to reset", map[string]any{"repository": repoName})
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName := d.Id()

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}

	if err := d.Set("repository", repo.GetName()); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", int(repo.GetID())); err != nil {
		return nil, err
	}
	d.SetId(repo.GetName())

	return []*schema.ResourceData{d}, nil
}

func expandRepositoryOIDCSubjectClaimCustomTemplate(d *schema.ResourceData) (github.OIDCSubjectClaimCustomTemplate, error) {
	useDefault := d.Get("use_default").(bool)
	includeClaimKeys, hasClaimKeys := d.GetOk("include_claim_keys")

	if useDefault && hasClaimKeys {
		return github.OIDCSubjectClaimCustomTemplate{}, errors.New("include_claim_keys cannot be set when use_default is true")
	}

	template := github.OIDCSubjectClaimCustomTemplate{
		UseDefault: new(useDefault),
	}

	if hasClaimKeys {
		template.IncludeClaimKeys = expandStringList(includeClaimKeys.([]any))
	}

	return template, nil
}

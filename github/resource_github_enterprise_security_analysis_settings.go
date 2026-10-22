package github

import (
	"context"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseSecurityAnalysisSettings() *schema.Resource {
	return &schema.Resource{
		Description:   "GitHub Enterprise Security Analysis Settings management.",
		CreateContext: resourceGithubEnterpriseSecurityAnalysisSettingsCreateOrUpdate,
		ReadContext:   resourceGithubEnterpriseSecurityAnalysisSettingsRead,
		UpdateContext: resourceGithubEnterpriseSecurityAnalysisSettingsCreateOrUpdate,
		DeleteContext: resourceGithubEnterpriseSecurityAnalysisSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"advanced_security_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether GitHub Advanced Security is automatically enabled for new repositories.",
			},
			"secret_scanning_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether secret scanning is automatically enabled for new repositories.",
			},
			"secret_scanning_push_protection_enabled_for_new_repositories": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether secret scanning push protection is automatically enabled for new repositories.",
			},
			"secret_scanning_push_protection_custom_link": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Custom URL for secret scanning push protection bypass instructions.",
			},
			"secret_scanning_validity_checks_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether secret scanning validity checks are enabled.",
			},
		},
	}
}

func resourceGithubEnterpriseSecurityAnalysisSettingsCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	d.SetId(enterpriseSlug)

	settings := &github.EnterpriseSecurityAnalysisSettings{}

	if v, ok := d.GetOk("advanced_security_enabled_for_new_repositories"); ok {
		settings.AdvancedSecurityEnabledForNewRepositories = new(v.(bool))
	}

	if v, ok := d.GetOk("secret_scanning_enabled_for_new_repositories"); ok {
		settings.SecretScanningEnabledForNewRepositories = new(v.(bool))
	}

	if v, ok := d.GetOk("secret_scanning_push_protection_enabled_for_new_repositories"); ok {
		settings.SecretScanningPushProtectionEnabledForNewRepositories = new(v.(bool))
	}

	if v, ok := d.GetOk("secret_scanning_push_protection_custom_link"); ok {
		settings.SecretScanningPushProtectionCustomLink = new(v.(string))
	}

	if v, ok := d.GetOk("secret_scanning_validity_checks_enabled"); ok {
		settings.SecretScanningValidityChecksEnabled = new(v.(bool))
	}
	tflog.Debug(ctx, "Updating security analysis settings for enterprise", map[string]any{"enterprise": enterpriseSlug})
	_, err := client.Enterprise.UpdateCodeSecurityAndAnalysis(ctx, enterpriseSlug, settings) //nolint:staticcheck // SA1019: UpdateCodeSecurityAndAnalysis is deprecated but still needed for legacy compatibility
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubEnterpriseSecurityAnalysisSettingsRead(ctx, d, meta)
}

func resourceGithubEnterpriseSecurityAnalysisSettingsRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Id()
	tflog.Debug(ctx, "Reading security analysis settings for enterprise", map[string]any{"enterprise": enterpriseSlug})

	settings, _, err := client.Enterprise.GetCodeSecurityAndAnalysis(ctx, enterpriseSlug) //nolint:staticcheck // SA1019: GetCodeSecurityAndAnalysis is deprecated but still needed for legacy compatibility
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("advanced_security_enabled_for_new_repositories", settings.AdvancedSecurityEnabledForNewRepositories); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret_scanning_enabled_for_new_repositories", settings.SecretScanningEnabledForNewRepositories); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret_scanning_push_protection_enabled_for_new_repositories", settings.SecretScanningPushProtectionEnabledForNewRepositories); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret_scanning_push_protection_custom_link", settings.SecretScanningPushProtectionCustomLink); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("secret_scanning_validity_checks_enabled", settings.SecretScanningValidityChecksEnabled); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseSecurityAnalysisSettingsDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Id()
	tflog.Debug(ctx, "Resetting security analysis settings to defaults for enterprise", map[string]any{"enterprise": enterpriseSlug})

	// Reset to safe defaults (all disabled)
	settings := &github.EnterpriseSecurityAnalysisSettings{
		AdvancedSecurityEnabledForNewRepositories:             new(false),
		SecretScanningEnabledForNewRepositories:               new(false),
		SecretScanningPushProtectionEnabledForNewRepositories: new(false),
		SecretScanningPushProtectionCustomLink:                new(""),
		SecretScanningValidityChecksEnabled:                   new(false),
	}

	_, err := client.Enterprise.UpdateCodeSecurityAndAnalysis(ctx, enterpriseSlug, settings) //nolint:staticcheck // SA1019: UpdateCodeSecurityAndAnalysis is deprecated but still needed for legacy compatibility
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

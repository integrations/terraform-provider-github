package github

import (
	"context"
	"log"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGithubEnterpriseSecurityAnalysisSettings() *schema.Resource {
	return &schema.Resource{
		Description: "GitHub Enterprise Security Analysis Settings management.",
		Create:      resourceGithubEnterpriseSecurityAnalysisSettingsCreateOrUpdate,
		Read:        resourceGithubEnterpriseSecurityAnalysisSettingsRead,
		Update:      resourceGithubEnterpriseSecurityAnalysisSettingsCreateOrUpdate,
		Delete:      resourceGithubEnterpriseSecurityAnalysisSettingsDelete,
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

func resourceGithubEnterpriseSecurityAnalysisSettingsCreateOrUpdate(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Get("enterprise_slug").(string)
	d.SetId(enterpriseSlug)

	settings := &github.EnterpriseSecurityAnalysisSettings{}

	if v, ok := d.GetOk("advanced_security_enabled_for_new_repositories"); ok {
		settings.AdvancedSecurityEnabledForNewRepositories = github.Bool(v.(bool))
	}

	if v, ok := d.GetOk("secret_scanning_enabled_for_new_repositories"); ok {
		settings.SecretScanningEnabledForNewRepositories = github.Bool(v.(bool))
	}

	if v, ok := d.GetOk("secret_scanning_push_protection_enabled_for_new_repositories"); ok {
		settings.SecretScanningPushProtectionEnabledForNewRepositories = github.Bool(v.(bool))
	}

	if v, ok := d.GetOk("secret_scanning_push_protection_custom_link"); ok {
		settings.SecretScanningPushProtectionCustomLink = github.String(v.(string))
	}

	if v, ok := d.GetOk("secret_scanning_validity_checks_enabled"); ok {
		settings.SecretScanningValidityChecksEnabled = github.Bool(v.(bool))
	}

	log.Printf("[DEBUG] Updating security analysis settings for enterprise: %s", enterpriseSlug)
	_, err := client.Enterprise.UpdateCodeSecurityAndAnalysis(ctx, enterpriseSlug, settings)
	if err != nil {
		return err
	}

	return resourceGithubEnterpriseSecurityAnalysisSettingsRead(d, meta)
}

func resourceGithubEnterpriseSecurityAnalysisSettingsRead(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Id()
	log.Printf("[DEBUG] Reading security analysis settings for enterprise: %s", enterpriseSlug)

	settings, _, err := client.Enterprise.GetCodeSecurityAndAnalysis(ctx, enterpriseSlug)
	if err != nil {
		return err
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return err
	}
	if err := d.Set("advanced_security_enabled_for_new_repositories", settings.AdvancedSecurityEnabledForNewRepositories); err != nil {
		return err
	}
	if err := d.Set("secret_scanning_enabled_for_new_repositories", settings.SecretScanningEnabledForNewRepositories); err != nil {
		return err
	}
	if err := d.Set("secret_scanning_push_protection_enabled_for_new_repositories", settings.SecretScanningPushProtectionEnabledForNewRepositories); err != nil {
		return err
	}
	if err := d.Set("secret_scanning_push_protection_custom_link", settings.SecretScanningPushProtectionCustomLink); err != nil {
		return err
	}
	if err := d.Set("secret_scanning_validity_checks_enabled", settings.SecretScanningValidityChecksEnabled); err != nil {
		return err
	}

	return nil
}

func resourceGithubEnterpriseSecurityAnalysisSettingsDelete(d *schema.ResourceData, meta any) error {
	client := meta.(*Owner).v3client
	ctx := context.Background()

	enterpriseSlug := d.Id()
	log.Printf("[DEBUG] Resetting security analysis settings to defaults for enterprise: %s", enterpriseSlug)

	// Reset to safe defaults (all disabled)
	settings := &github.EnterpriseSecurityAnalysisSettings{
		AdvancedSecurityEnabledForNewRepositories:             github.Bool(false),
		SecretScanningEnabledForNewRepositories:               github.Bool(false),
		SecretScanningPushProtectionEnabledForNewRepositories: github.Bool(false),
		SecretScanningPushProtectionCustomLink:                github.String(""),
		SecretScanningValidityChecksEnabled:                   github.Bool(false),
	}

	_, err := client.Enterprise.UpdateCodeSecurityAndAnalysis(ctx, enterpriseSlug, settings)
	if err != nil {
		return err
	}

	return nil
}

package github

import (
	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// flattenDependencyGraphAutosubmitActionOptions converts DependencyGraphAutosubmitActionOptions to a Terraform-compatible format
func flattenDependencyGraphAutosubmitActionOptions(options *github.DependencyGraphAutosubmitActionOptions) []any {
	if options == nil {
		return []any{}
	}
	autosubmitOpts := make(map[string]any)
	if options.LabeledRunners != nil {
		autosubmitOpts["labeled_runners"] = options.GetLabeledRunners()
	}
	return []any{autosubmitOpts}
}

// flattenCodeScanningDefaultSetupOptions converts CodeScanningDefaultSetupOptions to a Terraform-compatible format
func flattenCodeScanningDefaultSetupOptions(options *github.CodeScanningDefaultSetupOptions) []any {
	if options == nil {
		return []any{}
	}
	setupOpts := make(map[string]any)
	if options.RunnerType != "" {
		setupOpts["runner_type"] = options.RunnerType
	}
	if options.RunnerLabel != nil {
		setupOpts["runner_label"] = options.GetRunnerLabel()
	}
	return []any{setupOpts}
}

// flattenCodeScanningOptions converts CodeScanningOptions to a Terraform-compatible format
func flattenCodeScanningOptions(options *github.CodeScanningOptions) []any {
	if options == nil {
		return []any{}
	}
	scanOpts := make(map[string]any)
	if options.AllowAdvanced != nil {
		scanOpts["allow_advanced"] = options.GetAllowAdvanced()
	}
	return []any{scanOpts}
}

// flattenSecretScanningDelegatedBypassOptions converts SecretScanningDelegatedBypassOptions to a Terraform-compatible format
func flattenSecretScanningDelegatedBypassOptions(options *github.SecretScanningDelegatedBypassOptions) []any {
	if options == nil {
		return []any{}
	}
	bypassOpts := make(map[string]any)
	if options.Reviewers != nil {
		reviewers := make([]any, 0, len(options.Reviewers))
		for _, reviewer := range options.Reviewers {
			reviewerMap := make(map[string]any)
			reviewerMap["reviewer_id"] = reviewer.ReviewerID
			reviewerMap["reviewer_type"] = reviewer.ReviewerType
			reviewers = append(reviewers, reviewerMap)
		}
		bypassOpts["reviewers"] = reviewers
	}
	return []any{bypassOpts}
}

// expandCodeSecurityConfigurationCommon builds a CodeSecurityConfiguration from Terraform resource data.
// Used by both the organization and enterprise security configuration resources.
func expandCodeSecurityConfigurationCommon(d *schema.ResourceData) github.CodeSecurityConfiguration {
	config := github.CodeSecurityConfiguration{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	if val, ok := d.GetOk("advanced_security"); ok {
		config.AdvancedSecurity = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("dependency_graph"); ok {
		config.DependencyGraph = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("dependency_graph_autosubmit_action"); ok {
		config.DependencyGraphAutosubmitAction = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("dependabot_alerts"); ok {
		config.DependabotAlerts = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("dependabot_security_updates"); ok {
		config.DependabotSecurityUpdates = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("code_scanning_default_setup"); ok {
		config.CodeScanningDefaultSetup = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("code_scanning_delegated_alert_dismissal"); ok {
		config.CodeScanningDelegatedAlertDismissal = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("code_security"); ok {
		config.CodeSecurity = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning"); ok {
		config.SecretScanning = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_push_protection"); ok {
		config.SecretScanningPushProtection = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_delegated_bypass"); ok {
		config.SecretScanningDelegatedBypass = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_validity_checks"); ok {
		config.SecretScanningValidityChecks = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_non_provider_patterns"); ok {
		config.SecretScanningNonProviderPatterns = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_generic_secrets"); ok {
		config.SecretScanningGenericSecrets = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("secret_scanning_delegated_alert_dismissal"); ok {
		config.SecretScanningDelegatedAlertDismissal = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("secret_protection"); ok {
		config.SecretProtection = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("private_vulnerability_reporting"); ok {
		config.PrivateVulnerabilityReporting = github.Ptr(val.(string))
	}
	if val, ok := d.GetOk("enforcement"); ok {
		config.Enforcement = github.Ptr(val.(string))
	}

	if val, ok := d.GetOk("dependency_graph_autosubmit_action_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			autosubmitOpts := optionsList[0].(map[string]any)
			config.DependencyGraphAutosubmitActionOptions = &github.DependencyGraphAutosubmitActionOptions{
				LabeledRunners: github.Ptr(autosubmitOpts["labeled_runners"].(bool)),
			}
		}
	}

	if val, ok := d.GetOk("code_scanning_default_setup_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			setupOpts := optionsList[0].(map[string]any)
			config.CodeScanningDefaultSetupOptions = &github.CodeScanningDefaultSetupOptions{
				RunnerType: setupOpts["runner_type"].(string),
			}
			if runnerLabel, ok := setupOpts["runner_label"].(string); ok && runnerLabel != "" {
				config.CodeScanningDefaultSetupOptions.RunnerLabel = github.Ptr(runnerLabel)
			}
		}
	}

	if val, ok := d.GetOk("code_scanning_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			scanOpts := optionsList[0].(map[string]any)
			config.CodeScanningOptions = &github.CodeScanningOptions{
				AllowAdvanced: github.Ptr(scanOpts["allow_advanced"].(bool)),
			}
		}
	}

	if val, ok := d.GetOk("secret_scanning_delegated_bypass_options"); ok {
		optionsList := val.([]any)
		if len(optionsList) > 0 {
			bypassOpts := optionsList[0].(map[string]any)
			options := &github.SecretScanningDelegatedBypassOptions{}
			if reviewersVal, ok := bypassOpts["reviewers"]; ok {
				reviewersList := reviewersVal.([]any)
				reviewers := make([]*github.BypassReviewer, 0, len(reviewersList))
				for _, reviewerRaw := range reviewersList {
					reviewerMap := reviewerRaw.(map[string]any)
					reviewers = append(reviewers, &github.BypassReviewer{
						ReviewerID:   int64(reviewerMap["reviewer_id"].(int)),
						ReviewerType: reviewerMap["reviewer_type"].(string),
					})
				}
				options.Reviewers = reviewers
			}
			config.SecretScanningDelegatedBypassOptions = options
		}
	}

	return config
}

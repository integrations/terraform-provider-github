package github

import (
	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// expandOptionalString returns a pointer to the configured string at key, or nil when the value is
// unset. GetOk treats an empty string as unset, matching the Optional (non-computed) attributes.
func expandOptionalString(d *schema.ResourceData, key string) *string {
	v, ok := d.GetOk(key)
	if !ok {
		return nil
	}
	s, _ := v.(string)
	return new(s)
}

// flattenDependencyGraphAutosubmitActionOptions converts DependencyGraphAutosubmitActionOptions to a Terraform-compatible format.
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

// flattenCodeScanningDefaultSetupOptions converts CodeScanningDefaultSetupOptions to a Terraform-compatible format.
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

// flattenCodeScanningOptions converts CodeScanningOptions to a Terraform-compatible format.
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

// flattenSecretScanningDelegatedBypassOptions converts SecretScanningDelegatedBypassOptions to a Terraform-compatible format.
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

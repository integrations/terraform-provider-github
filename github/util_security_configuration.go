package github

import (
	"github.com/google/go-github/v82/github"
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
	setupOpts["runner_type"] = options.RunnerType
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

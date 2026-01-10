package github

import (
	"context"
	"fmt"
	"slices"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// branchTagOnlyRules contains rules that are only valid for branch and tag targets.
//
// These rules apply to ref-based operations (branches and tags) and are not supported
// for push rulesets which operate on file content.
//
// To verify/maintain this list:
//  1. Check the GitHub API documentation for organization rulesets:
//     https://docs.github.com/en/rest/orgs/rules?apiVersion=2022-11-28#create-an-organization-repository-ruleset
//  2. The API docs don't clearly separate push vs branch/tag rules. To verify,
//     attempt to create a push ruleset via API or UI with each rule type.
//     Push rulesets will reject branch/tag rules with "Invalid rule '<name>'" error.
//  3. Generally, push rules deal with file content (paths, sizes, extensions),
//     while branch/tag rules deal with ref lifecycle and merge requirements.
var branchTagOnlyRules = []string{
	"creation",
	"update",
	"deletion",
	"required_linear_history",
	"required_signatures",
	"pull_request",
	"required_status_checks",
	"non_fast_forward",
	"commit_message_pattern",
	"commit_author_email_pattern",
	"committer_email_pattern",
	"branch_name_pattern",
	"tag_name_pattern",
	"required_workflows",
	"required_code_scanning",
	"required_deployments",
	"merge_queue",
}

// pushOnlyRules contains rules that are only valid for push targets.
//
// These rules apply to push operations and control what content can be pushed
// to repositories. They are not supported for branch or tag rulesets.
//
// To verify/maintain this list:
//  1. Check the GitHub API documentation for organization rulesets:
//     https://docs.github.com/en/rest/orgs/rules?apiVersion=2022-11-28#create-an-organization-repository-ruleset
//  2. The API docs don't clearly separate push vs branch/tag rules. To verify,
//     attempt to create a branch ruleset via API or UI with each rule type.
//     Branch rulesets will reject push-only rules with an error.
//  3. Push rules control file content: paths, sizes, extensions, path lengths.
var pushOnlyRules = []string{
	"file_path_restriction",
	"max_file_path_length",
	"file_extension_restriction",
	"max_file_size",
}

func validateRulesForTarget(ctx context.Context, d *schema.ResourceDiff) error {
	target := d.Get("target").(string)
	tflog.Debug(ctx, "Validating rules for target", map[string]any{"target": target})

	switch target {
	case "push":
		return validateRulesForPushTarget(ctx, d)
	case "branch", "tag":
		return validateRulesForBranchTagTarget(ctx, d)
	}

	tflog.Debug(ctx, "Rules validation passed", map[string]any{"target": target})
	return nil
}

func validateRulesForPushTarget(ctx context.Context, d *schema.ResourceDiff) error {
	return validateRules(ctx, d, pushOnlyRules)
}

func validateRulesForBranchTagTarget(ctx context.Context, d *schema.ResourceDiff) error {
	return validateRules(ctx, d, branchTagOnlyRules)
}

func validateRules(ctx context.Context, d *schema.ResourceDiff, allowedRules []string) error {
	target := d.Get("target").(string)
	rules := d.Get("rules").([]any)[0].(map[string]any)
	for ruleName := range rules {
		ruleValue, exists := d.GetOk(fmt.Sprintf("rules.0.%s", ruleName))
		if !exists {
			continue
		}
		switch ruleValue := ruleValue.(type) {
		case []any:
			if len(ruleValue) == 0 {
				continue
			}
		case map[string]any:
			if len(ruleValue) == 0 {
				continue
			}
		case any:
			if ruleValue == nil {
				continue
			}
		}
		if slices.Contains(allowedRules, ruleName) {
			continue
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Invalid rule for %s target", target), map[string]any{"rule": ruleName, "value": ruleValue})
			return fmt.Errorf("rule %q is not valid for %[2]s target; %[2]s targets only support: %v", ruleName, target, allowedRules)
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Rules validation passed for %s target", target))
	return nil
}

func validateRepositoryRulesetConditionsFieldForBranchAndTagTargets(ctx context.Context, target string, conditions map[string]any) error {
	tflog.Debug(ctx, fmt.Sprintf("Validating conditions field for %s target", target), map[string]any{"target": target, "conditions": conditions})

	if conditions["ref_name"] == nil || len(conditions["ref_name"].([]any)) == 0 {
		tflog.Debug(ctx, fmt.Sprintf("Missing ref_name for %s target", target), map[string]any{"target": target})
		return fmt.Errorf("ref_name must be set for %s target", target)
	}

	tflog.Debug(ctx, fmt.Sprintf("Conditions validation passed for %s target", target))
	return nil
}

func validateConditionsFieldForBranchAndTagTargets(ctx context.Context, target string, conditions map[string]any) error {
	tflog.Debug(ctx, fmt.Sprintf("Validating conditions field for %s target", target), map[string]any{"target": target, "conditions": conditions})

	if conditions["ref_name"] == nil || len(conditions["ref_name"].([]any)) == 0 {
		tflog.Debug(ctx, fmt.Sprintf("Missing ref_name for %s target", target), map[string]any{"target": target})
		return fmt.Errorf("ref_name must be set for %s target", target)
	}

	if (conditions["repository_name"] == nil || len(conditions["repository_name"].([]any)) == 0) && (conditions["repository_id"] == nil || len(conditions["repository_id"].([]any)) == 0) {
		tflog.Debug(ctx, fmt.Sprintf("Missing repository_name or repository_id for %s target", target), map[string]any{"target": target})
		return fmt.Errorf("either repository_name or repository_id must be set for %s target", target)
	}
	tflog.Debug(ctx, fmt.Sprintf("Conditions validation passed for %s target", target))
	return nil
}

func validateConditionsFieldForPushTarget(ctx context.Context, conditions map[string]any) error {
	tflog.Debug(ctx, "Validating conditions field for push target", map[string]any{"target": "push", "conditions": conditions})

	if conditions["ref_name"] != nil && len(conditions["ref_name"].([]any)) > 0 {
		tflog.Debug(ctx, "Invalid ref_name for push target", map[string]any{"ref_name": conditions["ref_name"]})
		return fmt.Errorf("ref_name must not be set for push target")
	}
	tflog.Debug(ctx, "Conditions validation passed for push target")
	return nil
}

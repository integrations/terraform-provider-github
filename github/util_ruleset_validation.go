package github

import (
	"context"
	"fmt"
	"slices"

	"github.com/google/go-github/v83/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var operatorValidation = validation.ToDiagFunc(validation.StringInSlice([]string{"starts_with", "ends_with", "contains", "regex"}, false))

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
var branchTagOnlyRules = []github.RepositoryRuleType{
	github.RulesetRuleTypeCreation,
	github.RulesetRuleTypeUpdate,
	github.RulesetRuleTypeDeletion,
	github.RulesetRuleTypeRequiredLinearHistory,
	github.RulesetRuleTypeRequiredSignatures,
	github.RulesetRuleTypePullRequest,
	github.RulesetRuleTypeRequiredStatusChecks,
	github.RulesetRuleTypeNonFastForward,
	github.RulesetRuleTypeCommitMessagePattern,
	github.RulesetRuleTypeCommitAuthorEmailPattern,
	github.RulesetRuleTypeCommitterEmailPattern,
	github.RulesetRuleTypeBranchNamePattern,
	github.RulesetRuleTypeTagNamePattern,
	github.RulesetRuleTypeWorkflows,
	github.RulesetRuleTypeCodeScanning,
	github.RulesetRuleTypeRequiredDeployments,
	github.RulesetRuleTypeMergeQueue,
	github.RulesetRuleTypeCopilotCodeReview,
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
var pushOnlyRules = []github.RepositoryRuleType{
	github.RulesetRuleTypeFilePathRestriction,
	github.RulesetRuleTypeMaxFilePathLength,
	github.RulesetRuleTypeFileExtensionRestriction,
	github.RulesetRuleTypeMaxFileSize,
}

func validateRulesForTarget(ctx context.Context, d *schema.ResourceDiff) error {
	target := github.RulesetTarget(d.Get("target").(string))
	tflog.Debug(ctx, "Validating rules for target", map[string]any{"target": target})

	switch target {
	case github.RulesetTargetPush:
		return validateRulesForPushTarget(ctx, d)
	case github.RulesetTargetBranch, github.RulesetTargetTag:
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

func validateRules(ctx context.Context, d *schema.ResourceDiff, allowedRules []github.RepositoryRuleType) error {
	target := github.RulesetTarget(d.Get("target").(string))
	rules := d.Get("rules").([]any)[0].(map[string]any)
	for ruleName := range rules {
		ruleValue, exists := d.GetOk(fmt.Sprintf("rules.0.%s", ruleName))
		if !exists {
			continue
		}
		// These are the few rules which are not mapped to the same name in the API.
		switch ruleName {
		case "required_code_scanning":
			ruleName = string(github.RulesetRuleTypeCodeScanning)
		case "required_workflows":
			ruleName = string(github.RulesetRuleTypeWorkflows)
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
		if slices.Contains(allowedRules, github.RepositoryRuleType(ruleName)) {
			continue
		} else {
			tflog.Debug(ctx, fmt.Sprintf("Invalid rule for %s target", target), map[string]any{"rule": ruleName, "value": ruleValue})
			return fmt.Errorf("rule %q is not valid for %[2]s target; %[2]s targets only support: %v", ruleName, target, allowedRules)
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Rules validation passed for %s target", target))
	return nil
}

func validateRulesetConditions(ctx context.Context, d *schema.ResourceDiff, isOrg bool) error {
	target := github.RulesetTarget(d.Get("target").(string))
	tflog.Debug(ctx, "Validating conditions field based on target", map[string]any{"target": target})
	conditionsRaw := d.Get("conditions").([]any)

	if len(conditionsRaw) == 0 {
		tflog.Debug(ctx, "An empty conditions block, skipping validation.", map[string]any{"target": target})
		return nil
	}

	conditions := conditionsRaw[0].(map[string]any)

	switch target {
	case github.RulesetTargetBranch, github.RulesetTargetTag:
		return validateConditionsFieldForBranchAndTagTargets(ctx, target, conditions, isOrg)
	case github.RulesetTargetPush:
		return validateConditionsFieldForPushTarget(ctx, conditions)
	}
	return nil
}

func validateRulesetRules(ctx context.Context, d *schema.ResourceDiff) error {
	target := github.RulesetTarget(d.Get("target").(string))
	tflog.Debug(ctx, "Validating ruleset rules based on target", map[string]any{"target": target})

	rulesRaw := d.Get("rules").([]any)
	if len(rulesRaw) == 0 {
		tflog.Debug(ctx, "No rules block, skipping validation")
		return nil
	}

	return validateRulesForTarget(ctx, d)
}

func validateConditionsFieldForBranchAndTagTargets(ctx context.Context, target github.RulesetTarget, conditions map[string]any, isOrg bool) error {
	tflog.Debug(ctx, fmt.Sprintf("Validating conditions field for %s target", target), map[string]any{"target": target, "conditions": conditions, "isOrg": isOrg})

	if conditions["ref_name"] == nil || len(conditions["ref_name"].([]any)) == 0 {
		tflog.Debug(ctx, fmt.Sprintf("Missing ref_name for %s target", target), map[string]any{"target": target})
		return fmt.Errorf("ref_name must be set for %s target", target)
	}

	// Repository rulesets don't have repository_name or repository_id, only org rulesets do.
	if isOrg {
		if (conditions["repository_name"] == nil || len(conditions["repository_name"].([]any)) == 0) && (conditions["repository_id"] == nil || len(conditions["repository_id"].([]any)) == 0) {
			tflog.Debug(ctx, fmt.Sprintf("Missing repository_name or repository_id for %s target", target), map[string]any{"target": target})
			return fmt.Errorf("either repository_name or repository_id must be set for %s target", target)
		}
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

package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Repository target rules (enterprise only)
var repositoryTargetRules = []string{
	"repository_creation",
	"repository_deletion",
	"repository_transfer",
	"repository_name",
	"repository_visibility",
}

// resourceGithubEnterpriseRulesetCustomizeDiff validates enterprise ruleset configuration
func resourceGithubEnterpriseRulesetCustomizeDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	target := d.Get("target").(string)

	// Validate conditions
	if err := validateEnterpriseConditions(d, target); err != nil {
		return err
	}

	// Validate rules
	if err := validateEnterpriseRules(d, target); err != nil {
		return err
	}

	return nil
}

// validateEnterpriseConditions validates conditions based on target type
func validateEnterpriseConditions(d *schema.ResourceDiff, target string) error {
	conditions := d.Get("conditions").([]interface{})
	if len(conditions) == 0 {
		return nil
	}

	conditionsMap := conditions[0].(map[string]interface{})
	refName := conditionsMap["ref_name"].([]interface{})
	hasRefName := len(refName) > 0

	switch target {
	case "branch", "tag":
		if !hasRefName {
			return fmt.Errorf("'ref_name' condition is required when target is '%s'", target)
		}
	case "push", "repository":
		if hasRefName {
			return fmt.Errorf("'ref_name' condition must not be set when target is '%s'", target)
		}
	}

	return nil
}

// validateEnterpriseRules validates rules based on target type
func validateEnterpriseRules(d *schema.ResourceDiff, target string) error {
	rules := d.Get("rules").([]interface{})
	if len(rules) == 0 {
		return nil
	}

	rulesMap := rules[0].(map[string]interface{})

	// Repository rules only valid for repository target
	if target != "repository" {
		for _, rule := range repositoryTargetRules {
			if isRuleSet(rulesMap, rule) {
				return fmt.Errorf("rule '%s' is only valid for target 'repository', not '%s'", rule, target)
			}
		}
	}

	// Push rules only valid for push target
	pushRules := []string{"file_path_restriction", "max_file_size", "max_file_path_length", "file_extension_restriction"}
	if target != "push" {
		for _, rule := range pushRules {
			if isRuleSet(rulesMap, rule) {
				return fmt.Errorf("rule '%s' is only valid for target 'push', not '%s'", rule, target)
			}
		}
	}

	// Branch/tag rules not valid for push or repository targets
	if target == "push" || target == "repository" {
		branchTagRules := []string{
			"creation", "deletion", "update", "required_linear_history", "required_signatures",
			"pull_request", "required_status_checks", "non_fast_forward", "commit_message_pattern",
			"commit_author_email_pattern", "committer_email_pattern", "branch_name_pattern",
			"tag_name_pattern", "required_workflows", "required_code_scanning", "copilot_code_review",
		}
		for _, rule := range branchTagRules {
			if isRuleSet(rulesMap, rule) {
				return fmt.Errorf("rule '%s' is only valid for target 'branch' or 'tag', not '%s'", rule, target)
			}
		}
	}

	return nil
}

// isRuleSet checks if a rule is set in the rules map
func isRuleSet(rules map[string]interface{}, ruleName string) bool {
	if val, ok := rules[ruleName]; ok {
		switch v := val.(type) {
		case bool:
			return v
		case []interface{}:
			return len(v) > 0
		default:
			return val != nil
		}
	}
	return false
}

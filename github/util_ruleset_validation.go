package github

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateOrganizationRulesetConditions(ctx context.Context, d *schema.ResourceDiff, meta any) error {
	target := d.Get("target").(string)
	conditions := d.Get("conditions").([]any)

	// If no conditions, allow (repository-level rulesets)
	if len(conditions) == 0 {
		return nil
	}

	// Get the conditions block
	conditionsMap := conditions[0].(map[string]any)

	// Check ref_name
	refName, hasRefName := conditionsMap["ref_name"].([]any)
	hasRefNameValue := hasRefName && len(refName) > 0

	// Check repository targeting methods
	repoId, hasRepoId := conditionsMap["repository_id"].([]any)
	hasRepoIdValue := hasRepoId && len(repoId) > 0

	repoName, hasRepoName := conditionsMap["repository_name"].([]any)
	hasRepoNameValue := hasRepoName && len(repoName) > 0

	repoProp, hasRepoProp := conditionsMap["repository_property"].([]any)
	hasRepoPropValue := hasRepoProp && len(repoProp) > 0

	// Count repository targeting methods
	repoTargetingCount := 0
	if hasRepoIdValue {
		repoTargetingCount++
	}
	if hasRepoNameValue {
		repoTargetingCount++
	}
	if hasRepoPropValue {
		repoTargetingCount++
	}

	// Validate based on target type
	switch target {
	case "push":
		// Push rulesets must NOT have ref_name
		if hasRefNameValue {
			return fmt.Errorf("ref_name cannot be set for push rulesets")
		}
		// Push rulesets must have exactly one repository targeting method
		if repoTargetingCount != 1 {
			return fmt.Errorf("push rulesets require exactly one of repository_id, repository_name, or repository_property")
		}

	case "branch", "tag":
		// Branch/tag rulesets must have ref_name
		if !hasRefNameValue {
			return fmt.Errorf("%s rulesets require ref_name to be set", target)
		}
		// Branch/tag rulesets must have exactly one repository targeting method
		if repoTargetingCount != 1 {
			return fmt.Errorf("%s rulesets require exactly one of repository_id, repository_name, or repository_property", target)
		}
	}

	return nil
}

func validateRepositoryRulesetConditions(ctx context.Context, d *schema.ResourceDiff, meta any) error {
	target := d.Get("target").(string)
	conditions := d.Get("conditions").([]any)

	// If no conditions, allow (repository-level rulesets without conditions)
	if len(conditions) == 0 {
		return nil
	}

	// Get the conditions block
	conditionsMap := conditions[0].(map[string]any)

	// Check ref_name
	refName, hasRefName := conditionsMap["ref_name"].([]any)
	hasRefNameValue := hasRefName && len(refName) > 0

	// Validate based on target type
	switch target {
	case "push":
		// Push rulesets must NOT have ref_name
		if hasRefNameValue {
			return fmt.Errorf("ref_name cannot be set for push rulesets")
		}

	case "branch", "tag":
		// Branch/tag rulesets must have ref_name
		if !hasRefNameValue {
			return fmt.Errorf("%s rulesets require ref_name to be set", target)
		}
	}

	return nil
}

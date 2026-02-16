package github

import (
	"context"
	"testing"

	"github.com/google/go-github/v82/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestExpandRulesBasicRules(t *testing.T) {
	// Test expanding basic boolean rules with RepositoryRulesetRules
	rulesMap := map[string]any{
		"creation":                true,
		"deletion":                true,
		"required_linear_history": true,
		"required_signatures":     false,
		"non_fast_forward":        true,
	}

	input := []any{rulesMap}
	result := expandRules(input, false)

	if result == nil {
		t.Fatal("Expected result to not be nil")
		return
	}

	// Check boolean rules - they use EmptyRuleParameters and are nil when false
	if result.Creation == nil {
		t.Error("Expected Creation rule to be set")
	}

	if result.Deletion == nil {
		t.Error("Expected Deletion rule to be set")
	}

	if result.RequiredLinearHistory == nil {
		t.Error("Expected RequiredLinearHistory rule to be set")
	}

	if result.RequiredSignatures != nil {
		t.Error("Expected RequiredSignatures rule to be nil (false)")
	}

	if result.NonFastForward == nil {
		t.Error("Expected NonFastForward rule to be set")
	}
}

func TestFlattenRulesBasicRules(t *testing.T) {
	// Test flattening basic boolean rules with RepositoryRulesetRules
	rules := &github.RepositoryRulesetRules{
		Creation:              &github.EmptyRuleParameters{},
		Deletion:              &github.EmptyRuleParameters{},
		RequiredLinearHistory: &github.EmptyRuleParameters{},
		RequiredSignatures:    nil, // false means nil
		NonFastForward:        &github.EmptyRuleParameters{},
	}

	result := flattenRules(t.Context(), rules, false)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	rulesMap := result[0].(map[string]any)

	// Should contain the rules
	if !rulesMap["creation"].(bool) {
		t.Error("Expected creation rule to be true")
	}

	if !rulesMap["deletion"].(bool) {
		t.Error("Expected deletion rule to be true")
	}

	if !rulesMap["required_linear_history"].(bool) {
		t.Error("Expected required_linear_history rule to be true")
	}

	if rulesMap["required_signatures"].(bool) {
		t.Error("Expected required_signatures rule to be false")
	}

	if !rulesMap["non_fast_forward"].(bool) {
		t.Error("Expected non_fast_forward rule to be true")
	}
}

func TestExpandRulesMaxFilePathLength(t *testing.T) {
	// Test that max_file_path_length rule is properly expanded
	maxPathLength := 512

	rulesMap := map[string]any{
		"max_file_path_length": []any{
			map[string]any{
				"max_file_path_length": maxPathLength,
			},
		},
	}

	input := []any{rulesMap}
	result := expandRules(input, false)

	if result == nil {
		t.Fatal("Expected result to not be nil")
		return
	}

	if result.MaxFilePathLength == nil {
		t.Fatal("Expected MaxFilePathLength rule to be set")
		return
	}

	if result.MaxFilePathLength.MaxFilePathLength != maxPathLength {
		t.Errorf("Expected MaxFilePathLength to be %d, got %d", maxPathLength, result.MaxFilePathLength.MaxFilePathLength)
	}
}

func TestFlattenRulesMaxFilePathLength(t *testing.T) {
	// Test that max_file_path_length rule is properly flattened
	maxPathLength := 256
	rules := &github.RepositoryRulesetRules{
		MaxFilePathLength: &github.MaxFilePathLengthRuleParameters{
			MaxFilePathLength: maxPathLength,
		},
	}

	result := flattenRules(t.Context(), rules, false)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	rulesMap := result[0].(map[string]any)
	maxFilePathLengthRules := rulesMap["max_file_path_length"].([]map[string]any)

	if len(maxFilePathLengthRules) != 1 {
		t.Fatalf("Expected 1 max_file_path_length rule, got %d", len(maxFilePathLengthRules))
	}

	if maxFilePathLengthRules[0]["max_file_path_length"] != maxPathLength {
		t.Errorf("Expected max_file_path_length to be %d, got %v", maxPathLength, maxFilePathLengthRules[0]["max_file_path_length"])
	}
}

func TestRoundTripMaxFilePathLength(t *testing.T) {
	// Test that max_file_path_length rule survives expand -> flatten round trip
	maxPathLength := 1024

	// Start with terraform configuration
	rulesMap := map[string]any{
		"max_file_path_length": []any{
			map[string]any{
				"max_file_path_length": maxPathLength,
			},
		},
	}

	input := []any{rulesMap}

	// Expand to GitHub API format
	expandedRules := expandRules(input, false)

	if expandedRules == nil {
		t.Fatal("Expected expandedRules to not be nil")
	}

	// Flatten back to terraform format
	flattenedResult := flattenRules(t.Context(), expandedRules, false)

	if len(flattenedResult) != 1 {
		t.Fatalf("Expected 1 flattened result, got %d", len(flattenedResult))
	}

	flattenedRulesMap := flattenedResult[0].(map[string]any)
	maxFilePathLengthRules := flattenedRulesMap["max_file_path_length"].([]map[string]any)

	if len(maxFilePathLengthRules) != 1 {
		t.Fatalf("Expected 1 max_file_path_length rule after round trip, got %d", len(maxFilePathLengthRules))
	}

	if maxFilePathLengthRules[0]["max_file_path_length"] != maxPathLength {
		t.Errorf("Expected max_file_path_length to be %d after round trip, got %v", maxPathLength, maxFilePathLengthRules[0]["max_file_path_length"])
	}
}

func TestExpandRulesMaxFileSize(t *testing.T) {
	// Test that max_file_size rule is properly expanded
	maxFileSize := int64(1048576) // 1MB

	rulesMap := map[string]any{
		"max_file_size": []any{
			map[string]any{
				"max_file_size": float64(maxFileSize),
			},
		},
	}

	input := []any{rulesMap}
	result := expandRules(input, false)

	if result == nil {
		t.Fatal("Expected result to not be nil")
		return
	}

	if result.MaxFileSize == nil {
		t.Fatal("Expected MaxFileSize rule to be set")
		return
	}

	if result.MaxFileSize.MaxFileSize != maxFileSize {
		t.Errorf("Expected MaxFileSize to be %d, got %d", maxFileSize, result.MaxFileSize.MaxFileSize)
	}
}

func TestFlattenRulesMaxFileSize(t *testing.T) {
	// Test that max_file_size rule is properly flattened
	maxFileSize := int64(5242880) // 5MB
	rules := &github.RepositoryRulesetRules{
		MaxFileSize: &github.MaxFileSizeRuleParameters{
			MaxFileSize: maxFileSize,
		},
	}

	result := flattenRules(t.Context(), rules, false)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	rulesMap := result[0].(map[string]any)
	maxFileSizeRules := rulesMap["max_file_size"].([]map[string]any)

	if len(maxFileSizeRules) != 1 {
		t.Fatalf("Expected 1 max_file_size rule, got %d", len(maxFileSizeRules))
	}

	if maxFileSizeRules[0]["max_file_size"] != maxFileSize {
		t.Errorf("Expected max_file_size to be %d, got %v", maxFileSize, maxFileSizeRules[0]["max_file_size"])
	}
}

func TestExpandRulesFileExtensionRestriction(t *testing.T) {
	// Test that file_extension_restriction rule is properly expanded
	restrictedExtensions := []string{".exe", ".bat", ".com"}

	rulesMap := map[string]any{
		"file_extension_restriction": []any{
			map[string]any{
				"restricted_file_extensions": schema.NewSet(schema.HashString, []any{".exe", ".bat", ".com"}),
			},
		},
	}

	input := []any{rulesMap}
	result := expandRules(input, false)

	if result == nil {
		t.Fatal("Expected result to not be nil")
		return
	}

	if result.FileExtensionRestriction == nil {
		t.Fatal("Expected FileExtensionRestriction rule to be set")
		return
	}

	if len(result.FileExtensionRestriction.RestrictedFileExtensions) != len(restrictedExtensions) {
		t.Errorf("Expected %d restricted extensions, got %d", len(restrictedExtensions), len(result.FileExtensionRestriction.RestrictedFileExtensions))
	}

	resultExtensions := make(map[string]bool)
	for _, ext := range result.FileExtensionRestriction.RestrictedFileExtensions {
		resultExtensions[ext] = true
	}

	for _, expectedExt := range restrictedExtensions {
		if !resultExtensions[expectedExt] {
			t.Errorf("Expected extension %s not found in result", expectedExt)
		}
	}
}

func TestFlattenRulesFileExtensionRestriction(t *testing.T) {
	// Test that file_extension_restriction rule is properly flattened
	restrictedExtensions := []string{".exe", ".bat", ".com"}
	rules := &github.RepositoryRulesetRules{
		FileExtensionRestriction: &github.FileExtensionRestrictionRuleParameters{
			RestrictedFileExtensions: restrictedExtensions,
		},
	}

	result := flattenRules(t.Context(), rules, false)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	rulesMap := result[0].(map[string]any)
	fileExtensionRules := rulesMap["file_extension_restriction"].([]map[string]any)

	if len(fileExtensionRules) != 1 {
		t.Fatalf("Expected 1 file_extension_restriction rule, got %d", len(fileExtensionRules))
	}

	actualExtensions := fileExtensionRules[0]["restricted_file_extensions"].([]string)
	if len(actualExtensions) != len(restrictedExtensions) {
		t.Errorf("Expected %d restricted extensions, got %d", len(restrictedExtensions), len(actualExtensions))
	}

	for i, ext := range restrictedExtensions {
		if actualExtensions[i] != ext {
			t.Errorf("Expected extension %s at index %d, got %s", ext, i, actualExtensions[i])
		}
	}
}

func TestCompletePushRulesetSupport(t *testing.T) {
	// Test that all push-specific rules are supported together
	rulesMap := map[string]any{
		"file_path_restriction": []any{
			map[string]any{
				"restricted_file_paths": []any{"secrets/", "*.key", "private/"},
			},
		},
		"max_file_size": []any{
			map[string]any{
				"max_file_size": 5, // 5MB
			},
		},
		"max_file_path_length": []any{
			map[string]any{
				"max_file_path_length": 300,
			},
		},
		"file_extension_restriction": []any{
			map[string]any{
				"restricted_file_extensions": schema.NewSet(schema.HashString, []any{".exe", ".bat", ".sh"}),
			},
		},
	}

	input := []any{rulesMap}

	// Expand to GitHub API format
	expandedRules := expandRules(input, false)

	if expandedRules == nil {
		t.Fatal("Expected expandedRules to not be nil")
		return
	}

	// Count how many rules we have
	ruleCount := 0
	if expandedRules.FilePathRestriction != nil {
		ruleCount++
	}
	if expandedRules.MaxFileSize != nil {
		ruleCount++
	}
	if expandedRules.MaxFilePathLength != nil {
		ruleCount++
	}
	if expandedRules.FileExtensionRestriction != nil {
		ruleCount++
	}

	if ruleCount != 4 {
		t.Fatalf("Expected 4 expanded rules for complete push ruleset, got %d", ruleCount)
	}

	// Flatten back to terraform format
	flattenedResult := flattenRules(t.Context(), expandedRules, false)

	if len(flattenedResult) != 1 {
		t.Fatalf("Expected 1 flattened result, got %d", len(flattenedResult))
	}

	flattenedRulesMap := flattenedResult[0].(map[string]any)

	// Verify file_path_restriction
	filePathRules := flattenedRulesMap["file_path_restriction"].([]map[string]any)
	if len(filePathRules) != 1 {
		t.Fatalf("Expected 1 file_path_restriction rule, got %d", len(filePathRules))
	}
	restrictedPaths := filePathRules[0]["restricted_file_paths"].([]string)
	if len(restrictedPaths) != 3 {
		t.Errorf("Expected 3 restricted file paths, got %d", len(restrictedPaths))
	}

	// Verify max_file_size
	maxFileSizeRules := flattenedRulesMap["max_file_size"].([]map[string]any)
	if len(maxFileSizeRules) != 1 {
		t.Fatalf("Expected 1 max_file_size rule, got %d", len(maxFileSizeRules))
	}
	if maxFileSizeRules[0]["max_file_size"] != int64(5) {
		t.Errorf("Expected max_file_size to be 5, got %v", maxFileSizeRules[0]["max_file_size"])
	}

	// Verify max_file_path_length
	maxFilePathLengthRules := flattenedRulesMap["max_file_path_length"].([]map[string]any)
	if len(maxFilePathLengthRules) != 1 {
		t.Fatalf("Expected 1 max_file_path_length rule, got %d", len(maxFilePathLengthRules))
	}
	if maxFilePathLengthRules[0]["max_file_path_length"] != 300 {
		t.Errorf("Expected max_file_path_length to be 300, got %v", maxFilePathLengthRules[0]["max_file_path_length"])
	}

	// Verify file_extension_restriction
	fileExtRules := flattenedRulesMap["file_extension_restriction"].([]map[string]any)
	if len(fileExtRules) != 1 {
		t.Fatalf("Expected 1 file_extension_restriction rule, got %d", len(fileExtRules))
	}
	restrictedExts := fileExtRules[0]["restricted_file_extensions"].([]string)
	if len(restrictedExts) != 3 {
		t.Errorf("Expected 3 restricted file extensions, got %d", len(restrictedExts))
	}
}

func TestCopilotCodeReviewRoundTrip(t *testing.T) {
	// Test that copilot_code_review rule survives expand -> flatten round trip
	rulesMap := map[string]any{
		"copilot_code_review": []any{
			map[string]any{
				"review_on_push":             true,
				"review_draft_pull_requests": false,
			},
		},
	}

	input := []any{rulesMap}

	// Expand to GitHub API format
	expandedRules := expandRules(input, false)

	if expandedRules == nil {
		t.Fatal("Expected expandedRules to not be nil")
	}

	if expandedRules.CopilotCodeReview == nil {
		t.Fatal("Expected CopilotCodeReview rule to be set")
	}

	if expandedRules.CopilotCodeReview.ReviewOnPush != true {
		t.Errorf("Expected ReviewOnPush to be true, got %v", expandedRules.CopilotCodeReview.ReviewOnPush)
	}

	if expandedRules.CopilotCodeReview.ReviewDraftPullRequests != false {
		t.Errorf("Expected ReviewDraftPullRequests to be false, got %v", expandedRules.CopilotCodeReview.ReviewDraftPullRequests)
	}

	// Flatten back to terraform format
	flattenedResult := flattenRules(t.Context(), expandedRules, false)

	if len(flattenedResult) != 1 {
		t.Fatalf("Expected 1 flattened result, got %d", len(flattenedResult))
	}

	flattenedRulesMap := flattenedResult[0].(map[string]any)
	copilotRules := flattenedRulesMap["copilot_code_review"].([]map[string]any)

	if len(copilotRules) != 1 {
		t.Fatalf("Expected 1 copilot_code_review rule after round trip, got %d", len(copilotRules))
	}

	if copilotRules[0]["review_on_push"] != true {
		t.Errorf("Expected review_on_push to be true, got %v", copilotRules[0]["review_on_push"])
	}

	if copilotRules[0]["review_draft_pull_requests"] != false {
		t.Errorf("Expected review_draft_pull_requests to be false, got %v", copilotRules[0]["review_draft_pull_requests"])
	}
}

func TestFlattenConditions_PushRuleset_WithRepositoryNameOnly(t *testing.T) {
	// Push rulesets don't use ref_name - they only have repository_name or repository_id.
	// flattenConditions should return the conditions even when RefName is nil.
	conditions := &github.RepositoryRulesetConditions{
		RefName: nil, // Push rulesets don't have ref_name
		RepositoryName: &github.RepositoryRulesetRepositoryNamesConditionParameters{
			Include: []string{"~ALL"},
			Exclude: []string{},
		},
	}

	result := flattenConditions(t.Context(), conditions, true) // org=true for organization rulesets

	if len(result) != 1 {
		t.Fatalf("Expected 1 conditions block, got %d", len(result))
	}

	conditionsMap := result[0].(map[string]any)

	// ref_name should be empty for push rulesets
	refNameSlice := conditionsMap["ref_name"]
	if refNameSlice != nil {
		t.Fatalf("Expected ref_name to be nil, got %T", conditionsMap["ref_name"])
	}

	// repository_name should be present
	repoNameSlice, ok := conditionsMap["repository_name"].([]map[string]any)
	if !ok {
		t.Fatalf("Expected repository_name to be []map[string]any, got %T", conditionsMap["repository_name"])
	}
	if len(repoNameSlice) != 1 {
		t.Fatalf("Expected 1 repository_name block, got %d", len(repoNameSlice))
	}

	include, ok := repoNameSlice[0]["include"].([]string)
	if !ok {
		t.Fatalf("Expected include to be []string, got %T", repoNameSlice[0]["include"])
	}
	if len(include) != 1 || include[0] != "~ALL" {
		t.Errorf("Expected include to be [~ALL], got %v", include)
	}
}

func TestFlattenConditions_BranchRuleset_WithRefNameAndRepositoryName(t *testing.T) {
	// Branch/tag rulesets have both ref_name and repository_name.
	// This test ensures we didn't break the existing behavior.
	conditions := &github.RepositoryRulesetConditions{
		RefName: &github.RepositoryRulesetRefConditionParameters{
			Include: []string{"~DEFAULT_BRANCH", "refs/heads/main"},
			Exclude: []string{"refs/heads/experimental-*"},
		},
		RepositoryName: &github.RepositoryRulesetRepositoryNamesConditionParameters{
			Include: []string{"~ALL"},
			Exclude: []string{"test-*"},
		},
	}

	result := flattenConditions(t.Context(), conditions, true) // org=true for organization rulesets

	if len(result) != 1 {
		t.Fatalf("Expected 1 conditions block, got %d", len(result))
	}

	conditionsMap := result[0].(map[string]any)

	// ref_name should be present for branch/tag rulesets
	refNameSlice, ok := conditionsMap["ref_name"].([]map[string]any)
	if !ok {
		t.Fatalf("Expected ref_name to be []map[string]any, got %T", conditionsMap["ref_name"])
	}
	if len(refNameSlice) != 1 {
		t.Fatalf("Expected 1 ref_name block, got %d", len(refNameSlice))
	}

	refInclude, ok := refNameSlice[0]["include"].([]string)
	if !ok {
		t.Fatalf("Expected ref_name include to be []string, got %T", refNameSlice[0]["include"])
	}
	if len(refInclude) != 2 {
		t.Errorf("Expected 2 ref_name includes, got %d", len(refInclude))
	}

	refExclude, ok := refNameSlice[0]["exclude"].([]string)
	if !ok {
		t.Fatalf("Expected ref_name exclude to be []string, got %T", refNameSlice[0]["exclude"])
	}
	if len(refExclude) != 1 {
		t.Errorf("Expected 1 ref_name exclude, got %d", len(refExclude))
	}

	// repository_name should also be present
	repoNameSlice, ok := conditionsMap["repository_name"].([]map[string]any)
	if !ok {
		t.Fatalf("Expected repository_name to be []map[string]any, got %T", conditionsMap["repository_name"])
	}
	if len(repoNameSlice) != 1 {
		t.Fatalf("Expected 1 repository_name block, got %d", len(repoNameSlice))
	}

	repoInclude, ok := repoNameSlice[0]["include"].([]string)
	if !ok {
		t.Fatalf("Expected repository_name include to be []string, got %T", repoNameSlice[0]["include"])
	}
	if len(repoInclude) != 1 || repoInclude[0] != "~ALL" {
		t.Errorf("Expected repository_name include to be [~ALL], got %v", repoInclude)
	}

	repoExclude, ok := repoNameSlice[0]["exclude"].([]string)
	if !ok {
		t.Fatalf("Expected repository_name exclude to be []string, got %T", repoNameSlice[0]["exclude"])
	}
	if len(repoExclude) != 1 || repoExclude[0] != "test-*" {
		t.Errorf("Expected repository_name exclude to be [test-*], got %v", repoExclude)
	}
}

func TestFlattenConditions_PushRuleset_WithRepositoryIdOnly(t *testing.T) {
	// Push rulesets can also use repository_id instead of repository_name.
	conditions := &github.RepositoryRulesetConditions{
		RefName: nil, // Push rulesets don't have ref_name
		RepositoryID: &github.RepositoryRulesetRepositoryIDsConditionParameters{
			RepositoryIDs: []int64{12345, 67890},
		},
	}

	result := flattenConditions(t.Context(), conditions, true) // org=true for organization rulesets

	if len(result) != 1 {
		t.Fatalf("Expected 1 conditions block, got %d", len(result))
	}

	conditionsMap := result[0].(map[string]any)

	// ref_name should be nil for push rulesets
	refNameSlice := conditionsMap["ref_name"]
	if refNameSlice != nil {
		t.Fatalf("Expected ref_name to be nil, got %T", conditionsMap["ref_name"])
	}

	// repository_id should be present
	repoIDs, ok := conditionsMap["repository_id"].([]int64)
	if !ok {
		t.Fatalf("Expected repository_id to be []int64, got %T", conditionsMap["repository_id"])
	}
	if len(repoIDs) != 2 {
		t.Fatalf("Expected 2 repository IDs, got %d", len(repoIDs))
	}
	if repoIDs[0] != 12345 || repoIDs[1] != 67890 {
		t.Errorf("Expected repository IDs [12345, 67890], got %v", repoIDs)
	}
}

func TestExpandRequiredReviewers(t *testing.T) {
	input := []any{
		map[string]any{
			"reviewer": []any{
				map[string]any{
					"id":   12345,
					"type": "Team",
				},
			},
			"file_patterns":     []any{"*.go", "src/**/*.ts"},
			"minimum_approvals": 2,
		},
		map[string]any{
			"reviewer": []any{
				map[string]any{
					"id":   67890,
					"type": "Team",
				},
			},
			"file_patterns":     []any{"docs/**/*.md"},
			"minimum_approvals": 1,
		},
	}

	result := expandRequiredReviewers(input)

	if len(result) != 2 {
		t.Fatalf("Expected 2 reviewers, got %d", len(result))
	}

	// Check first reviewer
	if result[0].Reviewer == nil {
		t.Fatal("Expected first reviewer to have a Reviewer")
	}
	if *result[0].Reviewer.ID != 12345 {
		t.Errorf("Expected first reviewer ID to be 12345, got %d", *result[0].Reviewer.ID)
	}
	if *result[0].Reviewer.Type != github.RulesetReviewerTypeTeam {
		t.Errorf("Expected first reviewer type to be Team, got %s", *result[0].Reviewer.Type)
	}
	if *result[0].MinimumApprovals != 2 {
		t.Errorf("Expected first reviewer minimum approvals to be 2, got %d", *result[0].MinimumApprovals)
	}
	if len(result[0].FilePatterns) != 2 {
		t.Fatalf("Expected first reviewer to have 2 file patterns, got %d", len(result[0].FilePatterns))
	}
	if result[0].FilePatterns[0] != "*.go" || result[0].FilePatterns[1] != "src/**/*.ts" {
		t.Errorf("Unexpected file patterns for first reviewer: %v", result[0].FilePatterns)
	}

	// Check second reviewer
	if result[1].Reviewer == nil {
		t.Fatal("Expected second reviewer to have a Reviewer")
	}
	if *result[1].Reviewer.ID != 67890 {
		t.Errorf("Expected second reviewer ID to be 67890, got %d", *result[1].Reviewer.ID)
	}
	if *result[1].MinimumApprovals != 1 {
		t.Errorf("Expected second reviewer minimum approvals to be 1, got %d", *result[1].MinimumApprovals)
	}
}

func TestExpandRequiredReviewersEmpty(t *testing.T) {
	result := expandRequiredReviewers([]any{})
	if result != nil {
		t.Error("Expected nil for empty input")
	}

	result = expandRequiredReviewers(nil)
	if result != nil {
		t.Error("Expected nil for nil input")
	}
}

func TestFlattenRequiredReviewers(t *testing.T) {
	reviewerType := github.RulesetReviewerTypeTeam
	reviewers := []*github.RulesetRequiredReviewer{
		{
			MinimumApprovals: github.Ptr(2),
			FilePatterns:     []string{"*.go", "src/**/*.ts"},
			Reviewer: &github.RulesetReviewer{
				ID:   github.Ptr(int64(12345)),
				Type: &reviewerType,
			},
		},
		{
			MinimumApprovals: github.Ptr(1),
			FilePatterns:     []string{"docs/**/*.md"},
			Reviewer: &github.RulesetReviewer{
				ID:   github.Ptr(int64(67890)),
				Type: &reviewerType,
			},
		},
	}

	result := flattenRequiredReviewers(reviewers)

	if len(result) != 2 {
		t.Fatalf("Expected 2 reviewers, got %d", len(result))
	}

	// Check first reviewer
	if result[0]["minimum_approvals"] != 2 {
		t.Errorf("Expected first reviewer minimum approvals to be 2, got %v", result[0]["minimum_approvals"])
	}
	filePatterns := result[0]["file_patterns"].([]string)
	if len(filePatterns) != 2 {
		t.Fatalf("Expected first reviewer to have 2 file patterns, got %d", len(filePatterns))
	}
	if filePatterns[0] != "*.go" || filePatterns[1] != "src/**/*.ts" {
		t.Errorf("Unexpected file patterns for first reviewer: %v", filePatterns)
	}

	reviewerBlock := result[0]["reviewer"].([]map[string]any)
	if len(reviewerBlock) != 1 {
		t.Fatalf("Expected 1 reviewer block, got %d", len(reviewerBlock))
	}
	if reviewerBlock[0]["id"] != 12345 {
		t.Errorf("Expected first reviewer ID to be 12345, got %v", reviewerBlock[0]["id"])
	}
	if reviewerBlock[0]["type"] != "Team" {
		t.Errorf("Expected first reviewer type to be Team, got %v", reviewerBlock[0]["type"])
	}

	// Check second reviewer
	if result[1]["minimum_approvals"] != 1 {
		t.Errorf("Expected second reviewer minimum approvals to be 1, got %v", result[1]["minimum_approvals"])
	}
}

func TestFlattenRequiredReviewersEmpty(t *testing.T) {
	result := flattenRequiredReviewers(nil)
	if result != nil {
		t.Error("Expected nil for nil input")
	}

	result = flattenRequiredReviewers([]*github.RulesetRequiredReviewer{})
	if result != nil {
		t.Error("Expected nil for empty slice input")
	}
}

func TestRoundTripRequiredReviewers(t *testing.T) {
	// Start with Terraform-style input
	input := []any{
		map[string]any{
			"reviewer": []any{
				map[string]any{
					"id":   12345,
					"type": "Team",
				},
			},
			"file_patterns":     []any{"*.go", "src/**/*.ts"},
			"minimum_approvals": 2,
		},
	}

	// Expand to go-github types
	expanded := expandRequiredReviewers(input)

	// Flatten back to Terraform types
	flattened := flattenRequiredReviewers(expanded)

	// Verify the round trip maintains data
	if len(flattened) != 1 {
		t.Fatalf("Expected 1 reviewer after round trip, got %d", len(flattened))
	}

	if flattened[0]["minimum_approvals"] != 2 {
		t.Errorf("Expected minimum_approvals to be 2 after round trip, got %v", flattened[0]["minimum_approvals"])
	}

	filePatterns := flattened[0]["file_patterns"].([]string)
	if len(filePatterns) != 2 {
		t.Fatalf("Expected 2 file patterns after round trip, got %d", len(filePatterns))
	}

	reviewerBlock := flattened[0]["reviewer"].([]map[string]any)
	if len(reviewerBlock) != 1 {
		t.Fatalf("Expected 1 reviewer block after round trip, got %d", len(reviewerBlock))
	}
	if reviewerBlock[0]["id"] != 12345 {
		t.Errorf("Expected reviewer ID to be 12345 after round trip, got %v", reviewerBlock[0]["id"])
	}
	if reviewerBlock[0]["type"] != "Team" {
		t.Errorf("Expected reviewer type to be Team after round trip, got %v", reviewerBlock[0]["type"])
	}
}

// Tests for new condition types: organization_id

func TestExpandConditionsOrganizationID(t *testing.T) {
	// Test expanding organization_id condition
	conditionsMap := map[string]any{
		"ref_name": []any{
			map[string]any{
				"include": []any{"main", "develop"},
				"exclude": []any{"feature/*"},
			},
		},
		"organization_id": []any{123, 456, 789},
	}

	input := []any{conditionsMap}
	result := expandConditions(input, true) // org=true for enterprise rulesets

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	if result.OrganizationID == nil {
		t.Fatal("Expected OrganizationID to be set")
	}

	expectedIDs := []int64{123, 456, 789}
	if len(result.OrganizationID.OrganizationIDs) != len(expectedIDs) {
		t.Fatalf("Expected %d organization IDs, got %d", len(expectedIDs), len(result.OrganizationID.OrganizationIDs))
	}

	for i, expectedID := range expectedIDs {
		if result.OrganizationID.OrganizationIDs[i] != expectedID {
			t.Errorf("Expected organization ID %d at index %d, got %d", expectedID, i, result.OrganizationID.OrganizationIDs[i])
		}
	}
}

func TestFlattenConditionsOrganizationID(t *testing.T) {
	// Test flattening organization_id condition
	conditions := &github.RepositoryRulesetConditions{
		RefName: &github.RepositoryRulesetRefConditionParameters{
			Include: []string{"main"},
			Exclude: []string{},
		},
		OrganizationID: &github.RepositoryRulesetOrganizationIDsConditionParameters{
			OrganizationIDs: []int64{123, 456},
		},
	}

	result := flattenConditions(context.Background(), conditions, true)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	conditionsMap := result[0].(map[string]any)
	orgIDs := conditionsMap["organization_id"].([]int64)

	if len(orgIDs) != 2 {
		t.Fatalf("Expected 2 organization IDs, got %d", len(orgIDs))
	}

	if orgIDs[0] != 123 || orgIDs[1] != 456 {
		t.Errorf("Expected organization IDs [123, 456], got %v", orgIDs)
	}
}

func TestRoundTripConditionsWithAllProperties(t *testing.T) {
	// Test that organization_id condition survives expand -> flatten round trip
	conditionsMap := map[string]any{
		"ref_name": []any{
			map[string]any{
				"include": []any{"main", "develop"},
				"exclude": []any{"feature/*"},
			},
		},
		"organization_id": []any{123, 456},
	}

	input := []any{conditionsMap}

	// Expand to GitHub API format
	expandedConditions := expandConditions(input, true)

	if expandedConditions == nil {
		t.Fatal("Expected expandedConditions to not be nil")
	}

	// Flatten back to terraform format
	flattenedResult := flattenConditions(context.Background(), expandedConditions, true)

	if len(flattenedResult) != 1 {
		t.Fatalf("Expected 1 flattened result, got %d", len(flattenedResult))
	}

	flattenedConditionsMap := flattenedResult[0].(map[string]any)

	// Verify organization_id survived
	orgIDs := flattenedConditionsMap["organization_id"].([]int64)
	if len(orgIDs) != 2 || orgIDs[0] != 123 || orgIDs[1] != 456 {
		t.Errorf("Expected organization_id [123, 456] after round trip, got %v", orgIDs)
	}
}

func TestExpandConditionsRepositoryProperty(t *testing.T) {
	conditionsMap := map[string]any{
		"ref_name": []any{
			map[string]any{
				"include": []any{"main"},
				"exclude": []any{},
			},
		},
		"organization_id": []any{123},
		"repository_property": []any{
			map[string]any{
				"include": []any{
					map[string]any{
						"name":            "environment",
						"property_values": []any{"production", "staging"},
						"source":          "custom",
					},
				},
				"exclude": []any{
					map[string]any{
						"name":            "team",
						"property_values": []any{"experimental"},
						"source":          "",
					},
				},
			},
		},
	}

	input := []any{conditionsMap}
	result := expandConditions(input, true)

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	if result.RepositoryProperty == nil {
		t.Fatal("Expected RepositoryProperty to be set")
	}

	if len(result.RepositoryProperty.Include) != 1 {
		t.Fatalf("Expected 1 include target, got %d", len(result.RepositoryProperty.Include))
	}

	inc := result.RepositoryProperty.Include[0]
	if inc.Name != "environment" {
		t.Errorf("Expected include name to be 'environment', got %q", inc.Name)
	}
	if len(inc.PropertyValues) != 2 || inc.PropertyValues[0] != "production" || inc.PropertyValues[1] != "staging" {
		t.Errorf("Expected include property_values [production, staging], got %v", inc.PropertyValues)
	}
	if inc.Source == nil || *inc.Source != "custom" {
		t.Errorf("Expected include source to be 'custom', got %v", inc.Source)
	}

	if len(result.RepositoryProperty.Exclude) != 1 {
		t.Fatalf("Expected 1 exclude target, got %d", len(result.RepositoryProperty.Exclude))
	}

	exc := result.RepositoryProperty.Exclude[0]
	if exc.Name != "team" {
		t.Errorf("Expected exclude name to be 'team', got %q", exc.Name)
	}
	if exc.Source != nil {
		t.Errorf("Expected exclude source to be nil for empty string, got %v", exc.Source)
	}
}

func TestFlattenConditionsRepositoryProperty(t *testing.T) {
	conditions := &github.RepositoryRulesetConditions{
		RefName: &github.RepositoryRulesetRefConditionParameters{
			Include: []string{"main"},
			Exclude: []string{},
		},
		OrganizationID: &github.RepositoryRulesetOrganizationIDsConditionParameters{
			OrganizationIDs: []int64{123},
		},
		RepositoryProperty: &github.RepositoryRulesetRepositoryPropertyConditionParameters{
			Include: []*github.RepositoryRulesetRepositoryPropertyTargetParameters{
				{
					Name:           "environment",
					PropertyValues: []string{"production"},
					Source:         github.Ptr("custom"),
				},
			},
			Exclude: []*github.RepositoryRulesetRepositoryPropertyTargetParameters{
				{
					Name:           "team",
					PropertyValues: []string{"experimental"},
				},
			},
		},
	}

	result := flattenConditions(context.Background(), conditions, true)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	conditionsMap := result[0].(map[string]any)
	repoProp, ok := conditionsMap["repository_property"].([]map[string]any)
	if !ok {
		t.Fatalf("Expected repository_property to be []map[string]any, got %T", conditionsMap["repository_property"])
	}
	if len(repoProp) != 1 {
		t.Fatalf("Expected 1 repository_property block, got %d", len(repoProp))
	}

	includes := repoProp[0]["include"].([]map[string]any)
	if len(includes) != 1 {
		t.Fatalf("Expected 1 include, got %d", len(includes))
	}
	if includes[0]["name"] != "environment" {
		t.Errorf("Expected include name to be 'environment', got %v", includes[0]["name"])
	}
	if includes[0]["source"] != "custom" {
		t.Errorf("Expected include source to be 'custom', got %v", includes[0]["source"])
	}

	excludes := repoProp[0]["exclude"].([]map[string]any)
	if len(excludes) != 1 {
		t.Fatalf("Expected 1 exclude, got %d", len(excludes))
	}
	if excludes[0]["name"] != "team" {
		t.Errorf("Expected exclude name to be 'team', got %v", excludes[0]["name"])
	}
}

func TestRoundTripConditionsRepositoryProperty(t *testing.T) {
	conditionsMap := map[string]any{
		"ref_name": []any{
			map[string]any{
				"include": []any{"main"},
				"exclude": []any{},
			},
		},
		"organization_id": []any{123},
		"repository_property": []any{
			map[string]any{
				"include": []any{
					map[string]any{
						"name":            "environment",
						"property_values": []any{"production", "staging"},
						"source":          "custom",
					},
				},
				"exclude": []any{
					map[string]any{
						"name":            "team",
						"property_values": []any{"experimental"},
						"source":          "",
					},
				},
			},
		},
	}

	input := []any{conditionsMap}
	expanded := expandConditions(input, true)
	if expanded == nil {
		t.Fatal("Expected expanded conditions to not be nil")
	}

	flattened := flattenConditions(context.Background(), expanded, true)
	if len(flattened) != 1 {
		t.Fatalf("Expected 1 flattened result, got %d", len(flattened))
	}

	flatMap := flattened[0].(map[string]any)
	repoProp, ok := flatMap["repository_property"].([]map[string]any)
	if !ok {
		t.Fatalf("Expected repository_property after round trip, got %T", flatMap["repository_property"])
	}

	includes := repoProp[0]["include"].([]map[string]any)
	if len(includes) != 1 {
		t.Fatalf("Expected 1 include after round trip, got %d", len(includes))
	}
	if includes[0]["name"] != "environment" {
		t.Errorf("Expected include name 'environment' after round trip, got %v", includes[0]["name"])
	}

	propVals := includes[0]["property_values"].([]string)
	if len(propVals) != 2 || propVals[0] != "production" || propVals[1] != "staging" {
		t.Errorf("Expected property_values [production, staging] after round trip, got %v", propVals)
	}
}

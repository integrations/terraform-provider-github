package github

import (
	"testing"

	"github.com/google/go-github/v81/github"
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

	result := flattenRules(rules, false)

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

	result := flattenRules(rules, false)

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
	flattenedResult := flattenRules(expandedRules, false)

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

	result := flattenRules(rules, false)

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

	result := flattenRules(rules, false)

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
	flattenedResult := flattenRules(expandedRules, false)

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

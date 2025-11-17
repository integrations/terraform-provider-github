package github

import (
	"encoding/json"
	"testing"

	"github.com/google/go-github/v67/github"
)

func TestFlattenRulesHandlesUnknownTypes(t *testing.T) {
	// Create some test rules including an unknown type
	unknownParams := map[string]any{
		"some_parameter": "some_value",
	}
	unknownParamsJSON, _ := json.Marshal(unknownParams)
	unknownParamsRaw := json.RawMessage(unknownParamsJSON)

	rules := []*github.RepositoryRule{
		{
			Type: "creation",
		},
		{
			Type:       "unknown_copilot_rule",
			Parameters: &unknownParamsRaw,
		},
		{
			Type: "deletion",
		},
	}

	// This should not panic or fail, even with unknown rule types
	result := flattenRules(rules, false)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	rulesMap := result[0].(map[string]any)

	// Should contain the known rules
	if !rulesMap["creation"].(bool) {
		t.Error("Expected creation rule to be true")
	}

	if !rulesMap["deletion"].(bool) {
		t.Error("Expected deletion rule to be true")
	}

	// Should NOT contain the unknown rule type
	if _, exists := rulesMap["unknown_copilot_rule"]; exists {
		t.Error("Unknown rule type should not appear in flattened rules to avoid causing diffs")
	}
}

func TestFlattenRulesHandlesMaxFileSize(t *testing.T) {
	// Test that max_file_size rule is properly handled
	maxFileSize := int64(1024000)
	params := map[string]any{
		"max_file_size": maxFileSize,
	}
	paramsJSON, _ := json.Marshal(params)
	paramsRaw := json.RawMessage(paramsJSON)

	rules := []*github.RepositoryRule{
		{
			Type:       "max_file_size",
			Parameters: &paramsRaw,
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

func TestFlattenRulesHandlesFileExtensionRestriction(t *testing.T) {
	// Test that file_extension_restriction rule is properly handled
	restrictedExtensions := []string{".exe", ".bat", ".com"}
	params := map[string]any{
		"restricted_file_extensions": restrictedExtensions,
	}
	paramsJSON, _ := json.Marshal(params)
	paramsRaw := json.RawMessage(paramsJSON)

	rules := []*github.RepositoryRule{
		{
			Type:       "file_extension_restriction",
			Parameters: &paramsRaw,
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

func TestFlattenRulesHandlesMaxFilePathLength(t *testing.T) {
	// Test that max_file_path_length rule is properly handled
	maxPathLength := 256
	params := map[string]any{
		"max_file_path_length": maxPathLength,
	}
	paramsJSON, _ := json.Marshal(params)
	paramsRaw := json.RawMessage(paramsJSON)

	rules := []*github.RepositoryRule{
		{
			Type:       "max_file_path_length",
			Parameters: &paramsRaw,
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

func TestExpandRulesHandlesMaxFilePathLength(t *testing.T) {
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

	if len(result) != 1 {
		t.Fatalf("Expected 1 rule in result, got %d", len(result))
	}

	rule := result[0]
	if rule.Type != "max_file_path_length" {
		t.Errorf("Expected rule type to be 'max_file_path_length', got %s", rule.Type)
	}

	if rule.Parameters == nil {
		t.Fatal("Expected rule parameters to be set")
	}

	var params github.RuleMaxFilePathLengthParameters
	err := json.Unmarshal(*rule.Parameters, &params)
	if err != nil {
		t.Fatalf("Failed to unmarshal parameters: %v", err)
	}

	if params.MaxFilePathLength != maxPathLength {
		t.Errorf("Expected MaxFilePathLength to be %d, got %d", maxPathLength, params.MaxFilePathLength)
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

	if len(expandedRules) != 1 {
		t.Fatalf("Expected 1 expanded rule, got %d", len(expandedRules))
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

func TestMaxFilePathLengthWithOtherRules(t *testing.T) {
	// Test that max_file_path_length works correctly alongside other rules
	maxPathLength := 200

	rulesMap := map[string]any{
		"creation": true,
		"deletion": true,
		"max_file_path_length": []any{
			map[string]any{
				"max_file_path_length": maxPathLength,
			},
		},
		"max_file_size": []any{
			map[string]any{
				"max_file_size": float64(1048576), // 1MB
			},
		},
	}

	input := []any{rulesMap}

	// Expand to GitHub API format
	expandedRules := expandRules(input, false)

	if len(expandedRules) != 4 {
		t.Fatalf("Expected 4 expanded rules, got %d", len(expandedRules))
	}

	// Verify we have all expected rule types
	ruleTypes := make(map[string]bool)
	for _, rule := range expandedRules {
		ruleTypes[rule.Type] = true
	}

	expectedTypes := []string{"creation", "deletion", "max_file_path_length", "max_file_size"}
	for _, expectedType := range expectedTypes {
		if !ruleTypes[expectedType] {
			t.Errorf("Expected rule type %s not found in expanded rules", expectedType)
		}
	}

	// Flatten back and verify
	flattenedResult := flattenRules(expandedRules, false)
	flattenedRulesMap := flattenedResult[0].(map[string]any)

	// Check that all rules are preserved
	if !flattenedRulesMap["creation"].(bool) {
		t.Error("Expected creation rule to be true")
	}

	if !flattenedRulesMap["deletion"].(bool) {
		t.Error("Expected deletion rule to be true")
	}

	maxFilePathLengthRules := flattenedRulesMap["max_file_path_length"].([]map[string]any)
	if len(maxFilePathLengthRules) != 1 || maxFilePathLengthRules[0]["max_file_path_length"] != maxPathLength {
		t.Errorf("Expected max_file_path_length rule with value %d", maxPathLength)
	}

	maxFileSizeRules := flattenedRulesMap["max_file_size"].([]map[string]any)
	if len(maxFileSizeRules) != 1 || maxFileSizeRules[0]["max_file_size"] != int64(1048576) {
		t.Error("Expected max_file_size rule with value 1048576")
	}
}

func TestMaxFilePathLengthErrorHandling(t *testing.T) {
	// Test that malformed max_file_path_length parameters are handled gracefully
	malformedParams := []byte(`{"invalid_field": "invalid_value"}`)
	paramsRaw := json.RawMessage(malformedParams)

	rules := []*github.RepositoryRule{
		{
			Type:       "max_file_path_length",
			Parameters: &paramsRaw,
		},
	}

	// This should not panic, even with malformed parameters
	result := flattenRules(rules, false)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	rulesMap := result[0].(map[string]any)
	maxFilePathLengthRules, exists := rulesMap["max_file_path_length"]

	if !exists {
		t.Error("Expected max_file_path_length rule to be present even with malformed parameters")
	}

	// The rule should be present but may have default/zero values due to unmarshaling error
	rules_slice := maxFilePathLengthRules.([]map[string]any)
	if len(rules_slice) != 1 {
		t.Errorf("Expected 1 max_file_path_length rule, got %d", len(rules_slice))
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
				"max_file_size": float64(5242880), // 5MB
			},
		},
		"max_file_path_length": []any{
			map[string]any{
				"max_file_path_length": 300,
			},
		},
		"file_extension_restriction": []any{
			map[string]any{
				"restricted_file_extensions": []any{".exe", ".bat", ".sh"},
			},
		},
	}

	input := []any{rulesMap}

	// Expand to GitHub API format
	expandedRules := expandRules(input, false)

	if len(expandedRules) != 4 {
		t.Fatalf("Expected 4 expanded rules for complete push ruleset, got %d", len(expandedRules))
	}

	// Verify we have all expected push rule types
	ruleTypes := make(map[string]bool)
	for _, rule := range expandedRules {
		ruleTypes[rule.Type] = true
	}

	expectedPushRules := []string{"file_path_restriction", "max_file_size", "max_file_path_length", "file_extension_restriction"}
	for _, expectedType := range expectedPushRules {
		if !ruleTypes[expectedType] {
			t.Errorf("Expected push rule type %s not found in expanded rules", expectedType)
		}
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
	if maxFileSizeRules[0]["max_file_size"] != int64(5242880) {
		t.Errorf("Expected max_file_size to be 5242880, got %v", maxFileSizeRules[0]["max_file_size"])
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

func TestAllPushRulesWithUnknownRules(t *testing.T) {
	// Test that push rules work correctly even when unknown rules are present
	unknownParams := map[string]any{
		"some_copilot_parameter": "some_value",
	}
	unknownParamsJSON, _ := json.Marshal(unknownParams)
	unknownParamsRaw := json.RawMessage(unknownParamsJSON)

	maxPathLengthParams := map[string]any{
		"max_file_path_length": 100,
	}
	maxPathLengthParamsJSON, _ := json.Marshal(maxPathLengthParams)
	maxPathLengthParamsRaw := json.RawMessage(maxPathLengthParamsJSON)

	rules := []*github.RepositoryRule{
		{
			Type:       "max_file_path_length",
			Parameters: &maxPathLengthParamsRaw,
		},
		{
			Type:       "unknown_copilot_rule",
			Parameters: &unknownParamsRaw,
		},
		{
			Type: "creation",
		},
	}

	// This should not panic or fail, even with unknown rule types mixed in
	result := flattenRules(rules, false)

	if len(result) != 1 {
		t.Fatalf("Expected 1 element in result, got %d", len(result))
	}

	rulesMap := result[0].(map[string]any)

	// Should contain the known rules
	if !rulesMap["creation"].(bool) {
		t.Error("Expected creation rule to be true")
	}

	maxFilePathLengthRules := rulesMap["max_file_path_length"].([]map[string]any)
	if len(maxFilePathLengthRules) != 1 {
		t.Fatalf("Expected 1 max_file_path_length rule, got %d", len(maxFilePathLengthRules))
	}

	if maxFilePathLengthRules[0]["max_file_path_length"] != 100 {
		t.Errorf("Expected max_file_path_length to be 100, got %v", maxFilePathLengthRules[0]["max_file_path_length"])
	}

	// Should NOT contain the unknown rule type
	if _, exists := rulesMap["unknown_copilot_rule"]; exists {
		t.Error("Unknown rule type should not appear in flattened rules to avoid causing diffs")
	}
}

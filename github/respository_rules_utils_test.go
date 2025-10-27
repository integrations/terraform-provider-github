package github

import (
	"encoding/json"
	"testing"

	"github.com/google/go-github/v67/github"
)

func TestFlattenRulesHandlesUnknownTypes(t *testing.T) {
	// Create some test rules including an unknown type
	unknownParams := map[string]interface{}{
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

	rulesMap := result[0].(map[string]interface{})

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
	params := map[string]interface{}{
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

	rulesMap := result[0].(map[string]interface{})
	maxFileSizeRules := rulesMap["max_file_size"].([]map[string]interface{})

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
	params := map[string]interface{}{
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

	rulesMap := result[0].(map[string]interface{})
	fileExtensionRules := rulesMap["file_extension_restriction"].([]map[string]interface{})

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

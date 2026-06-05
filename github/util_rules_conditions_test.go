package github

import (
	"testing"

	"github.com/google/go-github/v88/github"
)

func TestExpandRepositoryPropertyConditions_SingleInclude(t *testing.T) {
	input := []any{
		map[string]any{
			"include": []any{
				map[string]any{
					"name":            "env",
					"source":          "custom",
					"property_values": []any{"prod"},
				},
			},
			"exclude": []any{},
		},
	}

	result := expandRepositoryPropertyConditions(input)

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	if len(result.Include) != 1 {
		t.Fatalf("Expected 1 include property, got %d", len(result.Include))
	}

	if len(result.Exclude) != 0 {
		t.Fatalf("Expected 0 exclude properties, got %d", len(result.Exclude))
	}

	prop := result.Include[0]
	if prop.Name != "env" {
		t.Errorf("Expected name to be 'env', got %s", prop.Name)
	}
	if prop.Source == nil || *prop.Source != "custom" {
		t.Errorf("Expected source to be 'custom', got %v", prop.Source)
	}
	if len(prop.PropertyValues) != 1 || prop.PropertyValues[0] != "prod" {
		t.Errorf("Expected property_values to be ['prod'], got %v", prop.PropertyValues)
	}
}

func TestExpandRepositoryPropertyConditions_IncludeAndExclude(t *testing.T) {
	input := []any{
		map[string]any{
			"include": []any{
				map[string]any{
					"name":            "env",
					"source":          "custom",
					"property_values": []any{"prod"},
				},
			},
			"exclude": []any{
				map[string]any{
					"name":            "tier",
					"source":          "system",
					"property_values": []any{"free"},
				},
			},
		},
	}

	result := expandRepositoryPropertyConditions(input)

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	if len(result.Include) != 1 {
		t.Fatalf("Expected 1 include property, got %d", len(result.Include))
	}

	if len(result.Exclude) != 1 {
		t.Fatalf("Expected 1 exclude property, got %d", len(result.Exclude))
	}

	includeProp := result.Include[0]
	if includeProp.Name != "env" {
		t.Errorf("Expected include name to be 'env', got %s", includeProp.Name)
	}
	if includeProp.Source == nil || *includeProp.Source != "custom" {
		t.Errorf("Expected include source to be 'custom', got %v", includeProp.Source)
	}

	excludeProp := result.Exclude[0]
	if excludeProp.Name != "tier" {
		t.Errorf("Expected exclude name to be 'tier', got %s", excludeProp.Name)
	}
	if excludeProp.Source == nil || *excludeProp.Source != "system" {
		t.Errorf("Expected exclude source to be 'system', got %v", excludeProp.Source)
	}
}

func TestExpandRepositoryPropertyConditions_MultipleValues(t *testing.T) {
	input := []any{
		map[string]any{
			"include": []any{
				map[string]any{
					"name":            "env",
					"source":          "custom",
					"property_values": []any{"prod", "staging", "dev"},
				},
			},
			"exclude": []any{},
		},
	}

	result := expandRepositoryPropertyConditions(input)

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	if len(result.Include) != 1 {
		t.Fatalf("Expected 1 include property, got %d", len(result.Include))
	}

	prop := result.Include[0]
	if len(prop.PropertyValues) != 3 {
		t.Fatalf("Expected 3 property values, got %d", len(prop.PropertyValues))
	}

	expectedValues := []string{"prod", "staging", "dev"}
	for i, expected := range expectedValues {
		if prop.PropertyValues[i] != expected {
			t.Errorf("Expected property_values[%d] to be '%s', got '%s'", i, expected, prop.PropertyValues[i])
		}
	}
}

func TestExpandRepositoryPropertyConditions_MultipleProperties(t *testing.T) {
	input := []any{
		map[string]any{
			"include": []any{
				map[string]any{
					"name":            "env",
					"source":          "custom",
					"property_values": []any{"prod"},
				},
				map[string]any{
					"name":            "tier",
					"source":          "system",
					"property_values": []any{"premium", "enterprise"},
				},
			},
			"exclude": []any{},
		},
	}

	result := expandRepositoryPropertyConditions(input)

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	if len(result.Include) != 2 {
		t.Fatalf("Expected 2 include properties, got %d", len(result.Include))
	}

	// Check first property
	if result.Include[0].Name != "env" {
		t.Errorf("Expected first property name to be 'env', got %s", result.Include[0].Name)
	}
	if len(result.Include[0].PropertyValues) != 1 {
		t.Errorf("Expected first property to have 1 value, got %d", len(result.Include[0].PropertyValues))
	}

	// Check second property
	if result.Include[1].Name != "tier" {
		t.Errorf("Expected second property name to be 'tier', got %s", result.Include[1].Name)
	}
	if len(result.Include[1].PropertyValues) != 2 {
		t.Errorf("Expected second property to have 2 values, got %d", len(result.Include[1].PropertyValues))
	}
}

func TestExpandRepositoryPropertyConditions_NilElements(t *testing.T) {
	input := []any{
		map[string]any{
			"include": []any{
				map[string]any{
					"name":            "env",
					"source":          "custom",
					"property_values": []any{"prod"},
				},
				nil,
				map[string]any{
					"name":            "tier",
					"source":          "system",
					"property_values": []any{"premium"},
				},
			},
			"exclude": []any{},
		},
	}

	result := expandRepositoryPropertyConditions(input)

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	// Nil element should be skipped, so we should have 2 properties
	if len(result.Include) != 2 {
		t.Fatalf("Expected 2 include properties (nil skipped), got %d", len(result.Include))
	}

	if result.Include[0].Name != "env" {
		t.Errorf("Expected first property name to be 'env', got %s", result.Include[0].Name)
	}
	if result.Include[1].Name != "tier" {
		t.Errorf("Expected second property name to be 'tier', got %s", result.Include[1].Name)
	}
}

func TestExpandRepositoryPropertyConditions_NilPropertyValues(t *testing.T) {
	input := []any{
		map[string]any{
			"include": []any{
				map[string]any{
					"name":            "env",
					"source":          "custom",
					"property_values": []any{"prod", nil, "staging"},
				},
			},
			"exclude": []any{},
		},
	}

	result := expandRepositoryPropertyConditions(input)

	if result == nil {
		t.Fatal("Expected result to not be nil")
	}

	if len(result.Include) != 1 {
		t.Fatalf("Expected 1 include property, got %d", len(result.Include))
	}

	prop := result.Include[0]
	// Nil value should be skipped, so we should have 2 values
	if len(prop.PropertyValues) != 2 {
		t.Fatalf("Expected 2 property values (nil skipped), got %d", len(prop.PropertyValues))
	}

	if prop.PropertyValues[0] != "prod" {
		t.Errorf("Expected first value to be 'prod', got '%s'", prop.PropertyValues[0])
	}
	if prop.PropertyValues[1] != "staging" {
		t.Errorf("Expected second value to be 'staging', got '%s'", prop.PropertyValues[1])
	}
}

func TestExpandConditions(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName           string
		input              []any
		org                bool
		wantNil            bool
		wantRefName        *github.RepositoryRulesetRefConditionParameters
		wantRepositoryName *github.RepositoryRulesetRepositoryNamesConditionParameters
	}{
		{
			testName: "returns_nil_for_empty_input",
			input:    []any{},
			wantNil:  true,
		},
		{
			testName: "returns_nil_for_nil_first_element",
			input:    []any{nil},
			wantNil:  true,
		},
		{
			testName: "returns_empty_conditions_for_empty_conditions_map",
			input:    []any{map[string]any{}},
		},
		{
			testName: "returns_empty_conditions_for_empty_ref_name",
			input:    []any{map[string]any{"ref_name": []any{}}},
		},
		{
			testName:    "returns_empty_conditions_for_empty_ref_name_arrays",
			input:       []any{map[string]any{"ref_name": []any{map[string]any{"include": []any{}, "exclude": []any{}}}}},
			wantRefName: &github.RepositoryRulesetRefConditionParameters{Include: []string{}, Exclude: []string{}},
		},
		{
			testName: "returns_empty_conditions_for_nil_ref_name_arrays",
			input:    []any{map[string]any{"ref_name": []any{nil}}},
		},
		{
			testName: "returns_empty_conditions_for_nil_repository_name_arrays",
			input:    []any{map[string]any{"repository_name": []any{nil}}},
			org:      true,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got := expandConditions(d.input, d.org)

			if d.wantNil {
				if got != nil {
					t.Fatalf("got %+v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("got nil, want conditions")
			}
			if d.wantRefName == nil && got.RefName != nil {
				t.Fatalf("got RefName %+v, want nil", got.RefName)
			}
			if d.wantRefName != nil {
				if got.RefName == nil {
					t.Fatal("got nil RefName, want ref_name conditions")
				}
				if len(got.RefName.Include) != len(d.wantRefName.Include) {
					t.Fatalf("got Include %v, want %v", got.RefName.Include, d.wantRefName.Include)
				}
				for i, v := range got.RefName.Include {
					if v != d.wantRefName.Include[i] {
						t.Fatalf("got Include %v, want %v", got.RefName.Include, d.wantRefName.Include)
					}
				}
				if len(got.RefName.Exclude) != len(d.wantRefName.Exclude) {
					t.Fatalf("got Exclude %v, want %v", got.RefName.Exclude, d.wantRefName.Exclude)
				}
				for i, v := range got.RefName.Exclude {
					if v != d.wantRefName.Exclude[i] {
						t.Fatalf("got Exclude %v, want %v", got.RefName.Exclude, d.wantRefName.Exclude)
					}
				}
			}
			if d.wantRepositoryName == nil && got.RepositoryName != nil {
				t.Fatalf("got RepositoryName %+v, want nil", got.RepositoryName)
			}
		})
	}
}

func TestFlattenRulesetRepositoryPropertyTargetParameters(t *testing.T) {
	input := []*github.RepositoryRulesetRepositoryPropertyTargetParameters{
		{
			Name:           "env",
			Source:         new("custom"),
			PropertyValues: []string{"prod", "staging"},
		},
		{
			Name:           "tier",
			Source:         new("system"),
			PropertyValues: []string{"premium"},
		},
	}

	result := flattenRulesetRepositoryPropertyTargetParameters(input)

	if len(result) != 2 {
		t.Fatalf("Expected 2 properties, got %d", len(result))
	}

	// Check first property
	if result[0]["name"] != "env" {
		t.Errorf("Expected first property name to be 'env', got %v", result[0]["name"])
	}
	if result[0]["source"] != "custom" {
		t.Errorf("Expected first property source to be 'custom', got %v", result[0]["source"])
	}
	values := result[0]["property_values"].([]string)
	if len(values) != 2 || values[0] != "prod" || values[1] != "staging" {
		t.Errorf("Expected first property values to be ['prod', 'staging'], got %v", values)
	}

	// Check second property
	if result[1]["name"] != "tier" {
		t.Errorf("Expected second property name to be 'tier', got %v", result[1]["name"])
	}
}

func TestFlattenRulesetRepositoryPropertyTargetParameters_EmptySource(t *testing.T) {
	input := []*github.RepositoryRulesetRepositoryPropertyTargetParameters{
		{
			Name:           "env",
			Source:         new(""),
			PropertyValues: []string{"prod"},
		},
	}

	result := flattenRulesetRepositoryPropertyTargetParameters(input)

	if len(result) != 1 {
		t.Fatalf("Expected 1 property, got %d", len(result))
	}

	// Empty source should default to "custom"
	if result[0]["source"] != "custom" {
		t.Errorf("Expected source to default to 'custom', got %v", result[0]["source"])
	}
}

func TestRoundTripRepositoryPropertyConditions(t *testing.T) {
	input := []any{
		map[string]any{
			"include": []any{
				map[string]any{
					"name":            "env",
					"source":          "custom",
					"property_values": []any{"prod", "staging"},
				},
				map[string]any{
					"name":            "tier",
					"source":          "system",
					"property_values": []any{"premium"},
				},
			},
			"exclude": []any{
				map[string]any{
					"name":            "region",
					"source":          "custom",
					"property_values": []any{"us-west"},
				},
			},
		},
	}

	// Expand
	expanded := expandRepositoryPropertyConditions(input)

	// Flatten
	flattenedInclude := flattenRulesetRepositoryPropertyTargetParameters(expanded.Include)
	flattenedExclude := flattenRulesetRepositoryPropertyTargetParameters(expanded.Exclude)

	// Verify include
	if len(flattenedInclude) != 2 {
		t.Fatalf("Expected 2 include properties after round trip, got %d", len(flattenedInclude))
	}

	if flattenedInclude[0]["name"] != "env" {
		t.Errorf("Expected first include name to be 'env', got %v", flattenedInclude[0]["name"])
	}
	if flattenedInclude[0]["source"] != "custom" {
		t.Errorf("Expected first include source to be 'custom', got %v", flattenedInclude[0]["source"])
	}
	includeValues := flattenedInclude[0]["property_values"].([]string)
	if len(includeValues) != 2 || includeValues[0] != "prod" || includeValues[1] != "staging" {
		t.Errorf("Expected first include values to be ['prod', 'staging'], got %v", includeValues)
	}

	if flattenedInclude[1]["name"] != "tier" {
		t.Errorf("Expected second include name to be 'tier', got %v", flattenedInclude[1]["name"])
	}

	// Verify exclude
	if len(flattenedExclude) != 1 {
		t.Fatalf("Expected 1 exclude property after round trip, got %d", len(flattenedExclude))
	}

	if flattenedExclude[0]["name"] != "region" {
		t.Errorf("Expected exclude name to be 'region', got %v", flattenedExclude[0]["name"])
	}
	excludeValues := flattenedExclude[0]["property_values"].([]string)
	if len(excludeValues) != 1 || excludeValues[0] != "us-west" {
		t.Errorf("Expected exclude values to be ['us-west'], got %v", excludeValues)
	}
}

func TestFlattenRulesetRepositoryPropertyTargetParameters_Empty(t *testing.T) {
	// Test nil input
	result := flattenRulesetRepositoryPropertyTargetParameters(nil)
	if len(result) != 0 {
		t.Errorf("Expected empty slice for nil input, got %v", result)
	}

	// Test empty slice input
	result = flattenRulesetRepositoryPropertyTargetParameters([]*github.RepositoryRulesetRepositoryPropertyTargetParameters{})
	if len(result) != 0 {
		t.Errorf("Expected empty slice for empty input, got %v", result)
	}
}

func TestFlattenRulesetRepositoryPropertyTargetParameters_SingleProperty(t *testing.T) {
	input := []*github.RepositoryRulesetRepositoryPropertyTargetParameters{
		{
			Name:           "env",
			Source:         new("system"),
			PropertyValues: []string{"prod", "staging"},
		},
	}

	result := flattenRulesetRepositoryPropertyTargetParameters(input)

	if len(result) != 1 {
		t.Fatalf("Expected 1 property, got %d", len(result))
	}

	if result[0]["name"] != "env" {
		t.Errorf("Expected name to be 'env', got %v", result[0]["name"])
	}

	if result[0]["source"] != "system" {
		t.Errorf("Expected source to be 'system', got %v", result[0]["source"])
	}

	values := result[0]["property_values"].([]string)
	if len(values) != 2 || values[0] != "prod" || values[1] != "staging" {
		t.Errorf("Expected property_values to be ['prod', 'staging'], got %v", values)
	}
}

func TestFlattenRulesetRepositoryPropertyTargetParameters_NilSource(t *testing.T) {
	input := []*github.RepositoryRulesetRepositoryPropertyTargetParameters{
		{
			Name:           "env",
			Source:         nil,
			PropertyValues: []string{"prod"},
		},
	}

	result := flattenRulesetRepositoryPropertyTargetParameters(input)

	if len(result) != 1 {
		t.Fatalf("Expected 1 property, got %d", len(result))
	}

	// Nil source should default to "custom"
	if result[0]["source"] != "custom" {
		t.Errorf("Expected source to default to 'custom' for nil source, got %v", result[0]["source"])
	}
}

func TestFlattenRulesetRepositoryPropertyTargetParameters_EmptyPropertyValues(t *testing.T) {
	input := []*github.RepositoryRulesetRepositoryPropertyTargetParameters{
		{
			Name:           "env",
			Source:         new("custom"),
			PropertyValues: []string{},
		},
	}

	result := flattenRulesetRepositoryPropertyTargetParameters(input)

	if len(result) != 1 {
		t.Fatalf("Expected 1 property, got %d", len(result))
	}

	values := result[0]["property_values"].([]string)
	if len(values) != 0 {
		t.Errorf("Expected property_values to be empty array, got %v", values)
	}
}

func TestFlattenRulesetRepositoryPropertyTargetParameters_NilPropertyValues(t *testing.T) {
	input := []*github.RepositoryRulesetRepositoryPropertyTargetParameters{
		{
			Name:           "env",
			Source:         new("custom"),
			PropertyValues: nil,
		},
	}

	result := flattenRulesetRepositoryPropertyTargetParameters(input)

	if len(result) != 1 {
		t.Fatalf("Expected 1 property, got %d", len(result))
	}

	// Nil PropertyValues should be preserved in the map
	values := result[0]["property_values"]
	if values != nil {
		if valSlice, ok := values.([]string); ok && len(valSlice) != 0 {
			t.Errorf("Expected property_values to be nil or empty, got %v", values)
		}
	}
}

package github

import (
	"testing"

	"github.com/google/go-github/v83/github"
)

func Test_validateConditionsFieldForPushTarget(t *testing.T) {
	tests := []struct {
		name        string
		conditions  map[string]any
		expectError bool
		errorMsg    string
	}{
		{
			name: "valid push target without ref_name",
			conditions: map[string]any{
				"repository_name": []any{map[string]any{"include": []any{"~ALL"}, "exclude": []any{}}},
			},
			expectError: false,
		},
		{
			name:        "valid push target with nil ref_name",
			conditions:  map[string]any{"ref_name": nil},
			expectError: false,
		},
		{
			name:        "valid push target with empty ref_name slice",
			conditions:  map[string]any{"ref_name": []any{}},
			expectError: false,
		},
		{
			name: "invalid push target with ref_name set",
			conditions: map[string]any{
				"ref_name": []any{map[string]any{"include": []any{"~ALL"}, "exclude": []any{}}},
			},
			expectError: true,
			errorMsg:    "ref_name must not be set for push target",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConditionsFieldForPushTarget(t.Context(), tt.conditions)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("expected error %q, got %q", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			}
		})
	}
}

func Test_validateRepositoryRulesetConditionsFieldForBranchAndTagTargets(t *testing.T) {
	tests := []struct {
		name        string
		target      github.RulesetTarget
		conditions  map[string]any
		expectError bool
		errorMsg    string
	}{
		{
			name:   "valid branch target with ref_name",
			target: github.RulesetTargetBranch,
			conditions: map[string]any{
				"ref_name": []any{map[string]any{"include": []any{"~DEFAULT_BRANCH"}, "exclude": []any{}}},
			},
			expectError: false,
		},
		{
			name:   "valid tag target with ref_name",
			target: github.RulesetTargetTag,
			conditions: map[string]any{
				"ref_name": []any{map[string]any{"include": []any{"v*"}, "exclude": []any{}}},
			},
			expectError: false,
		},
		{
			name:        "invalid branch target without ref_name",
			target:      github.RulesetTargetBranch,
			conditions:  map[string]any{},
			expectError: true,
			errorMsg:    "ref_name must be set for branch target",
		},
		{
			name:        "invalid tag target without ref_name",
			target:      github.RulesetTargetTag,
			conditions:  map[string]any{},
			expectError: true,
			errorMsg:    "ref_name must be set for tag target",
		},
		{
			name:        "invalid branch target with nil ref_name",
			target:      github.RulesetTargetBranch,
			conditions:  map[string]any{"ref_name": nil},
			expectError: true,
			errorMsg:    "ref_name must be set for branch target",
		},
		{
			name:        "invalid tag target with empty ref_name slice",
			target:      github.RulesetTargetTag,
			conditions:  map[string]any{"ref_name": []any{}},
			expectError: true,
			errorMsg:    "ref_name must be set for tag target",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConditionsFieldForBranchAndTagTargets(t.Context(), tt.target, tt.conditions, false)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("expected error %q, got %q", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			}
		})
	}
}

func Test_validateConditionsFieldForBranchAndTagTargets(t *testing.T) {
	tests := []struct {
		name        string
		target      github.RulesetTarget
		conditions  map[string]any
		expectError bool
		errorMsg    string
	}{
		{
			name:   "valid branch target with ref_name and repository_name",
			target: github.RulesetTargetBranch,
			conditions: map[string]any{
				"ref_name":        []any{map[string]any{"include": []any{"~DEFAULT_BRANCH"}, "exclude": []any{}}},
				"repository_name": []any{map[string]any{"include": []any{"~ALL"}, "exclude": []any{}}},
			},
			expectError: false,
		},
		{
			name:   "valid tag target with ref_name and repository_id",
			target: github.RulesetTargetTag,
			conditions: map[string]any{
				"ref_name":      []any{map[string]any{"include": []any{"v*"}, "exclude": []any{}}},
				"repository_id": []any{123, 456},
			},
			expectError: false,
		},
		{
			name:   "invalid branch target without ref_name",
			target: github.RulesetTargetBranch,
			conditions: map[string]any{
				"repository_name": []any{map[string]any{"include": []any{"~ALL"}, "exclude": []any{}}},
			},
			expectError: true,
			errorMsg:    "ref_name must be set for branch target",
		},
		{
			name:   "invalid branch target without repository_name or repository_id",
			target: github.RulesetTargetBranch,
			conditions: map[string]any{
				"ref_name": []any{map[string]any{"include": []any{"~DEFAULT_BRANCH"}, "exclude": []any{}}},
			},
			expectError: true,
			errorMsg:    "either repository_name or repository_id must be set for branch target",
		},
		{
			name:   "invalid tag target with nil repository_name and repository_id",
			target: github.RulesetTargetTag,
			conditions: map[string]any{
				"ref_name":        []any{map[string]any{"include": []any{"v*"}, "exclude": []any{}}},
				"repository_name": nil,
				"repository_id":   nil,
			},
			expectError: true,
			errorMsg:    "either repository_name or repository_id must be set for tag target",
		},
		{
			name:   "invalid branch target with empty repository_name and repository_id slices",
			target: github.RulesetTargetBranch,
			conditions: map[string]any{
				"ref_name":        []any{map[string]any{"include": []any{"~DEFAULT_BRANCH"}, "exclude": []any{}}},
				"repository_name": []any{},
				"repository_id":   []any{},
			},
			expectError: true,
			errorMsg:    "either repository_name or repository_id must be set for branch target",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConditionsFieldForBranchAndTagTargets(t.Context(), tt.target, tt.conditions, true)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("expected error %q, got %q", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			}
		})
	}
}

func Test_ruleListsDoNotOverlap(t *testing.T) {
	for _, pushRule := range pushOnlyRules {
		for _, branchTagRule := range branchTagOnlyRules {
			if pushRule == branchTagRule {
				t.Errorf("rule %q appears in both pushOnlyRules and branchTagOnlyRules", pushRule)
			}
		}
	}
}

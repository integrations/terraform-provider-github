package github

import (
	"strings"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func Test_validateConditionsFieldForPushTarget(t *testing.T) {
	t.Parallel()

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
			t.Parallel()

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
	t.Parallel()

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
			t.Parallel()

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
	t.Parallel()

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

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
	t.Parallel()

	for _, pushRule := range pushOnlyRules {
		for _, branchTagRule := range branchTagOnlyRules {
			if pushRule == branchTagRule {
				t.Errorf("rule %q appears in both pushOnlyRules and branchTagOnlyRules", pushRule)
			}
		}
	}
}

func Test_validateRulesetConditions_emptyBlockDiff(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name    string
		res     *schema.Resource
		raw     map[string]any
		wantErr string
	}{
		{
			name: "organization branch ruleset with empty conditions block requires ref_name",
			res:  resourceGithubOrganizationRuleset(),
			raw: map[string]any{
				"name":        "test",
				"target":      "branch",
				"enforcement": "active",
				"conditions":  []any{map[string]any{}},
				"rules":       []any{map[string]any{"creation": true}},
			},
			wantErr: "ref_name must be set for branch target",
		},
		{
			name: "organization tag ruleset with empty conditions block requires ref_name",
			res:  resourceGithubOrganizationRuleset(),
			raw: map[string]any{
				"name":        "test",
				"target":      "tag",
				"enforcement": "active",
				"conditions":  []any{map[string]any{}},
				"rules":       []any{map[string]any{"creation": true}},
			},
			wantErr: "ref_name must be set for tag target",
		},
		{
			name: "organization push ruleset with empty conditions block succeeds",
			res:  resourceGithubOrganizationRuleset(),
			raw: map[string]any{
				"name":        "test",
				"target":      "push",
				"enforcement": "active",
				"conditions":  []any{map[string]any{}},
				"rules":       []any{map[string]any{"max_file_size": []any{map[string]any{"max_file_size": 10}}}},
			},
		},
		{
			name: "organization branch ruleset with omitted conditions skips condition validation",
			res:  resourceGithubOrganizationRuleset(),
			raw: map[string]any{
				"name":        "test",
				"target":      "branch",
				"enforcement": "active",
				"rules":       []any{map[string]any{"creation": true}},
			},
		},
		{
			name: "repository branch ruleset with empty conditions block requires ref_name",
			res:  resourceGithubRepositoryRuleset(),
			raw: map[string]any{
				"name":        "test",
				"repository":  "my-repo",
				"target":      "branch",
				"enforcement": "active",
				"conditions":  []any{map[string]any{}},
				"rules":       []any{map[string]any{"creation": true}},
			},
			wantErr: "ref_name must be set for branch target",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := tt.res.Diff(t.Context(), nil, terraform.NewResourceConfigRaw(tt.raw), nil)
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("expected error containing %q, got %q", tt.wantErr, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

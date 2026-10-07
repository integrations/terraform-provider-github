package github

import (
	"testing"

	"github.com/google/go-github/v92/github"
)

func Test_validateConditionsFieldForRefLessTargets(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		target      github.RulesetTarget
		conditions  map[string]any
		expectError bool
		errorMsg    string
	}{
		{
			name:   "valid push target without ref_name",
			target: github.RulesetTargetPush,
			conditions: map[string]any{
				"repository_name": []any{map[string]any{"include": []any{"~ALL"}, "exclude": []any{}}},
			},
			expectError: false,
		},
		{
			name:        "valid push target with nil ref_name",
			target:      github.RulesetTargetPush,
			conditions:  map[string]any{"ref_name": nil},
			expectError: false,
		},
		{
			name:        "valid push target with empty ref_name slice",
			target:      github.RulesetTargetPush,
			conditions:  map[string]any{"ref_name": []any{}},
			expectError: false,
		},
		{
			name:   "invalid push target with ref_name set",
			target: github.RulesetTargetPush,
			conditions: map[string]any{
				"ref_name": []any{map[string]any{"include": []any{"~ALL"}, "exclude": []any{}}},
			},
			expectError: true,
			errorMsg:    "ref_name must not be set for push target",
		},
		{
			name:   "valid repository target without ref_name",
			target: github.RulesetTargetRepository,
			conditions: map[string]any{
				"organization_name": []any{map[string]any{"include": []any{"~ALL"}, "exclude": []any{}}},
			},
			expectError: false,
		},
		{
			name:   "invalid repository target with ref_name set",
			target: github.RulesetTargetRepository,
			conditions: map[string]any{
				"ref_name": []any{map[string]any{"include": []any{"~ALL"}, "exclude": []any{}}},
			},
			expectError: true,
			errorMsg:    "ref_name must not be set for repository target",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := validateConditionsFieldForRefLessTargets(t.Context(), tt.target, tt.conditions)
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
			name:   "invalid branch target with repository_name but without ref_name",
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

			err := validateConditionsFieldForBranchAndTagTargets(t.Context(), tt.target, tt.conditions)
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

	lists := map[string][]github.RepositoryRuleType{
		"branchTagOnlyRules":  branchTagOnlyRules,
		"pushOnlyRules":       pushOnlyRules,
		"repositoryOnlyRules": repositoryOnlyRules,
	}

	seen := make(map[github.RepositoryRuleType]string)
	for listName, rules := range lists {
		for _, rule := range rules {
			if other, ok := seen[rule]; ok {
				t.Errorf("rule %q appears in both %s and %s", rule, other, listName)
				continue
			}
			seen[rule] = listName
		}
	}
}

func Test_everyTargetHasAnAllowedRuleList(t *testing.T) {
	t.Parallel()

	targets := []github.RulesetTarget{
		github.RulesetTargetBranch,
		github.RulesetTargetTag,
		github.RulesetTargetPush,
		github.RulesetTargetRepository,
	}

	for _, target := range targets {
		if _, ok := rulesAllowedByTarget[target]; !ok {
			t.Errorf("target %q has no entry in rulesAllowedByTarget", target)
		}
	}
}

func Test_validateBypassActorsForTarget(t *testing.T) {
	t.Parallel()

	actor := func(actorType, bypassMode string) any {
		return map[string]any{"actor_type": actorType, "bypass_mode": bypassMode}
	}

	tests := []struct {
		name         string
		target       github.RulesetTarget
		bypassActors []any
		errorMsg     string
	}{
		{
			name:         "no bypass actors",
			target:       github.RulesetTargetRepository,
			bypassActors: []any{},
		},
		{
			name:         "always mode is valid for any target",
			target:       github.RulesetTargetRepository,
			bypassActors: []any{actor("EnterpriseOwner", "always")},
		},
		{
			name:         "exempt mode is valid for any target",
			target:       github.RulesetTargetPush,
			bypassActors: []any{actor(bypassActorTypeDeployKey, "exempt")},
		},
		{
			name:         "pull_request is valid for the branch target",
			target:       github.RulesetTargetBranch,
			bypassActors: []any{actor("Team", bypassModePullRequest)},
		},
		{
			name:         "pull_request is rejected for the tag target",
			target:       github.RulesetTargetTag,
			bypassActors: []any{actor("Team", bypassModePullRequest)},
			errorMsg:     `bypass_actors.0: bypass_mode "pull_request" is only valid for the branch target, got "tag"`,
		},
		{
			name:         "pull_request is rejected for the push target",
			target:       github.RulesetTargetPush,
			bypassActors: []any{actor("Team", bypassModePullRequest)},
			errorMsg:     `bypass_actors.0: bypass_mode "pull_request" is only valid for the branch target, got "push"`,
		},
		{
			name:         "pull_request is rejected for the repository target",
			target:       github.RulesetTargetRepository,
			bypassActors: []any{actor("Team", bypassModePullRequest)},
			errorMsg:     `bypass_actors.0: bypass_mode "pull_request" is only valid for the branch target, got "repository"`,
		},
		{
			name:         "pull_request is rejected for deploy keys on the branch target",
			target:       github.RulesetTargetBranch,
			bypassActors: []any{actor(bypassActorTypeDeployKey, bypassModePullRequest)},
			errorMsg:     `bypass_actors.0: bypass_mode "pull_request" is not valid for the "DeployKey" actor type`,
		},
		{
			name:         "the offending actor index is reported",
			target:       github.RulesetTargetBranch,
			bypassActors: []any{actor("Team", "always"), actor(bypassActorTypeDeployKey, bypassModePullRequest)},
			errorMsg:     `bypass_actors.1: bypass_mode "pull_request" is not valid for the "DeployKey" actor type`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := validateBypassActorsForTarget(t.Context(), tt.target, tt.bypassActors)

			if tt.errorMsg == "" {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("expected error %q, got nil", tt.errorMsg)
			}
			if err.Error() != tt.errorMsg {
				t.Errorf("expected error %q, got %q", tt.errorMsg, err.Error())
			}
		})
	}
}

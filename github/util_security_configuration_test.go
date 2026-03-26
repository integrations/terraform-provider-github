package github

import (
	"testing"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestFlattenDependencyGraphAutosubmitActionOptions(t *testing.T) {
	tests := []struct {
		name   string
		input  *github.DependencyGraphAutosubmitActionOptions
		expect func(t *testing.T, result []any)
	}{
		{
			name:  "returns empty slice when options is nil",
			input: nil,
			expect: func(t *testing.T, result []any) {
				if len(result) != 0 {
					t.Errorf("expected empty slice, got %v", result)
				}
			},
		},
		{
			name:  "omits labeled_runners key when LabeledRunners is nil",
			input: &github.DependencyGraphAutosubmitActionOptions{},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if _, ok := m["labeled_runners"]; ok {
					t.Errorf("labeled_runners should be absent when LabeledRunners is nil")
				}
			},
		},
		{
			name:  "sets labeled_runners when LabeledRunners is non-nil",
			input: &github.DependencyGraphAutosubmitActionOptions{LabeledRunners: github.Ptr(true)},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if m["labeled_runners"] != true {
					t.Errorf("expected labeled_runners true, got %v", m["labeled_runners"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenDependencyGraphAutosubmitActionOptions(tt.input)
			tt.expect(t, result)
		})
	}
}

func TestFlattenCodeScanningOptions(t *testing.T) {
	tests := []struct {
		name   string
		input  *github.CodeScanningOptions
		expect func(t *testing.T, result []any)
	}{
		{
			name:  "returns empty slice when options is nil",
			input: nil,
			expect: func(t *testing.T, result []any) {
				if len(result) != 0 {
					t.Errorf("expected empty slice, got %v", result)
				}
			},
		},
		{
			name:  "omits allow_advanced key when AllowAdvanced is nil",
			input: &github.CodeScanningOptions{},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if _, ok := m["allow_advanced"]; ok {
					t.Errorf("allow_advanced should be absent when AllowAdvanced is nil")
				}
			},
		},
		{
			name:  "sets allow_advanced when AllowAdvanced is true",
			input: &github.CodeScanningOptions{AllowAdvanced: github.Ptr(true)},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if m["allow_advanced"] != true {
					t.Errorf("expected allow_advanced true, got %v", m["allow_advanced"])
				}
			},
		},
		{
			name:  "sets allow_advanced when AllowAdvanced is false",
			input: &github.CodeScanningOptions{AllowAdvanced: github.Ptr(false)},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if m["allow_advanced"] != false {
					t.Errorf("expected allow_advanced false, got %v", m["allow_advanced"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenCodeScanningOptions(tt.input)
			tt.expect(t, result)
		})
	}
}

func TestFlattenSecretScanningDelegatedBypassOptions(t *testing.T) {
	tests := []struct {
		name   string
		input  *github.SecretScanningDelegatedBypassOptions
		expect func(t *testing.T, result []any)
	}{
		{
			name:  "returns empty slice when options is nil",
			input: nil,
			expect: func(t *testing.T, result []any) {
				if len(result) != 0 {
					t.Errorf("expected empty slice, got %v", result)
				}
			},
		},
		{
			name:  "omits reviewers key when Reviewers is nil",
			input: &github.SecretScanningDelegatedBypassOptions{},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if _, ok := m["reviewers"]; ok {
					t.Errorf("reviewers should be absent when Reviewers is nil")
				}
			},
		},
		{
			name: "sets reviewers when Reviewers is populated",
			input: &github.SecretScanningDelegatedBypassOptions{
				Reviewers: []*github.BypassReviewer{
					{ReviewerID: 42, ReviewerType: "TEAM"},
					{ReviewerID: 99, ReviewerType: "ROLE"},
				},
			},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				reviewers, ok := m["reviewers"].([]any)
				if !ok {
					t.Fatalf("expected reviewers to be []any, got %T", m["reviewers"])
				}
				if len(reviewers) != 2 {
					t.Fatalf("expected 2 reviewers, got %d", len(reviewers))
				}
				first := reviewers[0].(map[string]any)
				if first["reviewer_id"] != int64(42) {
					t.Errorf("expected reviewer_id 42, got %v", first["reviewer_id"])
				}
				if first["reviewer_type"] != "TEAM" {
					t.Errorf("expected reviewer_type TEAM, got %v", first["reviewer_type"])
				}
				second := reviewers[1].(map[string]any)
				if second["reviewer_id"] != int64(99) {
					t.Errorf("expected reviewer_id 99, got %v", second["reviewer_id"])
				}
				if second["reviewer_type"] != "ROLE" {
					t.Errorf("expected reviewer_type ROLE, got %v", second["reviewer_type"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenSecretScanningDelegatedBypassOptions(tt.input)
			tt.expect(t, result)
		})
	}
}

func TestFlattenCodeScanningDefaultSetupOptions(t *testing.T) {
	tests := []struct {
		name   string
		input  *github.CodeScanningDefaultSetupOptions
		expect func(t *testing.T, result []any)
	}{
		{
			name:  "returns empty slice when options is nil",
			input: nil,
			expect: func(t *testing.T, result []any) {
				if len(result) != 0 {
					t.Errorf("expected empty slice, got %v", result)
				}
			},
		},
		{
			name:  "omits runner_type key when RunnerType is empty string",
			input: &github.CodeScanningDefaultSetupOptions{RunnerType: ""},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if _, ok := m["runner_type"]; ok {
					t.Errorf("runner_type should be absent when RunnerType is empty, got %q", m["runner_type"])
				}
			},
		},
		{
			name:  "sets runner_type when RunnerType is non-empty",
			input: &github.CodeScanningDefaultSetupOptions{RunnerType: "standard"},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if m["runner_type"] != "standard" {
					t.Errorf("expected runner_type %q, got %q", "standard", m["runner_type"])
				}
			},
		},
		{
			name: "sets runner_label when RunnerLabel is non-nil",
			input: &github.CodeScanningDefaultSetupOptions{
				RunnerType:  "labeled",
				RunnerLabel: github.Ptr("my-runner"),
			},
			expect: func(t *testing.T, result []any) {
				if len(result) != 1 {
					t.Fatalf("expected 1 element, got %d", len(result))
				}
				m := result[0].(map[string]any)
				if m["runner_label"] != "my-runner" {
					t.Errorf("expected runner_label %q, got %q", "my-runner", m["runner_label"])
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := flattenCodeScanningDefaultSetupOptions(tt.input)
			tt.expect(t, result)
		})
	}
}

func TestExpandCodeSecurityConfigurationCommon(t *testing.T) {
	resourceSchema := resourceGithubOrganizationSecurityConfiguration().Schema

	tests := []struct {
		name   string
		input  map[string]any
		expect func(t *testing.T, config github.CodeSecurityConfiguration)
	}{
		{
			name: "minimal input sets only name",
			input: map[string]any{
				"name": "my-config",
			},
			expect: func(t *testing.T, config github.CodeSecurityConfiguration) {
				if config.Name != "my-config" {
					t.Errorf("expected name %q, got %q", "my-config", config.Name)
				}
				if config.AdvancedSecurity != nil {
					t.Errorf("expected AdvancedSecurity nil, got %v", *config.AdvancedSecurity)
				}
				if config.DependencyGraph != nil {
					t.Errorf("expected DependencyGraph nil, got %v", *config.DependencyGraph)
				}
				if config.Enforcement != nil {
					t.Errorf("expected Enforcement nil, got %v", *config.Enforcement)
				}
			},
		},
		{
			name: "sets all string fields",
			input: map[string]any{
				"name":                                    "full-config",
				"description":                             "A test config",
				"advanced_security":                       "enabled",
				"dependency_graph":                        "enabled",
				"dependency_graph_autosubmit_action":      "enabled",
				"dependabot_alerts":                       "enabled",
				"dependabot_security_updates":             "disabled",
				"code_scanning_default_setup":             "enabled",
				"code_scanning_delegated_alert_dismissal": "not_set",
				"code_security":                           "enabled",
				"secret_scanning":                         "enabled",
				"secret_scanning_push_protection":         "enabled",
				"secret_scanning_validity_checks":         "disabled",
				"secret_scanning_non_provider_patterns":   "not_set",
				"secret_scanning_generic_secrets":         "disabled",
				"secret_scanning_delegated_alert_dismissal": "not_set",
				"secret_protection":                         "enabled",
				"private_vulnerability_reporting":           "enabled",
				"enforcement":                               "enforced",
			},
			expect: func(t *testing.T, config github.CodeSecurityConfiguration) {
				if config.Name != "full-config" {
					t.Errorf("expected name %q, got %q", "full-config", config.Name)
				}
				if config.Description != "A test config" {
					t.Errorf("expected description %q, got %q", "A test config", config.Description)
				}
				if config.GetAdvancedSecurity() != "enabled" {
					t.Errorf("expected AdvancedSecurity %q, got %q", "enabled", config.GetAdvancedSecurity())
				}
				if config.GetDependencyGraph() != "enabled" {
					t.Errorf("expected DependencyGraph %q, got %q", "enabled", config.GetDependencyGraph())
				}
				if config.GetDependabotSecurityUpdates() != "disabled" {
					t.Errorf("expected DependabotSecurityUpdates %q, got %q", "disabled", config.GetDependabotSecurityUpdates())
				}
				if config.GetEnforcement() != "enforced" {
					t.Errorf("expected Enforcement %q, got %q", "enforced", config.GetEnforcement())
				}
				if config.GetSecretScanning() != "enabled" {
					t.Errorf("expected SecretScanning %q, got %q", "enabled", config.GetSecretScanning())
				}
				if config.GetPrivateVulnerabilityReporting() != "enabled" {
					t.Errorf("expected PrivateVulnerabilityReporting %q, got %q", "enabled", config.GetPrivateVulnerabilityReporting())
				}
			},
		},
		{
			name: "sets dependency_graph_autosubmit_action_options",
			input: map[string]any{
				"name": "with-autosubmit-opts",
				"dependency_graph_autosubmit_action_options": []any{
					map[string]any{
						"labeled_runners": true,
					},
				},
			},
			expect: func(t *testing.T, config github.CodeSecurityConfiguration) {
				if config.DependencyGraphAutosubmitActionOptions == nil {
					t.Fatal("expected DependencyGraphAutosubmitActionOptions to be set")
				}
				if !config.DependencyGraphAutosubmitActionOptions.GetLabeledRunners() {
					t.Errorf("expected LabeledRunners true, got false")
				}
			},
		},
		{
			name: "sets code_scanning_default_setup_options with runner_label",
			input: map[string]any{
				"name": "with-setup-opts",
				"code_scanning_default_setup_options": []any{
					map[string]any{
						"runner_type":  "labeled",
						"runner_label": "my-runner",
					},
				},
			},
			expect: func(t *testing.T, config github.CodeSecurityConfiguration) {
				if config.CodeScanningDefaultSetupOptions == nil {
					t.Fatal("expected CodeScanningDefaultSetupOptions to be set")
				}
				if config.CodeScanningDefaultSetupOptions.RunnerType != "labeled" {
					t.Errorf("expected RunnerType %q, got %q", "labeled", config.CodeScanningDefaultSetupOptions.RunnerType)
				}
				if config.CodeScanningDefaultSetupOptions.GetRunnerLabel() != "my-runner" {
					t.Errorf("expected RunnerLabel %q, got %q", "my-runner", config.CodeScanningDefaultSetupOptions.GetRunnerLabel())
				}
			},
		},
		{
			name: "sets code_scanning_options",
			input: map[string]any{
				"name": "with-scan-opts",
				"code_scanning_options": []any{
					map[string]any{
						"allow_advanced": true,
					},
				},
			},
			expect: func(t *testing.T, config github.CodeSecurityConfiguration) {
				if config.CodeScanningOptions == nil {
					t.Fatal("expected CodeScanningOptions to be set")
				}
				if !config.CodeScanningOptions.GetAllowAdvanced() {
					t.Errorf("expected AllowAdvanced true, got false")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, resourceSchema, tt.input)
			result := expandCodeSecurityConfigurationCommon(d)
			tt.expect(t, result)
		})
	}
}

func TestExpandSecretScanningDelegatedBypass(t *testing.T) {
	resourceSchema := resourceGithubOrganizationSecurityConfiguration().Schema

	tests := []struct {
		name   string
		input  map[string]any
		expect func(t *testing.T, config github.CodeSecurityConfiguration)
	}{
		{
			name: "no bypass fields leaves config unchanged",
			input: map[string]any{
				"name": "no-bypass",
			},
			expect: func(t *testing.T, config github.CodeSecurityConfiguration) {
				if config.SecretScanningDelegatedBypass != nil {
					t.Errorf("expected SecretScanningDelegatedBypass nil, got %v", *config.SecretScanningDelegatedBypass)
				}
				if config.SecretScanningDelegatedBypassOptions != nil {
					t.Errorf("expected SecretScanningDelegatedBypassOptions nil, got %v", config.SecretScanningDelegatedBypassOptions)
				}
			},
		},
		{
			name: "sets bypass string without options",
			input: map[string]any{
				"name":                              "bypass-only",
				"secret_scanning_delegated_bypass":  "enabled",
			},
			expect: func(t *testing.T, config github.CodeSecurityConfiguration) {
				if config.GetSecretScanningDelegatedBypass() != "enabled" {
					t.Errorf("expected SecretScanningDelegatedBypass %q, got %q", "enabled", config.GetSecretScanningDelegatedBypass())
				}
				if config.SecretScanningDelegatedBypassOptions != nil {
					t.Errorf("expected SecretScanningDelegatedBypassOptions nil, got %v", config.SecretScanningDelegatedBypassOptions)
				}
			},
		},
		{
			name: "sets bypass with reviewers",
			input: map[string]any{
				"name":                             "bypass-with-reviewers",
				"secret_scanning_delegated_bypass": "enabled",
				"secret_scanning_delegated_bypass_options": []any{
					map[string]any{
						"reviewers": []any{
							map[string]any{
								"reviewer_id":   42,
								"reviewer_type": "TEAM",
							},
							map[string]any{
								"reviewer_id":   99,
								"reviewer_type": "ROLE",
							},
						},
					},
				},
			},
			expect: func(t *testing.T, config github.CodeSecurityConfiguration) {
				if config.GetSecretScanningDelegatedBypass() != "enabled" {
					t.Errorf("expected SecretScanningDelegatedBypass %q, got %q", "enabled", config.GetSecretScanningDelegatedBypass())
				}
				if config.SecretScanningDelegatedBypassOptions == nil {
					t.Fatal("expected SecretScanningDelegatedBypassOptions to be set")
				}
				reviewers := config.SecretScanningDelegatedBypassOptions.Reviewers
				if len(reviewers) != 2 {
					t.Fatalf("expected 2 reviewers, got %d", len(reviewers))
				}
				if reviewers[0].ReviewerID != 42 {
					t.Errorf("expected first reviewer_id 42, got %d", reviewers[0].ReviewerID)
				}
				if reviewers[0].ReviewerType != "TEAM" {
					t.Errorf("expected first reviewer_type %q, got %q", "TEAM", reviewers[0].ReviewerType)
				}
				if reviewers[1].ReviewerID != 99 {
					t.Errorf("expected second reviewer_id 99, got %d", reviewers[1].ReviewerID)
				}
				if reviewers[1].ReviewerType != "ROLE" {
					t.Errorf("expected second reviewer_type %q, got %q", "ROLE", reviewers[1].ReviewerType)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, resourceSchema, tt.input)
			config := github.CodeSecurityConfiguration{Name: d.Get("name").(string)}
			expandSecretScanningDelegatedBypass(d, &config)
			tt.expect(t, config)
		})
	}
}

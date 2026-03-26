package github

import (
	"testing"

	"github.com/google/go-github/v84/github"
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

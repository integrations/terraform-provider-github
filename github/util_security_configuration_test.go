package github

import (
	"testing"

	"github.com/google/go-github/v84/github"
)

func TestFlattenDependencyGraphAutosubmitActionOptions(t *testing.T) {
	t.Run("returns empty slice when options is nil", func(t *testing.T) {
		result := flattenDependencyGraphAutosubmitActionOptions(nil)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("omits labeled_runners key when LabeledRunners is nil", func(t *testing.T) {
		opts := &github.DependencyGraphAutosubmitActionOptions{}
		result := flattenDependencyGraphAutosubmitActionOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if _, ok := m["labeled_runners"]; ok {
			t.Errorf("labeled_runners should be absent when LabeledRunners is nil")
		}
	})

	t.Run("sets labeled_runners when LabeledRunners is non-nil", func(t *testing.T) {
		opts := &github.DependencyGraphAutosubmitActionOptions{
			LabeledRunners: github.Ptr(true),
		}
		result := flattenDependencyGraphAutosubmitActionOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["labeled_runners"] != true {
			t.Errorf("expected labeled_runners true, got %v", m["labeled_runners"])
		}
	})
}

func TestFlattenCodeScanningOptions(t *testing.T) {
	t.Run("returns empty slice when options is nil", func(t *testing.T) {
		result := flattenCodeScanningOptions(nil)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("omits allow_advanced key when AllowAdvanced is nil", func(t *testing.T) {
		opts := &github.CodeScanningOptions{}
		result := flattenCodeScanningOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if _, ok := m["allow_advanced"]; ok {
			t.Errorf("allow_advanced should be absent when AllowAdvanced is nil")
		}
	})

	t.Run("sets allow_advanced when AllowAdvanced is true", func(t *testing.T) {
		opts := &github.CodeScanningOptions{
			AllowAdvanced: github.Ptr(true),
		}
		result := flattenCodeScanningOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["allow_advanced"] != true {
			t.Errorf("expected allow_advanced true, got %v", m["allow_advanced"])
		}
	})

	t.Run("sets allow_advanced when AllowAdvanced is false", func(t *testing.T) {
		opts := &github.CodeScanningOptions{
			AllowAdvanced: github.Ptr(false),
		}
		result := flattenCodeScanningOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["allow_advanced"] != false {
			t.Errorf("expected allow_advanced false, got %v", m["allow_advanced"])
		}
	})
}

func TestFlattenSecretScanningDelegatedBypassOptions(t *testing.T) {
	t.Run("returns empty slice when options is nil", func(t *testing.T) {
		result := flattenSecretScanningDelegatedBypassOptions(nil)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("omits reviewers key when Reviewers is nil", func(t *testing.T) {
		opts := &github.SecretScanningDelegatedBypassOptions{}
		result := flattenSecretScanningDelegatedBypassOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if _, ok := m["reviewers"]; ok {
			t.Errorf("reviewers should be absent when Reviewers is nil")
		}
	})

	t.Run("sets reviewers when Reviewers is populated", func(t *testing.T) {
		opts := &github.SecretScanningDelegatedBypassOptions{
			Reviewers: []*github.BypassReviewer{
				{ReviewerID: 42, ReviewerType: "TEAM"},
				{ReviewerID: 99, ReviewerType: "ROLE"},
			},
		}
		result := flattenSecretScanningDelegatedBypassOptions(opts)
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
	})
}

func TestFlattenCodeScanningDefaultSetupOptions(t *testing.T) {
	t.Run("returns empty slice when options is nil", func(t *testing.T) {
		result := flattenCodeScanningDefaultSetupOptions(nil)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("omits runner_type key when RunnerType is empty string", func(t *testing.T) {
		opts := &github.CodeScanningDefaultSetupOptions{
			RunnerType: "",
		}
		result := flattenCodeScanningDefaultSetupOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if _, ok := m["runner_type"]; ok {
			t.Errorf("runner_type should be absent when RunnerType is empty, got %q", m["runner_type"])
		}
	})

	t.Run("sets runner_type when RunnerType is non-empty", func(t *testing.T) {
		opts := &github.CodeScanningDefaultSetupOptions{
			RunnerType: "standard",
		}
		result := flattenCodeScanningDefaultSetupOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["runner_type"] != "standard" {
			t.Errorf("expected runner_type %q, got %q", "standard", m["runner_type"])
		}
	})

	t.Run("sets runner_label when RunnerLabel is non-nil", func(t *testing.T) {
		opts := &github.CodeScanningDefaultSetupOptions{
			RunnerType:  "labeled",
			RunnerLabel: github.Ptr("my-runner"),
		}
		result := flattenCodeScanningDefaultSetupOptions(opts)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["runner_label"] != "my-runner" {
			t.Errorf("expected runner_label %q, got %q", "my-runner", m["runner_label"])
		}
	})
}

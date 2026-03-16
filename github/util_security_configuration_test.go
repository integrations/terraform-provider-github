package github

import (
	"testing"

	"github.com/google/go-github/v83/github"
)

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

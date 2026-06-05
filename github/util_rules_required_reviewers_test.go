package github

import (
	"testing"

	"github.com/google/go-github/v88/github"
)

func TestExpandRequiredReviewers(t *testing.T) {
	input := []any{
		map[string]any{
			"reviewer": []any{
				map[string]any{
					"id":   12345,
					"type": "Team",
				},
			},
			"file_patterns":     []any{"*.go", "src/**/*.ts"},
			"minimum_approvals": 2,
		},
		map[string]any{
			"reviewer": []any{
				map[string]any{
					"id":   67890,
					"type": "Team",
				},
			},
			"file_patterns":     []any{"docs/**/*.md"},
			"minimum_approvals": 1,
		},
	}

	result := expandRequiredReviewers(input)

	if len(result) != 2 {
		t.Fatalf("Expected 2 reviewers, got %d", len(result))
	}

	// Check first reviewer
	if result[0].Reviewer == nil {
		t.Fatal("Expected first reviewer to have a Reviewer")
	}
	if *result[0].Reviewer.ID != 12345 {
		t.Errorf("Expected first reviewer ID to be 12345, got %d", *result[0].Reviewer.ID)
	}
	if *result[0].Reviewer.Type != github.RulesetReviewerTypeTeam {
		t.Errorf("Expected first reviewer type to be Team, got %s", *result[0].Reviewer.Type)
	}
	if *result[0].MinimumApprovals != 2 {
		t.Errorf("Expected first reviewer minimum approvals to be 2, got %d", *result[0].MinimumApprovals)
	}
	if len(result[0].FilePatterns) != 2 {
		t.Fatalf("Expected first reviewer to have 2 file patterns, got %d", len(result[0].FilePatterns))
	}
	if result[0].FilePatterns[0] != "*.go" || result[0].FilePatterns[1] != "src/**/*.ts" {
		t.Errorf("Unexpected file patterns for first reviewer: %v", result[0].FilePatterns)
	}

	// Check second reviewer
	if result[1].Reviewer == nil {
		t.Fatal("Expected second reviewer to have a Reviewer")
	}
	if *result[1].Reviewer.ID != 67890 {
		t.Errorf("Expected second reviewer ID to be 67890, got %d", *result[1].Reviewer.ID)
	}
	if *result[1].MinimumApprovals != 1 {
		t.Errorf("Expected second reviewer minimum approvals to be 1, got %d", *result[1].MinimumApprovals)
	}
}

func TestExpandRequiredReviewersEmpty(t *testing.T) {
	result := expandRequiredReviewers([]any{})
	if result != nil {
		t.Error("Expected nil for empty input")
	}

	result = expandRequiredReviewers(nil)
	if result != nil {
		t.Error("Expected nil for nil input")
	}
}

func TestFlattenRequiredReviewers(t *testing.T) {
	reviewerType := github.RulesetReviewerTypeTeam
	reviewers := []*github.RulesetRequiredReviewer{
		{
			MinimumApprovals: new(2),
			FilePatterns:     []string{"*.go", "src/**/*.ts"},
			Reviewer: &github.RulesetReviewer{
				ID:   new(int64(12345)),
				Type: &reviewerType,
			},
		},
		{
			MinimumApprovals: new(1),
			FilePatterns:     []string{"docs/**/*.md"},
			Reviewer: &github.RulesetReviewer{
				ID:   new(int64(67890)),
				Type: &reviewerType,
			},
		},
	}

	result := flattenRequiredReviewers(reviewers)

	if len(result) != 2 {
		t.Fatalf("Expected 2 reviewers, got %d", len(result))
	}

	// Check first reviewer
	if result[0]["minimum_approvals"] != 2 {
		t.Errorf("Expected first reviewer minimum approvals to be 2, got %v", result[0]["minimum_approvals"])
	}
	filePatterns := result[0]["file_patterns"].([]string)
	if len(filePatterns) != 2 {
		t.Fatalf("Expected first reviewer to have 2 file patterns, got %d", len(filePatterns))
	}
	if filePatterns[0] != "*.go" || filePatterns[1] != "src/**/*.ts" {
		t.Errorf("Unexpected file patterns for first reviewer: %v", filePatterns)
	}

	reviewerBlock := result[0]["reviewer"].([]map[string]any)
	if len(reviewerBlock) != 1 {
		t.Fatalf("Expected 1 reviewer block, got %d", len(reviewerBlock))
	}
	if reviewerBlock[0]["id"] != 12345 {
		t.Errorf("Expected first reviewer ID to be 12345, got %v", reviewerBlock[0]["id"])
	}
	if reviewerBlock[0]["type"] != "Team" {
		t.Errorf("Expected first reviewer type to be Team, got %v", reviewerBlock[0]["type"])
	}

	// Check second reviewer
	if result[1]["minimum_approvals"] != 1 {
		t.Errorf("Expected second reviewer minimum approvals to be 1, got %v", result[1]["minimum_approvals"])
	}
}

func TestFlattenRequiredReviewersEmpty(t *testing.T) {
	result := flattenRequiredReviewers(nil)
	if result != nil {
		t.Error("Expected nil for nil input")
	}

	result = flattenRequiredReviewers([]*github.RulesetRequiredReviewer{})
	if result != nil {
		t.Error("Expected nil for empty slice input")
	}
}

func TestRoundTripRequiredReviewers(t *testing.T) {
	// Start with Terraform-style input
	input := []any{
		map[string]any{
			"reviewer": []any{
				map[string]any{
					"id":   12345,
					"type": "Team",
				},
			},
			"file_patterns":     []any{"*.go", "src/**/*.ts"},
			"minimum_approvals": 2,
		},
	}

	// Expand to go-github types
	expanded := expandRequiredReviewers(input)

	// Flatten back to Terraform types
	flattened := flattenRequiredReviewers(expanded)

	// Verify the round trip maintains data
	if len(flattened) != 1 {
		t.Fatalf("Expected 1 reviewer after round trip, got %d", len(flattened))
	}

	if flattened[0]["minimum_approvals"] != 2 {
		t.Errorf("Expected minimum_approvals to be 2 after round trip, got %v", flattened[0]["minimum_approvals"])
	}

	filePatterns := flattened[0]["file_patterns"].([]string)
	if len(filePatterns) != 2 {
		t.Fatalf("Expected 2 file patterns after round trip, got %d", len(filePatterns))
	}

	reviewerBlock := flattened[0]["reviewer"].([]map[string]any)
	if len(reviewerBlock) != 1 {
		t.Fatalf("Expected 1 reviewer block after round trip, got %d", len(reviewerBlock))
	}
	if reviewerBlock[0]["id"] != 12345 {
		t.Errorf("Expected reviewer ID to be 12345 after round trip, got %v", reviewerBlock[0]["id"])
	}
	if reviewerBlock[0]["type"] != "Team" {
		t.Errorf("Expected reviewer type to be Team after round trip, got %v", reviewerBlock[0]["type"])
	}
}

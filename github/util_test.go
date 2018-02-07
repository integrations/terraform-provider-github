package github

import (
	"testing"
)

func TestAccGithubUtilRole_validation(t *testing.T) {
	cases := []struct {
		Value    string
		ErrCount int
	}{
		{
			Value:    "invalid",
			ErrCount: 1,
		},
		{
			Value:    "valid_one",
			ErrCount: 0,
		},
		{
			Value:    "valid_two",
			ErrCount: 0,
		},
	}

	validationFunc := validateValueFunc([]string{"valid_one", "valid_two"})

	for _, tc := range cases {
		_, errors := validationFunc(tc.Value, "test_arg")

		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected 1 validation error")
		}
	}
}

func TestAccGithubUtilTwoPartID(t *testing.T) {
	partOne, partTwo := "foo", "bar"

	id := buildTwoPartID(&partOne, &partTwo)

	if id != "foo:bar" {
		t.Fatalf("Expected two part id to be foo:bar, actual: %s", id)
	}

	parsedPartOne, parsedPartTwo := parseTwoPartID(id)

	if parsedPartOne != "foo" {
		t.Fatalf("Expected parsed part one foo, actual: %s", parsedPartOne)
	}

	if parsedPartTwo != "bar" {
		t.Fatalf("Expected parsed part two bar, actual: %s", parsedPartTwo)
	}
}

func TestAccValidateTwoPartID(t *testing.T) {
	cases := []struct {
		name        string
		id          string
		expectedErr string
	}{
		{
			name: "valid",
			id:   "foo:bar",
		},
		{
			name:        "blank ID",
			id:          "",
			expectedErr: "no ID supplied. Please supply an ID format matching organization:username",
		},
		{
			name:        "not enough parts",
			id:          "foo",
			expectedErr: "incorrectly formatted ID \"foo\". Please supply an ID format matching organization:username",
		},
		{
			name:        "too many parts",
			id:          "foo:bar:baz",
			expectedErr: "incorrectly formatted ID \"foo:bar:baz\". Please supply an ID format matching organization:username",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateTwoPartID(tc.id)
			switch {
			case err != nil && tc.expectedErr == "":
				t.Fatalf("expected no error, got %q", err)
			case err != nil && tc.expectedErr != "":
				if err.Error() != tc.expectedErr {
					t.Fatalf("expected error to be %q, got %q", tc.expectedErr, err.Error())
				}
			}
		})
	}
}

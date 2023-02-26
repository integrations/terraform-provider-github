package github

import (
	"strings"
	"testing"
	"unicode"
)

func TestAccValidateTeamIDFunc(t *testing.T) {
	// warnings, errors := validateTeamIDFunc(interface{"1234567"})

	cases := []struct {
		TeamID   interface{}
		ErrCount int
	}{
		{

			TeamID:   "1234567",
			ErrCount: 0,
		},
		{
			// an int cannot be cast to a string
			TeamID:   1234567,
			ErrCount: 1,
		},
		{
			TeamID:   "notAnInt",
			ErrCount: 1,
		},
	}

	for _, tc := range cases {
		_, errors := validateTeamIDFunc(tc.TeamID, "keyName")
		if len(errors) != tc.ErrCount {
			t.Fatalf("Expected %d validation error but got %d", tc.ErrCount, len(errors))
		}
	}
}

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

	id := buildTwoPartID(partOne, partTwo)

	if id != "foo:bar" {
		t.Fatalf("Expected two part id to be foo:bar, actual: %s", id)
	}

	parsedPartOne, parsedPartTwo, err := parseTwoPartID(id, "left", "right")
	if err != nil {
		t.Fatal(err)
	}

	if parsedPartOne != "foo" {
		t.Fatalf("Expected parsed part one foo, actual: %s", parsedPartOne)
	}

	if parsedPartTwo != "bar" {
		t.Fatalf("Expected parsed part two bar, actual: %s", parsedPartTwo)
	}
}

func flipUsernameCase(username string) string {
	oc := []rune(username)

	for i, ch := range oc {
		if unicode.IsLetter(ch) {

			if unicode.IsUpper(ch) {
				oc[i] = unicode.ToLower(ch)
			} else {
				oc[i] = unicode.ToUpper(ch)
			}
			break
		}

	}
	return string(oc)
}

func TestAccGithubUtilValidateSecretName(t *testing.T) {
	cases := []struct {
		Name  string
		Error bool
	}{
		{
			Name: "valid",
		},
		{
			Name: "v",
		},
		{
			Name: "_valid_underscore_",
		},
		{
			Name: "valid_digit_1",
		},
		{
			Name:  "invalid-dashed",
			Error: true,
		},
		{
			Name:  "1_invalid_leading_digit",
			Error: true,
		},
		{
			Name:  "GITHUB_PREFIX",
			Error: true,
		},
		{
			Name:  "github_prefix",
			Error: true,
		},
	}

	for _, tc := range cases {
		var name interface{} = tc.Name
		_, errors := validateSecretNameFunc(name, "")

		if tc.Error != (len(errors) != 0) {
			if tc.Error {
				t.Fatalf("expected error, got none (%s)", tc.Name)
			} else {
				t.Fatalf("unexpected error(s): %s (%s)", errors, tc.Name)
			}
		}
	}
}

func TestValidateVariableName(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		provided interface{}
		exp      string
	}{
		{
			name:     "not_string",
			provided: 123,
			exp:      `expected field to be a string, got int`,
		},
		{
			name:     "GITHUB_prefix",
			provided: "GITHUB_foo",
			exp:      `field cannot start with "GITHUB_"`,
		},
		{
			name:     "github_prefix",
			provided: "github_foo",
			exp:      `field cannot start with "GITHUB_"`,
		},
		{
			name:     "start_with_number",
			provided: "1_abc",
			exp:      `field must start with a letter`,
		},
		{
			name:     "emoji",
			provided: "abc🤘",
			exp:      `invalid character '🤘' for field at position 3`,
		},
		{
			name:     "international",
			provided: "abcü",
			exp:      `invalid character 'ü'`,
		},
		{
			name:     "lowercase",
			provided: "abc",
		},
		{
			name:     "uppercase",
			provided: "ABC",
		},
		{
			name:     "mixed_case",
			provided: "aBcDEfg",
		},
		{
			name:     "everything",
			provided: "aBcD2_4Efg",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, errs := validateVariableName(tc.provided, "field")

			if tc.exp == "" && len(errs) > 0 {
				t.Errorf("expected no errors, got %q", errs)
			}

			strs := make([]string, 0, len(errs))
			for _, err := range errs {
				strs = append(strs, err.Error())
			}
			combined := strings.Join(strs, ", ")

			if !strings.Contains(combined, tc.exp) {
				t.Errorf("expected %q to contain %q", combined, tc.exp)
			}
		})
	}
}

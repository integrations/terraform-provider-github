package github

import (
	"testing"
	"unicode"

	"github.com/hashicorp/go-cty/cty"
)

func Test_escapeIDPart(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName string
		input    string
		expected string
	}{
		{
			testName: "no_separator",
			input:    "part1",
			expected: "part1",
		},
		{
			testName: "with_separator",
			input:    "part:1",
			expected: "part??1",
		},
		{
			testName: "multiple_separators",
			input:    "part:1:subpart:2",
			expected: "part??1??subpart??2",
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got := escapeIDPart(d.input)

			if got != d.expected {
				t.Fatalf("expected escaped part %q but got %q", d.expected, got)
			}
		})
	}
}

func Test_unescapeIDPart(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName string
		input    string
		expected string
	}{
		{
			testName: "no_escaped_separator",
			input:    "part1",
			expected: "part1",
		},
		{
			testName: "with_escaped_separator",
			input:    "part??1",
			expected: "part:1",
		},
		{
			testName: "multiple_escaped_separators",
			input:    "part??1??subpart??2",
			expected: "part:1:subpart:2",
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got := unescapeIDPart(d.input)

			if got != d.expected {
				t.Fatalf("expected unescaped part %q but got %q", d.expected, got)
			}
		})
	}
}

func Test_buildID(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName string
		parts    []string
		expect   string
		hasError bool
	}{
		{
			testName: "one_part",
			parts:    []string{"part1"},
			expect:   "part1",
			hasError: false,
		},
		{
			testName: "two_parts",
			parts:    []string{"part1", "part2"},
			expect:   "part1:part2",
			hasError: false,
		},
		{
			testName: "three_parts",
			parts:    []string{"part1", "part2", "part3"},
			expect:   "part1:part2:part3",
			hasError: false,
		},
		{
			testName: "part_with_unescaped_separator",
			parts:    []string{"part1", "part:2", "part3"},
			expect:   "",
			hasError: true,
		},
		{
			testName: "no_parts",
			parts:    []string{},
			expect:   "",
			hasError: true,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got, err := buildID(d.parts...)

			if d.hasError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !d.hasError && err != nil {
				t.Fatalf("did not expect error but got: %v", err)
			}
			if got != d.expect {
				t.Fatalf("expected id %q but got %q", d.expect, got)
			}
		})
	}
}

func Test_parseID(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName string
		id       string
		count    int
		expect   []string
		hasError bool
	}{
		{
			testName: "two_parts_expected_two",
			id:       "part1:part2",
			count:    2,
			expect:   []string{"part1", "part2"},
			hasError: false,
		},
		{
			testName: "three_parts_expected_three",
			id:       "part1:part2:part3",
			count:    3,
			expect:   []string{"part1", "part2", "part3"},
			hasError: false,
		},
		{
			testName: "two_parts_expected_three",
			id:       "part1:part2",
			count:    3,
			expect:   nil,
			hasError: true,
		},
		{
			testName: "three_parts_expected_two",
			id:       "part1:part2:part3",
			count:    2,
			expect:   nil,
			hasError: true,
		},
		{
			testName: "empty_id",
			id:       "",
			count:    0,
			expect:   nil,
			hasError: true,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got, err := parseID(d.id, d.count)

			if d.hasError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !d.hasError && err != nil {
				t.Fatalf("did not expect error but got: %v", err)
			}
			if !d.hasError {
				for i, part := range d.expect {
					if got[i] != part {
						t.Fatalf("expected part %d to be %q but got %q", i, part, got[i])
					}
				}
			}
		})
	}
}

func Test_parseID2(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName string
		id       string
		expect1  string
		expect2  string
		hasError bool
	}{
		{
			testName: "valid_two_parts",
			id:       "part1:part2",
			expect1:  "part1",
			expect2:  "part2",
			hasError: false,
		},
		{
			testName: "invalid_three_parts",
			id:       "part1:part2:part3",
			expect1:  "",
			expect2:  "",
			hasError: true,
		},
		{
			testName: "invalid_one_part",
			id:       "part1",
			expect1:  "",
			expect2:  "",
			hasError: true,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got1, got2, err := parseID2(d.id)

			if d.hasError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !d.hasError && err != nil {
				t.Fatalf("did not expect error but got: %v", err)
			}
			if !d.hasError {
				if got1 != d.expect1 {
					t.Fatalf("expected part 1 to be %q but got %q", d.expect1, got1)
				}
				if got2 != d.expect2 {
					t.Fatalf("expected part 2 to be %q but got %q", d.expect2, got2)
				}
			}
		})
	}
}

func Test_parseID3(t *testing.T) {
	t.Parallel()

	for _, d := range []struct {
		testName string
		id       string
		expect1  string
		expect2  string
		expect3  string
		hasError bool
	}{
		{
			testName: "valid_three_parts",
			id:       "part1:part2:part3",
			expect1:  "part1",
			expect2:  "part2",
			expect3:  "part3",
			hasError: false,
		},
		{
			testName: "invalid_two_parts",
			id:       "part1:part2",
			expect1:  "",
			expect2:  "",
			expect3:  "",
			hasError: true,
		},
		{
			testName: "invalid_four_parts",
			id:       "part1:part2:part3:part4",
			expect1:  "",
			expect2:  "",
			expect3:  "",
			hasError: true,
		},
	} {
		t.Run(d.testName, func(t *testing.T) {
			t.Parallel()

			got1, got2, got3, err := parseID3(d.id)

			if d.hasError && err == nil {
				t.Fatalf("expected error but got none")
			}
			if !d.hasError && err != nil {
				t.Fatalf("did not expect error but got: %v", err)
			}
			if !d.hasError {
				if got1 != d.expect1 {
					t.Fatalf("expected part 1 to be %q but got %q", d.expect1, got1)
				}
				if got2 != d.expect2 {
					t.Fatalf("expected part 2 to be %q but got %q", d.expect2, got2)
				}
				if got3 != d.expect3 {
					t.Fatalf("expected part 3 to be %q but got %q", d.expect3, got3)
				}
			}
		})
	}
}

func TestGithubUtilRole_validation(t *testing.T) {
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
		diags := validationFunc(tc.Value, cty.Path{cty.GetAttrStep{Name: "test_arg"}})

		if len(diags) != tc.ErrCount {
			t.Fatalf("Expected 1 validation error")
		}
	}
}

func TestGithubUtilTwoPartID(t *testing.T) {
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

func TestGithubUtilValidateSecretName(t *testing.T) {
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
		var name any = tc.Name
		diags := validateSecretNameFunc(name, cty.Path{cty.GetAttrStep{Name: ""}})

		if tc.Error != (len(diags) != 0) {
			if tc.Error {
				t.Fatalf("expected error, got none (%s)", tc.Name)
			} else {
				t.Fatalf("unexpected error(s): %v (%s)", diags, tc.Name)
			}
		}
	}
}

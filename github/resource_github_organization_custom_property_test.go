package github

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccGithubOrganizationCustomProperty_CustomizeDiff_Validations(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		config      string
		expectError *regexp.Regexp
	}{
		{
			name: "required true but default_value missing",
			config: `
resource "github_organization_custom_property" "test" {
  name    = "required_without_default"
  type    = "string"
  required = true
}
`,
			expectError: regexp.MustCompile("default_value can not be empty"),
		},
		{
			name: "default_value set but required false",
			config: `
resource "github_organization_custom_property" "test" {
  name    = "default_without_required"
  type    = "string"
  required = false
  default_value = ["foo"]
}
`,
			expectError: regexp.MustCompile("default_value is only allowed if required is true"),
		},
		{
			name: "default_value not in allowed_values with SINGLE_SELECT",
			config: `
resource "github_organization_custom_property" "test" {
  name = "invalid_default"
  type = "single_select"
  required = true
  allowed_values = ["foo", "bar"]
  default_value  = ["baz"]
}
`,
			expectError: regexp.MustCompile("default_value must be a subset of allowed_values"),
		},
		{
			name: "default_value not in allowed_values with MULTI_SELECT",
			config: `
resource "github_organization_custom_property" "test" {
  name = "invalid_default"
  type = "multi_select"
  required = true
  allowed_values = ["foo", "bar"]
  default_value  = ["foo", "baz"]
}
`,
			expectError: regexp.MustCompile("default_value must be a subset of allowed_values"),
		},
		{
			name: "allowed_values used with STRING",
			config: `
resource "github_organization_custom_property" "test" {
  name = "string_with_allowed"
  type = "string"
  allowed_values = ["not", "allowed"]
}
`,
			expectError: regexp.MustCompile("allowed_values must be empty when type is STRING or TRUE_FALSE"),
		},
		{
			name: "allowed_values used with TRUE_FALSE",
			config: `
resource "github_organization_custom_property" "test" {
  name = "true_false_with_allowed"
  type = "true_false"
  allowed_values = ["true"]
}
`,
			expectError: regexp.MustCompile("allowed_values must be empty when type is STRING or TRUE_FALSE"),
		},
		{
			name: "multiple default_value with SINGLE_SELECT",
			config: `
resource "github_organization_custom_property" "test" {
  name = "multi_default_single_select"
  type = "single_select"
  required = true
  allowed_values = ["foo"]
  default_value  = ["foo", "bar"]
}
`,
			expectError: regexp.MustCompile("defaultValue must contain zero or one item when type is SINGLE_SELECT or STRING"),
		},
		{
			name: "multiple default_value with STRING",
			config: `
resource "github_organization_custom_property" "test" {
  name = "multi_default_string"
  type = "string"
  required = true
  default_value  = ["foo", "bar"]
}
`,
			expectError: regexp.MustCompile("defaultValue must contain zero or one item when type is SINGLE_SELECT or STRING"),
		},
		{
			name: "invalid TRUE_FALSE default",
			config: `
resource "github_organization_custom_property" "test" {
  name = "invalid_true_false"
  type = "true_false"
  required = true
  default_value = ["maybe"]
}
`,
			expectError: regexp.MustCompile("default_value must be either \"true\" or \"false\" when type is TRUE_FALSE"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck:  func() { skipUnlessMode(t, organization) },
				Providers: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      tc.config,
						ExpectError: tc.expectError,
					},
				},
			})
		})
	}
}

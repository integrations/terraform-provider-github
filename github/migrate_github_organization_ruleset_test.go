package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestResourceGithubOrganizationRulesetMigrateState(t *testing.T) {
	cases := map[string]struct {
		StateVersion int
		Attributes   map[string]string
		Expected     map[string]string
		Meta         any
		ExpectError  bool
	}{
		"v1_to_v2": {
			StateVersion: 1,
			Attributes: map[string]string{
				"name":        "test-org-ruleset",
				"target":      "branch",
				"enforcement": "active",
			},
			Expected: map[string]string{
				"name":        "test-org-ruleset",
				"target":      "branch",
				"enforcement": "active",
			},
		},
		"v1_to_v2_empty_state": {
			StateVersion: 1,
			Attributes:   map[string]string{},
			Expected:     map[string]string{},
		},
		"v1_to_v2_with_conditions": {
			StateVersion: 1,
			Attributes: map[string]string{
				"name":                              "org-ruleset-with-conditions",
				"target":                            "branch",
				"enforcement":                       "evaluate",
				"conditions.0.repository_id.0":      "12345",
				"conditions.0.ref_name.0.include.0": "main",
			},
			Expected: map[string]string{
				"name":                              "org-ruleset-with-conditions",
				"target":                            "branch",
				"enforcement":                       "evaluate",
				"conditions.0.repository_id.0":      "12345",
				"conditions.0.ref_name.0.include.0": "main",
			},
		},
		"unsupported_version": {
			StateVersion: 2,
			Attributes: map[string]string{
				"name": "test",
			},
			ExpectError: true,
		},
	}

	for tn, tc := range cases {
		is := &terraform.InstanceState{
			ID:         "test",
			Attributes: tc.Attributes,
		}

		is, err := resourceGithubOrganizationRulesetMigrateState(tc.StateVersion, is, tc.Meta)

		if tc.ExpectError {
			if err == nil {
				t.Fatalf("bad: %s, expected error but got none", tn)
			}
			continue
		}

		if err != nil {
			t.Fatalf("bad: %s, err: %#v", tn, err)
		}

		for k, v := range tc.Expected {
			actual := is.Attributes[k]
			if actual != v {
				t.Fatalf("bad: %s\n\nexpected: %#v -> %#v\n     got: %#v -> %#v\n in: %#v",
					tn, k, v, k, actual, is.Attributes)
			}
		}
	}
}

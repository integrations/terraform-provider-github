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
	}

	for tn, tc := range cases {
		is := &terraform.InstanceState{
			ID:         "test",
			Attributes: tc.Attributes,
		}

		is, err := resourceGithubOrganizationRulesetMigrateState(tc.StateVersion, is, tc.Meta)
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

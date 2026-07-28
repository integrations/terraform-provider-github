package github

import (
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestExpandRequiredStatusChecks(t *testing.T) {
	appID := int64(15368)
	anyApp := int64(-1)

	testCases := map[string]struct {
		contexts []any
		checks   []any
		want     []*github.RequiredStatusCheck
	}{
		"contexts only": {
			contexts: []any{"ci/test"},
			want: []*github.RequiredStatusCheck{
				{Context: "ci/test", AppID: &anyApp},
			},
		},
		"checks only": {
			checks: []any{"ci/test:15368"},
			want: []*github.RequiredStatusCheck{
				{Context: "ci/test", AppID: &appID},
			},
		},
		// A read populates both fields from the API response, so state can hold the
		// same context twice even though configuration cannot set both fields.
		"same context in both fields is sent once": {
			contexts: []any{"ci/test", "ci/build"},
			checks:   []any{"ci/test", "ci/build"},
			want: []*github.RequiredStatusCheck{
				{Context: "ci/test", AppID: &anyApp},
				{Context: "ci/build", AppID: &anyApp},
			},
		},
		"app_id from checks wins over the context duplicate": {
			contexts: []any{"ci/test"},
			checks:   []any{"ci/test:15368"},
			want: []*github.RequiredStatusCheck{
				{Context: "ci/test", AppID: &appID},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			statusChecks := map[string]any{"strict": true}
			if tc.contexts != nil {
				statusChecks["contexts"] = schema.NewSet(schema.HashString, tc.contexts)
			}
			if tc.checks != nil {
				statusChecks["checks"] = schema.NewSet(schema.HashString, tc.checks)
			}

			d := schema.TestResourceDataRaw(t, resourceGithubBranchProtectionV3().Schema, map[string]any{
				"repository": "test",
				"branch":     "main",
			})
			if err := d.Set("required_status_checks", []any{statusChecks}); err != nil {
				t.Fatalf("failed to set required_status_checks: %v", err)
			}

			rsc, err := expandRequiredStatusChecks(d)
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if !rsc.Strict {
				t.Error("expected strict to be true")
			}

			got := map[string]*int64{}
			for _, check := range *rsc.Checks {
				if _, ok := got[check.Context]; ok {
					t.Fatalf("context %q sent more than once, GitHub rejects duplicates", check.Context)
				}
				got[check.Context] = check.AppID
			}

			if len(got) != len(tc.want) {
				t.Fatalf("expected %d checks, got %d: %v", len(tc.want), len(got), got)
			}
			for _, want := range tc.want {
				gotAppID, ok := got[want.Context]
				if !ok {
					t.Fatalf("expected context %q to be present", want.Context)
				}
				if gotAppID == nil || *gotAppID != *want.AppID {
					t.Errorf("context %q: expected app_id %d, got %v", want.Context, *want.AppID, gotAppID)
				}
			}
		})
	}
}

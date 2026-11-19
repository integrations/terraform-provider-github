package github

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Matching computed contexts intentionally normalize null App IDs to -1 (#3585).
// Omission selects the last providing App; -1 allows any App.
// These tests catch mixing refreshed Optional+Computed fields in an update.
// Expectations describe the wire representation, including the distinction
// between an omitted app_id and -1 (accept checks from any App).
func TestBranchProtectionV3ConfiguredStatusChecks(t *testing.T) {
	tests := []struct {
		name       string
		configured map[string]any
		refreshed  map[string]any
		want       string
	}{
		{
			"configured checks discard computed contexts",
			map[string]any{"checks": []any{"build:15368"}},
			map[string]any{"contexts": []any{"old"}, "checks": []any{"build:15368"}},
			`{"strict":false,"checks":[{"context":"build","app_id":15368}]}`,
		},
		{
			"configured contexts discard computed checks during rollback",
			map[string]any{"contexts": []any{"build"}},
			map[string]any{"contexts": []any{"build"}, "checks": []any{"build:15368", "stale:42"}},
			`{"strict":false,"checks":[{"context":"build","app_id":-1}]}`,
		},
		{
			"explicit empty checks clear computed contexts",
			map[string]any{"checks": []any{}},
			map[string]any{"contexts": []any{"old"}, "checks": []any{}},
			`{"strict":false,"checks":[]}`,
		},
		{
			"explicit empty contexts clear computed checks",
			map[string]any{"contexts": []any{}},
			map[string]any{"contexts": []any{}, "checks": []any{"old:42"}},
			`{"strict":false,"checks":[]}`,
		},
		{
			"omitted families preserve computed app bindings",
			map[string]any{"strict": true},
			map[string]any{"strict": true, "contexts": []any{"build"}, "checks": []any{"build:42"}},
			`{"strict":true,"checks":[{"context":"build","app_id":42}]}`,
		},
		{
			"omitted families preserve contexts-only legacy state",
			map[string]any{"strict": true},
			map[string]any{"strict": true, "contexts": []any{"legacy"}, "checks": []any{}},
			`{"strict":true,"checks":[{"context":"legacy","app_id":-1}]}`,
		},
		{
			"explicit null checks behave as omitted",
			map[string]any{"checks": nil, "strict": true},
			map[string]any{"strict": true, "contexts": []any{"build"}, "checks": []any{"build:42"}},
			`{"strict":true,"checks":[{"context":"build","app_id":42}]}`,
		},
		{
			"explicit null contexts behave as omitted",
			map[string]any{"contexts": nil, "strict": true},
			map[string]any{"strict": true, "contexts": []any{"build"}, "checks": []any{"build:42"}},
			`{"strict":true,"checks":[{"context":"build","app_id":42}]}`,
		},
		{
			"configured checks without matching context omit app id",
			map[string]any{"checks": []any{"build"}},
			map[string]any{"contexts": []any{}, "checks": []any{"build"}},
			`{"strict":false,"checks":[{"context":"build"}]}`,
		},
		{
			"configured checks suppress stale context and preserve matching null app metadata",
			map[string]any{"checks": []any{"build"}},
			map[string]any{"contexts": []any{"build", "stale"}, "checks": []any{"build"}},
			`{"strict":false,"checks":[{"context":"build","app_id":-1}]}`,
		},
		{
			"configured checks preserve matching null app metadata",
			map[string]any{"checks": []any{"build"}},
			map[string]any{"contexts": []any{"build"}, "checks": []any{"build"}},
			`{"strict":false,"checks":[{"context":"build","app_id":-1}]}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := resourceGithubBranchProtectionV3()
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
				"repository": "test", "branch": "main",
				"required_status_checks": []any{tt.configured},
			})
			// Match Core's separate raw configuration and refreshed state. The
			// legacy TestResourceDataRaw helper does not retain RawConfig.
			raw, err := schema.JSONMapToStateValue(map[string]any{
				"repository": "test", "branch": "main",
				"required_status_checks": []any{tt.configured},
			}, r.CoreConfigSchema())
			if err != nil {
				t.Fatal(err)
			}
			d.SetId("test:main")
			state := d.State()
			state.RawConfig = raw
			d = r.Data(state)
			// Simulate state populated by Read.
			if err := d.Set("required_status_checks", []any{tt.refreshed}); err != nil {
				t.Fatal(err)
			}
			request, err := expandRequiredStatusChecks(d)
			if err != nil {
				t.Fatal(err)
			}
			got, err := json.Marshal(request)
			if err != nil {
				t.Fatal(err)
			}
			var gotJSON, wantJSON any
			if err := json.Unmarshal(got, &gotJSON); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(tt.want), &wantJSON); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(gotJSON, wantJSON) {
				t.Errorf("request = %s; want %s", got, tt.want)
			}
		})
	}
}

// Current upstream compatibility: absent/unknown config uses refreshed deduplication.
func TestBranchProtectionV3UnavailableConfigurationPreservesChecks(t *testing.T) {
	r := resourceGithubBranchProtectionV3()
	raw, err := schema.JSONMapToStateValue(map[string]any{
		"repository": "test", "branch": "main", "required_status_checks": []any{map[string]any{"strict": true}},
	}, r.CoreConfigSchema())
	if err != nil {
		t.Fatal(err)
	}
	attrs := raw.AsValueMap()
	blocksType := attrs["required_status_checks"].Type()
	attrs["required_status_checks"] = cty.UnknownVal(blocksType)
	unknownBlocks := cty.ObjectVal(attrs)
	attrs["required_status_checks"] = cty.ListVal([]cty.Value{cty.UnknownVal(blocksType.ElementType())})
	unknownBlock := cty.ObjectVal(attrs)
	for name, config := range map[string]cty.Value{
		"null raw config":    cty.NullVal(raw.Type()),
		"unknown raw config": cty.UnknownVal(raw.Type()),
		"unknown blocks":     unknownBlocks,
		"unknown block":      unknownBlock,
	} {
		t.Run(name, func(t *testing.T) {
			d := schema.TestResourceDataRaw(t, r.Schema, map[string]any{
				"repository": "test", "branch": "main", "required_status_checks": []any{map[string]any{
					"contexts": []any{"build"}, "checks": []any{"build:42"},
				}},
			})
			d.SetId("test:main")
			state := d.State()
			state.RawConfig = config
			d = r.Data(state)
			request, err := expandRequiredStatusChecks(d)
			if err != nil {
				t.Fatal(err)
			}
			got, err := json.Marshal(request)
			if err != nil {
				t.Fatal(err)
			}
			const want = `{"strict":false,"checks":[{"context":"build","app_id":42}]}`
			if string(got) != want {
				t.Errorf("request = %s; want %s", got, want)
			}
		})
	}
}

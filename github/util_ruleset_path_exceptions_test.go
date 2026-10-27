package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestRulesetPathExceptionsLifecycle(t *testing.T) {
	t.Parallel()

	for _, org := range []bool{false, true} {
		t.Run(fmt.Sprintf("organization=%t", org), func(t *testing.T) {
			t.Parallel()
			res := resourceGithubRepositoryRuleset()
			path := "/repos/owner/repo/rulesets"
			if org {
				res = resourceGithubOrganizationRuleset()
				path = "/orgs/owner/rulesets"
			}
			var responseRules string
			var expectedMethod string
			var expectedPaths []string
			requests := 0
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/repos/owner/repo" {
					mustWrite(w, `{"name":"repo","archived":false}`)
					return
				}
				requests++
				if r.Method != expectedMethod {
					t.Errorf("method = %s, want %s", r.Method, expectedMethod)
				}
				wantPath := path + "/123"
				if r.Method == http.MethodPost {
					wantPath = path
				}
				if r.URL.Path != wantPath {
					t.Errorf("path = %s, want %s", r.URL.Path, wantPath)
				}
				if !org && r.Method == http.MethodGet && r.URL.Query().Get("includes_parents") != "false" {
					t.Error("repository reads must exclude parent rulesets")
				}
				if r.Method != http.MethodGet {
					var body struct {
						Rules []struct {
							Type       string         `json:"type"`
							Parameters map[string]any `json:"parameters"`
						} `json:"rules"`
					}
					if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
						t.Fatal(err)
					}
					got := make(map[string]map[string]any)
					for _, rule := range body.Rules {
						got[rule.Type] = rule.Parameters
					}
					paths := make([]any, len(expectedPaths))
					for i, p := range expectedPaths {
						paths[i] = p
					}
					want := map[string]map[string]any{
						"file_path_restriction": {"restricted_file_paths": []any{"**/*.jsonl"}, "ignored_file_paths": paths},
						"max_file_size":         {"max_file_size": float64(10), "ignored_file_paths": paths},
						"max_file_path_length":  {"max_file_path_length": float64(255)},
					}
					if diff := cmp.Diff(want, got); diff != "" {
						t.Errorf("request rules mismatch (-want +got):\n%s", diff)
					}
				}
				w.Header().Set("ETag", `"test-etag"`)
				mustWrite(w, fmt.Sprintf(`{"id":123,"node_id":"ruleset-node","name":"test","target":"push","enforcement":"active","rules":%s}`, responseRules))
			})
			client := mustCreateTestGitHubClient(t, "https://api.github.test/", github.WithHTTPClient(&http.Client{Transport: localRoundTripper{handler: handler}}))
			owner := &Owner{name: "owner", v3client: client}
			config := map[string]any{"name": "test", "target": "push", "enforcement": "active"}
			if !org {
				config["repository"] = "repo"
			}
			d := schema.TestResourceDataRaw(t, res.Schema, config)

			for i, paths := range [][]string{{"**/*mock*.jsonl", "fixtures/**"}, {"testdata/**"}, nil, {"fixtures/**"}, {}} {
				expectedPaths = paths
				rules := map[string]any{
					"file_path_restriction": []any{map[string]any{"restricted_file_paths": []string{"**/*.jsonl"}}},
					"max_file_size":         []any{map[string]any{"max_file_size": 10}},
					"max_file_path_length":  []any{map[string]any{"max_file_path_length": 255}},
				}
				if paths != nil {
					for _, ruleType := range []string{"file_path_restriction", "max_file_size"} {
						blocks, _ := rules[ruleType].([]any)
						block, _ := blocks[0].(map[string]any)
						block["ignored_file_paths"] = paths
					}
				}
				if err := d.Set("rules", []any{rules}); err != nil {
					t.Fatal(err)
				}
				pathsJSON, err := json.Marshal(slices.Concat([]string{}, paths))
				if err != nil {
					t.Fatal(err)
				}
				responseRules = fmt.Sprintf(`[
					{"type":"file_path_restriction","parameters":{"restricted_file_paths":["**/*.jsonl"],"ignored_file_paths":%s}},
					{"type":"max_file_size","parameters":{"max_file_size":10,"ignored_file_paths":%s}},
					{"type":"max_file_path_length","parameters":{"max_file_path_length":255}}
				]`, pathsJSON, pathsJSON)
				expectedMethod = http.MethodPut
				if i == 0 {
					expectedMethod = http.MethodPost
					if diags := res.CreateContext(t.Context(), d, owner); diags.HasError() {
						t.Fatal(diags)
					}
				} else if diags := res.UpdateContext(t.Context(), d, owner); diags.HasError() {
					t.Fatal(diags)
				}
				expectedMethod = http.MethodGet
				if diags := res.ReadContext(t.Context(), d, owner); diags.HasError() {
					t.Fatal(diags)
				}
				if i == 0 {
					imported := schema.TestResourceDataRaw(t, res.Schema, nil)
					imported.SetId("repo:123")
					if org {
						imported.SetId("123")
					}
					if _, err := res.Importer.StateContext(t.Context(), imported, owner); err != nil {
						t.Fatal(err)
					}
					if diags := res.ReadContext(t.Context(), imported, owner); diags.HasError() {
						t.Fatal(diags)
					}
					if diff := cmp.Diff(d.Get("rules"), imported.Get("rules")); diff != "" {
						t.Errorf("import lost rules (-want +got):\n%s", diff)
					}
				}
				for _, ruleType := range []string{"file_path_restriction", "max_file_size"} {
					got, _ := d.Get("rules.0." + ruleType + ".0.ignored_file_paths").([]any)
					if !slices.Equal(expandStringList(got), paths) {
						t.Errorf("%s state = %v, want %v", ruleType, got, paths)
					}
				}
				if d.Id() != "123" || d.Get("node_id") != "ruleset-node" || d.Get("etag") != `"test-etag"` {
					t.Errorf("unexpected resource identity or etag: %s, %v, %v", d.Id(), d.Get("node_id"), d.Get("etag"))
				}
			}
			if requests != 12 {
				t.Errorf("ruleset requests = %d, want 12", requests)
			}
		})
	}
}

func TestRulesetPathExceptionsPreservesOtherRules(t *testing.T) {
	t.Parallel()

	for _, body := range []string{
		`[]`,
		`[{"type":"creation"},{"type":"update","parameters":{"update_allows_fetch_and_merge":true}},{"type":"pull_request","parameters":{"dismiss_stale_reviews_on_push":true,"require_code_owner_review":true,"required_approving_review_count":2}}]`,
		`[{"type":"deletion"},{"type":"non_fast_forward"},{"type":"tag_name_pattern","parameters":{"operator":"starts_with","pattern":"v"}}]`,
	} {
		t.Run(body, func(t *testing.T) {
			t.Parallel()
			var sdk github.RepositoryRulesetRules
			var extended rulesWithPathExceptions
			if err := json.Unmarshal([]byte(body), &sdk); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(body), &extended); err != nil {
				t.Fatal(err)
			}
			want, err := json.Marshal(sdk)
			if err != nil {
				t.Fatal(err)
			}
			got, err := json.Marshal(extended)
			if err != nil {
				t.Fatal(err)
			}
			var wantJSON, gotJSON any
			if err := json.Unmarshal(want, &wantJSON); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(got, &gotJSON); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(wantJSON, gotJSON); diff != "" {
				t.Errorf("rules encoding changed (-want +got):\n%s", diff)
			}
		})
	}
}

func TestRulesetPathExceptionsReadCompatibility(t *testing.T) {
	t.Parallel()

	for _, exceptions := range []string{``, `,"ignored_file_paths":null`, `,"ignored_file_paths":[]`} {
		t.Run(exceptions, func(t *testing.T) {
			t.Parallel()
			var rules rulesWithPathExceptions
			body := fmt.Sprintf(`[{"type":"file_path_restriction","parameters":{"restricted_file_paths":["**/*.jsonl"]%s}},{"type":"max_file_size","parameters":{"max_file_size":10%s}}]`, exceptions, exceptions)
			if err := json.Unmarshal([]byte(body), &rules); err != nil {
				t.Fatal(err)
			}
			res := resourceGithubRepositoryRuleset()
			d := schema.TestResourceDataRaw(t, res.Schema, nil)
			if err := d.Set("rules", []any{map[string]any{
				"file_path_restriction": []any{map[string]any{"restricted_file_paths": []string{"**/*.jsonl"}, "ignored_file_paths": []string{"old/**"}}},
				"max_file_size":         []any{map[string]any{"max_file_size": 10, "ignored_file_paths": []string{"old/**"}}},
			}}); err != nil {
				t.Fatal(err)
			}
			if err := d.Set("rules", flattenRulesWithPathExceptions(t.Context(), &rules, false)); err != nil {
				t.Fatal(err)
			}
			for _, ruleType := range []string{"file_path_restriction", "max_file_size"} {
				if got := d.Get("rules.0." + ruleType + ".0.ignored_file_paths.#"); got != 0 {
					t.Errorf("%s retains stale exceptions: %v", ruleType, got)
				}
			}
		})
	}
}

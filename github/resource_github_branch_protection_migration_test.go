package github

import (
	"maps"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/shurcooL/githubv4"
)

func testGithubBranchProtectionStateDataV1() map[string]any {
	return map[string]any{
		"blocks_creations":  true,
		"push_restrictions": [...]string{"/example-user"},
	}
}

func testGithubBranchProtectionStateDataV2() map[string]any {
	restrictions := []any{map[string]any{
		"blocks_creations": true,
		"push_allowances":  [...]string{"/example-user"},
	}}
	return map[string]any{
		"restrict_pushes": restrictions,
	}
}

func Test_resourceGithubBranchProtectionStateUpgradeV1(t *testing.T) {
	// IMPORTANT: This test is not parallelized because it uses global state.

	for _, tt := range []struct {
		name     string
		rawState map[string]any
		expected map[string]any
	}{
		{
			name:     "true blocks_creations",
			rawState: testGithubBranchProtectionStateDataV1(),
			expected: testGithubBranchProtectionStateDataV2(),
		},
		{
			name: "null blocks_creations defaults to false",
			rawState: map[string]any{
				"blocks_creations":  nil,
				"push_restrictions": [...]string{"/example-user"},
			},
			expected: map[string]any{
				"restrict_pushes": []any{map[string]any{
					"blocks_creations": false,
					"push_allowances":  [...]string{"/example-user"},
				}},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := resourceGithubBranchProtectionUpgradeV1(t.Context(), tt.rawState, nil)
			if err != nil {
				t.Fatalf("error migrating state: %s", err)
			}

			if !reflect.DeepEqual(tt.expected, actual) {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", tt.expected, actual)
			}
		})
	}
}

func Test_resourceGithubBranchProtectionStateUpgradeV1_malformedBlocksCreations(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name  string
		value any
	}{
		{name: "string", value: "true"},
		{name: "int", value: 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rawState := map[string]any{
				"blocks_creations":  tt.value,
				"push_restrictions": []any{"/example-user"},
			}
			before := maps.Clone(rawState)

			got, err := resourceGithubBranchProtectionUpgradeV1(t.Context(), rawState, nil)
			if err == nil {
				t.Fatal("expected error for malformed blocks_creations, got nil")
			}
			if !strings.Contains(err.Error(), "blocks_creations") {
				t.Fatalf("expected error mentioning blocks_creations, got %q", err.Error())
			}
			if got != nil {
				t.Fatalf("expected nil state, got %#v", got)
			}
			if !reflect.DeepEqual(rawState, before) {
				t.Fatalf("expected input state unchanged: want %#v, got %#v", before, rawState)
			}
		})
	}
}

func Test_resourceGithubBranchProtectionStateUpgradeV0_unconfiguredMeta(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name string
		meta any
	}{
		{name: "nil meta", meta: nil},
		{name: "wrong meta type", meta: "not-an-owner"},
		{name: "nil owner pointer", meta: (*Owner)(nil)},
		{name: "owner without v4client", meta: &Owner{name: "test-org"}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rawState := map[string]any{
				"repository": "my-repo",
				"branch":     "main",
			}
			before := maps.Clone(rawState)

			got, err := resourceGithubBranchProtectionUpgradeV0(t.Context(), rawState, tt.meta)
			if err == nil {
				t.Fatal("expected error for unconfigured provider meta, got nil")
			}
			if got != nil {
				t.Fatalf("expected nil state, got %#v", got)
			}
			if !reflect.DeepEqual(rawState, before) {
				t.Fatalf("expected input state unchanged: want %#v, got %#v", before, rawState)
			}
		})
	}
}

func Test_resourceGithubBranchProtectionStateUpgradeV0_malformedState(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name         string
		rawState     map[string]any
		wantErr      string
		wantAPICalls int
	}{
		{
			name:         "missing repository",
			rawState:     map[string]any{"branch": "main"},
			wantErr:      "repository not found or is not a string",
			wantAPICalls: 0,
		},
		{
			name:         "nil repository",
			rawState:     map[string]any{"repository": nil, "branch": "main"},
			wantErr:      "repository not found or is not a string",
			wantAPICalls: 0,
		},
		{
			name:         "non-string repository",
			rawState:     map[string]any{"repository": 123, "branch": "main"},
			wantErr:      "repository not found or is not a string",
			wantAPICalls: 0,
		},
		{
			name:         "missing branch",
			rawState:     map[string]any{"repository": "repo-node-id"},
			wantErr:      "branch not found or is not a string",
			wantAPICalls: 1,
		},
		{
			name:         "nil branch",
			rawState:     map[string]any{"repository": "repo-node-id", "branch": nil},
			wantErr:      "branch not found or is not a string",
			wantAPICalls: 1,
		},
		{
			name:         "non-string branch",
			rawState:     map[string]any{"repository": "repo-node-id", "branch": 456},
			wantErr:      "branch not found or is not a string",
			wantAPICalls: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var apiCalls atomic.Int32
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				apiCalls.Add(1)
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"data":{"node":{"id":"repo-node-id"}}}`))
			}))
			t.Cleanup(srv.Close)

			meta := &Owner{
				name:     "test-org",
				v4client: githubv4.NewEnterpriseClient(srv.URL, srv.Client()),
			}
			before := maps.Clone(tt.rawState)

			got, err := resourceGithubBranchProtectionUpgradeV0(t.Context(), tt.rawState, meta)
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("expected error containing %q, got %q", tt.wantErr, err.Error())
			}
			if got != nil {
				t.Fatalf("expected nil state, got %#v", got)
			}
			if !reflect.DeepEqual(tt.rawState, before) {
				t.Fatalf("expected input state unchanged: want %#v, got %#v", before, tt.rawState)
			}
			if gotCalls := int(apiCalls.Load()); gotCalls != tt.wantAPICalls {
				t.Fatalf("expected %d API call(s), got %d", tt.wantAPICalls, gotCalls)
			}
		})
	}
}

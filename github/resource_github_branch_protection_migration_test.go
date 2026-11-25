package github

import (
	"maps"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
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

	expected := testGithubBranchProtectionStateDataV2()
	actual, err := resourceGithubBranchProtectionUpgradeV1(t.Context(), testGithubBranchProtectionStateDataV1(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
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

			apiCalls := 0
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				apiCalls++
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
			if apiCalls != tt.wantAPICalls {
				t.Fatalf("expected %d API call(s), got %d", tt.wantAPICalls, apiCalls)
			}
		})
	}
}

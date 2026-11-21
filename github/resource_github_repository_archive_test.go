package github

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// fakeArchivingRepoServer behaves like the GitHub REST API for one repository:
// once a PATCH archives it, every later PATCH fails with 403 because the
// repository is read-only.
type fakeArchivingRepoServer struct {
	mu         sync.Mutex
	archived   bool
	visibility string
	patches    []map[string]any
}

func (f *fakeArchivingRepoServer) handler(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.URL.Path == "/repos/test-owner/test-repo" && r.Method == http.MethodPatch:
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.patches = append(f.patches, body)
		if f.archived {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"message":"Repository was archived so is read-only."}`))
			return
		}
		if v, ok := body["visibility"].(string); ok {
			f.visibility = v
		}
		if a, ok := body["archived"].(bool); ok && a {
			f.archived = true
		}
		f.writeRepo(w)
	case r.URL.Path == "/repos/test-owner/test-repo" && r.Method == http.MethodGet:
		f.writeRepo(w)
	default:
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"Not Found"}`))
	}
}

func (f *fakeArchivingRepoServer) writeRepo(w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(map[string]any{
		"name":       "test-repo",
		"full_name":  "test-owner/test-repo",
		"visibility": f.visibility,
		"private":    f.visibility != "public",
		"archived":   f.archived,
		"owner":      map[string]any{"login": "test-owner"},
	})
}

func TestResourceGithubRepositoryUpdateArchivesLast(t *testing.T) {
	fake := &fakeArchivingRepoServer{visibility: "private"}
	ts := httptest.NewServer(http.HandlerFunc(fake.handler))
	defer ts.Close()

	baseURL := ts.URL + "/"
	client, err := github.NewClient(github.WithURLs(&baseURL, &baseURL))
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	meta := &Owner{name: "test-owner", v3client: client, IsOrganization: true}

	r := resourceGithubRepository()
	state := &terraform.InstanceState{
		ID: "test-repo",
		Attributes: map[string]string{
			"name":       "test-repo",
			"visibility": "private",
			"private":    "true",
			"archived":   "false",
		},
	}
	config := terraform.NewResourceConfigRaw(map[string]any{
		"name":       "test-repo",
		"visibility": "internal",
		"archived":   true,
	})
	diff, err := r.Diff(context.Background(), state, config, meta)
	if err != nil {
		t.Fatalf("diff: %v", err)
	}
	d, err := schema.InternalMap(r.SchemaMap()).Data(state, diff)
	if err != nil {
		t.Fatalf("data: %v", err)
	}

	if diags := resourceGithubRepositoryUpdate(context.Background(), d, meta); diags.HasError() {
		t.Fatalf("update failed: %v", diags)
	}

	if !fake.archived || fake.visibility != "internal" {
		t.Fatalf("got archived=%v visibility=%q, want archived=true visibility=\"internal\"", fake.archived, fake.visibility)
	}
	last := fake.patches[len(fake.patches)-1]
	if len(last) != 1 || last["archived"] != true {
		t.Errorf("last PATCH = %v, want only {\"archived\": true}", last)
	}
}

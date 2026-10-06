package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAgentsPublicKey(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/repos/test/repository/agents/secrets/public-key" {
			t.Errorf("unexpected public key scope: %s", r.URL.Path)
		}
		if _, err := fmt.Fprint(w, `{"key_id":"agent-key","key":"cHVibGljLWtleQ=="}`); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	meta := &Owner{name: "test", v3client: mustCreateTestGitHubClient(t, server.URL)}
	d := schema.TestResourceDataRaw(t, dataSourceGithubAgentsPublicKey().Schema, map[string]any{"repository": "repository"})
	if diagnostics := dataSourceGithubAgentsPublicKeyRead(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	if d.Id() != "agent-key" || d.Get("key_id") != "agent-key" || d.Get("key") != "cHVibGljLWtleQ==" {
		t.Fatal("public key metadata was not populated")
	}
}

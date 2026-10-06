package github

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAgentsOrganizationPublicKey(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/orgs/test/agents/secrets/public-key" {
			t.Errorf("unexpected public key scope: %s", r.URL.Path)
		}
		if _, err := fmt.Fprint(w, `{"key_id":"agent-key","key":"cHVibGljLWtleQ=="}`); err != nil {
			t.Error(err)
		}
	}))
	defer server.Close()
	meta := &Owner{name: "test", IsOrganization: true, v3client: mustCreateTestGitHubClient(t, server.URL)}
	d := schema.TestResourceDataRaw(t, dataSourceGithubAgentsOrganizationPublicKey().Schema, nil)
	if diagnostics := dataSourceGithubAgentsOrganizationPublicKeyRead(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	if d.Id() != "agent-key" || d.Get("key_id") != "agent-key" || d.Get("key") != "cHVibGljLWtleQ==" {
		t.Fatal("organization public key metadata was not populated")
	}
}

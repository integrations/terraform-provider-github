package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAgentsSecretsPagination(t *testing.T) {
	t.Parallel()
	meta := agentSecretsPaginationTestOwner(t)
	d := schema.TestResourceDataRaw(t, dataSourceGithubAgentsSecrets().Schema, map[string]any{"name": "repository"})
	if diagnostics := dataSourceGithubAgentsSecretsRead(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	secrets, _ := d.Get("secrets").([]any)
	if len(secrets) != 2 {
		t.Fatal("secret listing did not follow pagination")
	}
}

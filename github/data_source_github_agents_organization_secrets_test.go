package github

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestAgentsOrganizationSecretsPagination(t *testing.T) {
	t.Parallel()
	meta := agentSecretsPaginationTestOwner(t)
	d := schema.TestResourceDataRaw(t, dataSourceGithubAgentsOrganizationSecrets().Schema, nil)
	if diagnostics := dataSourceGithubAgentsOrganizationSecretsRead(t.Context(), d, meta); diagnostics.HasError() {
		t.Fatal(diagnostics)
	}
	secrets, _ := d.Get("secrets").([]any)
	if len(secrets) != 2 {
		t.Fatal("organization secret listing did not follow pagination")
	}
}

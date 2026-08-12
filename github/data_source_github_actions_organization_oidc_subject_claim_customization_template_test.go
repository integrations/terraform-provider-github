package github

import (
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsOrganizationOIDCSubjectClaimCustomizationTemplateDataSource(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		config := `
data "github_actions_organization_oidc_subject_claim_customization_template" "test" {}
`

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnlessHasOrgs(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_actions_organization_oidc_subject_claim_customization_template.test", tfjsonpath.New("include_claim_keys"), knownvalue.NotNull()),
					},
				},
				{
					PreConfig: func() {
						if _, err := testAccConf.meta.v3client.Actions.SetOrgOIDCSubjectClaimCustomTemplate(t.Context(), testAccConf.meta.name, github.OIDCSubjectClaimCustomTemplate{IncludeClaimKeys: []string{"actor", "actor_id", "head_ref", "repository"}}); err != nil {
							t.Fatalf("failed to set org OIDC subject claim custom template: %v", err)
						}
					},
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_actions_organization_oidc_subject_claim_customization_template.test", tfjsonpath.New("include_claim_keys"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("actor"),
							knownvalue.StringExact("actor_id"),
							knownvalue.StringExact("head_ref"),
							knownvalue.StringExact("repository"),
						})),
					},
					PostApplyFunc: func() {
						if _, err := testAccConf.meta.v3client.Actions.SetOrgOIDCSubjectClaimCustomTemplate(t.Context(), testAccConf.meta.name, github.OIDCSubjectClaimCustomTemplate{IncludeClaimKeys: []string{"repo", "context"}}); err != nil {
							t.Fatalf("failed to set org OIDC subject claim custom template: %v", err)
						}
					},
				},
			},
		})
	})
}

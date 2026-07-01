package github

import (
	"fmt"
	"testing"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccGithubActionsRepositoryOIDCSubjectClaimCustomizationTemplateDataSource(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		repo := mustCreateTestRepository(t)

		config := fmt.Sprintf(`
data "github_actions_repository_oidc_subject_claim_customization_template" "test" {
  name = "%s"
}
`, repo.GetName())

		resource.Test(t, resource.TestCase{
			PreCheck:          func() { skipUnauthenticated(t) },
			ProviderFactories: providerFactories,
			Steps: []resource.TestStep{
				{
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_actions_repository_oidc_subject_claim_customization_template.test", tfjsonpath.New("use_default"), knownvalue.Bool(true)),
						statecheck.ExpectKnownValue("data.github_actions_repository_oidc_subject_claim_customization_template.test", tfjsonpath.New("include_claim_keys"), knownvalue.NotNull()),
					},
				},
				{
					PreConfig: func() {
						if _, err := testAccConf.meta.v3client.Actions.SetRepoOIDCSubjectClaimCustomTemplate(t.Context(), testAccConf.meta.name, repo.GetName(), &github.OIDCSubjectClaimCustomTemplate{UseDefault: new(false), IncludeClaimKeys: []string{"actor", "actor_id", "head_ref", "repository"}}); err != nil {
							t.Fatalf("failed to set repo OIDC subject claim custom template: %v", err)
						}
					},
					Config: config,
					ConfigStateChecks: []statecheck.StateCheck{
						statecheck.ExpectKnownValue("data.github_actions_repository_oidc_subject_claim_customization_template.test", tfjsonpath.New("use_default"), knownvalue.Bool(false)),
						statecheck.ExpectKnownValue("data.github_actions_repository_oidc_subject_claim_customization_template.test", tfjsonpath.New("include_claim_keys"), knownvalue.ListExact([]knownvalue.Check{
							knownvalue.StringExact("actor"),
							knownvalue.StringExact("actor_id"),
							knownvalue.StringExact("head_ref"),
							knownvalue.StringExact("repository"),
						})),
					},
				},
			},
		})
	})
}

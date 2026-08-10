package github

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

type createTestEnterpriseRulesetOptionsFunc func(*github.RepositoryRuleset)

// mustCreateTestEnterpriseRuleset creates a minimal active branch ruleset on the test
// enterprise and removes it again once the test finishes.
func mustCreateTestEnterpriseRuleset(t *testing.T, f ...createTestEnterpriseRulesetOptionsFunc) *github.RepositoryRuleset {
	t.Helper()

	name := fmt.Sprintf("%s%s", testResourcePrefix, acctest.RandString(testRandomIDLength))
	target := github.RulesetTargetBranch
	sourceType := github.RulesetSourceTypeEnterprise

	req := github.RepositoryRuleset{
		Name:        name,
		Target:      &target,
		Source:      testAccConf.enterpriseSlug,
		SourceType:  &sourceType,
		Enforcement: github.RulesetEnforcementActive,
		Conditions: &github.RepositoryRulesetConditions{
			OrganizationName: &github.RepositoryRulesetOrganizationNamesConditionParameters{
				Include: []string{"~ALL"},
				Exclude: []string{},
			},
			RepositoryName: &github.RepositoryRulesetRepositoryNamesConditionParameters{
				Include: []string{"~ALL"},
				Exclude: []string{},
			},
			RefName: &github.RepositoryRulesetRefConditionParameters{
				Include: []string{"~ALL"},
				Exclude: []string{},
			},
		},
		Rules: &github.RepositoryRulesetRules{
			Creation: &github.EmptyRuleParameters{},
		},
	}

	for _, fn := range f {
		if fn != nil {
			fn(&req)
		}
	}

	ruleset, _, err := testAccConf.meta.v3client.Enterprise.CreateRepositoryRuleset(t.Context(), testAccConf.enterpriseSlug, req)
	if err != nil {
		t.Fatalf("failed to create test enterprise ruleset: %v", err)
	}

	t.Cleanup(func() {
		if _, err := testAccConf.meta.v3client.Enterprise.DeleteRepositoryRuleset(context.Background(), testAccConf.enterpriseSlug, ruleset.GetID()); err != nil {
			if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
				return
			}
			t.Logf("failed to delete test enterprise ruleset %s: %v", ruleset.Name, err)
		}
	})

	return ruleset
}

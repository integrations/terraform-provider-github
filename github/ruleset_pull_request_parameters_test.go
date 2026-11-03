package github

import (
	"encoding/json"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestRulesetPullRequestProtectionParameters(t *testing.T) {
	for _, org := range []bool{false, true} {
		for _, extraApproval := range []bool{false, true} {
			t.Run(testRulesetProtectionName(org, extraApproval), func(t *testing.T) {
				resource := resourceGithubRepositoryRuleset()
				if org {
					resource = resourceGithubOrganizationRuleset()
				}
				pullRequest := map[string]any{
					"require_extra_approval_for_unattributed_changes": extraApproval,
					"required_approving_review_count":                 1,
					"require_code_owner_review":                       true,
					"allowed_merge_methods":                           []any{"squash"},
					"dismissal_restriction": []any{map[string]any{
						"enabled": true,
						"allowed_actors": []any{map[string]any{
							"actor_id": 1234, "actor_type": "IntegrationInstallation",
						}},
					}},
				}
				data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
					"rules": []any{map[string]any{"pull_request": []any{pullRequest}}},
				})
				expanded := expandRules(data.Get("rules").([]any), org)
				if expanded.PullRequest.RequireExtraApprovalForUnattributedChanges != extraApproval {
					t.Fatalf("extra approval was not expanded: %+v", expanded.PullRequest)
				}
				if expanded.PullRequest.DismissalRestriction == nil || !expanded.PullRequest.DismissalRestriction.Enabled {
					t.Fatal("enabled dismissal restriction was not expanded")
				}
				actors := expanded.PullRequest.DismissalRestriction.AllowedActors
				if len(actors) != 1 || actors[0].ID != 1234 || actors[0].Type != github.DismissalRestrictionActorTypeIntegrationInstallation {
					t.Fatalf("dismissal actor identity lost: %+v", actors)
				}
				payload, err := json.Marshal(expanded.PullRequest)
				if err != nil {
					t.Fatal(err)
				}
				var encoded map[string]any
				if err := json.Unmarshal(payload, &encoded); err != nil {
					t.Fatal(err)
				}
				if encoded["require_extra_approval_for_unattributed_changes"] != extraApproval {
					t.Fatalf("API payload lost extra approval: %s", payload)
				}
				flattened := flattenRules(t.Context(), expanded, org)[0].(map[string]any)["pull_request"].([]map[string]any)[0]
				if flattened["require_extra_approval_for_unattributed_changes"] != extraApproval {
					t.Fatal("extra approval was not flattened")
				}
				restrictions := flattened["dismissal_restriction"].([]map[string]any)
				if len(restrictions) != 1 || restrictions[0]["enabled"] != true {
					t.Fatal("dismissal restriction was not flattened")
				}
			})
		}
	}
}

func testRulesetProtectionName(org, extraApproval bool) string {
	name := "repository"
	if org {
		name = "organization"
	}
	if extraApproval {
		return name + "_extra_approval"
	}
	return name + "_no_extra_approval"
}

func TestRulesetDismissalRestrictionUnsetDoesNotDrift(t *testing.T) {
	for _, restriction := range []*github.DismissalRestriction{nil, {Enabled: false}} {
		for _, org := range []bool{false, true} {
			flattened := flattenRules(t.Context(), &github.RepositoryRulesetRules{
				PullRequest: &github.PullRequestRuleParameters{DismissalRestriction: restriction},
			}, org)[0].(map[string]any)["pull_request"].([]map[string]any)[0]
			value, exists := flattened["dismissal_restriction"]
			if exists && len(value.([]map[string]any)) != 0 {
				t.Fatalf("unset restriction must not introduce a block: %+v", value)
			}
		}
	}
}

func TestRulesetDismissalRestrictionWithNoActors(t *testing.T) {
	for _, org := range []bool{false, true} {
		resource := resourceGithubRepositoryRuleset()
		if org {
			resource = resourceGithubOrganizationRuleset()
		}
		data := schema.TestResourceDataRaw(t, resource.Schema, map[string]any{
			"rules": []any{map[string]any{"pull_request": []any{map[string]any{
				"dismissal_restriction": []any{map[string]any{"enabled": true}},
			}}}},
		})
		expanded := expandRules(data.Get("rules").([]any), org)
		if expanded.PullRequest.DismissalRestriction == nil || !expanded.PullRequest.DismissalRestriction.Enabled {
			t.Fatal("enabled restriction without actors was lost")
		}
		if len(expanded.PullRequest.DismissalRestriction.AllowedActors) != 0 {
			t.Fatal("unexpected dismissal actor")
		}
		flattened := flattenRules(t.Context(), expanded, org)[0].(map[string]any)["pull_request"].([]map[string]any)[0]
		if len(flattened["dismissal_restriction"].([]map[string]any)) != 1 {
			t.Fatal("enabled restriction without actors was not flattened")
		}
	}
}

func TestRulesetDismissalActorTypes(t *testing.T) {
	for _, resource := range []*schema.Resource{resourceGithubRepositoryRuleset(), resourceGithubOrganizationRuleset()} {
		pullRequest := resource.Schema["rules"].Elem.(*schema.Resource).Schema["pull_request"].Elem.(*schema.Resource)
		restriction, exists := pullRequest.Schema["dismissal_restriction"]
		if !exists {
			t.Fatal("dismissal restriction schema missing")
		}
		actor := restriction.Elem.(*schema.Resource).Schema["allowed_actors"].Elem.(*schema.Resource)
		validate := actor.Schema["actor_type"].ValidateDiagFunc
		for _, valid := range []string{"User", "Team", "IntegrationInstallation", "RepositoryRole"} {
			if diags := validate(valid, nil); diags.HasError() {
				t.Fatalf("valid dismissal actor %q rejected: %v", valid, diags)
			}
		}
		if diags := validate("Integration", nil); !diags.HasError() {
			t.Fatal("bypass-only actor type Integration accepted for dismissal")
		}
	}
}

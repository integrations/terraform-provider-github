package github

import (
	"encoding/json"
	"testing"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestRulesetPullRequestImportUpdateRoundTrip(t *testing.T) {
	// Anonymous fixture: no organization policy or credentials in upstream tests.
	payload := []byte(`{"rules":[{"type":"deletion"},{"type":"non_fast_forward"},{"type":"pull_request","parameters":{"required_approving_review_count":1,"dismiss_stale_reviews_on_push":false,"require_code_owner_review":true,"require_last_push_approval":false,"required_review_thread_resolution":false,"require_extra_approval_for_unattributed_changes":true,"allowed_merge_methods":["merge","squash","rebase"],"dismissal_restriction":{"enabled":true,"allowed_actors":[{"id":1234,"type":"Team"}]}}}]}`)
	var live github.RepositoryRuleset
	if err := json.Unmarshal(payload, &live); err != nil {
		t.Fatal(err)
	}
	for _, org := range []bool{false, true} {
		flattened := flattenRules(t.Context(), live.Rules, org)
		resource := resourceGithubRepositoryRuleset()
		if org {
			resource = resourceGithubOrganizationRuleset()
		}
		data := schema.TestResourceDataRaw(t, resource.Schema, nil)
		if err := data.Set("rules", flattened); err != nil {
			t.Fatal(err)
		}
		expanded := expandRules(data.Get("rules").([]any), org)
		before, err := json.Marshal(live.Rules.PullRequest)
		if err != nil {
			t.Fatal(err)
		}
		after, err := json.Marshal(expanded.PullRequest)
		if err != nil {
			t.Fatal(err)
		}
		if string(before) != string(after) {
			t.Fatalf("import/update changed protection parameters:\nbefore %s\nafter  %s", before, after)
		}
		if expanded.Deletion == nil || expanded.NonFastForward == nil {
			t.Fatal("ref safety rules were lost")
		}
	}
}

package github

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// TestResourceGithubRepositoryEnvironmentRead_OutOfBandReviewerRemoval reproduces
// https://github.com/integrations/terraform-provider-github/issues/3609: once the last
// required reviewer is removed outside Terraform, the API stops returning a
// "required_reviewers" entry in protection_rules. reviewers and prevent_self_review are
// only ever set inside that case, so without a pre-loop reset the prior state values
// survive a refresh untouched.
func TestResourceGithubRepositoryEnvironmentRead_OutOfBandReviewerRemoval(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/repos/test-owner/test-repo/environments/production":
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{
				"id": 1,
				"name": "production",
				"can_admins_bypass": true,
				"protection_rules": [],
				"deployment_branch_policy": null
			}`)
			return
		case "/users/test-owner":
			w.WriteHeader(http.StatusOK)
			_, _ = fmt.Fprint(w, `{"login": "test-owner", "id": 1, "type": "Organization"}`)
			return
		}

		w.WriteHeader(http.StatusNotFound)
	}))
	t.Cleanup(ts.Close)

	meta, err := configureProviderMeta(context.Background(), "test", &Config{
		LegacyClient: true,
		Owner:        "test-owner",
		Token:        "test-token",
		BaseURL:      mustNewURL(t, ts.URL),
	})
	if err != nil {
		t.Fatalf("failed to configure provider meta: %s", err)
	}

	d := schema.TestResourceDataRaw(t, resourceGithubRepositoryEnvironment().Schema, map[string]any{
		"repository":          "test-repo",
		"environment":         "production",
		"prevent_self_review": true,
		"reviewers": []any{
			map[string]any{
				"users": []any{78364628},
			},
		},
	})

	if diags := resourceGithubRepositoryEnvironmentRead(context.Background(), d, meta); diags.HasError() {
		t.Fatalf("unexpected error from read: %v", diags)
	}

	if got := d.Get("prevent_self_review").(bool); got {
		t.Errorf("prevent_self_review = %v, want false once the required_reviewers rule is removed out-of-band", got)
	}

	if reviewers := d.Get("reviewers").([]any); len(reviewers) != 0 {
		t.Errorf("reviewers = %v, want empty once the required_reviewers rule is removed out-of-band", reviewers)
	}
}

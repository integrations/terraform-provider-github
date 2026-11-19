package ghratelimit

import (
	"net/http"
	"testing"
)

func Test_lookupRateLimitCategory(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		method string
		path   string
		want   string
	}{
		// default
		{name: "default_no_match", method: http.MethodGet, path: "/repos/o/r", want: rateLimitCategoryCore},
		{name: "default_empty", method: "", path: "", want: rateLimitCategoryCore},

		// anchoring: segments in the middle of a path must not match
		{name: "anchored_graphql_as_repo_name", method: http.MethodGet, path: "/repos/graphql-go/graphql", want: rateLimitCategoryCore},
		{name: "anchored_search_as_repo_name", method: http.MethodGet, path: "/repos/o/search", want: rateLimitCategoryCore},
		{name: "anchored_search_in_contents_path", method: http.MethodGet, path: "/repos/o/r/contents/src/search/index.ts", want: rateLimitCategoryCore},

		// actions_runner_registration: POST only; repos|orgs|enterprises; registration|remove
		{name: "actions_runner_registration_repos_registration", method: http.MethodPost, path: "/repos/o/r/actions/runners/registration-token", want: "actions_runner_registration"},
		{name: "actions_runner_registration_repos_remove", method: http.MethodPost, path: "/repos/o/r/actions/runners/remove-token", want: "actions_runner_registration"},
		{name: "actions_runner_registration_orgs_registration", method: http.MethodPost, path: "/orgs/o/actions/runners/registration-token", want: "actions_runner_registration"},
		{name: "actions_runner_registration_orgs_remove", method: http.MethodPost, path: "/orgs/o/actions/runners/remove-token", want: "actions_runner_registration"},
		{name: "actions_runner_registration_enterprises_registration", method: http.MethodPost, path: "/enterprises/e/actions/runners/registration-token", want: "actions_runner_registration"},
		{name: "actions_runner_registration_enterprises_remove", method: http.MethodPost, path: "/enterprises/e/actions/runners/remove-token", want: "actions_runner_registration"},
		{name: "actions_runner_registration_method_mismatch", method: http.MethodGet, path: "/orgs/o/actions/runners/registration-token", want: rateLimitCategoryCore},
		{name: "actions_runner_registration_path_mismatch", method: http.MethodPost, path: "/orgs/o/actions/runners", want: rateLimitCategoryCore},

		// audit_log: GET only; orgs|enterprises; optional sub path
		{name: "audit_log_orgs", method: http.MethodGet, path: "/orgs/o/audit-log", want: "audit_log"},
		{name: "audit_log_enterprises", method: http.MethodGet, path: "/enterprises/e/audit-log", want: "audit_log"},
		{name: "audit_log_sub_path", method: http.MethodGet, path: "/orgs/o/audit-log/extra", want: "audit_log"},
		{name: "audit_log_method_mismatch_post", method: http.MethodPost, path: "/orgs/o/audit-log", want: rateLimitCategoryCore},
		{name: "audit_log_method_mismatch_head", method: http.MethodHead, path: "/orgs/o/audit-log", want: rateLimitCategoryCore},
		{name: "audit_log_path_mismatch", method: http.MethodGet, path: "/repos/o/audit-log", want: rateLimitCategoryCore},

		// code_scanning_upload: POST only
		{name: "code_scanning_upload", method: http.MethodPost, path: "/repos/o/r/code-scanning/sarifs", want: "code_scanning_upload"},
		{name: "code_scanning_upload_method_mismatch", method: http.MethodGet, path: "/repos/o/r/code-scanning/sarifs", want: rateLimitCategoryCore},
		{name: "code_scanning_upload_path_mismatch", method: http.MethodPost, path: "/repos/o/r/code-scanning/sarifs/1", want: rateLimitCategoryCore},

		// code_search: GET or HEAD; optional sub path; takes precedence over search
		{name: "code_search_get", method: http.MethodGet, path: "/search/code", want: "code_search"},
		{name: "code_search_head", method: http.MethodHead, path: "/search/code", want: "code_search"},
		{name: "code_search_sub_path", method: http.MethodGet, path: "/search/code/extra", want: "code_search"},
		{name: "code_search_method_mismatch", method: http.MethodPost, path: "/search/code", want: rateLimitCategoryCore},
		{name: "code_search_path_mismatch_falls_to_search", method: http.MethodGet, path: "/search/codeql", want: "search"},

		// copilot_usage_records: GET only; orgs|enterprises; usage|billing; optional sub path
		{name: "copilot_usage_records_orgs_usage", method: http.MethodGet, path: "/orgs/o/copilot/usage", want: "copilot_usage_records"},
		{name: "copilot_usage_records_orgs_billing", method: http.MethodGet, path: "/orgs/o/copilot/billing", want: "copilot_usage_records"},
		{name: "copilot_usage_records_enterprises_usage", method: http.MethodGet, path: "/enterprises/e/copilot/usage", want: "copilot_usage_records"},
		{name: "copilot_usage_records_enterprises_billing", method: http.MethodGet, path: "/enterprises/e/copilot/billing", want: "copilot_usage_records"},
		{name: "copilot_usage_records_sub_path", method: http.MethodGet, path: "/orgs/o/copilot/usage/extra", want: "copilot_usage_records"},
		{name: "copilot_usage_records_method_mismatch_post", method: http.MethodPost, path: "/orgs/o/copilot/usage", want: rateLimitCategoryCore},
		{name: "copilot_usage_records_method_mismatch_head", method: http.MethodHead, path: "/orgs/o/copilot/usage", want: rateLimitCategoryCore},
		{name: "copilot_usage_records_path_mismatch", method: http.MethodGet, path: "/orgs/o/copilot/other", want: rateLimitCategoryCore},

		// dependency_sbom: GET only
		{name: "dependency_sbom", method: http.MethodGet, path: "/repos/o/r/dependency-graph/sbom", want: "dependency_sbom"},
		{name: "dependency_sbom_method_mismatch_post", method: http.MethodPost, path: "/repos/o/r/dependency-graph/sbom", want: rateLimitCategoryCore},
		{name: "dependency_sbom_method_mismatch_head", method: http.MethodHead, path: "/repos/o/r/dependency-graph/sbom", want: rateLimitCategoryCore},
		{name: "dependency_sbom_path_mismatch", method: http.MethodGet, path: "/repos/o/r/dependency-graph/sbom/extra", want: rateLimitCategoryCore},

		// dependency_snapshots: POST only
		{name: "dependency_snapshots", method: http.MethodPost, path: "/repos/o/r/dependency-graph/snapshots", want: "dependency_snapshots"},
		{name: "dependency_snapshots_method_mismatch", method: http.MethodGet, path: "/repos/o/r/dependency-graph/snapshots", want: rateLimitCategoryCore},
		{name: "dependency_snapshots_path_mismatch", method: http.MethodPost, path: "/repos/o/r/dependency-graph/snapshots/1", want: rateLimitCategoryCore},

		// graphql: any method
		{name: "graphql_get", method: http.MethodGet, path: "/graphql", want: "graphql"},
		{name: "graphql_post", method: http.MethodPost, path: "/graphql", want: "graphql"},
		{name: "graphql_empty_method", method: "", path: "/graphql", want: "graphql"},
		{name: "graphql_path_mismatch", method: http.MethodPost, path: "/graphql/extra", want: rateLimitCategoryCore},

		// integration_manifest: POST only
		{name: "integration_manifest", method: http.MethodPost, path: "/app-manifests/code/conversions", want: "integration_manifest"},
		{name: "integration_manifest_method_mismatch", method: http.MethodGet, path: "/app-manifests/code/conversions", want: rateLimitCategoryCore},
		{name: "integration_manifest_path_mismatch", method: http.MethodPost, path: "/app-manifests/conversions", want: rateLimitCategoryCore},

		// scim: any method; optional sub path
		{name: "scim_root_get", method: http.MethodGet, path: "/scim/v2", want: "scim"},
		{name: "scim_sub_path_get", method: http.MethodGet, path: "/scim/v2/extra", want: "scim"},
		{name: "scim_sub_path_post", method: http.MethodPost, path: "/scim/v2/extra", want: "scim"},
		{name: "scim_empty_method", method: "", path: "/scim/v2", want: "scim"},
		{name: "scim_path_mismatch", method: http.MethodGet, path: "/scim/v22", want: rateLimitCategoryCore},

		// search: GET or HEAD; optional sub path
		{name: "search_root_get", method: http.MethodGet, path: "/search", want: "search"},
		{name: "search_root_head", method: http.MethodHead, path: "/search", want: "search"},
		{name: "search_sub_path", method: http.MethodGet, path: "/search/issues", want: "search"},
		{name: "search_method_mismatch", method: http.MethodPost, path: "/search/issues", want: rateLimitCategoryCore},
		{name: "search_path_mismatch", method: http.MethodGet, path: "/searches", want: rateLimitCategoryCore},

		// source_import: any method; optional sub path
		{name: "source_import_get", method: http.MethodGet, path: "/repos/o/r/import", want: "source_import"},
		{name: "source_import_put", method: http.MethodPut, path: "/repos/o/r/import", want: "source_import"},
		{name: "source_import_empty_method", method: "", path: "/repos/o/r/import", want: "source_import"},
		{name: "source_import_sub_path", method: http.MethodGet, path: "/repos/o/r/import/authors", want: "source_import"},
		{name: "source_import_path_mismatch", method: http.MethodGet, path: "/repos/o/r/imports", want: rateLimitCategoryCore},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := lookupRateLimitCategory(tt.method, tt.path)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}

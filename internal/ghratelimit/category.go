package ghratelimit

import (
	"net/http"
	"regexp"
)

const (
	rateLimitCategoryActionsRunnerRegistration = "actions_runner_registration"
	rateLimitCategoryAuditLog                  = "audit_log"
	rateLimitCategoryCodeScanningUpload        = "code_scanning_upload"
	rateLimitCategoryCodeSearch                = "code_search"
	rateLimitCategoryCopilotUsageRecords       = "copilot_usage_records"
	rateLimitCategoryCore                      = "core"
	rateLimitCategoryDependencySbom            = "dependency_sbom"
	rateLimitCategoryDependencySnapshots       = "dependency_snapshots"
	rateLimitCategoryGraphQL                   = "graphql"
	rateLimitCategoryIntegrationManifest       = "integration_manifest"
	rateLimitCategorySCIM                      = "scim"
	rateLimitCategorySearch                    = "search"
	rateLimitCategorySecondary                 = "secondary"
	rateLimitCategorySourceImport              = "source_import"
)

var (
	reRateLimitCategoryActionsRunnerRegistration = regexp.MustCompile(`^/(?:repos/[^/]+/[^/]+|orgs/[^/]+|enterprises/[^/]+)/actions/runners/(?:registration|remove)-token$`)
	reRateLimitCategoryAuditLog                  = regexp.MustCompile(`^/(?:orgs|enterprises)/[^/]+/audit-log(?:/.*)?$`)
	reRateLimitCategoryCodeScanningUpload        = regexp.MustCompile(`^/repos/[^/]+/[^/]+/code-scanning/sarifs$`)
	reRateLimitCategoryCodeSearch                = regexp.MustCompile(`^/search/code(?:/.*)?$`)
	reRateLimitCategoryCopilotUsageRecords       = regexp.MustCompile(`^/(?:orgs|enterprises)/[^/]+/copilot/(?:usage|billing)(?:/.*)?$`)
	reRateLimitCategoryDependencySbom            = regexp.MustCompile(`^/repos/[^/]+/[^/]+/dependency-graph/sbom$`)
	reRateLimitCategoryDependencySnapshots       = regexp.MustCompile(`^/repos/[^/]+/[^/]+/dependency-graph/snapshots$`)
	reRateLimitCategoryGraphQL                   = regexp.MustCompile(`^/graphql$`)
	reRateLimitCategoryIntegrationManifest       = regexp.MustCompile(`^/app-manifests/[^/]+/conversions$`)
	reRateLimitCategorySCIM                      = regexp.MustCompile(`^/scim/v2(?:/.*)?$`)
	reRateLimitCategorySearch                    = regexp.MustCompile(`^/search(?:/.*)?$`)
	reRateLimitCategorySourceImport              = regexp.MustCompile(`^/repos/[^/]+/[^/]+/import(?:/.*)?$`)
)

// lookupRateLimitCategory returns the rate limit category for the given HTTP method and path, based on predefined patterns. It returns a string representing the category, which can be used to determine how to handle rate limiting for that request.
func lookupRateLimitCategory(method, path string) string {
	switch {
	case method == http.MethodPost && reRateLimitCategoryActionsRunnerRegistration.MatchString(path):
		return rateLimitCategoryActionsRunnerRegistration

	case method == http.MethodGet && reRateLimitCategoryAuditLog.MatchString(path):
		return rateLimitCategoryAuditLog

	case method == http.MethodPost && reRateLimitCategoryCodeScanningUpload.MatchString(path):
		return rateLimitCategoryCodeScanningUpload

	case (method == http.MethodGet || method == http.MethodHead) && reRateLimitCategoryCodeSearch.MatchString(path):
		return rateLimitCategoryCodeSearch

	case method == http.MethodGet && reRateLimitCategoryCopilotUsageRecords.MatchString(path):
		return rateLimitCategoryCopilotUsageRecords

	case method == http.MethodGet && reRateLimitCategoryDependencySbom.MatchString(path):
		return rateLimitCategoryDependencySbom

	case method == http.MethodPost && reRateLimitCategoryDependencySnapshots.MatchString(path):
		return rateLimitCategoryDependencySnapshots

	case reRateLimitCategoryGraphQL.MatchString(path):
		return rateLimitCategoryGraphQL

	case method == http.MethodPost && reRateLimitCategoryIntegrationManifest.MatchString(path):
		return rateLimitCategoryIntegrationManifest

	case reRateLimitCategorySCIM.MatchString(path):
		return rateLimitCategorySCIM

	case (method == http.MethodGet || method == http.MethodHead) && reRateLimitCategorySearch.MatchString(path):
		return rateLimitCategorySearch

	case reRateLimitCategorySourceImport.MatchString(path):
		return rateLimitCategorySourceImport

	default:
		return rateLimitCategoryCore
	}
}

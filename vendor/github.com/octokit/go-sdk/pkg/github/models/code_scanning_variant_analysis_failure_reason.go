package models
import (
    "errors"
)
// The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
type CodeScanningVariantAnalysis_failure_reason int

const (
    NO_REPOS_QUERIED_CODESCANNINGVARIANTANALYSIS_FAILURE_REASON CodeScanningVariantAnalysis_failure_reason = iota
    ACTIONS_WORKFLOW_RUN_FAILED_CODESCANNINGVARIANTANALYSIS_FAILURE_REASON
    INTERNAL_ERROR_CODESCANNINGVARIANTANALYSIS_FAILURE_REASON
)

func (i CodeScanningVariantAnalysis_failure_reason) String() string {
    return []string{"no_repos_queried", "actions_workflow_run_failed", "internal_error"}[i]
}
func ParseCodeScanningVariantAnalysis_failure_reason(v string) (any, error) {
    result := NO_REPOS_QUERIED_CODESCANNINGVARIANTANALYSIS_FAILURE_REASON
    switch v {
        case "no_repos_queried":
            result = NO_REPOS_QUERIED_CODESCANNINGVARIANTANALYSIS_FAILURE_REASON
        case "actions_workflow_run_failed":
            result = ACTIONS_WORKFLOW_RUN_FAILED_CODESCANNINGVARIANTANALYSIS_FAILURE_REASON
        case "internal_error":
            result = INTERNAL_ERROR_CODESCANNINGVARIANTANALYSIS_FAILURE_REASON
        default:
            return 0, errors.New("Unknown CodeScanningVariantAnalysis_failure_reason value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningVariantAnalysis_failure_reason(values []CodeScanningVariantAnalysis_failure_reason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningVariantAnalysis_failure_reason) isMultiValue() bool {
    return false
}

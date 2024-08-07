package models
import (
    "errors"
)
// The new status of the CodeQL variant analysis repository task.
type CodeScanningVariantAnalysisStatus int

const (
    PENDING_CODESCANNINGVARIANTANALYSISSTATUS CodeScanningVariantAnalysisStatus = iota
    IN_PROGRESS_CODESCANNINGVARIANTANALYSISSTATUS
    SUCCEEDED_CODESCANNINGVARIANTANALYSISSTATUS
    FAILED_CODESCANNINGVARIANTANALYSISSTATUS
    CANCELED_CODESCANNINGVARIANTANALYSISSTATUS
    TIMED_OUT_CODESCANNINGVARIANTANALYSISSTATUS
)

func (i CodeScanningVariantAnalysisStatus) String() string {
    return []string{"pending", "in_progress", "succeeded", "failed", "canceled", "timed_out"}[i]
}
func ParseCodeScanningVariantAnalysisStatus(v string) (any, error) {
    result := PENDING_CODESCANNINGVARIANTANALYSISSTATUS
    switch v {
        case "pending":
            result = PENDING_CODESCANNINGVARIANTANALYSISSTATUS
        case "in_progress":
            result = IN_PROGRESS_CODESCANNINGVARIANTANALYSISSTATUS
        case "succeeded":
            result = SUCCEEDED_CODESCANNINGVARIANTANALYSISSTATUS
        case "failed":
            result = FAILED_CODESCANNINGVARIANTANALYSISSTATUS
        case "canceled":
            result = CANCELED_CODESCANNINGVARIANTANALYSISSTATUS
        case "timed_out":
            result = TIMED_OUT_CODESCANNINGVARIANTANALYSISSTATUS
        default:
            return 0, errors.New("Unknown CodeScanningVariantAnalysisStatus value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningVariantAnalysisStatus(values []CodeScanningVariantAnalysisStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningVariantAnalysisStatus) isMultiValue() bool {
    return false
}

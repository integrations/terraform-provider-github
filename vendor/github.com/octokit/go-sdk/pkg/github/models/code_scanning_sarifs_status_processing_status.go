package models
import (
    "errors"
)
// `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
type CodeScanningSarifsStatus_processing_status int

const (
    PENDING_CODESCANNINGSARIFSSTATUS_PROCESSING_STATUS CodeScanningSarifsStatus_processing_status = iota
    COMPLETE_CODESCANNINGSARIFSSTATUS_PROCESSING_STATUS
    FAILED_CODESCANNINGSARIFSSTATUS_PROCESSING_STATUS
)

func (i CodeScanningSarifsStatus_processing_status) String() string {
    return []string{"pending", "complete", "failed"}[i]
}
func ParseCodeScanningSarifsStatus_processing_status(v string) (any, error) {
    result := PENDING_CODESCANNINGSARIFSSTATUS_PROCESSING_STATUS
    switch v {
        case "pending":
            result = PENDING_CODESCANNINGSARIFSSTATUS_PROCESSING_STATUS
        case "complete":
            result = COMPLETE_CODESCANNINGSARIFSSTATUS_PROCESSING_STATUS
        case "failed":
            result = FAILED_CODESCANNINGSARIFSSTATUS_PROCESSING_STATUS
        default:
            return 0, errors.New("Unknown CodeScanningSarifsStatus_processing_status value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningSarifsStatus_processing_status(values []CodeScanningSarifsStatus_processing_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningSarifsStatus_processing_status) isMultiValue() bool {
    return false
}

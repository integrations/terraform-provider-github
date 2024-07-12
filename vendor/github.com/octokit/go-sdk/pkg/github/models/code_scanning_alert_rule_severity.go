package models
import (
    "errors"
)
// The severity of the alert.
type CodeScanningAlertRule_severity int

const (
    NONE_CODESCANNINGALERTRULE_SEVERITY CodeScanningAlertRule_severity = iota
    NOTE_CODESCANNINGALERTRULE_SEVERITY
    WARNING_CODESCANNINGALERTRULE_SEVERITY
    ERROR_CODESCANNINGALERTRULE_SEVERITY
)

func (i CodeScanningAlertRule_severity) String() string {
    return []string{"none", "note", "warning", "error"}[i]
}
func ParseCodeScanningAlertRule_severity(v string) (any, error) {
    result := NONE_CODESCANNINGALERTRULE_SEVERITY
    switch v {
        case "none":
            result = NONE_CODESCANNINGALERTRULE_SEVERITY
        case "note":
            result = NOTE_CODESCANNINGALERTRULE_SEVERITY
        case "warning":
            result = WARNING_CODESCANNINGALERTRULE_SEVERITY
        case "error":
            result = ERROR_CODESCANNINGALERTRULE_SEVERITY
        default:
            return 0, errors.New("Unknown CodeScanningAlertRule_severity value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertRule_severity(values []CodeScanningAlertRule_severity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertRule_severity) isMultiValue() bool {
    return false
}

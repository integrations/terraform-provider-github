package models
import (
    "errors"
)
// Severity of a code scanning alert.
type CodeScanningAlertSeverity int

const (
    CRITICAL_CODESCANNINGALERTSEVERITY CodeScanningAlertSeverity = iota
    HIGH_CODESCANNINGALERTSEVERITY
    MEDIUM_CODESCANNINGALERTSEVERITY
    LOW_CODESCANNINGALERTSEVERITY
    WARNING_CODESCANNINGALERTSEVERITY
    NOTE_CODESCANNINGALERTSEVERITY
    ERROR_CODESCANNINGALERTSEVERITY
)

func (i CodeScanningAlertSeverity) String() string {
    return []string{"critical", "high", "medium", "low", "warning", "note", "error"}[i]
}
func ParseCodeScanningAlertSeverity(v string) (any, error) {
    result := CRITICAL_CODESCANNINGALERTSEVERITY
    switch v {
        case "critical":
            result = CRITICAL_CODESCANNINGALERTSEVERITY
        case "high":
            result = HIGH_CODESCANNINGALERTSEVERITY
        case "medium":
            result = MEDIUM_CODESCANNINGALERTSEVERITY
        case "low":
            result = LOW_CODESCANNINGALERTSEVERITY
        case "warning":
            result = WARNING_CODESCANNINGALERTSEVERITY
        case "note":
            result = NOTE_CODESCANNINGALERTSEVERITY
        case "error":
            result = ERROR_CODESCANNINGALERTSEVERITY
        default:
            return 0, errors.New("Unknown CodeScanningAlertSeverity value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertSeverity(values []CodeScanningAlertSeverity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertSeverity) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The severity of the alert.
type CodeScanningAlertRuleSummary_severity int

const (
    NONE_CODESCANNINGALERTRULESUMMARY_SEVERITY CodeScanningAlertRuleSummary_severity = iota
    NOTE_CODESCANNINGALERTRULESUMMARY_SEVERITY
    WARNING_CODESCANNINGALERTRULESUMMARY_SEVERITY
    ERROR_CODESCANNINGALERTRULESUMMARY_SEVERITY
)

func (i CodeScanningAlertRuleSummary_severity) String() string {
    return []string{"none", "note", "warning", "error"}[i]
}
func ParseCodeScanningAlertRuleSummary_severity(v string) (any, error) {
    result := NONE_CODESCANNINGALERTRULESUMMARY_SEVERITY
    switch v {
        case "none":
            result = NONE_CODESCANNINGALERTRULESUMMARY_SEVERITY
        case "note":
            result = NOTE_CODESCANNINGALERTRULESUMMARY_SEVERITY
        case "warning":
            result = WARNING_CODESCANNINGALERTRULESUMMARY_SEVERITY
        case "error":
            result = ERROR_CODESCANNINGALERTRULESUMMARY_SEVERITY
        default:
            return 0, errors.New("Unknown CodeScanningAlertRuleSummary_severity value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertRuleSummary_severity(values []CodeScanningAlertRuleSummary_severity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertRuleSummary_severity) isMultiValue() bool {
    return false
}

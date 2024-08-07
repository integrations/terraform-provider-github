package models
import (
    "errors"
)
// The security severity of the alert.
type CodeScanningAlertRuleSummary_security_severity_level int

const (
    LOW_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL CodeScanningAlertRuleSummary_security_severity_level = iota
    MEDIUM_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL
    HIGH_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL
    CRITICAL_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL
)

func (i CodeScanningAlertRuleSummary_security_severity_level) String() string {
    return []string{"low", "medium", "high", "critical"}[i]
}
func ParseCodeScanningAlertRuleSummary_security_severity_level(v string) (any, error) {
    result := LOW_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL
    switch v {
        case "low":
            result = LOW_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL
        case "medium":
            result = MEDIUM_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL
        case "high":
            result = HIGH_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL
        case "critical":
            result = CRITICAL_CODESCANNINGALERTRULESUMMARY_SECURITY_SEVERITY_LEVEL
        default:
            return 0, errors.New("Unknown CodeScanningAlertRuleSummary_security_severity_level value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertRuleSummary_security_severity_level(values []CodeScanningAlertRuleSummary_security_severity_level) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertRuleSummary_security_severity_level) isMultiValue() bool {
    return false
}

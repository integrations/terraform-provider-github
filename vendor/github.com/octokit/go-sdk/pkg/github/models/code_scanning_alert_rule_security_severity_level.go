package models
import (
    "errors"
)
// The security severity of the alert.
type CodeScanningAlertRule_security_severity_level int

const (
    LOW_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL CodeScanningAlertRule_security_severity_level = iota
    MEDIUM_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL
    HIGH_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL
    CRITICAL_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL
)

func (i CodeScanningAlertRule_security_severity_level) String() string {
    return []string{"low", "medium", "high", "critical"}[i]
}
func ParseCodeScanningAlertRule_security_severity_level(v string) (any, error) {
    result := LOW_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL
    switch v {
        case "low":
            result = LOW_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL
        case "medium":
            result = MEDIUM_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL
        case "high":
            result = HIGH_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL
        case "critical":
            result = CRITICAL_CODESCANNINGALERTRULE_SECURITY_SEVERITY_LEVEL
        default:
            return 0, errors.New("Unknown CodeScanningAlertRule_security_severity_level value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertRule_security_severity_level(values []CodeScanningAlertRule_security_severity_level) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertRule_security_severity_level) isMultiValue() bool {
    return false
}

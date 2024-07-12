package models
import (
    "errors"
)
// The severity of the advisory.
type GlobalAdvisory_severity int

const (
    CRITICAL_GLOBALADVISORY_SEVERITY GlobalAdvisory_severity = iota
    HIGH_GLOBALADVISORY_SEVERITY
    MEDIUM_GLOBALADVISORY_SEVERITY
    LOW_GLOBALADVISORY_SEVERITY
    UNKNOWN_GLOBALADVISORY_SEVERITY
)

func (i GlobalAdvisory_severity) String() string {
    return []string{"critical", "high", "medium", "low", "unknown"}[i]
}
func ParseGlobalAdvisory_severity(v string) (any, error) {
    result := CRITICAL_GLOBALADVISORY_SEVERITY
    switch v {
        case "critical":
            result = CRITICAL_GLOBALADVISORY_SEVERITY
        case "high":
            result = HIGH_GLOBALADVISORY_SEVERITY
        case "medium":
            result = MEDIUM_GLOBALADVISORY_SEVERITY
        case "low":
            result = LOW_GLOBALADVISORY_SEVERITY
        case "unknown":
            result = UNKNOWN_GLOBALADVISORY_SEVERITY
        default:
            return 0, errors.New("Unknown GlobalAdvisory_severity value: " + v)
    }
    return &result, nil
}
func SerializeGlobalAdvisory_severity(values []GlobalAdvisory_severity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GlobalAdvisory_severity) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The severity of the advisory.
type DependabotAlertSecurityAdvisory_severity int

const (
    LOW_DEPENDABOTALERTSECURITYADVISORY_SEVERITY DependabotAlertSecurityAdvisory_severity = iota
    MEDIUM_DEPENDABOTALERTSECURITYADVISORY_SEVERITY
    HIGH_DEPENDABOTALERTSECURITYADVISORY_SEVERITY
    CRITICAL_DEPENDABOTALERTSECURITYADVISORY_SEVERITY
)

func (i DependabotAlertSecurityAdvisory_severity) String() string {
    return []string{"low", "medium", "high", "critical"}[i]
}
func ParseDependabotAlertSecurityAdvisory_severity(v string) (any, error) {
    result := LOW_DEPENDABOTALERTSECURITYADVISORY_SEVERITY
    switch v {
        case "low":
            result = LOW_DEPENDABOTALERTSECURITYADVISORY_SEVERITY
        case "medium":
            result = MEDIUM_DEPENDABOTALERTSECURITYADVISORY_SEVERITY
        case "high":
            result = HIGH_DEPENDABOTALERTSECURITYADVISORY_SEVERITY
        case "critical":
            result = CRITICAL_DEPENDABOTALERTSECURITYADVISORY_SEVERITY
        default:
            return 0, errors.New("Unknown DependabotAlertSecurityAdvisory_severity value: " + v)
    }
    return &result, nil
}
func SerializeDependabotAlertSecurityAdvisory_severity(values []DependabotAlertSecurityAdvisory_severity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependabotAlertSecurityAdvisory_severity) isMultiValue() bool {
    return false
}

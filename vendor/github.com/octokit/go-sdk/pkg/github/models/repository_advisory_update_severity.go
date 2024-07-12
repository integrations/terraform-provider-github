package models
import (
    "errors"
)
// The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
type RepositoryAdvisoryUpdate_severity int

const (
    CRITICAL_REPOSITORYADVISORYUPDATE_SEVERITY RepositoryAdvisoryUpdate_severity = iota
    HIGH_REPOSITORYADVISORYUPDATE_SEVERITY
    MEDIUM_REPOSITORYADVISORYUPDATE_SEVERITY
    LOW_REPOSITORYADVISORYUPDATE_SEVERITY
)

func (i RepositoryAdvisoryUpdate_severity) String() string {
    return []string{"critical", "high", "medium", "low"}[i]
}
func ParseRepositoryAdvisoryUpdate_severity(v string) (any, error) {
    result := CRITICAL_REPOSITORYADVISORYUPDATE_SEVERITY
    switch v {
        case "critical":
            result = CRITICAL_REPOSITORYADVISORYUPDATE_SEVERITY
        case "high":
            result = HIGH_REPOSITORYADVISORYUPDATE_SEVERITY
        case "medium":
            result = MEDIUM_REPOSITORYADVISORYUPDATE_SEVERITY
        case "low":
            result = LOW_REPOSITORYADVISORYUPDATE_SEVERITY
        default:
            return 0, errors.New("Unknown RepositoryAdvisoryUpdate_severity value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryAdvisoryUpdate_severity(values []RepositoryAdvisoryUpdate_severity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryAdvisoryUpdate_severity) isMultiValue() bool {
    return false
}

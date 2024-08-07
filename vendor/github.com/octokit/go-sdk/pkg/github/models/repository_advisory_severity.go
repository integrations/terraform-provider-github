package models
import (
    "errors"
)
// The severity of the advisory.
type RepositoryAdvisory_severity int

const (
    CRITICAL_REPOSITORYADVISORY_SEVERITY RepositoryAdvisory_severity = iota
    HIGH_REPOSITORYADVISORY_SEVERITY
    MEDIUM_REPOSITORYADVISORY_SEVERITY
    LOW_REPOSITORYADVISORY_SEVERITY
)

func (i RepositoryAdvisory_severity) String() string {
    return []string{"critical", "high", "medium", "low"}[i]
}
func ParseRepositoryAdvisory_severity(v string) (any, error) {
    result := CRITICAL_REPOSITORYADVISORY_SEVERITY
    switch v {
        case "critical":
            result = CRITICAL_REPOSITORYADVISORY_SEVERITY
        case "high":
            result = HIGH_REPOSITORYADVISORY_SEVERITY
        case "medium":
            result = MEDIUM_REPOSITORYADVISORY_SEVERITY
        case "low":
            result = LOW_REPOSITORYADVISORY_SEVERITY
        default:
            return 0, errors.New("Unknown RepositoryAdvisory_severity value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryAdvisory_severity(values []RepositoryAdvisory_severity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryAdvisory_severity) isMultiValue() bool {
    return false
}

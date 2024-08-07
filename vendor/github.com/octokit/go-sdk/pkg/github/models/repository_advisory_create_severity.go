package models
import (
    "errors"
)
// The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
type RepositoryAdvisoryCreate_severity int

const (
    CRITICAL_REPOSITORYADVISORYCREATE_SEVERITY RepositoryAdvisoryCreate_severity = iota
    HIGH_REPOSITORYADVISORYCREATE_SEVERITY
    MEDIUM_REPOSITORYADVISORYCREATE_SEVERITY
    LOW_REPOSITORYADVISORYCREATE_SEVERITY
)

func (i RepositoryAdvisoryCreate_severity) String() string {
    return []string{"critical", "high", "medium", "low"}[i]
}
func ParseRepositoryAdvisoryCreate_severity(v string) (any, error) {
    result := CRITICAL_REPOSITORYADVISORYCREATE_SEVERITY
    switch v {
        case "critical":
            result = CRITICAL_REPOSITORYADVISORYCREATE_SEVERITY
        case "high":
            result = HIGH_REPOSITORYADVISORYCREATE_SEVERITY
        case "medium":
            result = MEDIUM_REPOSITORYADVISORYCREATE_SEVERITY
        case "low":
            result = LOW_REPOSITORYADVISORYCREATE_SEVERITY
        default:
            return 0, errors.New("Unknown RepositoryAdvisoryCreate_severity value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryAdvisoryCreate_severity(values []RepositoryAdvisoryCreate_severity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryAdvisoryCreate_severity) isMultiValue() bool {
    return false
}

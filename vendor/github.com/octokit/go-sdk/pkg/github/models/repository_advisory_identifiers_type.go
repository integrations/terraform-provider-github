package models
import (
    "errors"
)
// The type of identifier.
type RepositoryAdvisory_identifiers_type int

const (
    CVE_REPOSITORYADVISORY_IDENTIFIERS_TYPE RepositoryAdvisory_identifiers_type = iota
    GHSA_REPOSITORYADVISORY_IDENTIFIERS_TYPE
)

func (i RepositoryAdvisory_identifiers_type) String() string {
    return []string{"CVE", "GHSA"}[i]
}
func ParseRepositoryAdvisory_identifiers_type(v string) (any, error) {
    result := CVE_REPOSITORYADVISORY_IDENTIFIERS_TYPE
    switch v {
        case "CVE":
            result = CVE_REPOSITORYADVISORY_IDENTIFIERS_TYPE
        case "GHSA":
            result = GHSA_REPOSITORYADVISORY_IDENTIFIERS_TYPE
        default:
            return 0, errors.New("Unknown RepositoryAdvisory_identifiers_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryAdvisory_identifiers_type(values []RepositoryAdvisory_identifiers_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryAdvisory_identifiers_type) isMultiValue() bool {
    return false
}

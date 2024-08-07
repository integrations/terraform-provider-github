package models
import (
    "errors"
)
// The type of advisory identifier.
type DependabotAlertSecurityAdvisory_identifiers_type int

const (
    CVE_DEPENDABOTALERTSECURITYADVISORY_IDENTIFIERS_TYPE DependabotAlertSecurityAdvisory_identifiers_type = iota
    GHSA_DEPENDABOTALERTSECURITYADVISORY_IDENTIFIERS_TYPE
)

func (i DependabotAlertSecurityAdvisory_identifiers_type) String() string {
    return []string{"CVE", "GHSA"}[i]
}
func ParseDependabotAlertSecurityAdvisory_identifiers_type(v string) (any, error) {
    result := CVE_DEPENDABOTALERTSECURITYADVISORY_IDENTIFIERS_TYPE
    switch v {
        case "CVE":
            result = CVE_DEPENDABOTALERTSECURITYADVISORY_IDENTIFIERS_TYPE
        case "GHSA":
            result = GHSA_DEPENDABOTALERTSECURITYADVISORY_IDENTIFIERS_TYPE
        default:
            return 0, errors.New("Unknown DependabotAlertSecurityAdvisory_identifiers_type value: " + v)
    }
    return &result, nil
}
func SerializeDependabotAlertSecurityAdvisory_identifiers_type(values []DependabotAlertSecurityAdvisory_identifiers_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DependabotAlertSecurityAdvisory_identifiers_type) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The type of identifier.
type GlobalAdvisory_identifiers_type int

const (
    CVE_GLOBALADVISORY_IDENTIFIERS_TYPE GlobalAdvisory_identifiers_type = iota
    GHSA_GLOBALADVISORY_IDENTIFIERS_TYPE
)

func (i GlobalAdvisory_identifiers_type) String() string {
    return []string{"CVE", "GHSA"}[i]
}
func ParseGlobalAdvisory_identifiers_type(v string) (any, error) {
    result := CVE_GLOBALADVISORY_IDENTIFIERS_TYPE
    switch v {
        case "CVE":
            result = CVE_GLOBALADVISORY_IDENTIFIERS_TYPE
        case "GHSA":
            result = GHSA_GLOBALADVISORY_IDENTIFIERS_TYPE
        default:
            return 0, errors.New("Unknown GlobalAdvisory_identifiers_type value: " + v)
    }
    return &result, nil
}
func SerializeGlobalAdvisory_identifiers_type(values []GlobalAdvisory_identifiers_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GlobalAdvisory_identifiers_type) isMultiValue() bool {
    return false
}

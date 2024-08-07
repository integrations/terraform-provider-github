package models
import (
    "errors"
)
// The state if the domain is verified
type Page_protected_domain_state int

const (
    PENDING_PAGE_PROTECTED_DOMAIN_STATE Page_protected_domain_state = iota
    VERIFIED_PAGE_PROTECTED_DOMAIN_STATE
    UNVERIFIED_PAGE_PROTECTED_DOMAIN_STATE
)

func (i Page_protected_domain_state) String() string {
    return []string{"pending", "verified", "unverified"}[i]
}
func ParsePage_protected_domain_state(v string) (any, error) {
    result := PENDING_PAGE_PROTECTED_DOMAIN_STATE
    switch v {
        case "pending":
            result = PENDING_PAGE_PROTECTED_DOMAIN_STATE
        case "verified":
            result = VERIFIED_PAGE_PROTECTED_DOMAIN_STATE
        case "unverified":
            result = UNVERIFIED_PAGE_PROTECTED_DOMAIN_STATE
        default:
            return 0, errors.New("Unknown Page_protected_domain_state value: " + v)
    }
    return &result, nil
}
func SerializePage_protected_domain_state(values []Page_protected_domain_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Page_protected_domain_state) isMultiValue() bool {
    return false
}

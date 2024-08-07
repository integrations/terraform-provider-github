package models
import (
    "errors"
)
// The user's membership type in the organization.
type OrgMembership_role int

const (
    ADMIN_ORGMEMBERSHIP_ROLE OrgMembership_role = iota
    MEMBER_ORGMEMBERSHIP_ROLE
    BILLING_MANAGER_ORGMEMBERSHIP_ROLE
)

func (i OrgMembership_role) String() string {
    return []string{"admin", "member", "billing_manager"}[i]
}
func ParseOrgMembership_role(v string) (any, error) {
    result := ADMIN_ORGMEMBERSHIP_ROLE
    switch v {
        case "admin":
            result = ADMIN_ORGMEMBERSHIP_ROLE
        case "member":
            result = MEMBER_ORGMEMBERSHIP_ROLE
        case "billing_manager":
            result = BILLING_MANAGER_ORGMEMBERSHIP_ROLE
        default:
            return 0, errors.New("Unknown OrgMembership_role value: " + v)
    }
    return &result, nil
}
func SerializeOrgMembership_role(values []OrgMembership_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrgMembership_role) isMultiValue() bool {
    return false
}

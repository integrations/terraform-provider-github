package models
import (
    "errors"
)
// The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
type OrgMembership_state int

const (
    ACTIVE_ORGMEMBERSHIP_STATE OrgMembership_state = iota
    PENDING_ORGMEMBERSHIP_STATE
)

func (i OrgMembership_state) String() string {
    return []string{"active", "pending"}[i]
}
func ParseOrgMembership_state(v string) (any, error) {
    result := ACTIVE_ORGMEMBERSHIP_STATE
    switch v {
        case "active":
            result = ACTIVE_ORGMEMBERSHIP_STATE
        case "pending":
            result = PENDING_ORGMEMBERSHIP_STATE
        default:
            return 0, errors.New("Unknown OrgMembership_state value: " + v)
    }
    return &result, nil
}
func SerializeOrgMembership_state(values []OrgMembership_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrgMembership_state) isMultiValue() bool {
    return false
}

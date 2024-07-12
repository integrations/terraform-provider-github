package models
import (
    "errors"
)
// The state of the user's membership in the team.
type TeamMembership_state int

const (
    ACTIVE_TEAMMEMBERSHIP_STATE TeamMembership_state = iota
    PENDING_TEAMMEMBERSHIP_STATE
)

func (i TeamMembership_state) String() string {
    return []string{"active", "pending"}[i]
}
func ParseTeamMembership_state(v string) (any, error) {
    result := ACTIVE_TEAMMEMBERSHIP_STATE
    switch v {
        case "active":
            result = ACTIVE_TEAMMEMBERSHIP_STATE
        case "pending":
            result = PENDING_TEAMMEMBERSHIP_STATE
        default:
            return 0, errors.New("Unknown TeamMembership_state value: " + v)
    }
    return &result, nil
}
func SerializeTeamMembership_state(values []TeamMembership_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamMembership_state) isMultiValue() bool {
    return false
}

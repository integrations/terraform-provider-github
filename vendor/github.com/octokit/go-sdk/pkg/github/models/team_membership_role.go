package models
import (
    "errors"
)
// The role of the user in the team.
type TeamMembership_role int

const (
    MEMBER_TEAMMEMBERSHIP_ROLE TeamMembership_role = iota
    MAINTAINER_TEAMMEMBERSHIP_ROLE
)

func (i TeamMembership_role) String() string {
    return []string{"member", "maintainer"}[i]
}
func ParseTeamMembership_role(v string) (any, error) {
    result := MEMBER_TEAMMEMBERSHIP_ROLE
    switch v {
        case "member":
            result = MEMBER_TEAMMEMBERSHIP_ROLE
        case "maintainer":
            result = MAINTAINER_TEAMMEMBERSHIP_ROLE
        default:
            return 0, errors.New("Unknown TeamMembership_role value: " + v)
    }
    return &result, nil
}
func SerializeTeamMembership_role(values []TeamMembership_role) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamMembership_role) isMultiValue() bool {
    return false
}

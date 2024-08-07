package models
import (
    "errors"
)
// The level of permission to grant the access token to manage team discussions and related comments.
type AppPermissions_team_discussions int

const (
    READ_APPPERMISSIONS_TEAM_DISCUSSIONS AppPermissions_team_discussions = iota
    WRITE_APPPERMISSIONS_TEAM_DISCUSSIONS
)

func (i AppPermissions_team_discussions) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_team_discussions(v string) (any, error) {
    result := READ_APPPERMISSIONS_TEAM_DISCUSSIONS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_TEAM_DISCUSSIONS
        case "write":
            result = WRITE_APPPERMISSIONS_TEAM_DISCUSSIONS
        default:
            return 0, errors.New("Unknown AppPermissions_team_discussions value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_team_discussions(values []AppPermissions_team_discussions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_team_discussions) isMultiValue() bool {
    return false
}

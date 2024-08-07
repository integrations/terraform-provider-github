package models
import (
    "errors"
)
// The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
type AppPermissions_actions int

const (
    READ_APPPERMISSIONS_ACTIONS AppPermissions_actions = iota
    WRITE_APPPERMISSIONS_ACTIONS
)

func (i AppPermissions_actions) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_actions(v string) (any, error) {
    result := READ_APPPERMISSIONS_ACTIONS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ACTIONS
        case "write":
            result = WRITE_APPPERMISSIONS_ACTIONS
        default:
            return 0, errors.New("Unknown AppPermissions_actions value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_actions(values []AppPermissions_actions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_actions) isMultiValue() bool {
    return false
}

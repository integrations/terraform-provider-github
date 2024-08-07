package models
import (
    "errors"
)
// The level of permission to grant the access token to view and manage interaction limits on a repository.
type AppPermissions_interaction_limits int

const (
    READ_APPPERMISSIONS_INTERACTION_LIMITS AppPermissions_interaction_limits = iota
    WRITE_APPPERMISSIONS_INTERACTION_LIMITS
)

func (i AppPermissions_interaction_limits) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_interaction_limits(v string) (any, error) {
    result := READ_APPPERMISSIONS_INTERACTION_LIMITS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_INTERACTION_LIMITS
        case "write":
            result = WRITE_APPPERMISSIONS_INTERACTION_LIMITS
        default:
            return 0, errors.New("Unknown AppPermissions_interaction_limits value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_interaction_limits(values []AppPermissions_interaction_limits) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_interaction_limits) isMultiValue() bool {
    return false
}

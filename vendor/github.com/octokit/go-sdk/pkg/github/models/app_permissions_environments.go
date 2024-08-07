package models
import (
    "errors"
)
// The level of permission to grant the access token for managing repository environments.
type AppPermissions_environments int

const (
    READ_APPPERMISSIONS_ENVIRONMENTS AppPermissions_environments = iota
    WRITE_APPPERMISSIONS_ENVIRONMENTS
)

func (i AppPermissions_environments) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_environments(v string) (any, error) {
    result := READ_APPPERMISSIONS_ENVIRONMENTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ENVIRONMENTS
        case "write":
            result = WRITE_APPPERMISSIONS_ENVIRONMENTS
        default:
            return 0, errors.New("Unknown AppPermissions_environments value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_environments(values []AppPermissions_environments) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_environments) isMultiValue() bool {
    return false
}

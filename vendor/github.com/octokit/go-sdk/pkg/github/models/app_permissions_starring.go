package models
import (
    "errors"
)
// The level of permission to grant the access token to list and manage repositories a user is starring.
type AppPermissions_starring int

const (
    READ_APPPERMISSIONS_STARRING AppPermissions_starring = iota
    WRITE_APPPERMISSIONS_STARRING
)

func (i AppPermissions_starring) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_starring(v string) (any, error) {
    result := READ_APPPERMISSIONS_STARRING
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_STARRING
        case "write":
            result = WRITE_APPPERMISSIONS_STARRING
        default:
            return 0, errors.New("Unknown AppPermissions_starring value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_starring(values []AppPermissions_starring) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_starring) isMultiValue() bool {
    return false
}

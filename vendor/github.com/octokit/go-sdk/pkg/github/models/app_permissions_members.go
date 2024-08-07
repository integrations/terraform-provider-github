package models
import (
    "errors"
)
// The level of permission to grant the access token for organization teams and members.
type AppPermissions_members int

const (
    READ_APPPERMISSIONS_MEMBERS AppPermissions_members = iota
    WRITE_APPPERMISSIONS_MEMBERS
)

func (i AppPermissions_members) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_members(v string) (any, error) {
    result := READ_APPPERMISSIONS_MEMBERS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_MEMBERS
        case "write":
            result = WRITE_APPPERMISSIONS_MEMBERS
        default:
            return 0, errors.New("Unknown AppPermissions_members value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_members(values []AppPermissions_members) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_members) isMultiValue() bool {
    return false
}

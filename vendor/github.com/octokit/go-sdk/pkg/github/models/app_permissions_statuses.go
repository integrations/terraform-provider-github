package models
import (
    "errors"
)
// The level of permission to grant the access token for commit statuses.
type AppPermissions_statuses int

const (
    READ_APPPERMISSIONS_STATUSES AppPermissions_statuses = iota
    WRITE_APPPERMISSIONS_STATUSES
)

func (i AppPermissions_statuses) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_statuses(v string) (any, error) {
    result := READ_APPPERMISSIONS_STATUSES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_STATUSES
        case "write":
            result = WRITE_APPPERMISSIONS_STATUSES
        default:
            return 0, errors.New("Unknown AppPermissions_statuses value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_statuses(values []AppPermissions_statuses) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_statuses) isMultiValue() bool {
    return false
}

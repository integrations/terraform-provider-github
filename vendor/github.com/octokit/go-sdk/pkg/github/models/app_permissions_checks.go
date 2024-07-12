package models
import (
    "errors"
)
// The level of permission to grant the access token for checks on code.
type AppPermissions_checks int

const (
    READ_APPPERMISSIONS_CHECKS AppPermissions_checks = iota
    WRITE_APPPERMISSIONS_CHECKS
)

func (i AppPermissions_checks) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_checks(v string) (any, error) {
    result := READ_APPPERMISSIONS_CHECKS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_CHECKS
        case "write":
            result = WRITE_APPPERMISSIONS_CHECKS
        default:
            return 0, errors.New("Unknown AppPermissions_checks value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_checks(values []AppPermissions_checks) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_checks) isMultiValue() bool {
    return false
}

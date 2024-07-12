package models
import (
    "errors"
)
// The level of permission to grant the access token to manage just a single file.
type AppPermissions_single_file int

const (
    READ_APPPERMISSIONS_SINGLE_FILE AppPermissions_single_file = iota
    WRITE_APPPERMISSIONS_SINGLE_FILE
)

func (i AppPermissions_single_file) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_single_file(v string) (any, error) {
    result := READ_APPPERMISSIONS_SINGLE_FILE
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_SINGLE_FILE
        case "write":
            result = WRITE_APPPERMISSIONS_SINGLE_FILE
        default:
            return 0, errors.New("Unknown AppPermissions_single_file value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_single_file(values []AppPermissions_single_file) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_single_file) isMultiValue() bool {
    return false
}

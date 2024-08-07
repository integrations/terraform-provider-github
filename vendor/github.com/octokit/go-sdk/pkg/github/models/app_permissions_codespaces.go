package models
import (
    "errors"
)
// The level of permission to grant the access token to create, edit, delete, and list Codespaces.
type AppPermissions_codespaces int

const (
    READ_APPPERMISSIONS_CODESPACES AppPermissions_codespaces = iota
    WRITE_APPPERMISSIONS_CODESPACES
)

func (i AppPermissions_codespaces) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_codespaces(v string) (any, error) {
    result := READ_APPPERMISSIONS_CODESPACES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_CODESPACES
        case "write":
            result = WRITE_APPPERMISSIONS_CODESPACES
        default:
            return 0, errors.New("Unknown AppPermissions_codespaces value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_codespaces(values []AppPermissions_codespaces) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_codespaces) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
type AppPermissions_metadata int

const (
    READ_APPPERMISSIONS_METADATA AppPermissions_metadata = iota
    WRITE_APPPERMISSIONS_METADATA
)

func (i AppPermissions_metadata) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_metadata(v string) (any, error) {
    result := READ_APPPERMISSIONS_METADATA
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_METADATA
        case "write":
            result = WRITE_APPPERMISSIONS_METADATA
        default:
            return 0, errors.New("Unknown AppPermissions_metadata value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_metadata(values []AppPermissions_metadata) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_metadata) isMultiValue() bool {
    return false
}

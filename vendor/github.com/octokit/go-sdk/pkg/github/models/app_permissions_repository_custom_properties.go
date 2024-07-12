package models
import (
    "errors"
)
// The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
type AppPermissions_repository_custom_properties int

const (
    READ_APPPERMISSIONS_REPOSITORY_CUSTOM_PROPERTIES AppPermissions_repository_custom_properties = iota
    WRITE_APPPERMISSIONS_REPOSITORY_CUSTOM_PROPERTIES
)

func (i AppPermissions_repository_custom_properties) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_repository_custom_properties(v string) (any, error) {
    result := READ_APPPERMISSIONS_REPOSITORY_CUSTOM_PROPERTIES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_REPOSITORY_CUSTOM_PROPERTIES
        case "write":
            result = WRITE_APPPERMISSIONS_REPOSITORY_CUSTOM_PROPERTIES
        default:
            return 0, errors.New("Unknown AppPermissions_repository_custom_properties value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_repository_custom_properties(values []AppPermissions_repository_custom_properties) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_repository_custom_properties) isMultiValue() bool {
    return false
}

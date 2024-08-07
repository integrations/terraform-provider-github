package models
import (
    "errors"
)
// The level of permission to grant the access token for custom property management.
type AppPermissions_organization_custom_properties int

const (
    READ_APPPERMISSIONS_ORGANIZATION_CUSTOM_PROPERTIES AppPermissions_organization_custom_properties = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_CUSTOM_PROPERTIES
    ADMIN_APPPERMISSIONS_ORGANIZATION_CUSTOM_PROPERTIES
)

func (i AppPermissions_organization_custom_properties) String() string {
    return []string{"read", "write", "admin"}[i]
}
func ParseAppPermissions_organization_custom_properties(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_CUSTOM_PROPERTIES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_CUSTOM_PROPERTIES
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_CUSTOM_PROPERTIES
        case "admin":
            result = ADMIN_APPPERMISSIONS_ORGANIZATION_CUSTOM_PROPERTIES
        default:
            return 0, errors.New("Unknown AppPermissions_organization_custom_properties value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_custom_properties(values []AppPermissions_organization_custom_properties) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_custom_properties) isMultiValue() bool {
    return false
}

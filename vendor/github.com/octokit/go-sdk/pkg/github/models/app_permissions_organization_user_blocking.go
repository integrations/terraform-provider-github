package models
import (
    "errors"
)
// The level of permission to grant the access token to view and manage users blocked by the organization.
type AppPermissions_organization_user_blocking int

const (
    READ_APPPERMISSIONS_ORGANIZATION_USER_BLOCKING AppPermissions_organization_user_blocking = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_USER_BLOCKING
)

func (i AppPermissions_organization_user_blocking) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_user_blocking(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_USER_BLOCKING
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_USER_BLOCKING
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_USER_BLOCKING
        default:
            return 0, errors.New("Unknown AppPermissions_organization_user_blocking value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_user_blocking(values []AppPermissions_organization_user_blocking) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_user_blocking) isMultiValue() bool {
    return false
}

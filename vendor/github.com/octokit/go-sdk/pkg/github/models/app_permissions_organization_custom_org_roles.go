package models
import (
    "errors"
)
// The level of permission to grant the access token for custom organization roles management.
type AppPermissions_organization_custom_org_roles int

const (
    READ_APPPERMISSIONS_ORGANIZATION_CUSTOM_ORG_ROLES AppPermissions_organization_custom_org_roles = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_CUSTOM_ORG_ROLES
)

func (i AppPermissions_organization_custom_org_roles) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_custom_org_roles(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_CUSTOM_ORG_ROLES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_CUSTOM_ORG_ROLES
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_CUSTOM_ORG_ROLES
        default:
            return 0, errors.New("Unknown AppPermissions_organization_custom_org_roles value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_custom_org_roles(values []AppPermissions_organization_custom_org_roles) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_custom_org_roles) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The level of permission to grant the access token to manage access to an organization.
type AppPermissions_organization_administration int

const (
    READ_APPPERMISSIONS_ORGANIZATION_ADMINISTRATION AppPermissions_organization_administration = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_ADMINISTRATION
)

func (i AppPermissions_organization_administration) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_administration(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_ADMINISTRATION
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_ADMINISTRATION
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_ADMINISTRATION
        default:
            return 0, errors.New("Unknown AppPermissions_organization_administration value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_administration(values []AppPermissions_organization_administration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_administration) isMultiValue() bool {
    return false
}

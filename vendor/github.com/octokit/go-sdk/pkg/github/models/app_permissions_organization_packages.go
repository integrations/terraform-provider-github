package models
import (
    "errors"
)
// The level of permission to grant the access token for organization packages published to GitHub Packages.
type AppPermissions_organization_packages int

const (
    READ_APPPERMISSIONS_ORGANIZATION_PACKAGES AppPermissions_organization_packages = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_PACKAGES
)

func (i AppPermissions_organization_packages) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_packages(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_PACKAGES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_PACKAGES
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_PACKAGES
        default:
            return 0, errors.New("Unknown AppPermissions_organization_packages value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_packages(values []AppPermissions_organization_packages) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_packages) isMultiValue() bool {
    return false
}

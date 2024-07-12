package models
import (
    "errors"
)
// The level of permission to grant the access token for packages published to GitHub Packages.
type AppPermissions_packages int

const (
    READ_APPPERMISSIONS_PACKAGES AppPermissions_packages = iota
    WRITE_APPPERMISSIONS_PACKAGES
)

func (i AppPermissions_packages) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_packages(v string) (any, error) {
    result := READ_APPPERMISSIONS_PACKAGES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_PACKAGES
        case "write":
            result = WRITE_APPPERMISSIONS_PACKAGES
        default:
            return 0, errors.New("Unknown AppPermissions_packages value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_packages(values []AppPermissions_packages) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_packages) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The level of permission to grant the access token to manage organization secrets.
type AppPermissions_organization_secrets int

const (
    READ_APPPERMISSIONS_ORGANIZATION_SECRETS AppPermissions_organization_secrets = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_SECRETS
)

func (i AppPermissions_organization_secrets) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_secrets(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_SECRETS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_SECRETS
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_SECRETS
        default:
            return 0, errors.New("Unknown AppPermissions_organization_secrets value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_secrets(values []AppPermissions_organization_secrets) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_secrets) isMultiValue() bool {
    return false
}

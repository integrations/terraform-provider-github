package models
import (
    "errors"
)
// The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
type AppPermissions_organization_personal_access_tokens int

const (
    READ_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKENS AppPermissions_organization_personal_access_tokens = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKENS
)

func (i AppPermissions_organization_personal_access_tokens) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_personal_access_tokens(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKENS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKENS
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKENS
        default:
            return 0, errors.New("Unknown AppPermissions_organization_personal_access_tokens value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_personal_access_tokens(values []AppPermissions_organization_personal_access_tokens) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_personal_access_tokens) isMultiValue() bool {
    return false
}

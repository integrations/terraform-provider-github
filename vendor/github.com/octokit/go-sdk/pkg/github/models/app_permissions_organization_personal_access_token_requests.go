package models
import (
    "errors"
)
// The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
type AppPermissions_organization_personal_access_token_requests int

const (
    READ_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKEN_REQUESTS AppPermissions_organization_personal_access_token_requests = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKEN_REQUESTS
)

func (i AppPermissions_organization_personal_access_token_requests) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_personal_access_token_requests(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKEN_REQUESTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKEN_REQUESTS
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_PERSONAL_ACCESS_TOKEN_REQUESTS
        default:
            return 0, errors.New("Unknown AppPermissions_organization_personal_access_token_requests value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_personal_access_token_requests(values []AppPermissions_organization_personal_access_token_requests) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_personal_access_token_requests) isMultiValue() bool {
    return false
}

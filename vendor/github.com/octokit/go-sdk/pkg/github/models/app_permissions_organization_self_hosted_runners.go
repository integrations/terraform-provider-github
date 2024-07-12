package models
import (
    "errors"
)
// The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
type AppPermissions_organization_self_hosted_runners int

const (
    READ_APPPERMISSIONS_ORGANIZATION_SELF_HOSTED_RUNNERS AppPermissions_organization_self_hosted_runners = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_SELF_HOSTED_RUNNERS
)

func (i AppPermissions_organization_self_hosted_runners) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_self_hosted_runners(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_SELF_HOSTED_RUNNERS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_SELF_HOSTED_RUNNERS
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_SELF_HOSTED_RUNNERS
        default:
            return 0, errors.New("Unknown AppPermissions_organization_self_hosted_runners value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_self_hosted_runners(values []AppPermissions_organization_self_hosted_runners) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_self_hosted_runners) isMultiValue() bool {
    return false
}

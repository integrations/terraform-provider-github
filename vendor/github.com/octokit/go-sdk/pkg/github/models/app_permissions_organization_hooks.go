package models
import (
    "errors"
)
// The level of permission to grant the access token to manage the post-receive hooks for an organization.
type AppPermissions_organization_hooks int

const (
    READ_APPPERMISSIONS_ORGANIZATION_HOOKS AppPermissions_organization_hooks = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_HOOKS
)

func (i AppPermissions_organization_hooks) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_hooks(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_HOOKS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_HOOKS
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_HOOKS
        default:
            return 0, errors.New("Unknown AppPermissions_organization_hooks value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_hooks(values []AppPermissions_organization_hooks) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_hooks) isMultiValue() bool {
    return false
}

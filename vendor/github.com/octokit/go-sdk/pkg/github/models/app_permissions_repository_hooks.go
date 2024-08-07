package models
import (
    "errors"
)
// The level of permission to grant the access token to manage the post-receive hooks for a repository.
type AppPermissions_repository_hooks int

const (
    READ_APPPERMISSIONS_REPOSITORY_HOOKS AppPermissions_repository_hooks = iota
    WRITE_APPPERMISSIONS_REPOSITORY_HOOKS
)

func (i AppPermissions_repository_hooks) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_repository_hooks(v string) (any, error) {
    result := READ_APPPERMISSIONS_REPOSITORY_HOOKS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_REPOSITORY_HOOKS
        case "write":
            result = WRITE_APPPERMISSIONS_REPOSITORY_HOOKS
        default:
            return 0, errors.New("Unknown AppPermissions_repository_hooks value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_repository_hooks(values []AppPermissions_repository_hooks) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_repository_hooks) isMultiValue() bool {
    return false
}

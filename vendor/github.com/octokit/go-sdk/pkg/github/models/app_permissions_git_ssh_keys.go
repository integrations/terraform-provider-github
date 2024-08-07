package models
import (
    "errors"
)
// The level of permission to grant the access token to manage git SSH keys.
type AppPermissions_git_ssh_keys int

const (
    READ_APPPERMISSIONS_GIT_SSH_KEYS AppPermissions_git_ssh_keys = iota
    WRITE_APPPERMISSIONS_GIT_SSH_KEYS
)

func (i AppPermissions_git_ssh_keys) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_git_ssh_keys(v string) (any, error) {
    result := READ_APPPERMISSIONS_GIT_SSH_KEYS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_GIT_SSH_KEYS
        case "write":
            result = WRITE_APPPERMISSIONS_GIT_SSH_KEYS
        default:
            return 0, errors.New("Unknown AppPermissions_git_ssh_keys value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_git_ssh_keys(values []AppPermissions_git_ssh_keys) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_git_ssh_keys) isMultiValue() bool {
    return false
}

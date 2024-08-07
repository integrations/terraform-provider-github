package models
import (
    "errors"
)
// The level of permission to grant the access token to view and manage GPG keys belonging to a user.
type AppPermissions_gpg_keys int

const (
    READ_APPPERMISSIONS_GPG_KEYS AppPermissions_gpg_keys = iota
    WRITE_APPPERMISSIONS_GPG_KEYS
)

func (i AppPermissions_gpg_keys) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_gpg_keys(v string) (any, error) {
    result := READ_APPPERMISSIONS_GPG_KEYS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_GPG_KEYS
        case "write":
            result = WRITE_APPPERMISSIONS_GPG_KEYS
        default:
            return 0, errors.New("Unknown AppPermissions_gpg_keys value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_gpg_keys(values []AppPermissions_gpg_keys) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_gpg_keys) isMultiValue() bool {
    return false
}

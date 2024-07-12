package models
import (
    "errors"
)
// The leve of permission to grant the access token to manage Dependabot secrets.
type AppPermissions_dependabot_secrets int

const (
    READ_APPPERMISSIONS_DEPENDABOT_SECRETS AppPermissions_dependabot_secrets = iota
    WRITE_APPPERMISSIONS_DEPENDABOT_SECRETS
)

func (i AppPermissions_dependabot_secrets) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_dependabot_secrets(v string) (any, error) {
    result := READ_APPPERMISSIONS_DEPENDABOT_SECRETS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_DEPENDABOT_SECRETS
        case "write":
            result = WRITE_APPPERMISSIONS_DEPENDABOT_SECRETS
        default:
            return 0, errors.New("Unknown AppPermissions_dependabot_secrets value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_dependabot_secrets(values []AppPermissions_dependabot_secrets) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_dependabot_secrets) isMultiValue() bool {
    return false
}

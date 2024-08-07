package models
import (
    "errors"
)
// The level of permission to grant the access token for deployments and deployment statuses.
type AppPermissions_deployments int

const (
    READ_APPPERMISSIONS_DEPLOYMENTS AppPermissions_deployments = iota
    WRITE_APPPERMISSIONS_DEPLOYMENTS
)

func (i AppPermissions_deployments) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_deployments(v string) (any, error) {
    result := READ_APPPERMISSIONS_DEPLOYMENTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_DEPLOYMENTS
        case "write":
            result = WRITE_APPPERMISSIONS_DEPLOYMENTS
        default:
            return 0, errors.New("Unknown AppPermissions_deployments value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_deployments(values []AppPermissions_deployments) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_deployments) isMultiValue() bool {
    return false
}

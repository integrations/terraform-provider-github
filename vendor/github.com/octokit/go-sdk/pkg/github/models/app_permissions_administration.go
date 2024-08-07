package models
import (
    "errors"
)
// The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
type AppPermissions_administration int

const (
    READ_APPPERMISSIONS_ADMINISTRATION AppPermissions_administration = iota
    WRITE_APPPERMISSIONS_ADMINISTRATION
)

func (i AppPermissions_administration) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_administration(v string) (any, error) {
    result := READ_APPPERMISSIONS_ADMINISTRATION
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ADMINISTRATION
        case "write":
            result = WRITE_APPPERMISSIONS_ADMINISTRATION
        default:
            return 0, errors.New("Unknown AppPermissions_administration value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_administration(values []AppPermissions_administration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_administration) isMultiValue() bool {
    return false
}

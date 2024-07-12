package models
import (
    "errors"
)
// The level of permission to grant the access token to manage the profile settings belonging to a user.
type AppPermissions_profile int

const (
    WRITE_APPPERMISSIONS_PROFILE AppPermissions_profile = iota
)

func (i AppPermissions_profile) String() string {
    return []string{"write"}[i]
}
func ParseAppPermissions_profile(v string) (any, error) {
    result := WRITE_APPPERMISSIONS_PROFILE
    switch v {
        case "write":
            result = WRITE_APPPERMISSIONS_PROFILE
        default:
            return 0, errors.New("Unknown AppPermissions_profile value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_profile(values []AppPermissions_profile) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_profile) isMultiValue() bool {
    return false
}

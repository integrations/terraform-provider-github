package models
import (
    "errors"
)
// The level of permission to grant the access token to manage the followers belonging to a user.
type AppPermissions_followers int

const (
    READ_APPPERMISSIONS_FOLLOWERS AppPermissions_followers = iota
    WRITE_APPPERMISSIONS_FOLLOWERS
)

func (i AppPermissions_followers) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_followers(v string) (any, error) {
    result := READ_APPPERMISSIONS_FOLLOWERS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_FOLLOWERS
        case "write":
            result = WRITE_APPPERMISSIONS_FOLLOWERS
        default:
            return 0, errors.New("Unknown AppPermissions_followers value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_followers(values []AppPermissions_followers) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_followers) isMultiValue() bool {
    return false
}

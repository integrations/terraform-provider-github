package models
import (
    "errors"
)
// The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
type AppPermissions_pull_requests int

const (
    READ_APPPERMISSIONS_PULL_REQUESTS AppPermissions_pull_requests = iota
    WRITE_APPPERMISSIONS_PULL_REQUESTS
)

func (i AppPermissions_pull_requests) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_pull_requests(v string) (any, error) {
    result := READ_APPPERMISSIONS_PULL_REQUESTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_PULL_REQUESTS
        case "write":
            result = WRITE_APPPERMISSIONS_PULL_REQUESTS
        default:
            return 0, errors.New("Unknown AppPermissions_pull_requests value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_pull_requests(values []AppPermissions_pull_requests) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_pull_requests) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
type AppPermissions_issues int

const (
    READ_APPPERMISSIONS_ISSUES AppPermissions_issues = iota
    WRITE_APPPERMISSIONS_ISSUES
)

func (i AppPermissions_issues) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_issues(v string) (any, error) {
    result := READ_APPPERMISSIONS_ISSUES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ISSUES
        case "write":
            result = WRITE_APPPERMISSIONS_ISSUES
        default:
            return 0, errors.New("Unknown AppPermissions_issues value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_issues(values []AppPermissions_issues) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_issues) isMultiValue() bool {
    return false
}

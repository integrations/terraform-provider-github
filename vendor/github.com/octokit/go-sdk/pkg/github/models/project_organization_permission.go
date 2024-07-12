package models
import (
    "errors"
)
// The baseline permission that all organization members have on this project. Only present if owner is an organization.
type Project_organization_permission int

const (
    READ_PROJECT_ORGANIZATION_PERMISSION Project_organization_permission = iota
    WRITE_PROJECT_ORGANIZATION_PERMISSION
    ADMIN_PROJECT_ORGANIZATION_PERMISSION
    NONE_PROJECT_ORGANIZATION_PERMISSION
)

func (i Project_organization_permission) String() string {
    return []string{"read", "write", "admin", "none"}[i]
}
func ParseProject_organization_permission(v string) (any, error) {
    result := READ_PROJECT_ORGANIZATION_PERMISSION
    switch v {
        case "read":
            result = READ_PROJECT_ORGANIZATION_PERMISSION
        case "write":
            result = WRITE_PROJECT_ORGANIZATION_PERMISSION
        case "admin":
            result = ADMIN_PROJECT_ORGANIZATION_PERMISSION
        case "none":
            result = NONE_PROJECT_ORGANIZATION_PERMISSION
        default:
            return 0, errors.New("Unknown Project_organization_permission value: " + v)
    }
    return &result, nil
}
func SerializeProject_organization_permission(values []Project_organization_permission) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Project_organization_permission) isMultiValue() bool {
    return false
}

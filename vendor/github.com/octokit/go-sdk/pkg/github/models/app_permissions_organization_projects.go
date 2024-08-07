package models
import (
    "errors"
)
// The level of permission to grant the access token to manage organization projects and projects beta (where available).
type AppPermissions_organization_projects int

const (
    READ_APPPERMISSIONS_ORGANIZATION_PROJECTS AppPermissions_organization_projects = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_PROJECTS
    ADMIN_APPPERMISSIONS_ORGANIZATION_PROJECTS
)

func (i AppPermissions_organization_projects) String() string {
    return []string{"read", "write", "admin"}[i]
}
func ParseAppPermissions_organization_projects(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_PROJECTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_PROJECTS
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_PROJECTS
        case "admin":
            result = ADMIN_APPPERMISSIONS_ORGANIZATION_PROJECTS
        default:
            return 0, errors.New("Unknown AppPermissions_organization_projects value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_projects(values []AppPermissions_organization_projects) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_projects) isMultiValue() bool {
    return false
}

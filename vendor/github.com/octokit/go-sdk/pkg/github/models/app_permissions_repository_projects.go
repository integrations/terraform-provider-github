package models
import (
    "errors"
)
// The level of permission to grant the access token to manage repository projects, columns, and cards.
type AppPermissions_repository_projects int

const (
    READ_APPPERMISSIONS_REPOSITORY_PROJECTS AppPermissions_repository_projects = iota
    WRITE_APPPERMISSIONS_REPOSITORY_PROJECTS
    ADMIN_APPPERMISSIONS_REPOSITORY_PROJECTS
)

func (i AppPermissions_repository_projects) String() string {
    return []string{"read", "write", "admin"}[i]
}
func ParseAppPermissions_repository_projects(v string) (any, error) {
    result := READ_APPPERMISSIONS_REPOSITORY_PROJECTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_REPOSITORY_PROJECTS
        case "write":
            result = WRITE_APPPERMISSIONS_REPOSITORY_PROJECTS
        case "admin":
            result = ADMIN_APPPERMISSIONS_REPOSITORY_PROJECTS
        default:
            return 0, errors.New("Unknown AppPermissions_repository_projects value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_repository_projects(values []AppPermissions_repository_projects) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_repository_projects) isMultiValue() bool {
    return false
}

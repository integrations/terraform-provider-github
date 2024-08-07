package models
import (
    "errors"
)
// The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
type AppPermissions_pages int

const (
    READ_APPPERMISSIONS_PAGES AppPermissions_pages = iota
    WRITE_APPPERMISSIONS_PAGES
)

func (i AppPermissions_pages) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_pages(v string) (any, error) {
    result := READ_APPPERMISSIONS_PAGES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_PAGES
        case "write":
            result = WRITE_APPPERMISSIONS_PAGES
        default:
            return 0, errors.New("Unknown AppPermissions_pages value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_pages(values []AppPermissions_pages) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_pages) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The level of permission to grant the access token to view events triggered by an activity in an organization.
type AppPermissions_organization_events int

const (
    READ_APPPERMISSIONS_ORGANIZATION_EVENTS AppPermissions_organization_events = iota
)

func (i AppPermissions_organization_events) String() string {
    return []string{"read"}[i]
}
func ParseAppPermissions_organization_events(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_EVENTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_EVENTS
        default:
            return 0, errors.New("Unknown AppPermissions_organization_events value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_events(values []AppPermissions_organization_events) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_events) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The level of permission to grant the access token for viewing an organization's plan.
type AppPermissions_organization_plan int

const (
    READ_APPPERMISSIONS_ORGANIZATION_PLAN AppPermissions_organization_plan = iota
)

func (i AppPermissions_organization_plan) String() string {
    return []string{"read"}[i]
}
func ParseAppPermissions_organization_plan(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_PLAN
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_PLAN
        default:
            return 0, errors.New("Unknown AppPermissions_organization_plan value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_plan(values []AppPermissions_organization_plan) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_plan) isMultiValue() bool {
    return false
}

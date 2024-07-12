package models
import (
    "errors"
)
// The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in beta and is subject to change.
type AppPermissions_organization_copilot_seat_management int

const (
    WRITE_APPPERMISSIONS_ORGANIZATION_COPILOT_SEAT_MANAGEMENT AppPermissions_organization_copilot_seat_management = iota
)

func (i AppPermissions_organization_copilot_seat_management) String() string {
    return []string{"write"}[i]
}
func ParseAppPermissions_organization_copilot_seat_management(v string) (any, error) {
    result := WRITE_APPPERMISSIONS_ORGANIZATION_COPILOT_SEAT_MANAGEMENT
    switch v {
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_COPILOT_SEAT_MANAGEMENT
        default:
            return 0, errors.New("Unknown AppPermissions_organization_copilot_seat_management value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_copilot_seat_management(values []AppPermissions_organization_copilot_seat_management) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_copilot_seat_management) isMultiValue() bool {
    return false
}

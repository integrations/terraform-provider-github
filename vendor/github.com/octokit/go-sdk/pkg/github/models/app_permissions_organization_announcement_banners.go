package models
import (
    "errors"
)
// The level of permission to grant the access token to view and manage announcement banners for an organization.
type AppPermissions_organization_announcement_banners int

const (
    READ_APPPERMISSIONS_ORGANIZATION_ANNOUNCEMENT_BANNERS AppPermissions_organization_announcement_banners = iota
    WRITE_APPPERMISSIONS_ORGANIZATION_ANNOUNCEMENT_BANNERS
)

func (i AppPermissions_organization_announcement_banners) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_organization_announcement_banners(v string) (any, error) {
    result := READ_APPPERMISSIONS_ORGANIZATION_ANNOUNCEMENT_BANNERS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_ORGANIZATION_ANNOUNCEMENT_BANNERS
        case "write":
            result = WRITE_APPPERMISSIONS_ORGANIZATION_ANNOUNCEMENT_BANNERS
        default:
            return 0, errors.New("Unknown AppPermissions_organization_announcement_banners value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_organization_announcement_banners(values []AppPermissions_organization_announcement_banners) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_organization_announcement_banners) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The level of permission to grant the access token to view and manage secret scanning alerts.
type AppPermissions_secret_scanning_alerts int

const (
    READ_APPPERMISSIONS_SECRET_SCANNING_ALERTS AppPermissions_secret_scanning_alerts = iota
    WRITE_APPPERMISSIONS_SECRET_SCANNING_ALERTS
)

func (i AppPermissions_secret_scanning_alerts) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_secret_scanning_alerts(v string) (any, error) {
    result := READ_APPPERMISSIONS_SECRET_SCANNING_ALERTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_SECRET_SCANNING_ALERTS
        case "write":
            result = WRITE_APPPERMISSIONS_SECRET_SCANNING_ALERTS
        default:
            return 0, errors.New("Unknown AppPermissions_secret_scanning_alerts value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_secret_scanning_alerts(values []AppPermissions_secret_scanning_alerts) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_secret_scanning_alerts) isMultiValue() bool {
    return false
}

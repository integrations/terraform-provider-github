package models
import (
    "errors"
)
// The level of permission to grant the access token to view and manage security events like code scanning alerts.
type AppPermissions_security_events int

const (
    READ_APPPERMISSIONS_SECURITY_EVENTS AppPermissions_security_events = iota
    WRITE_APPPERMISSIONS_SECURITY_EVENTS
)

func (i AppPermissions_security_events) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_security_events(v string) (any, error) {
    result := READ_APPPERMISSIONS_SECURITY_EVENTS
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_SECURITY_EVENTS
        case "write":
            result = WRITE_APPPERMISSIONS_SECURITY_EVENTS
        default:
            return 0, errors.New("Unknown AppPermissions_security_events value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_security_events(values []AppPermissions_security_events) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_security_events) isMultiValue() bool {
    return false
}

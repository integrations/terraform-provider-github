package models
import (
    "errors"
)
// The level of permission to grant the access token to manage the email addresses belonging to a user.
type AppPermissions_email_addresses int

const (
    READ_APPPERMISSIONS_EMAIL_ADDRESSES AppPermissions_email_addresses = iota
    WRITE_APPPERMISSIONS_EMAIL_ADDRESSES
)

func (i AppPermissions_email_addresses) String() string {
    return []string{"read", "write"}[i]
}
func ParseAppPermissions_email_addresses(v string) (any, error) {
    result := READ_APPPERMISSIONS_EMAIL_ADDRESSES
    switch v {
        case "read":
            result = READ_APPPERMISSIONS_EMAIL_ADDRESSES
        case "write":
            result = WRITE_APPPERMISSIONS_EMAIL_ADDRESSES
        default:
            return 0, errors.New("Unknown AppPermissions_email_addresses value: " + v)
    }
    return &result, nil
}
func SerializeAppPermissions_email_addresses(values []AppPermissions_email_addresses) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppPermissions_email_addresses) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The token status as of the latest validity check.
type OrganizationSecretScanningAlert_validity int

const (
    ACTIVE_ORGANIZATIONSECRETSCANNINGALERT_VALIDITY OrganizationSecretScanningAlert_validity = iota
    INACTIVE_ORGANIZATIONSECRETSCANNINGALERT_VALIDITY
    UNKNOWN_ORGANIZATIONSECRETSCANNINGALERT_VALIDITY
)

func (i OrganizationSecretScanningAlert_validity) String() string {
    return []string{"active", "inactive", "unknown"}[i]
}
func ParseOrganizationSecretScanningAlert_validity(v string) (any, error) {
    result := ACTIVE_ORGANIZATIONSECRETSCANNINGALERT_VALIDITY
    switch v {
        case "active":
            result = ACTIVE_ORGANIZATIONSECRETSCANNINGALERT_VALIDITY
        case "inactive":
            result = INACTIVE_ORGANIZATIONSECRETSCANNINGALERT_VALIDITY
        case "unknown":
            result = UNKNOWN_ORGANIZATIONSECRETSCANNINGALERT_VALIDITY
        default:
            return 0, errors.New("Unknown OrganizationSecretScanningAlert_validity value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationSecretScanningAlert_validity(values []OrganizationSecretScanningAlert_validity) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrganizationSecretScanningAlert_validity) isMultiValue() bool {
    return false
}

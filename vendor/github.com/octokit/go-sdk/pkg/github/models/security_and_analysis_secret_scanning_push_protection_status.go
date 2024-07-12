package models
import (
    "errors"
)
type SecurityAndAnalysis_secret_scanning_push_protection_status int

const (
    ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_PUSH_PROTECTION_STATUS SecurityAndAnalysis_secret_scanning_push_protection_status = iota
    DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_PUSH_PROTECTION_STATUS
)

func (i SecurityAndAnalysis_secret_scanning_push_protection_status) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseSecurityAndAnalysis_secret_scanning_push_protection_status(v string) (any, error) {
    result := ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_PUSH_PROTECTION_STATUS
    switch v {
        case "enabled":
            result = ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_PUSH_PROTECTION_STATUS
        case "disabled":
            result = DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_PUSH_PROTECTION_STATUS
        default:
            return 0, errors.New("Unknown SecurityAndAnalysis_secret_scanning_push_protection_status value: " + v)
    }
    return &result, nil
}
func SerializeSecurityAndAnalysis_secret_scanning_push_protection_status(values []SecurityAndAnalysis_secret_scanning_push_protection_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAndAnalysis_secret_scanning_push_protection_status) isMultiValue() bool {
    return false
}

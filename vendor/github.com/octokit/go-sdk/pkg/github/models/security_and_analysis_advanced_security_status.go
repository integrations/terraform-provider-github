package models
import (
    "errors"
)
type SecurityAndAnalysis_advanced_security_status int

const (
    ENABLED_SECURITYANDANALYSIS_ADVANCED_SECURITY_STATUS SecurityAndAnalysis_advanced_security_status = iota
    DISABLED_SECURITYANDANALYSIS_ADVANCED_SECURITY_STATUS
)

func (i SecurityAndAnalysis_advanced_security_status) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseSecurityAndAnalysis_advanced_security_status(v string) (any, error) {
    result := ENABLED_SECURITYANDANALYSIS_ADVANCED_SECURITY_STATUS
    switch v {
        case "enabled":
            result = ENABLED_SECURITYANDANALYSIS_ADVANCED_SECURITY_STATUS
        case "disabled":
            result = DISABLED_SECURITYANDANALYSIS_ADVANCED_SECURITY_STATUS
        default:
            return 0, errors.New("Unknown SecurityAndAnalysis_advanced_security_status value: " + v)
    }
    return &result, nil
}
func SerializeSecurityAndAnalysis_advanced_security_status(values []SecurityAndAnalysis_advanced_security_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAndAnalysis_advanced_security_status) isMultiValue() bool {
    return false
}

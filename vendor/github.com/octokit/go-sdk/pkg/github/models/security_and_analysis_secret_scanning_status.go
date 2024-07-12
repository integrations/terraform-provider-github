package models
import (
    "errors"
)
type SecurityAndAnalysis_secret_scanning_status int

const (
    ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_STATUS SecurityAndAnalysis_secret_scanning_status = iota
    DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_STATUS
)

func (i SecurityAndAnalysis_secret_scanning_status) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseSecurityAndAnalysis_secret_scanning_status(v string) (any, error) {
    result := ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_STATUS
    switch v {
        case "enabled":
            result = ENABLED_SECURITYANDANALYSIS_SECRET_SCANNING_STATUS
        case "disabled":
            result = DISABLED_SECURITYANDANALYSIS_SECRET_SCANNING_STATUS
        default:
            return 0, errors.New("Unknown SecurityAndAnalysis_secret_scanning_status value: " + v)
    }
    return &result, nil
}
func SerializeSecurityAndAnalysis_secret_scanning_status(values []SecurityAndAnalysis_secret_scanning_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAndAnalysis_secret_scanning_status) isMultiValue() bool {
    return false
}

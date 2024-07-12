package models
import (
    "errors"
)
// The enablement status of Dependabot security updates for the repository.
type SecurityAndAnalysis_dependabot_security_updates_status int

const (
    ENABLED_SECURITYANDANALYSIS_DEPENDABOT_SECURITY_UPDATES_STATUS SecurityAndAnalysis_dependabot_security_updates_status = iota
    DISABLED_SECURITYANDANALYSIS_DEPENDABOT_SECURITY_UPDATES_STATUS
)

func (i SecurityAndAnalysis_dependabot_security_updates_status) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseSecurityAndAnalysis_dependabot_security_updates_status(v string) (any, error) {
    result := ENABLED_SECURITYANDANALYSIS_DEPENDABOT_SECURITY_UPDATES_STATUS
    switch v {
        case "enabled":
            result = ENABLED_SECURITYANDANALYSIS_DEPENDABOT_SECURITY_UPDATES_STATUS
        case "disabled":
            result = DISABLED_SECURITYANDANALYSIS_DEPENDABOT_SECURITY_UPDATES_STATUS
        default:
            return 0, errors.New("Unknown SecurityAndAnalysis_dependabot_security_updates_status value: " + v)
    }
    return &result, nil
}
func SerializeSecurityAndAnalysis_dependabot_security_updates_status(values []SecurityAndAnalysis_dependabot_security_updates_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAndAnalysis_dependabot_security_updates_status) isMultiValue() bool {
    return false
}

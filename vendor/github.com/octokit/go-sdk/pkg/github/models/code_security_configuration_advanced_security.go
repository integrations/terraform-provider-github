package models
import (
    "errors"
)
// The enablement status of GitHub Advanced Security
type CodeSecurityConfiguration_advanced_security int

const (
    ENABLED_CODESECURITYCONFIGURATION_ADVANCED_SECURITY CodeSecurityConfiguration_advanced_security = iota
    DISABLED_CODESECURITYCONFIGURATION_ADVANCED_SECURITY
)

func (i CodeSecurityConfiguration_advanced_security) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseCodeSecurityConfiguration_advanced_security(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_ADVANCED_SECURITY
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_ADVANCED_SECURITY
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_ADVANCED_SECURITY
        default:
            return 0, errors.New("Unknown CodeSecurityConfiguration_advanced_security value: " + v)
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_advanced_security(values []CodeSecurityConfiguration_advanced_security) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_advanced_security) isMultiValue() bool {
    return false
}

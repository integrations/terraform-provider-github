package models
import (
    "errors"
)
// The enablement status of Dependabot security updates
type CodeSecurityConfiguration_dependabot_security_updates int

const (
    ENABLED_CODESECURITYCONFIGURATION_DEPENDABOT_SECURITY_UPDATES CodeSecurityConfiguration_dependabot_security_updates = iota
    DISABLED_CODESECURITYCONFIGURATION_DEPENDABOT_SECURITY_UPDATES
    NOT_SET_CODESECURITYCONFIGURATION_DEPENDABOT_SECURITY_UPDATES
)

func (i CodeSecurityConfiguration_dependabot_security_updates) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_dependabot_security_updates(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_DEPENDABOT_SECURITY_UPDATES
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_DEPENDABOT_SECURITY_UPDATES
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_DEPENDABOT_SECURITY_UPDATES
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_DEPENDABOT_SECURITY_UPDATES
        default:
            return 0, errors.New("Unknown CodeSecurityConfiguration_dependabot_security_updates value: " + v)
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_dependabot_security_updates(values []CodeSecurityConfiguration_dependabot_security_updates) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_dependabot_security_updates) isMultiValue() bool {
    return false
}

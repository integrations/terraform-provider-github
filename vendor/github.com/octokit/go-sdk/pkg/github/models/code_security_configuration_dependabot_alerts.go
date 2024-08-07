package models
import (
    "errors"
)
// The enablement status of Dependabot alerts
type CodeSecurityConfiguration_dependabot_alerts int

const (
    ENABLED_CODESECURITYCONFIGURATION_DEPENDABOT_ALERTS CodeSecurityConfiguration_dependabot_alerts = iota
    DISABLED_CODESECURITYCONFIGURATION_DEPENDABOT_ALERTS
    NOT_SET_CODESECURITYCONFIGURATION_DEPENDABOT_ALERTS
)

func (i CodeSecurityConfiguration_dependabot_alerts) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_dependabot_alerts(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_DEPENDABOT_ALERTS
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_DEPENDABOT_ALERTS
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_DEPENDABOT_ALERTS
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_DEPENDABOT_ALERTS
        default:
            return 0, errors.New("Unknown CodeSecurityConfiguration_dependabot_alerts value: " + v)
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_dependabot_alerts(values []CodeSecurityConfiguration_dependabot_alerts) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_dependabot_alerts) isMultiValue() bool {
    return false
}

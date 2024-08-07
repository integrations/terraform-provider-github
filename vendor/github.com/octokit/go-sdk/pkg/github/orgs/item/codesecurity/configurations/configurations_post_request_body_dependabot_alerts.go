package configurations
import (
    "errors"
)
// The enablement status of Dependabot alerts
type ConfigurationsPostRequestBody_dependabot_alerts int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_ALERTS ConfigurationsPostRequestBody_dependabot_alerts = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_ALERTS
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_ALERTS
)

func (i ConfigurationsPostRequestBody_dependabot_alerts) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_dependabot_alerts(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_ALERTS
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_ALERTS
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_ALERTS
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_ALERTS
        default:
            return 0, errors.New("Unknown ConfigurationsPostRequestBody_dependabot_alerts value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_dependabot_alerts(values []ConfigurationsPostRequestBody_dependabot_alerts) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_dependabot_alerts) isMultiValue() bool {
    return false
}

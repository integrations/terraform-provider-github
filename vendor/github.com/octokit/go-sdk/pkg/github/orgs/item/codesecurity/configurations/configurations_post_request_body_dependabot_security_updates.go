package configurations
import (
    "errors"
)
// The enablement status of Dependabot security updates
type ConfigurationsPostRequestBody_dependabot_security_updates int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_SECURITY_UPDATES ConfigurationsPostRequestBody_dependabot_security_updates = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
)

func (i ConfigurationsPostRequestBody_dependabot_security_updates) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_dependabot_security_updates(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_DEPENDABOT_SECURITY_UPDATES
        default:
            return 0, errors.New("Unknown ConfigurationsPostRequestBody_dependabot_security_updates value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_dependabot_security_updates(values []ConfigurationsPostRequestBody_dependabot_security_updates) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_dependabot_security_updates) isMultiValue() bool {
    return false
}

package configurations
import (
    "errors"
)
// The enablement status of GitHub Advanced Security
type ConfigurationsPostRequestBody_advanced_security int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_ADVANCED_SECURITY ConfigurationsPostRequestBody_advanced_security = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_ADVANCED_SECURITY
)

func (i ConfigurationsPostRequestBody_advanced_security) String() string {
    return []string{"enabled", "disabled"}[i]
}
func ParseConfigurationsPostRequestBody_advanced_security(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_ADVANCED_SECURITY
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_ADVANCED_SECURITY
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_ADVANCED_SECURITY
        default:
            return 0, errors.New("Unknown ConfigurationsPostRequestBody_advanced_security value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_advanced_security(values []ConfigurationsPostRequestBody_advanced_security) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_advanced_security) isMultiValue() bool {
    return false
}

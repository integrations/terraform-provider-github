package configurations
import (
    "errors"
)
// The enablement status of secret scanning push protection
type ConfigurationsPostRequestBody_secret_scanning_push_protection int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION ConfigurationsPostRequestBody_secret_scanning_push_protection = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
)

func (i ConfigurationsPostRequestBody_secret_scanning_push_protection) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_secret_scanning_push_protection(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING_PUSH_PROTECTION
        default:
            return 0, errors.New("Unknown ConfigurationsPostRequestBody_secret_scanning_push_protection value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_secret_scanning_push_protection(values []ConfigurationsPostRequestBody_secret_scanning_push_protection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_secret_scanning_push_protection) isMultiValue() bool {
    return false
}

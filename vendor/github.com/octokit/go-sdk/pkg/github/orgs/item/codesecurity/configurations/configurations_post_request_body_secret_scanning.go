package configurations
import (
    "errors"
)
// The enablement status of secret scanning
type ConfigurationsPostRequestBody_secret_scanning int

const (
    ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING ConfigurationsPostRequestBody_secret_scanning = iota
    DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING
    NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING
)

func (i ConfigurationsPostRequestBody_secret_scanning) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseConfigurationsPostRequestBody_secret_scanning(v string) (any, error) {
    result := ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING
    switch v {
        case "enabled":
            result = ENABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING
        case "disabled":
            result = DISABLED_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING
        case "not_set":
            result = NOT_SET_CONFIGURATIONSPOSTREQUESTBODY_SECRET_SCANNING
        default:
            return 0, errors.New("Unknown ConfigurationsPostRequestBody_secret_scanning value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationsPostRequestBody_secret_scanning(values []ConfigurationsPostRequestBody_secret_scanning) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConfigurationsPostRequestBody_secret_scanning) isMultiValue() bool {
    return false
}

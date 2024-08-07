package models
import (
    "errors"
)
// The enablement status of code scanning default setup
type CodeSecurityConfiguration_code_scanning_default_setup int

const (
    ENABLED_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP CodeSecurityConfiguration_code_scanning_default_setup = iota
    DISABLED_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP
    NOT_SET_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP
)

func (i CodeSecurityConfiguration_code_scanning_default_setup) String() string {
    return []string{"enabled", "disabled", "not_set"}[i]
}
func ParseCodeSecurityConfiguration_code_scanning_default_setup(v string) (any, error) {
    result := ENABLED_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP
    switch v {
        case "enabled":
            result = ENABLED_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP
        case "disabled":
            result = DISABLED_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP
        case "not_set":
            result = NOT_SET_CODESECURITYCONFIGURATION_CODE_SCANNING_DEFAULT_SETUP
        default:
            return 0, errors.New("Unknown CodeSecurityConfiguration_code_scanning_default_setup value: " + v)
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_code_scanning_default_setup(values []CodeSecurityConfiguration_code_scanning_default_setup) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_code_scanning_default_setup) isMultiValue() bool {
    return false
}

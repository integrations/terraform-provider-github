package models
import (
    "errors"
)
// The type of the code security configuration.
type CodeSecurityConfiguration_target_type int

const (
    GLOBAL_CODESECURITYCONFIGURATION_TARGET_TYPE CodeSecurityConfiguration_target_type = iota
    ORGANIZATION_CODESECURITYCONFIGURATION_TARGET_TYPE
)

func (i CodeSecurityConfiguration_target_type) String() string {
    return []string{"global", "organization"}[i]
}
func ParseCodeSecurityConfiguration_target_type(v string) (any, error) {
    result := GLOBAL_CODESECURITYCONFIGURATION_TARGET_TYPE
    switch v {
        case "global":
            result = GLOBAL_CODESECURITYCONFIGURATION_TARGET_TYPE
        case "organization":
            result = ORGANIZATION_CODESECURITYCONFIGURATION_TARGET_TYPE
        default:
            return 0, errors.New("Unknown CodeSecurityConfiguration_target_type value: " + v)
    }
    return &result, nil
}
func SerializeCodeSecurityConfiguration_target_type(values []CodeSecurityConfiguration_target_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeSecurityConfiguration_target_type) isMultiValue() bool {
    return false
}

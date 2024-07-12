package models
import (
    "errors"
)
// Code scanning default setup has been configured or not.
type CodeScanningDefaultSetup_state int

const (
    CONFIGURED_CODESCANNINGDEFAULTSETUP_STATE CodeScanningDefaultSetup_state = iota
    NOTCONFIGURED_CODESCANNINGDEFAULTSETUP_STATE
)

func (i CodeScanningDefaultSetup_state) String() string {
    return []string{"configured", "not-configured"}[i]
}
func ParseCodeScanningDefaultSetup_state(v string) (any, error) {
    result := CONFIGURED_CODESCANNINGDEFAULTSETUP_STATE
    switch v {
        case "configured":
            result = CONFIGURED_CODESCANNINGDEFAULTSETUP_STATE
        case "not-configured":
            result = NOTCONFIGURED_CODESCANNINGDEFAULTSETUP_STATE
        default:
            return 0, errors.New("Unknown CodeScanningDefaultSetup_state value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetup_state(values []CodeScanningDefaultSetup_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetup_state) isMultiValue() bool {
    return false
}

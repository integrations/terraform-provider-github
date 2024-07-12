package models
import (
    "errors"
)
// The desired state of code scanning default setup.
type CodeScanningDefaultSetupUpdate_state int

const (
    CONFIGURED_CODESCANNINGDEFAULTSETUPUPDATE_STATE CodeScanningDefaultSetupUpdate_state = iota
    NOTCONFIGURED_CODESCANNINGDEFAULTSETUPUPDATE_STATE
)

func (i CodeScanningDefaultSetupUpdate_state) String() string {
    return []string{"configured", "not-configured"}[i]
}
func ParseCodeScanningDefaultSetupUpdate_state(v string) (any, error) {
    result := CONFIGURED_CODESCANNINGDEFAULTSETUPUPDATE_STATE
    switch v {
        case "configured":
            result = CONFIGURED_CODESCANNINGDEFAULTSETUPUPDATE_STATE
        case "not-configured":
            result = NOTCONFIGURED_CODESCANNINGDEFAULTSETUPUPDATE_STATE
        default:
            return 0, errors.New("Unknown CodeScanningDefaultSetupUpdate_state value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetupUpdate_state(values []CodeScanningDefaultSetupUpdate_state) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetupUpdate_state) isMultiValue() bool {
    return false
}

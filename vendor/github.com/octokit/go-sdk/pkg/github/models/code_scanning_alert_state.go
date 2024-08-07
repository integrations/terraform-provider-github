package models
import (
    "errors"
)
// State of a code scanning alert.
type CodeScanningAlertState int

const (
    OPEN_CODESCANNINGALERTSTATE CodeScanningAlertState = iota
    DISMISSED_CODESCANNINGALERTSTATE
    FIXED_CODESCANNINGALERTSTATE
)

func (i CodeScanningAlertState) String() string {
    return []string{"open", "dismissed", "fixed"}[i]
}
func ParseCodeScanningAlertState(v string) (any, error) {
    result := OPEN_CODESCANNINGALERTSTATE
    switch v {
        case "open":
            result = OPEN_CODESCANNINGALERTSTATE
        case "dismissed":
            result = DISMISSED_CODESCANNINGALERTSTATE
        case "fixed":
            result = FIXED_CODESCANNINGALERTSTATE
        default:
            return 0, errors.New("Unknown CodeScanningAlertState value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertState(values []CodeScanningAlertState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertState) isMultiValue() bool {
    return false
}

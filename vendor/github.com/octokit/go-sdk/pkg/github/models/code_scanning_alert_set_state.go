package models
import (
    "errors"
)
// Sets the state of the code scanning alert. You must provide `dismissed_reason` when you set the state to `dismissed`.
type CodeScanningAlertSetState int

const (
    OPEN_CODESCANNINGALERTSETSTATE CodeScanningAlertSetState = iota
    DISMISSED_CODESCANNINGALERTSETSTATE
)

func (i CodeScanningAlertSetState) String() string {
    return []string{"open", "dismissed"}[i]
}
func ParseCodeScanningAlertSetState(v string) (any, error) {
    result := OPEN_CODESCANNINGALERTSETSTATE
    switch v {
        case "open":
            result = OPEN_CODESCANNINGALERTSETSTATE
        case "dismissed":
            result = DISMISSED_CODESCANNINGALERTSETSTATE
        default:
            return 0, errors.New("Unknown CodeScanningAlertSetState value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertSetState(values []CodeScanningAlertSetState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertSetState) isMultiValue() bool {
    return false
}

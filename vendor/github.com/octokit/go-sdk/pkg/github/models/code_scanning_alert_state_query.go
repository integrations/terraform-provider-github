package models
import (
    "errors"
)
// State of a code scanning alert.
type CodeScanningAlertStateQuery int

const (
    OPEN_CODESCANNINGALERTSTATEQUERY CodeScanningAlertStateQuery = iota
    CLOSED_CODESCANNINGALERTSTATEQUERY
    DISMISSED_CODESCANNINGALERTSTATEQUERY
    FIXED_CODESCANNINGALERTSTATEQUERY
)

func (i CodeScanningAlertStateQuery) String() string {
    return []string{"open", "closed", "dismissed", "fixed"}[i]
}
func ParseCodeScanningAlertStateQuery(v string) (any, error) {
    result := OPEN_CODESCANNINGALERTSTATEQUERY
    switch v {
        case "open":
            result = OPEN_CODESCANNINGALERTSTATEQUERY
        case "closed":
            result = CLOSED_CODESCANNINGALERTSTATEQUERY
        case "dismissed":
            result = DISMISSED_CODESCANNINGALERTSTATEQUERY
        case "fixed":
            result = FIXED_CODESCANNINGALERTSTATEQUERY
        default:
            return 0, errors.New("Unknown CodeScanningAlertStateQuery value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningAlertStateQuery(values []CodeScanningAlertStateQuery) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningAlertStateQuery) isMultiValue() bool {
    return false
}

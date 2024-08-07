package models
import (
    "errors"
)
// CodeQL query suite to be used.
type CodeScanningDefaultSetup_query_suite int

const (
    DEFAULT_CODESCANNINGDEFAULTSETUP_QUERY_SUITE CodeScanningDefaultSetup_query_suite = iota
    EXTENDED_CODESCANNINGDEFAULTSETUP_QUERY_SUITE
)

func (i CodeScanningDefaultSetup_query_suite) String() string {
    return []string{"default", "extended"}[i]
}
func ParseCodeScanningDefaultSetup_query_suite(v string) (any, error) {
    result := DEFAULT_CODESCANNINGDEFAULTSETUP_QUERY_SUITE
    switch v {
        case "default":
            result = DEFAULT_CODESCANNINGDEFAULTSETUP_QUERY_SUITE
        case "extended":
            result = EXTENDED_CODESCANNINGDEFAULTSETUP_QUERY_SUITE
        default:
            return 0, errors.New("Unknown CodeScanningDefaultSetup_query_suite value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetup_query_suite(values []CodeScanningDefaultSetup_query_suite) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetup_query_suite) isMultiValue() bool {
    return false
}

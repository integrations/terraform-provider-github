package models
import (
    "errors"
)
// CodeQL query suite to be used.
type CodeScanningDefaultSetupUpdate_query_suite int

const (
    DEFAULT_CODESCANNINGDEFAULTSETUPUPDATE_QUERY_SUITE CodeScanningDefaultSetupUpdate_query_suite = iota
    EXTENDED_CODESCANNINGDEFAULTSETUPUPDATE_QUERY_SUITE
)

func (i CodeScanningDefaultSetupUpdate_query_suite) String() string {
    return []string{"default", "extended"}[i]
}
func ParseCodeScanningDefaultSetupUpdate_query_suite(v string) (any, error) {
    result := DEFAULT_CODESCANNINGDEFAULTSETUPUPDATE_QUERY_SUITE
    switch v {
        case "default":
            result = DEFAULT_CODESCANNINGDEFAULTSETUPUPDATE_QUERY_SUITE
        case "extended":
            result = EXTENDED_CODESCANNINGDEFAULTSETUPUPDATE_QUERY_SUITE
        default:
            return 0, errors.New("Unknown CodeScanningDefaultSetupUpdate_query_suite value: " + v)
    }
    return &result, nil
}
func SerializeCodeScanningDefaultSetupUpdate_query_suite(values []CodeScanningDefaultSetupUpdate_query_suite) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CodeScanningDefaultSetupUpdate_query_suite) isMultiValue() bool {
    return false
}

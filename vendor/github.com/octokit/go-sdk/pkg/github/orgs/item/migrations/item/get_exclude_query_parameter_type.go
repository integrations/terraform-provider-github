package item
import (
    "errors"
)
// Allowed values that can be passed to the exclude param.
type GetExcludeQueryParameterType int

const (
    REPOSITORIES_GETEXCLUDEQUERYPARAMETERTYPE GetExcludeQueryParameterType = iota
)

func (i GetExcludeQueryParameterType) String() string {
    return []string{"repositories"}[i]
}
func ParseGetExcludeQueryParameterType(v string) (any, error) {
    result := REPOSITORIES_GETEXCLUDEQUERYPARAMETERTYPE
    switch v {
        case "repositories":
            result = REPOSITORIES_GETEXCLUDEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetExcludeQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetExcludeQueryParameterType(values []GetExcludeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetExcludeQueryParameterType) isMultiValue() bool {
    return false
}

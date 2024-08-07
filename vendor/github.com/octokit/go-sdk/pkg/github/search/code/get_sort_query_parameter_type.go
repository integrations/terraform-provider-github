package code
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    INDEXED_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
)

func (i GetSortQueryParameterType) String() string {
    return []string{"indexed"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := INDEXED_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "indexed":
            result = INDEXED_GETSORTQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetSortQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetSortQueryParameterType(values []GetSortQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSortQueryParameterType) isMultiValue() bool {
    return false
}

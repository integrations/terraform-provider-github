package caches
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    CREATED_AT_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    LAST_ACCESSED_AT_GETSORTQUERYPARAMETERTYPE
    SIZE_IN_BYTES_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"created_at", "last_accessed_at", "size_in_bytes"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := CREATED_AT_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "created_at":
            result = CREATED_AT_GETSORTQUERYPARAMETERTYPE
        case "last_accessed_at":
            result = LAST_ACCESSED_AT_GETSORTQUERYPARAMETERTYPE
        case "size_in_bytes":
            result = SIZE_IN_BYTES_GETSORTQUERYPARAMETERTYPE
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

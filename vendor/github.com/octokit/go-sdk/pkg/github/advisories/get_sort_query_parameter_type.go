package advisories
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    UPDATED_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    PUBLISHED_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"updated", "published"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := UPDATED_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "updated":
            result = UPDATED_GETSORTQUERYPARAMETERTYPE
        case "published":
            result = PUBLISHED_GETSORTQUERYPARAMETERTYPE
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

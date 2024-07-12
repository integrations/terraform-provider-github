package personalaccesstokens
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    CREATED_AT_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
)

func (i GetSortQueryParameterType) String() string {
    return []string{"created_at"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := CREATED_AT_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "created_at":
            result = CREATED_AT_GETSORTQUERYPARAMETERTYPE
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

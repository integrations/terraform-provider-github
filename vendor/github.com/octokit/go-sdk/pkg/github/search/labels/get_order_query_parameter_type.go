package labels
import (
    "errors"
)
type GetOrderQueryParameterType int

const (
    DESC_GETORDERQUERYPARAMETERTYPE GetOrderQueryParameterType = iota
    ASC_GETORDERQUERYPARAMETERTYPE
)

func (i GetOrderQueryParameterType) String() string {
    return []string{"desc", "asc"}[i]
}
func ParseGetOrderQueryParameterType(v string) (any, error) {
    result := DESC_GETORDERQUERYPARAMETERTYPE
    switch v {
        case "desc":
            result = DESC_GETORDERQUERYPARAMETERTYPE
        case "asc":
            result = ASC_GETORDERQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetOrderQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetOrderQueryParameterType(values []GetOrderQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetOrderQueryParameterType) isMultiValue() bool {
    return false
}

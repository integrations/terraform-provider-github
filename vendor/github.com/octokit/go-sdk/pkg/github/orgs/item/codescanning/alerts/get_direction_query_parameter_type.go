package alerts
import (
    "errors"
)
type GetDirectionQueryParameterType int

const (
    ASC_GETDIRECTIONQUERYPARAMETERTYPE GetDirectionQueryParameterType = iota
    DESC_GETDIRECTIONQUERYPARAMETERTYPE
)

func (i GetDirectionQueryParameterType) String() string {
    return []string{"asc", "desc"}[i]
}
func ParseGetDirectionQueryParameterType(v string) (any, error) {
    result := ASC_GETDIRECTIONQUERYPARAMETERTYPE
    switch v {
        case "asc":
            result = ASC_GETDIRECTIONQUERYPARAMETERTYPE
        case "desc":
            result = DESC_GETDIRECTIONQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetDirectionQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetDirectionQueryParameterType(values []GetDirectionQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetDirectionQueryParameterType) isMultiValue() bool {
    return false
}

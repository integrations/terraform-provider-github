package jobs
import (
    "errors"
)
type GetFilterQueryParameterType int

const (
    LATEST_GETFILTERQUERYPARAMETERTYPE GetFilterQueryParameterType = iota
    ALL_GETFILTERQUERYPARAMETERTYPE
)

func (i GetFilterQueryParameterType) String() string {
    return []string{"latest", "all"}[i]
}
func ParseGetFilterQueryParameterType(v string) (any, error) {
    result := LATEST_GETFILTERQUERYPARAMETERTYPE
    switch v {
        case "latest":
            result = LATEST_GETFILTERQUERYPARAMETERTYPE
        case "all":
            result = ALL_GETFILTERQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetFilterQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetFilterQueryParameterType(values []GetFilterQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetFilterQueryParameterType) isMultiValue() bool {
    return false
}

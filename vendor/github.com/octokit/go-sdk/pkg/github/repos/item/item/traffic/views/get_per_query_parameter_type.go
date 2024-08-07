package views
import (
    "errors"
)
type GetPerQueryParameterType int

const (
    DAY_GETPERQUERYPARAMETERTYPE GetPerQueryParameterType = iota
    WEEK_GETPERQUERYPARAMETERTYPE
)

func (i GetPerQueryParameterType) String() string {
    return []string{"day", "week"}[i]
}
func ParseGetPerQueryParameterType(v string) (any, error) {
    result := DAY_GETPERQUERYPARAMETERTYPE
    switch v {
        case "day":
            result = DAY_GETPERQUERYPARAMETERTYPE
        case "week":
            result = WEEK_GETPERQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetPerQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetPerQueryParameterType(values []GetPerQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetPerQueryParameterType) isMultiValue() bool {
    return false
}

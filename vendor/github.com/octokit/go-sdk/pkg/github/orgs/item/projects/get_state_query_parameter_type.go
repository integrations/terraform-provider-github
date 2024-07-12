package projects
import (
    "errors"
)
type GetStateQueryParameterType int

const (
    OPEN_GETSTATEQUERYPARAMETERTYPE GetStateQueryParameterType = iota
    CLOSED_GETSTATEQUERYPARAMETERTYPE
    ALL_GETSTATEQUERYPARAMETERTYPE
)

func (i GetStateQueryParameterType) String() string {
    return []string{"open", "closed", "all"}[i]
}
func ParseGetStateQueryParameterType(v string) (any, error) {
    result := OPEN_GETSTATEQUERYPARAMETERTYPE
    switch v {
        case "open":
            result = OPEN_GETSTATEQUERYPARAMETERTYPE
        case "closed":
            result = CLOSED_GETSTATEQUERYPARAMETERTYPE
        case "all":
            result = ALL_GETSTATEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetStateQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetStateQueryParameterType(values []GetStateQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetStateQueryParameterType) isMultiValue() bool {
    return false
}

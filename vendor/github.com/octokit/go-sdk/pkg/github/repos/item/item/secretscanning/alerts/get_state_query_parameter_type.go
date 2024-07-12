package alerts
import (
    "errors"
)
type GetStateQueryParameterType int

const (
    OPEN_GETSTATEQUERYPARAMETERTYPE GetStateQueryParameterType = iota
    RESOLVED_GETSTATEQUERYPARAMETERTYPE
)

func (i GetStateQueryParameterType) String() string {
    return []string{"open", "resolved"}[i]
}
func ParseGetStateQueryParameterType(v string) (any, error) {
    result := OPEN_GETSTATEQUERYPARAMETERTYPE
    switch v {
        case "open":
            result = OPEN_GETSTATEQUERYPARAMETERTYPE
        case "resolved":
            result = RESOLVED_GETSTATEQUERYPARAMETERTYPE
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

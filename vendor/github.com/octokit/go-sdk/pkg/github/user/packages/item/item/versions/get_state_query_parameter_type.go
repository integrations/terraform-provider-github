package versions
import (
    "errors"
)
type GetStateQueryParameterType int

const (
    ACTIVE_GETSTATEQUERYPARAMETERTYPE GetStateQueryParameterType = iota
    DELETED_GETSTATEQUERYPARAMETERTYPE
)

func (i GetStateQueryParameterType) String() string {
    return []string{"active", "deleted"}[i]
}
func ParseGetStateQueryParameterType(v string) (any, error) {
    result := ACTIVE_GETSTATEQUERYPARAMETERTYPE
    switch v {
        case "active":
            result = ACTIVE_GETSTATEQUERYPARAMETERTYPE
        case "deleted":
            result = DELETED_GETSTATEQUERYPARAMETERTYPE
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

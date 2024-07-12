package configurations
import (
    "errors"
)
type GetTarget_typeQueryParameterType int

const (
    GLOBAL_GETTARGET_TYPEQUERYPARAMETERTYPE GetTarget_typeQueryParameterType = iota
    ALL_GETTARGET_TYPEQUERYPARAMETERTYPE
)

func (i GetTarget_typeQueryParameterType) String() string {
    return []string{"global", "all"}[i]
}
func ParseGetTarget_typeQueryParameterType(v string) (any, error) {
    result := GLOBAL_GETTARGET_TYPEQUERYPARAMETERTYPE
    switch v {
        case "global":
            result = GLOBAL_GETTARGET_TYPEQUERYPARAMETERTYPE
        case "all":
            result = ALL_GETTARGET_TYPEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetTarget_typeQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetTarget_typeQueryParameterType(values []GetTarget_typeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetTarget_typeQueryParameterType) isMultiValue() bool {
    return false
}

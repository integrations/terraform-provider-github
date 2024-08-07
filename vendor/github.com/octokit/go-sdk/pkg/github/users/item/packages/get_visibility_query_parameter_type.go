package packages
import (
    "errors"
)
type GetVisibilityQueryParameterType int

const (
    PUBLIC_GETVISIBILITYQUERYPARAMETERTYPE GetVisibilityQueryParameterType = iota
    PRIVATE_GETVISIBILITYQUERYPARAMETERTYPE
    INTERNAL_GETVISIBILITYQUERYPARAMETERTYPE
)

func (i GetVisibilityQueryParameterType) String() string {
    return []string{"public", "private", "internal"}[i]
}
func ParseGetVisibilityQueryParameterType(v string) (any, error) {
    result := PUBLIC_GETVISIBILITYQUERYPARAMETERTYPE
    switch v {
        case "public":
            result = PUBLIC_GETVISIBILITYQUERYPARAMETERTYPE
        case "private":
            result = PRIVATE_GETVISIBILITYQUERYPARAMETERTYPE
        case "internal":
            result = INTERNAL_GETVISIBILITYQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetVisibilityQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetVisibilityQueryParameterType(values []GetVisibilityQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetVisibilityQueryParameterType) isMultiValue() bool {
    return false
}

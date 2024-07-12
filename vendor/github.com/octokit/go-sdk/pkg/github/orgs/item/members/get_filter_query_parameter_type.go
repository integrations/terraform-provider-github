package members
import (
    "errors"
)
type GetFilterQueryParameterType int

const (
    TWOFA_DISABLED_GETFILTERQUERYPARAMETERTYPE GetFilterQueryParameterType = iota
    ALL_GETFILTERQUERYPARAMETERTYPE
)

func (i GetFilterQueryParameterType) String() string {
    return []string{"2fa_disabled", "all"}[i]
}
func ParseGetFilterQueryParameterType(v string) (any, error) {
    result := TWOFA_DISABLED_GETFILTERQUERYPARAMETERTYPE
    switch v {
        case "2fa_disabled":
            result = TWOFA_DISABLED_GETFILTERQUERYPARAMETERTYPE
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

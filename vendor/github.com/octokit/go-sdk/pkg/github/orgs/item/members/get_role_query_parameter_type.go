package members
import (
    "errors"
)
type GetRoleQueryParameterType int

const (
    ALL_GETROLEQUERYPARAMETERTYPE GetRoleQueryParameterType = iota
    ADMIN_GETROLEQUERYPARAMETERTYPE
    MEMBER_GETROLEQUERYPARAMETERTYPE
)

func (i GetRoleQueryParameterType) String() string {
    return []string{"all", "admin", "member"}[i]
}
func ParseGetRoleQueryParameterType(v string) (any, error) {
    result := ALL_GETROLEQUERYPARAMETERTYPE
    switch v {
        case "all":
            result = ALL_GETROLEQUERYPARAMETERTYPE
        case "admin":
            result = ADMIN_GETROLEQUERYPARAMETERTYPE
        case "member":
            result = MEMBER_GETROLEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetRoleQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetRoleQueryParameterType(values []GetRoleQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetRoleQueryParameterType) isMultiValue() bool {
    return false
}

package members
import (
    "errors"
)
type GetRoleQueryParameterType int

const (
    MEMBER_GETROLEQUERYPARAMETERTYPE GetRoleQueryParameterType = iota
    MAINTAINER_GETROLEQUERYPARAMETERTYPE
    ALL_GETROLEQUERYPARAMETERTYPE
)

func (i GetRoleQueryParameterType) String() string {
    return []string{"member", "maintainer", "all"}[i]
}
func ParseGetRoleQueryParameterType(v string) (any, error) {
    result := MEMBER_GETROLEQUERYPARAMETERTYPE
    switch v {
        case "member":
            result = MEMBER_GETROLEQUERYPARAMETERTYPE
        case "maintainer":
            result = MAINTAINER_GETROLEQUERYPARAMETERTYPE
        case "all":
            result = ALL_GETROLEQUERYPARAMETERTYPE
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

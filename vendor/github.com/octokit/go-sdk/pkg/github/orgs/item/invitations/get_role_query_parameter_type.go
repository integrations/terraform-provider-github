package invitations
import (
    "errors"
)
type GetRoleQueryParameterType int

const (
    ALL_GETROLEQUERYPARAMETERTYPE GetRoleQueryParameterType = iota
    ADMIN_GETROLEQUERYPARAMETERTYPE
    DIRECT_MEMBER_GETROLEQUERYPARAMETERTYPE
    BILLING_MANAGER_GETROLEQUERYPARAMETERTYPE
    HIRING_MANAGER_GETROLEQUERYPARAMETERTYPE
)

func (i GetRoleQueryParameterType) String() string {
    return []string{"all", "admin", "direct_member", "billing_manager", "hiring_manager"}[i]
}
func ParseGetRoleQueryParameterType(v string) (any, error) {
    result := ALL_GETROLEQUERYPARAMETERTYPE
    switch v {
        case "all":
            result = ALL_GETROLEQUERYPARAMETERTYPE
        case "admin":
            result = ADMIN_GETROLEQUERYPARAMETERTYPE
        case "direct_member":
            result = DIRECT_MEMBER_GETROLEQUERYPARAMETERTYPE
        case "billing_manager":
            result = BILLING_MANAGER_GETROLEQUERYPARAMETERTYPE
        case "hiring_manager":
            result = HIRING_MANAGER_GETROLEQUERYPARAMETERTYPE
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

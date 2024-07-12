package collaborators
import (
    "errors"
)
type GetPermissionQueryParameterType int

const (
    PULL_GETPERMISSIONQUERYPARAMETERTYPE GetPermissionQueryParameterType = iota
    TRIAGE_GETPERMISSIONQUERYPARAMETERTYPE
    PUSH_GETPERMISSIONQUERYPARAMETERTYPE
    MAINTAIN_GETPERMISSIONQUERYPARAMETERTYPE
    ADMIN_GETPERMISSIONQUERYPARAMETERTYPE
)

func (i GetPermissionQueryParameterType) String() string {
    return []string{"pull", "triage", "push", "maintain", "admin"}[i]
}
func ParseGetPermissionQueryParameterType(v string) (any, error) {
    result := PULL_GETPERMISSIONQUERYPARAMETERTYPE
    switch v {
        case "pull":
            result = PULL_GETPERMISSIONQUERYPARAMETERTYPE
        case "triage":
            result = TRIAGE_GETPERMISSIONQUERYPARAMETERTYPE
        case "push":
            result = PUSH_GETPERMISSIONQUERYPARAMETERTYPE
        case "maintain":
            result = MAINTAIN_GETPERMISSIONQUERYPARAMETERTYPE
        case "admin":
            result = ADMIN_GETPERMISSIONQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetPermissionQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetPermissionQueryParameterType(values []GetPermissionQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetPermissionQueryParameterType) isMultiValue() bool {
    return false
}

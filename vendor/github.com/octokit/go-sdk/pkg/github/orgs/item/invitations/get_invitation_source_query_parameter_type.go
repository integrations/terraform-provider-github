package invitations
import (
    "errors"
)
type GetInvitation_sourceQueryParameterType int

const (
    ALL_GETINVITATION_SOURCEQUERYPARAMETERTYPE GetInvitation_sourceQueryParameterType = iota
    MEMBER_GETINVITATION_SOURCEQUERYPARAMETERTYPE
    SCIM_GETINVITATION_SOURCEQUERYPARAMETERTYPE
)

func (i GetInvitation_sourceQueryParameterType) String() string {
    return []string{"all", "member", "scim"}[i]
}
func ParseGetInvitation_sourceQueryParameterType(v string) (any, error) {
    result := ALL_GETINVITATION_SOURCEQUERYPARAMETERTYPE
    switch v {
        case "all":
            result = ALL_GETINVITATION_SOURCEQUERYPARAMETERTYPE
        case "member":
            result = MEMBER_GETINVITATION_SOURCEQUERYPARAMETERTYPE
        case "scim":
            result = SCIM_GETINVITATION_SOURCEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetInvitation_sourceQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetInvitation_sourceQueryParameterType(values []GetInvitation_sourceQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetInvitation_sourceQueryParameterType) isMultiValue() bool {
    return false
}

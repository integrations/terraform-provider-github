package issues
import (
    "errors"
)
type GetFilterQueryParameterType int

const (
    ASSIGNED_GETFILTERQUERYPARAMETERTYPE GetFilterQueryParameterType = iota
    CREATED_GETFILTERQUERYPARAMETERTYPE
    MENTIONED_GETFILTERQUERYPARAMETERTYPE
    SUBSCRIBED_GETFILTERQUERYPARAMETERTYPE
    REPOS_GETFILTERQUERYPARAMETERTYPE
    ALL_GETFILTERQUERYPARAMETERTYPE
)

func (i GetFilterQueryParameterType) String() string {
    return []string{"assigned", "created", "mentioned", "subscribed", "repos", "all"}[i]
}
func ParseGetFilterQueryParameterType(v string) (any, error) {
    result := ASSIGNED_GETFILTERQUERYPARAMETERTYPE
    switch v {
        case "assigned":
            result = ASSIGNED_GETFILTERQUERYPARAMETERTYPE
        case "created":
            result = CREATED_GETFILTERQUERYPARAMETERTYPE
        case "mentioned":
            result = MENTIONED_GETFILTERQUERYPARAMETERTYPE
        case "subscribed":
            result = SUBSCRIBED_GETFILTERQUERYPARAMETERTYPE
        case "repos":
            result = REPOS_GETFILTERQUERYPARAMETERTYPE
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

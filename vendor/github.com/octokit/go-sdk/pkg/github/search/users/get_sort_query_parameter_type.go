package users
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    FOLLOWERS_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    REPOSITORIES_GETSORTQUERYPARAMETERTYPE
    JOINED_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"followers", "repositories", "joined"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := FOLLOWERS_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "followers":
            result = FOLLOWERS_GETSORTQUERYPARAMETERTYPE
        case "repositories":
            result = REPOSITORIES_GETSORTQUERYPARAMETERTYPE
        case "joined":
            result = JOINED_GETSORTQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetSortQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetSortQueryParameterType(values []GetSortQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSortQueryParameterType) isMultiValue() bool {
    return false
}

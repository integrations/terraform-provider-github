package forks
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    NEWEST_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    OLDEST_GETSORTQUERYPARAMETERTYPE
    STARGAZERS_GETSORTQUERYPARAMETERTYPE
    WATCHERS_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"newest", "oldest", "stargazers", "watchers"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := NEWEST_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "newest":
            result = NEWEST_GETSORTQUERYPARAMETERTYPE
        case "oldest":
            result = OLDEST_GETSORTQUERYPARAMETERTYPE
        case "stargazers":
            result = STARGAZERS_GETSORTQUERYPARAMETERTYPE
        case "watchers":
            result = WATCHERS_GETSORTQUERYPARAMETERTYPE
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

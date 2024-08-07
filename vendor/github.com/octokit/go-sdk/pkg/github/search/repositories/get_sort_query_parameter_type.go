package repositories
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    STARS_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    FORKS_GETSORTQUERYPARAMETERTYPE
    HELPWANTEDISSUES_GETSORTQUERYPARAMETERTYPE
    UPDATED_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"stars", "forks", "help-wanted-issues", "updated"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := STARS_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "stars":
            result = STARS_GETSORTQUERYPARAMETERTYPE
        case "forks":
            result = FORKS_GETSORTQUERYPARAMETERTYPE
        case "help-wanted-issues":
            result = HELPWANTEDISSUES_GETSORTQUERYPARAMETERTYPE
        case "updated":
            result = UPDATED_GETSORTQUERYPARAMETERTYPE
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

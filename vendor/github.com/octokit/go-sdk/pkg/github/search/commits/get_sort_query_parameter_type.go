package commits
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    AUTHORDATE_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    COMMITTERDATE_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"author-date", "committer-date"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := AUTHORDATE_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "author-date":
            result = AUTHORDATE_GETSORTQUERYPARAMETERTYPE
        case "committer-date":
            result = COMMITTERDATE_GETSORTQUERYPARAMETERTYPE
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

package milestones
import (
    "errors"
)
type GetSortQueryParameterType int

const (
    DUE_ON_GETSORTQUERYPARAMETERTYPE GetSortQueryParameterType = iota
    COMPLETENESS_GETSORTQUERYPARAMETERTYPE
)

func (i GetSortQueryParameterType) String() string {
    return []string{"due_on", "completeness"}[i]
}
func ParseGetSortQueryParameterType(v string) (any, error) {
    result := DUE_ON_GETSORTQUERYPARAMETERTYPE
    switch v {
        case "due_on":
            result = DUE_ON_GETSORTQUERYPARAMETERTYPE
        case "completeness":
            result = COMPLETENESS_GETSORTQUERYPARAMETERTYPE
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

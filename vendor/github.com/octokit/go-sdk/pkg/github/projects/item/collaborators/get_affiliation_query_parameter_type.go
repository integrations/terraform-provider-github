package collaborators
import (
    "errors"
)
type GetAffiliationQueryParameterType int

const (
    OUTSIDE_GETAFFILIATIONQUERYPARAMETERTYPE GetAffiliationQueryParameterType = iota
    DIRECT_GETAFFILIATIONQUERYPARAMETERTYPE
    ALL_GETAFFILIATIONQUERYPARAMETERTYPE
)

func (i GetAffiliationQueryParameterType) String() string {
    return []string{"outside", "direct", "all"}[i]
}
func ParseGetAffiliationQueryParameterType(v string) (any, error) {
    result := OUTSIDE_GETAFFILIATIONQUERYPARAMETERTYPE
    switch v {
        case "outside":
            result = OUTSIDE_GETAFFILIATIONQUERYPARAMETERTYPE
        case "direct":
            result = DIRECT_GETAFFILIATIONQUERYPARAMETERTYPE
        case "all":
            result = ALL_GETAFFILIATIONQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetAffiliationQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetAffiliationQueryParameterType(values []GetAffiliationQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetAffiliationQueryParameterType) isMultiValue() bool {
    return false
}

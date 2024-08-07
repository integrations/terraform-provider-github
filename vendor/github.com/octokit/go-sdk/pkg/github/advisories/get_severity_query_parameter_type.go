package advisories
import (
    "errors"
)
type GetSeverityQueryParameterType int

const (
    UNKNOWN_GETSEVERITYQUERYPARAMETERTYPE GetSeverityQueryParameterType = iota
    LOW_GETSEVERITYQUERYPARAMETERTYPE
    MEDIUM_GETSEVERITYQUERYPARAMETERTYPE
    HIGH_GETSEVERITYQUERYPARAMETERTYPE
    CRITICAL_GETSEVERITYQUERYPARAMETERTYPE
)

func (i GetSeverityQueryParameterType) String() string {
    return []string{"unknown", "low", "medium", "high", "critical"}[i]
}
func ParseGetSeverityQueryParameterType(v string) (any, error) {
    result := UNKNOWN_GETSEVERITYQUERYPARAMETERTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_GETSEVERITYQUERYPARAMETERTYPE
        case "low":
            result = LOW_GETSEVERITYQUERYPARAMETERTYPE
        case "medium":
            result = MEDIUM_GETSEVERITYQUERYPARAMETERTYPE
        case "high":
            result = HIGH_GETSEVERITYQUERYPARAMETERTYPE
        case "critical":
            result = CRITICAL_GETSEVERITYQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetSeverityQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetSeverityQueryParameterType(values []GetSeverityQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSeverityQueryParameterType) isMultiValue() bool {
    return false
}

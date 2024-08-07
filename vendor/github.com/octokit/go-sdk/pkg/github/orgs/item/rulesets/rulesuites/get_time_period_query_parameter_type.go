package rulesuites
import (
    "errors"
)
type GetTime_periodQueryParameterType int

const (
    HOUR_GETTIME_PERIODQUERYPARAMETERTYPE GetTime_periodQueryParameterType = iota
    DAY_GETTIME_PERIODQUERYPARAMETERTYPE
    WEEK_GETTIME_PERIODQUERYPARAMETERTYPE
    MONTH_GETTIME_PERIODQUERYPARAMETERTYPE
)

func (i GetTime_periodQueryParameterType) String() string {
    return []string{"hour", "day", "week", "month"}[i]
}
func ParseGetTime_periodQueryParameterType(v string) (any, error) {
    result := HOUR_GETTIME_PERIODQUERYPARAMETERTYPE
    switch v {
        case "hour":
            result = HOUR_GETTIME_PERIODQUERYPARAMETERTYPE
        case "day":
            result = DAY_GETTIME_PERIODQUERYPARAMETERTYPE
        case "week":
            result = WEEK_GETTIME_PERIODQUERYPARAMETERTYPE
        case "month":
            result = MONTH_GETTIME_PERIODQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetTime_periodQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetTime_periodQueryParameterType(values []GetTime_periodQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetTime_periodQueryParameterType) isMultiValue() bool {
    return false
}

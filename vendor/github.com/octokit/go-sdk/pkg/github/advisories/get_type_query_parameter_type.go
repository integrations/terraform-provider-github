package advisories
import (
    "errors"
)
type GetTypeQueryParameterType int

const (
    REVIEWED_GETTYPEQUERYPARAMETERTYPE GetTypeQueryParameterType = iota
    MALWARE_GETTYPEQUERYPARAMETERTYPE
    UNREVIEWED_GETTYPEQUERYPARAMETERTYPE
)

func (i GetTypeQueryParameterType) String() string {
    return []string{"reviewed", "malware", "unreviewed"}[i]
}
func ParseGetTypeQueryParameterType(v string) (any, error) {
    result := REVIEWED_GETTYPEQUERYPARAMETERTYPE
    switch v {
        case "reviewed":
            result = REVIEWED_GETTYPEQUERYPARAMETERTYPE
        case "malware":
            result = MALWARE_GETTYPEQUERYPARAMETERTYPE
        case "unreviewed":
            result = UNREVIEWED_GETTYPEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetTypeQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetTypeQueryParameterType(values []GetTypeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetTypeQueryParameterType) isMultiValue() bool {
    return false
}

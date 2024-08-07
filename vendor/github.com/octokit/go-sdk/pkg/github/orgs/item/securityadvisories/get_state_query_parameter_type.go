package securityadvisories
import (
    "errors"
)
type GetStateQueryParameterType int

const (
    TRIAGE_GETSTATEQUERYPARAMETERTYPE GetStateQueryParameterType = iota
    DRAFT_GETSTATEQUERYPARAMETERTYPE
    PUBLISHED_GETSTATEQUERYPARAMETERTYPE
    CLOSED_GETSTATEQUERYPARAMETERTYPE
)

func (i GetStateQueryParameterType) String() string {
    return []string{"triage", "draft", "published", "closed"}[i]
}
func ParseGetStateQueryParameterType(v string) (any, error) {
    result := TRIAGE_GETSTATEQUERYPARAMETERTYPE
    switch v {
        case "triage":
            result = TRIAGE_GETSTATEQUERYPARAMETERTYPE
        case "draft":
            result = DRAFT_GETSTATEQUERYPARAMETERTYPE
        case "published":
            result = PUBLISHED_GETSTATEQUERYPARAMETERTYPE
        case "closed":
            result = CLOSED_GETSTATEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetStateQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetStateQueryParameterType(values []GetStateQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetStateQueryParameterType) isMultiValue() bool {
    return false
}

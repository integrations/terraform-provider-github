package reactions
import (
    "errors"
)
type GetContentQueryParameterType int

const (
    PLUS_1_GETCONTENTQUERYPARAMETERTYPE GetContentQueryParameterType = iota
    MINUS_1_GETCONTENTQUERYPARAMETERTYPE
    LAUGH_GETCONTENTQUERYPARAMETERTYPE
    CONFUSED_GETCONTENTQUERYPARAMETERTYPE
    HEART_GETCONTENTQUERYPARAMETERTYPE
    HOORAY_GETCONTENTQUERYPARAMETERTYPE
    ROCKET_GETCONTENTQUERYPARAMETERTYPE
    EYES_GETCONTENTQUERYPARAMETERTYPE
)

func (i GetContentQueryParameterType) String() string {
    return []string{"+1", "-1", "laugh", "confused", "heart", "hooray", "rocket", "eyes"}[i]
}
func ParseGetContentQueryParameterType(v string) (any, error) {
    result := PLUS_1_GETCONTENTQUERYPARAMETERTYPE
    switch v {
        case "+1":
            result = PLUS_1_GETCONTENTQUERYPARAMETERTYPE
        case "-1":
            result = MINUS_1_GETCONTENTQUERYPARAMETERTYPE
        case "laugh":
            result = LAUGH_GETCONTENTQUERYPARAMETERTYPE
        case "confused":
            result = CONFUSED_GETCONTENTQUERYPARAMETERTYPE
        case "heart":
            result = HEART_GETCONTENTQUERYPARAMETERTYPE
        case "hooray":
            result = HOORAY_GETCONTENTQUERYPARAMETERTYPE
        case "rocket":
            result = ROCKET_GETCONTENTQUERYPARAMETERTYPE
        case "eyes":
            result = EYES_GETCONTENTQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetContentQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetContentQueryParameterType(values []GetContentQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetContentQueryParameterType) isMultiValue() bool {
    return false
}

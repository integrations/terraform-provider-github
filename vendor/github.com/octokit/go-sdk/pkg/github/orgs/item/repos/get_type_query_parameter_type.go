package repos
import (
    "errors"
)
type GetTypeQueryParameterType int

const (
    ALL_GETTYPEQUERYPARAMETERTYPE GetTypeQueryParameterType = iota
    PUBLIC_GETTYPEQUERYPARAMETERTYPE
    PRIVATE_GETTYPEQUERYPARAMETERTYPE
    FORKS_GETTYPEQUERYPARAMETERTYPE
    SOURCES_GETTYPEQUERYPARAMETERTYPE
    MEMBER_GETTYPEQUERYPARAMETERTYPE
)

func (i GetTypeQueryParameterType) String() string {
    return []string{"all", "public", "private", "forks", "sources", "member"}[i]
}
func ParseGetTypeQueryParameterType(v string) (any, error) {
    result := ALL_GETTYPEQUERYPARAMETERTYPE
    switch v {
        case "all":
            result = ALL_GETTYPEQUERYPARAMETERTYPE
        case "public":
            result = PUBLIC_GETTYPEQUERYPARAMETERTYPE
        case "private":
            result = PRIVATE_GETTYPEQUERYPARAMETERTYPE
        case "forks":
            result = FORKS_GETTYPEQUERYPARAMETERTYPE
        case "sources":
            result = SOURCES_GETTYPEQUERYPARAMETERTYPE
        case "member":
            result = MEMBER_GETTYPEQUERYPARAMETERTYPE
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

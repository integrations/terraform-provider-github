package hovercard
import (
    "errors"
)
type GetSubject_typeQueryParameterType int

const (
    ORGANIZATION_GETSUBJECT_TYPEQUERYPARAMETERTYPE GetSubject_typeQueryParameterType = iota
    REPOSITORY_GETSUBJECT_TYPEQUERYPARAMETERTYPE
    ISSUE_GETSUBJECT_TYPEQUERYPARAMETERTYPE
    PULL_REQUEST_GETSUBJECT_TYPEQUERYPARAMETERTYPE
)

func (i GetSubject_typeQueryParameterType) String() string {
    return []string{"organization", "repository", "issue", "pull_request"}[i]
}
func ParseGetSubject_typeQueryParameterType(v string) (any, error) {
    result := ORGANIZATION_GETSUBJECT_TYPEQUERYPARAMETERTYPE
    switch v {
        case "organization":
            result = ORGANIZATION_GETSUBJECT_TYPEQUERYPARAMETERTYPE
        case "repository":
            result = REPOSITORY_GETSUBJECT_TYPEQUERYPARAMETERTYPE
        case "issue":
            result = ISSUE_GETSUBJECT_TYPEQUERYPARAMETERTYPE
        case "pull_request":
            result = PULL_REQUEST_GETSUBJECT_TYPEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetSubject_typeQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetSubject_typeQueryParameterType(values []GetSubject_typeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetSubject_typeQueryParameterType) isMultiValue() bool {
    return false
}

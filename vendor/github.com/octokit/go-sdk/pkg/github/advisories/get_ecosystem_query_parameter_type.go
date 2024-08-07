package advisories
import (
    "errors"
)
type GetEcosystemQueryParameterType int

const (
    ACTIONS_GETECOSYSTEMQUERYPARAMETERTYPE GetEcosystemQueryParameterType = iota
    COMPOSER_GETECOSYSTEMQUERYPARAMETERTYPE
    ERLANG_GETECOSYSTEMQUERYPARAMETERTYPE
    GO_GETECOSYSTEMQUERYPARAMETERTYPE
    MAVEN_GETECOSYSTEMQUERYPARAMETERTYPE
    NPM_GETECOSYSTEMQUERYPARAMETERTYPE
    NUGET_GETECOSYSTEMQUERYPARAMETERTYPE
    OTHER_GETECOSYSTEMQUERYPARAMETERTYPE
    PIP_GETECOSYSTEMQUERYPARAMETERTYPE
    PUB_GETECOSYSTEMQUERYPARAMETERTYPE
    RUBYGEMS_GETECOSYSTEMQUERYPARAMETERTYPE
    RUST_GETECOSYSTEMQUERYPARAMETERTYPE
)

func (i GetEcosystemQueryParameterType) String() string {
    return []string{"actions", "composer", "erlang", "go", "maven", "npm", "nuget", "other", "pip", "pub", "rubygems", "rust"}[i]
}
func ParseGetEcosystemQueryParameterType(v string) (any, error) {
    result := ACTIONS_GETECOSYSTEMQUERYPARAMETERTYPE
    switch v {
        case "actions":
            result = ACTIONS_GETECOSYSTEMQUERYPARAMETERTYPE
        case "composer":
            result = COMPOSER_GETECOSYSTEMQUERYPARAMETERTYPE
        case "erlang":
            result = ERLANG_GETECOSYSTEMQUERYPARAMETERTYPE
        case "go":
            result = GO_GETECOSYSTEMQUERYPARAMETERTYPE
        case "maven":
            result = MAVEN_GETECOSYSTEMQUERYPARAMETERTYPE
        case "npm":
            result = NPM_GETECOSYSTEMQUERYPARAMETERTYPE
        case "nuget":
            result = NUGET_GETECOSYSTEMQUERYPARAMETERTYPE
        case "other":
            result = OTHER_GETECOSYSTEMQUERYPARAMETERTYPE
        case "pip":
            result = PIP_GETECOSYSTEMQUERYPARAMETERTYPE
        case "pub":
            result = PUB_GETECOSYSTEMQUERYPARAMETERTYPE
        case "rubygems":
            result = RUBYGEMS_GETECOSYSTEMQUERYPARAMETERTYPE
        case "rust":
            result = RUST_GETECOSYSTEMQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetEcosystemQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetEcosystemQueryParameterType(values []GetEcosystemQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetEcosystemQueryParameterType) isMultiValue() bool {
    return false
}

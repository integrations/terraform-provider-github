package packages
import (
    "errors"
)
type GetPackage_typeQueryParameterType int

const (
    NPM_GETPACKAGE_TYPEQUERYPARAMETERTYPE GetPackage_typeQueryParameterType = iota
    MAVEN_GETPACKAGE_TYPEQUERYPARAMETERTYPE
    RUBYGEMS_GETPACKAGE_TYPEQUERYPARAMETERTYPE
    DOCKER_GETPACKAGE_TYPEQUERYPARAMETERTYPE
    NUGET_GETPACKAGE_TYPEQUERYPARAMETERTYPE
    CONTAINER_GETPACKAGE_TYPEQUERYPARAMETERTYPE
)

func (i GetPackage_typeQueryParameterType) String() string {
    return []string{"npm", "maven", "rubygems", "docker", "nuget", "container"}[i]
}
func ParseGetPackage_typeQueryParameterType(v string) (any, error) {
    result := NPM_GETPACKAGE_TYPEQUERYPARAMETERTYPE
    switch v {
        case "npm":
            result = NPM_GETPACKAGE_TYPEQUERYPARAMETERTYPE
        case "maven":
            result = MAVEN_GETPACKAGE_TYPEQUERYPARAMETERTYPE
        case "rubygems":
            result = RUBYGEMS_GETPACKAGE_TYPEQUERYPARAMETERTYPE
        case "docker":
            result = DOCKER_GETPACKAGE_TYPEQUERYPARAMETERTYPE
        case "nuget":
            result = NUGET_GETPACKAGE_TYPEQUERYPARAMETERTYPE
        case "container":
            result = CONTAINER_GETPACKAGE_TYPEQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown GetPackage_typeQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializeGetPackage_typeQueryParameterType(values []GetPackage_typeQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetPackage_typeQueryParameterType) isMultiValue() bool {
    return false
}

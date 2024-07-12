package models
import (
    "errors"
)
type Package_package_type int

const (
    NPM_PACKAGE_PACKAGE_TYPE Package_package_type = iota
    MAVEN_PACKAGE_PACKAGE_TYPE
    RUBYGEMS_PACKAGE_PACKAGE_TYPE
    DOCKER_PACKAGE_PACKAGE_TYPE
    NUGET_PACKAGE_PACKAGE_TYPE
    CONTAINER_PACKAGE_PACKAGE_TYPE
)

func (i Package_package_type) String() string {
    return []string{"npm", "maven", "rubygems", "docker", "nuget", "container"}[i]
}
func ParsePackage_package_type(v string) (any, error) {
    result := NPM_PACKAGE_PACKAGE_TYPE
    switch v {
        case "npm":
            result = NPM_PACKAGE_PACKAGE_TYPE
        case "maven":
            result = MAVEN_PACKAGE_PACKAGE_TYPE
        case "rubygems":
            result = RUBYGEMS_PACKAGE_PACKAGE_TYPE
        case "docker":
            result = DOCKER_PACKAGE_PACKAGE_TYPE
        case "nuget":
            result = NUGET_PACKAGE_PACKAGE_TYPE
        case "container":
            result = CONTAINER_PACKAGE_PACKAGE_TYPE
        default:
            return 0, errors.New("Unknown Package_package_type value: " + v)
    }
    return &result, nil
}
func SerializePackage_package_type(values []Package_package_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Package_package_type) isMultiValue() bool {
    return false
}

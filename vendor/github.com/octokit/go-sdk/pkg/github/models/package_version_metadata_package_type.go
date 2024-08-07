package models
import (
    "errors"
)
type PackageVersion_metadata_package_type int

const (
    NPM_PACKAGEVERSION_METADATA_PACKAGE_TYPE PackageVersion_metadata_package_type = iota
    MAVEN_PACKAGEVERSION_METADATA_PACKAGE_TYPE
    RUBYGEMS_PACKAGEVERSION_METADATA_PACKAGE_TYPE
    DOCKER_PACKAGEVERSION_METADATA_PACKAGE_TYPE
    NUGET_PACKAGEVERSION_METADATA_PACKAGE_TYPE
    CONTAINER_PACKAGEVERSION_METADATA_PACKAGE_TYPE
)

func (i PackageVersion_metadata_package_type) String() string {
    return []string{"npm", "maven", "rubygems", "docker", "nuget", "container"}[i]
}
func ParsePackageVersion_metadata_package_type(v string) (any, error) {
    result := NPM_PACKAGEVERSION_METADATA_PACKAGE_TYPE
    switch v {
        case "npm":
            result = NPM_PACKAGEVERSION_METADATA_PACKAGE_TYPE
        case "maven":
            result = MAVEN_PACKAGEVERSION_METADATA_PACKAGE_TYPE
        case "rubygems":
            result = RUBYGEMS_PACKAGEVERSION_METADATA_PACKAGE_TYPE
        case "docker":
            result = DOCKER_PACKAGEVERSION_METADATA_PACKAGE_TYPE
        case "nuget":
            result = NUGET_PACKAGEVERSION_METADATA_PACKAGE_TYPE
        case "container":
            result = CONTAINER_PACKAGEVERSION_METADATA_PACKAGE_TYPE
        default:
            return 0, errors.New("Unknown PackageVersion_metadata_package_type value: " + v)
    }
    return &result, nil
}
func SerializePackageVersion_metadata_package_type(values []PackageVersion_metadata_package_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PackageVersion_metadata_package_type) isMultiValue() bool {
    return false
}

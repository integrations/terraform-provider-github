package models
import (
    "errors"
)
// The package's language or package management ecosystem.
type SecurityAdvisoryEcosystems int

const (
    RUBYGEMS_SECURITYADVISORYECOSYSTEMS SecurityAdvisoryEcosystems = iota
    NPM_SECURITYADVISORYECOSYSTEMS
    PIP_SECURITYADVISORYECOSYSTEMS
    MAVEN_SECURITYADVISORYECOSYSTEMS
    NUGET_SECURITYADVISORYECOSYSTEMS
    COMPOSER_SECURITYADVISORYECOSYSTEMS
    GO_SECURITYADVISORYECOSYSTEMS
    RUST_SECURITYADVISORYECOSYSTEMS
    ERLANG_SECURITYADVISORYECOSYSTEMS
    ACTIONS_SECURITYADVISORYECOSYSTEMS
    PUB_SECURITYADVISORYECOSYSTEMS
    OTHER_SECURITYADVISORYECOSYSTEMS
    SWIFT_SECURITYADVISORYECOSYSTEMS
)

func (i SecurityAdvisoryEcosystems) String() string {
    return []string{"rubygems", "npm", "pip", "maven", "nuget", "composer", "go", "rust", "erlang", "actions", "pub", "other", "swift"}[i]
}
func ParseSecurityAdvisoryEcosystems(v string) (any, error) {
    result := RUBYGEMS_SECURITYADVISORYECOSYSTEMS
    switch v {
        case "rubygems":
            result = RUBYGEMS_SECURITYADVISORYECOSYSTEMS
        case "npm":
            result = NPM_SECURITYADVISORYECOSYSTEMS
        case "pip":
            result = PIP_SECURITYADVISORYECOSYSTEMS
        case "maven":
            result = MAVEN_SECURITYADVISORYECOSYSTEMS
        case "nuget":
            result = NUGET_SECURITYADVISORYECOSYSTEMS
        case "composer":
            result = COMPOSER_SECURITYADVISORYECOSYSTEMS
        case "go":
            result = GO_SECURITYADVISORYECOSYSTEMS
        case "rust":
            result = RUST_SECURITYADVISORYECOSYSTEMS
        case "erlang":
            result = ERLANG_SECURITYADVISORYECOSYSTEMS
        case "actions":
            result = ACTIONS_SECURITYADVISORYECOSYSTEMS
        case "pub":
            result = PUB_SECURITYADVISORYECOSYSTEMS
        case "other":
            result = OTHER_SECURITYADVISORYECOSYSTEMS
        case "swift":
            result = SWIFT_SECURITYADVISORYECOSYSTEMS
        default:
            return 0, errors.New("Unknown SecurityAdvisoryEcosystems value: " + v)
    }
    return &result, nil
}
func SerializeSecurityAdvisoryEcosystems(values []SecurityAdvisoryEcosystems) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityAdvisoryEcosystems) isMultiValue() bool {
    return false
}

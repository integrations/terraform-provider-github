package models
import (
    "errors"
)
type RepositoryRuleRequiredDeployments_type int

const (
    REQUIRED_DEPLOYMENTS_REPOSITORYRULEREQUIREDDEPLOYMENTS_TYPE RepositoryRuleRequiredDeployments_type = iota
)

func (i RepositoryRuleRequiredDeployments_type) String() string {
    return []string{"required_deployments"}[i]
}
func ParseRepositoryRuleRequiredDeployments_type(v string) (any, error) {
    result := REQUIRED_DEPLOYMENTS_REPOSITORYRULEREQUIREDDEPLOYMENTS_TYPE
    switch v {
        case "required_deployments":
            result = REQUIRED_DEPLOYMENTS_REPOSITORYRULEREQUIREDDEPLOYMENTS_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleRequiredDeployments_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleRequiredDeployments_type(values []RepositoryRuleRequiredDeployments_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleRequiredDeployments_type) isMultiValue() bool {
    return false
}

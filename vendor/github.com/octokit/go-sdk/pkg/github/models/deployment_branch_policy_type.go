package models
import (
    "errors"
)
// Whether this rule targets a branch or tag.
type DeploymentBranchPolicy_type int

const (
    BRANCH_DEPLOYMENTBRANCHPOLICY_TYPE DeploymentBranchPolicy_type = iota
    TAG_DEPLOYMENTBRANCHPOLICY_TYPE
)

func (i DeploymentBranchPolicy_type) String() string {
    return []string{"branch", "tag"}[i]
}
func ParseDeploymentBranchPolicy_type(v string) (any, error) {
    result := BRANCH_DEPLOYMENTBRANCHPOLICY_TYPE
    switch v {
        case "branch":
            result = BRANCH_DEPLOYMENTBRANCHPOLICY_TYPE
        case "tag":
            result = TAG_DEPLOYMENTBRANCHPOLICY_TYPE
        default:
            return 0, errors.New("Unknown DeploymentBranchPolicy_type value: " + v)
    }
    return &result, nil
}
func SerializeDeploymentBranchPolicy_type(values []DeploymentBranchPolicy_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeploymentBranchPolicy_type) isMultiValue() bool {
    return false
}

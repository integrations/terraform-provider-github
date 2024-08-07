package models
import (
    "errors"
)
// Whether this rule targets a branch or tag
type DeploymentBranchPolicyNamePatternWithType_type int

const (
    BRANCH_DEPLOYMENTBRANCHPOLICYNAMEPATTERNWITHTYPE_TYPE DeploymentBranchPolicyNamePatternWithType_type = iota
    TAG_DEPLOYMENTBRANCHPOLICYNAMEPATTERNWITHTYPE_TYPE
)

func (i DeploymentBranchPolicyNamePatternWithType_type) String() string {
    return []string{"branch", "tag"}[i]
}
func ParseDeploymentBranchPolicyNamePatternWithType_type(v string) (any, error) {
    result := BRANCH_DEPLOYMENTBRANCHPOLICYNAMEPATTERNWITHTYPE_TYPE
    switch v {
        case "branch":
            result = BRANCH_DEPLOYMENTBRANCHPOLICYNAMEPATTERNWITHTYPE_TYPE
        case "tag":
            result = TAG_DEPLOYMENTBRANCHPOLICYNAMEPATTERNWITHTYPE_TYPE
        default:
            return 0, errors.New("Unknown DeploymentBranchPolicyNamePatternWithType_type value: " + v)
    }
    return &result, nil
}
func SerializeDeploymentBranchPolicyNamePatternWithType_type(values []DeploymentBranchPolicyNamePatternWithType_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeploymentBranchPolicyNamePatternWithType_type) isMultiValue() bool {
    return false
}

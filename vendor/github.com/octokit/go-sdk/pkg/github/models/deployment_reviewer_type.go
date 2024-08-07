package models
import (
    "errors"
)
// The type of reviewer.
type DeploymentReviewerType int

const (
    USER_DEPLOYMENTREVIEWERTYPE DeploymentReviewerType = iota
    TEAM_DEPLOYMENTREVIEWERTYPE
)

func (i DeploymentReviewerType) String() string {
    return []string{"User", "Team"}[i]
}
func ParseDeploymentReviewerType(v string) (any, error) {
    result := USER_DEPLOYMENTREVIEWERTYPE
    switch v {
        case "User":
            result = USER_DEPLOYMENTREVIEWERTYPE
        case "Team":
            result = TEAM_DEPLOYMENTREVIEWERTYPE
        default:
            return 0, errors.New("Unknown DeploymentReviewerType value: " + v)
    }
    return &result, nil
}
func SerializeDeploymentReviewerType(values []DeploymentReviewerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeploymentReviewerType) isMultiValue() bool {
    return false
}

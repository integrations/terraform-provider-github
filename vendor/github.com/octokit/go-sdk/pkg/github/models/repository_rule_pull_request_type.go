package models
import (
    "errors"
)
type RepositoryRulePullRequest_type int

const (
    PULL_REQUEST_REPOSITORYRULEPULLREQUEST_TYPE RepositoryRulePullRequest_type = iota
)

func (i RepositoryRulePullRequest_type) String() string {
    return []string{"pull_request"}[i]
}
func ParseRepositoryRulePullRequest_type(v string) (any, error) {
    result := PULL_REQUEST_REPOSITORYRULEPULLREQUEST_TYPE
    switch v {
        case "pull_request":
            result = PULL_REQUEST_REPOSITORYRULEPULLREQUEST_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRulePullRequest_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRulePullRequest_type(values []RepositoryRulePullRequest_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRulePullRequest_type) isMultiValue() bool {
    return false
}

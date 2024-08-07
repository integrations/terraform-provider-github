package models
import (
    "errors"
)
type RepositoryRuleCommitMessagePattern_type int

const (
    COMMIT_MESSAGE_PATTERN_REPOSITORYRULECOMMITMESSAGEPATTERN_TYPE RepositoryRuleCommitMessagePattern_type = iota
)

func (i RepositoryRuleCommitMessagePattern_type) String() string {
    return []string{"commit_message_pattern"}[i]
}
func ParseRepositoryRuleCommitMessagePattern_type(v string) (any, error) {
    result := COMMIT_MESSAGE_PATTERN_REPOSITORYRULECOMMITMESSAGEPATTERN_TYPE
    switch v {
        case "commit_message_pattern":
            result = COMMIT_MESSAGE_PATTERN_REPOSITORYRULECOMMITMESSAGEPATTERN_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleCommitMessagePattern_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleCommitMessagePattern_type(values []RepositoryRuleCommitMessagePattern_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleCommitMessagePattern_type) isMultiValue() bool {
    return false
}

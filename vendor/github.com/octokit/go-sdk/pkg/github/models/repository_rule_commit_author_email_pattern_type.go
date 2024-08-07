package models
import (
    "errors"
)
type RepositoryRuleCommitAuthorEmailPattern_type int

const (
    COMMIT_AUTHOR_EMAIL_PATTERN_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_TYPE RepositoryRuleCommitAuthorEmailPattern_type = iota
)

func (i RepositoryRuleCommitAuthorEmailPattern_type) String() string {
    return []string{"commit_author_email_pattern"}[i]
}
func ParseRepositoryRuleCommitAuthorEmailPattern_type(v string) (any, error) {
    result := COMMIT_AUTHOR_EMAIL_PATTERN_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_TYPE
    switch v {
        case "commit_author_email_pattern":
            result = COMMIT_AUTHOR_EMAIL_PATTERN_REPOSITORYRULECOMMITAUTHOREMAILPATTERN_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleCommitAuthorEmailPattern_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleCommitAuthorEmailPattern_type(values []RepositoryRuleCommitAuthorEmailPattern_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleCommitAuthorEmailPattern_type) isMultiValue() bool {
    return false
}

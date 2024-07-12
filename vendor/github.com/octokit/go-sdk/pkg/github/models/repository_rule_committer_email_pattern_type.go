package models
import (
    "errors"
)
type RepositoryRuleCommitterEmailPattern_type int

const (
    COMMITTER_EMAIL_PATTERN_REPOSITORYRULECOMMITTEREMAILPATTERN_TYPE RepositoryRuleCommitterEmailPattern_type = iota
)

func (i RepositoryRuleCommitterEmailPattern_type) String() string {
    return []string{"committer_email_pattern"}[i]
}
func ParseRepositoryRuleCommitterEmailPattern_type(v string) (any, error) {
    result := COMMITTER_EMAIL_PATTERN_REPOSITORYRULECOMMITTEREMAILPATTERN_TYPE
    switch v {
        case "committer_email_pattern":
            result = COMMITTER_EMAIL_PATTERN_REPOSITORYRULECOMMITTEREMAILPATTERN_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleCommitterEmailPattern_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleCommitterEmailPattern_type(values []RepositoryRuleCommitterEmailPattern_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleCommitterEmailPattern_type) isMultiValue() bool {
    return false
}

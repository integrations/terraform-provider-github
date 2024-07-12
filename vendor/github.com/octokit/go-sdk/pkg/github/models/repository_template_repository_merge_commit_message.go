package models
import (
    "errors"
)
// The default value for a merge commit message.- `PR_TITLE` - default to the pull request's title.- `PR_BODY` - default to the pull request's body.- `BLANK` - default to a blank commit message.
type Repository_template_repository_merge_commit_message int

const (
    PR_BODY_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_MESSAGE Repository_template_repository_merge_commit_message = iota
    PR_TITLE_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_MESSAGE
    BLANK_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_MESSAGE
)

func (i Repository_template_repository_merge_commit_message) String() string {
    return []string{"PR_BODY", "PR_TITLE", "BLANK"}[i]
}
func ParseRepository_template_repository_merge_commit_message(v string) (any, error) {
    result := PR_BODY_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_MESSAGE
    switch v {
        case "PR_BODY":
            result = PR_BODY_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_MESSAGE
        case "PR_TITLE":
            result = PR_TITLE_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_MESSAGE
        case "BLANK":
            result = BLANK_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_MESSAGE
        default:
            return 0, errors.New("Unknown Repository_template_repository_merge_commit_message value: " + v)
    }
    return &result, nil
}
func SerializeRepository_template_repository_merge_commit_message(values []Repository_template_repository_merge_commit_message) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Repository_template_repository_merge_commit_message) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The default value for a merge commit title.- `PR_TITLE` - default to the pull request's title.- `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
type Repository_template_repository_merge_commit_title int

const (
    PR_TITLE_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_TITLE Repository_template_repository_merge_commit_title = iota
    MERGE_MESSAGE_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_TITLE
)

func (i Repository_template_repository_merge_commit_title) String() string {
    return []string{"PR_TITLE", "MERGE_MESSAGE"}[i]
}
func ParseRepository_template_repository_merge_commit_title(v string) (any, error) {
    result := PR_TITLE_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_TITLE
    switch v {
        case "PR_TITLE":
            result = PR_TITLE_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_TITLE
        case "MERGE_MESSAGE":
            result = MERGE_MESSAGE_REPOSITORY_TEMPLATE_REPOSITORY_MERGE_COMMIT_TITLE
        default:
            return 0, errors.New("Unknown Repository_template_repository_merge_commit_title value: " + v)
    }
    return &result, nil
}
func SerializeRepository_template_repository_merge_commit_title(values []Repository_template_repository_merge_commit_title) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Repository_template_repository_merge_commit_title) isMultiValue() bool {
    return false
}

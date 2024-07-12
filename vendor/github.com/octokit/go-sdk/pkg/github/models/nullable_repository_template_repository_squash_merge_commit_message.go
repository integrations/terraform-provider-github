package models
import (
    "errors"
)
// The default value for a squash merge commit message:- `PR_BODY` - default to the pull request's body.- `COMMIT_MESSAGES` - default to the branch's commit messages.- `BLANK` - default to a blank commit message.
type NullableRepository_template_repository_squash_merge_commit_message int

const (
    PR_BODY_NULLABLEREPOSITORY_TEMPLATE_REPOSITORY_SQUASH_MERGE_COMMIT_MESSAGE NullableRepository_template_repository_squash_merge_commit_message = iota
    COMMIT_MESSAGES_NULLABLEREPOSITORY_TEMPLATE_REPOSITORY_SQUASH_MERGE_COMMIT_MESSAGE
    BLANK_NULLABLEREPOSITORY_TEMPLATE_REPOSITORY_SQUASH_MERGE_COMMIT_MESSAGE
)

func (i NullableRepository_template_repository_squash_merge_commit_message) String() string {
    return []string{"PR_BODY", "COMMIT_MESSAGES", "BLANK"}[i]
}
func ParseNullableRepository_template_repository_squash_merge_commit_message(v string) (any, error) {
    result := PR_BODY_NULLABLEREPOSITORY_TEMPLATE_REPOSITORY_SQUASH_MERGE_COMMIT_MESSAGE
    switch v {
        case "PR_BODY":
            result = PR_BODY_NULLABLEREPOSITORY_TEMPLATE_REPOSITORY_SQUASH_MERGE_COMMIT_MESSAGE
        case "COMMIT_MESSAGES":
            result = COMMIT_MESSAGES_NULLABLEREPOSITORY_TEMPLATE_REPOSITORY_SQUASH_MERGE_COMMIT_MESSAGE
        case "BLANK":
            result = BLANK_NULLABLEREPOSITORY_TEMPLATE_REPOSITORY_SQUASH_MERGE_COMMIT_MESSAGE
        default:
            return 0, errors.New("Unknown NullableRepository_template_repository_squash_merge_commit_message value: " + v)
    }
    return &result, nil
}
func SerializeNullableRepository_template_repository_squash_merge_commit_message(values []NullableRepository_template_repository_squash_merge_commit_message) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i NullableRepository_template_repository_squash_merge_commit_message) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The default value for a squash merge commit title:- `PR_TITLE` - default to the pull request's title.- `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
type PullRequestWebhook_squash_merge_commit_title int

const (
    PR_TITLE_PULLREQUESTWEBHOOK_SQUASH_MERGE_COMMIT_TITLE PullRequestWebhook_squash_merge_commit_title = iota
    COMMIT_OR_PR_TITLE_PULLREQUESTWEBHOOK_SQUASH_MERGE_COMMIT_TITLE
)

func (i PullRequestWebhook_squash_merge_commit_title) String() string {
    return []string{"PR_TITLE", "COMMIT_OR_PR_TITLE"}[i]
}
func ParsePullRequestWebhook_squash_merge_commit_title(v string) (any, error) {
    result := PR_TITLE_PULLREQUESTWEBHOOK_SQUASH_MERGE_COMMIT_TITLE
    switch v {
        case "PR_TITLE":
            result = PR_TITLE_PULLREQUESTWEBHOOK_SQUASH_MERGE_COMMIT_TITLE
        case "COMMIT_OR_PR_TITLE":
            result = COMMIT_OR_PR_TITLE_PULLREQUESTWEBHOOK_SQUASH_MERGE_COMMIT_TITLE
        default:
            return 0, errors.New("Unknown PullRequestWebhook_squash_merge_commit_title value: " + v)
    }
    return &result, nil
}
func SerializePullRequestWebhook_squash_merge_commit_title(values []PullRequestWebhook_squash_merge_commit_title) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PullRequestWebhook_squash_merge_commit_title) isMultiValue() bool {
    return false
}

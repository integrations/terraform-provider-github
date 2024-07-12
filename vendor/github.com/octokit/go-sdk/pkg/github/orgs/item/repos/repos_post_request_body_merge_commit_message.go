package repos
import (
    "errors"
)
// The default value for a merge commit message.- `PR_TITLE` - default to the pull request's title.- `PR_BODY` - default to the pull request's body.- `BLANK` - default to a blank commit message.
type ReposPostRequestBody_merge_commit_message int

const (
    PR_BODY_REPOSPOSTREQUESTBODY_MERGE_COMMIT_MESSAGE ReposPostRequestBody_merge_commit_message = iota
    PR_TITLE_REPOSPOSTREQUESTBODY_MERGE_COMMIT_MESSAGE
    BLANK_REPOSPOSTREQUESTBODY_MERGE_COMMIT_MESSAGE
)

func (i ReposPostRequestBody_merge_commit_message) String() string {
    return []string{"PR_BODY", "PR_TITLE", "BLANK"}[i]
}
func ParseReposPostRequestBody_merge_commit_message(v string) (any, error) {
    result := PR_BODY_REPOSPOSTREQUESTBODY_MERGE_COMMIT_MESSAGE
    switch v {
        case "PR_BODY":
            result = PR_BODY_REPOSPOSTREQUESTBODY_MERGE_COMMIT_MESSAGE
        case "PR_TITLE":
            result = PR_TITLE_REPOSPOSTREQUESTBODY_MERGE_COMMIT_MESSAGE
        case "BLANK":
            result = BLANK_REPOSPOSTREQUESTBODY_MERGE_COMMIT_MESSAGE
        default:
            return 0, errors.New("Unknown ReposPostRequestBody_merge_commit_message value: " + v)
    }
    return &result, nil
}
func SerializeReposPostRequestBody_merge_commit_message(values []ReposPostRequestBody_merge_commit_message) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReposPostRequestBody_merge_commit_message) isMultiValue() bool {
    return false
}

package comments
import (
    "errors"
)
// In a split diff view, the side of the diff that the pull request's changes appear on. Can be `LEFT` or `RIGHT`. Use `LEFT` for deletions that appear in red. Use `RIGHT` for additions that appear in green or unchanged lines that appear in white and are shown for context. For a multi-line comment, side represents whether the last line of the comment range is a deletion or addition. For more information, see "[Diff view options](https://docs.github.com/articles/about-comparing-branches-in-pull-requests#diff-view-options)" in the GitHub Help documentation.
type CommentsPostRequestBody_side int

const (
    LEFT_COMMENTSPOSTREQUESTBODY_SIDE CommentsPostRequestBody_side = iota
    RIGHT_COMMENTSPOSTREQUESTBODY_SIDE
)

func (i CommentsPostRequestBody_side) String() string {
    return []string{"LEFT", "RIGHT"}[i]
}
func ParseCommentsPostRequestBody_side(v string) (any, error) {
    result := LEFT_COMMENTSPOSTREQUESTBODY_SIDE
    switch v {
        case "LEFT":
            result = LEFT_COMMENTSPOSTREQUESTBODY_SIDE
        case "RIGHT":
            result = RIGHT_COMMENTSPOSTREQUESTBODY_SIDE
        default:
            return 0, errors.New("Unknown CommentsPostRequestBody_side value: " + v)
    }
    return &result, nil
}
func SerializeCommentsPostRequestBody_side(values []CommentsPostRequestBody_side) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CommentsPostRequestBody_side) isMultiValue() bool {
    return false
}

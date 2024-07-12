package comments
import (
    "errors"
)
// **Required when using multi-line comments unless using `in_reply_to`**. The `start_side` is the starting side of the diff that the comment applies to. Can be `LEFT` or `RIGHT`. To learn more about multi-line comments, see "[Commenting on a pull request](https://docs.github.com/articles/commenting-on-a-pull-request#adding-line-comments-to-a-pull-request)" in the GitHub Help documentation. See `side` in this table for additional context.
type CommentsPostRequestBody_start_side int

const (
    LEFT_COMMENTSPOSTREQUESTBODY_START_SIDE CommentsPostRequestBody_start_side = iota
    RIGHT_COMMENTSPOSTREQUESTBODY_START_SIDE
    SIDE_COMMENTSPOSTREQUESTBODY_START_SIDE
)

func (i CommentsPostRequestBody_start_side) String() string {
    return []string{"LEFT", "RIGHT", "side"}[i]
}
func ParseCommentsPostRequestBody_start_side(v string) (any, error) {
    result := LEFT_COMMENTSPOSTREQUESTBODY_START_SIDE
    switch v {
        case "LEFT":
            result = LEFT_COMMENTSPOSTREQUESTBODY_START_SIDE
        case "RIGHT":
            result = RIGHT_COMMENTSPOSTREQUESTBODY_START_SIDE
        case "side":
            result = SIDE_COMMENTSPOSTREQUESTBODY_START_SIDE
        default:
            return 0, errors.New("Unknown CommentsPostRequestBody_start_side value: " + v)
    }
    return &result, nil
}
func SerializeCommentsPostRequestBody_start_side(values []CommentsPostRequestBody_start_side) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CommentsPostRequestBody_start_side) isMultiValue() bool {
    return false
}

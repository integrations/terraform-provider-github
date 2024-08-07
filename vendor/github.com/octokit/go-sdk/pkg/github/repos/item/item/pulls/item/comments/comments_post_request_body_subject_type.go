package comments
import (
    "errors"
)
// The level at which the comment is targeted.
type CommentsPostRequestBody_subject_type int

const (
    LINE_COMMENTSPOSTREQUESTBODY_SUBJECT_TYPE CommentsPostRequestBody_subject_type = iota
    FILE_COMMENTSPOSTREQUESTBODY_SUBJECT_TYPE
)

func (i CommentsPostRequestBody_subject_type) String() string {
    return []string{"line", "file"}[i]
}
func ParseCommentsPostRequestBody_subject_type(v string) (any, error) {
    result := LINE_COMMENTSPOSTREQUESTBODY_SUBJECT_TYPE
    switch v {
        case "line":
            result = LINE_COMMENTSPOSTREQUESTBODY_SUBJECT_TYPE
        case "file":
            result = FILE_COMMENTSPOSTREQUESTBODY_SUBJECT_TYPE
        default:
            return 0, errors.New("Unknown CommentsPostRequestBody_subject_type value: " + v)
    }
    return &result, nil
}
func SerializeCommentsPostRequestBody_subject_type(values []CommentsPostRequestBody_subject_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CommentsPostRequestBody_subject_type) isMultiValue() bool {
    return false
}

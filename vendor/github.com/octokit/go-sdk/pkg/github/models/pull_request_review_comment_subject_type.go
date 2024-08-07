package models
import (
    "errors"
)
// The level at which the comment is targeted, can be a diff line or a file.
type PullRequestReviewComment_subject_type int

const (
    LINE_PULLREQUESTREVIEWCOMMENT_SUBJECT_TYPE PullRequestReviewComment_subject_type = iota
    FILE_PULLREQUESTREVIEWCOMMENT_SUBJECT_TYPE
)

func (i PullRequestReviewComment_subject_type) String() string {
    return []string{"line", "file"}[i]
}
func ParsePullRequestReviewComment_subject_type(v string) (any, error) {
    result := LINE_PULLREQUESTREVIEWCOMMENT_SUBJECT_TYPE
    switch v {
        case "line":
            result = LINE_PULLREQUESTREVIEWCOMMENT_SUBJECT_TYPE
        case "file":
            result = FILE_PULLREQUESTREVIEWCOMMENT_SUBJECT_TYPE
        default:
            return 0, errors.New("Unknown PullRequestReviewComment_subject_type value: " + v)
    }
    return &result, nil
}
func SerializePullRequestReviewComment_subject_type(values []PullRequestReviewComment_subject_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PullRequestReviewComment_subject_type) isMultiValue() bool {
    return false
}

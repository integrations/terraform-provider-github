package models
import (
    "errors"
)
// The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
type PullRequestReviewComment_side int

const (
    LEFT_PULLREQUESTREVIEWCOMMENT_SIDE PullRequestReviewComment_side = iota
    RIGHT_PULLREQUESTREVIEWCOMMENT_SIDE
)

func (i PullRequestReviewComment_side) String() string {
    return []string{"LEFT", "RIGHT"}[i]
}
func ParsePullRequestReviewComment_side(v string) (any, error) {
    result := LEFT_PULLREQUESTREVIEWCOMMENT_SIDE
    switch v {
        case "LEFT":
            result = LEFT_PULLREQUESTREVIEWCOMMENT_SIDE
        case "RIGHT":
            result = RIGHT_PULLREQUESTREVIEWCOMMENT_SIDE
        default:
            return 0, errors.New("Unknown PullRequestReviewComment_side value: " + v)
    }
    return &result, nil
}
func SerializePullRequestReviewComment_side(values []PullRequestReviewComment_side) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PullRequestReviewComment_side) isMultiValue() bool {
    return false
}

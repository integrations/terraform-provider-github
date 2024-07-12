package models
import (
    "errors"
)
// The side of the first line of the range for a multi-line comment.
type PullRequestReviewComment_start_side int

const (
    LEFT_PULLREQUESTREVIEWCOMMENT_START_SIDE PullRequestReviewComment_start_side = iota
    RIGHT_PULLREQUESTREVIEWCOMMENT_START_SIDE
)

func (i PullRequestReviewComment_start_side) String() string {
    return []string{"LEFT", "RIGHT"}[i]
}
func ParsePullRequestReviewComment_start_side(v string) (any, error) {
    result := LEFT_PULLREQUESTREVIEWCOMMENT_START_SIDE
    switch v {
        case "LEFT":
            result = LEFT_PULLREQUESTREVIEWCOMMENT_START_SIDE
        case "RIGHT":
            result = RIGHT_PULLREQUESTREVIEWCOMMENT_START_SIDE
        default:
            return 0, errors.New("Unknown PullRequestReviewComment_start_side value: " + v)
    }
    return &result, nil
}
func SerializePullRequestReviewComment_start_side(values []PullRequestReviewComment_start_side) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PullRequestReviewComment_start_side) isMultiValue() bool {
    return false
}

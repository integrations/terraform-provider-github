package models
import (
    "errors"
)
type CommitComparison_status int

const (
    DIVERGED_COMMITCOMPARISON_STATUS CommitComparison_status = iota
    AHEAD_COMMITCOMPARISON_STATUS
    BEHIND_COMMITCOMPARISON_STATUS
    IDENTICAL_COMMITCOMPARISON_STATUS
)

func (i CommitComparison_status) String() string {
    return []string{"diverged", "ahead", "behind", "identical"}[i]
}
func ParseCommitComparison_status(v string) (any, error) {
    result := DIVERGED_COMMITCOMPARISON_STATUS
    switch v {
        case "diverged":
            result = DIVERGED_COMMITCOMPARISON_STATUS
        case "ahead":
            result = AHEAD_COMMITCOMPARISON_STATUS
        case "behind":
            result = BEHIND_COMMITCOMPARISON_STATUS
        case "identical":
            result = IDENTICAL_COMMITCOMPARISON_STATUS
        default:
            return 0, errors.New("Unknown CommitComparison_status value: " + v)
    }
    return &result, nil
}
func SerializeCommitComparison_status(values []CommitComparison_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CommitComparison_status) isMultiValue() bool {
    return false
}

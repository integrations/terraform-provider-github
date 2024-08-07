package models
import (
    "errors"
)
type MergedUpstream_merge_type int

const (
    MERGE_MERGEDUPSTREAM_MERGE_TYPE MergedUpstream_merge_type = iota
    FASTFORWARD_MERGEDUPSTREAM_MERGE_TYPE
    NONE_MERGEDUPSTREAM_MERGE_TYPE
)

func (i MergedUpstream_merge_type) String() string {
    return []string{"merge", "fast-forward", "none"}[i]
}
func ParseMergedUpstream_merge_type(v string) (any, error) {
    result := MERGE_MERGEDUPSTREAM_MERGE_TYPE
    switch v {
        case "merge":
            result = MERGE_MERGEDUPSTREAM_MERGE_TYPE
        case "fast-forward":
            result = FASTFORWARD_MERGEDUPSTREAM_MERGE_TYPE
        case "none":
            result = NONE_MERGEDUPSTREAM_MERGE_TYPE
        default:
            return 0, errors.New("Unknown MergedUpstream_merge_type value: " + v)
    }
    return &result, nil
}
func SerializeMergedUpstream_merge_type(values []MergedUpstream_merge_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MergedUpstream_merge_type) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
type ContentSymlink_type int

const (
    SYMLINK_CONTENTSYMLINK_TYPE ContentSymlink_type = iota
)

func (i ContentSymlink_type) String() string {
    return []string{"symlink"}[i]
}
func ParseContentSymlink_type(v string) (any, error) {
    result := SYMLINK_CONTENTSYMLINK_TYPE
    switch v {
        case "symlink":
            result = SYMLINK_CONTENTSYMLINK_TYPE
        default:
            return 0, errors.New("Unknown ContentSymlink_type value: " + v)
    }
    return &result, nil
}
func SerializeContentSymlink_type(values []ContentSymlink_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ContentSymlink_type) isMultiValue() bool {
    return false
}

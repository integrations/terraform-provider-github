package models
import (
    "errors"
)
type ContentFile_type int

const (
    FILE_CONTENTFILE_TYPE ContentFile_type = iota
)

func (i ContentFile_type) String() string {
    return []string{"file"}[i]
}
func ParseContentFile_type(v string) (any, error) {
    result := FILE_CONTENTFILE_TYPE
    switch v {
        case "file":
            result = FILE_CONTENTFILE_TYPE
        default:
            return 0, errors.New("Unknown ContentFile_type value: " + v)
    }
    return &result, nil
}
func SerializeContentFile_type(values []ContentFile_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ContentFile_type) isMultiValue() bool {
    return false
}

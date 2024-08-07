package models
import (
    "errors"
)
type File_extension_restriction_type int

const (
    FILE_EXTENSION_RESTRICTION_FILE_EXTENSION_RESTRICTION_TYPE File_extension_restriction_type = iota
)

func (i File_extension_restriction_type) String() string {
    return []string{"file_extension_restriction"}[i]
}
func ParseFile_extension_restriction_type(v string) (any, error) {
    result := FILE_EXTENSION_RESTRICTION_FILE_EXTENSION_RESTRICTION_TYPE
    switch v {
        case "file_extension_restriction":
            result = FILE_EXTENSION_RESTRICTION_FILE_EXTENSION_RESTRICTION_TYPE
        default:
            return 0, errors.New("Unknown File_extension_restriction_type value: " + v)
    }
    return &result, nil
}
func SerializeFile_extension_restriction_type(values []File_extension_restriction_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i File_extension_restriction_type) isMultiValue() bool {
    return false
}

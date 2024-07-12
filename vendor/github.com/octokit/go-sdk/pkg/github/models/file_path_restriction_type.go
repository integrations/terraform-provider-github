package models
import (
    "errors"
)
type File_path_restriction_type int

const (
    FILE_PATH_RESTRICTION_FILE_PATH_RESTRICTION_TYPE File_path_restriction_type = iota
)

func (i File_path_restriction_type) String() string {
    return []string{"file_path_restriction"}[i]
}
func ParseFile_path_restriction_type(v string) (any, error) {
    result := FILE_PATH_RESTRICTION_FILE_PATH_RESTRICTION_TYPE
    switch v {
        case "file_path_restriction":
            result = FILE_PATH_RESTRICTION_FILE_PATH_RESTRICTION_TYPE
        default:
            return 0, errors.New("Unknown File_path_restriction_type value: " + v)
    }
    return &result, nil
}
func SerializeFile_path_restriction_type(values []File_path_restriction_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i File_path_restriction_type) isMultiValue() bool {
    return false
}

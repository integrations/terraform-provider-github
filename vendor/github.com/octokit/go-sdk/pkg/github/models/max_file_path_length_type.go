package models
import (
    "errors"
)
type Max_file_path_length_type int

const (
    MAX_FILE_PATH_LENGTH_MAX_FILE_PATH_LENGTH_TYPE Max_file_path_length_type = iota
)

func (i Max_file_path_length_type) String() string {
    return []string{"max_file_path_length"}[i]
}
func ParseMax_file_path_length_type(v string) (any, error) {
    result := MAX_FILE_PATH_LENGTH_MAX_FILE_PATH_LENGTH_TYPE
    switch v {
        case "max_file_path_length":
            result = MAX_FILE_PATH_LENGTH_MAX_FILE_PATH_LENGTH_TYPE
        default:
            return 0, errors.New("Unknown Max_file_path_length_type value: " + v)
    }
    return &result, nil
}
func SerializeMax_file_path_length_type(values []Max_file_path_length_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Max_file_path_length_type) isMultiValue() bool {
    return false
}

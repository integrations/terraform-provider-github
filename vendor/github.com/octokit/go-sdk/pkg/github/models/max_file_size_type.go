package models
import (
    "errors"
)
type Max_file_size_type int

const (
    MAX_FILE_SIZE_MAX_FILE_SIZE_TYPE Max_file_size_type = iota
)

func (i Max_file_size_type) String() string {
    return []string{"max_file_size"}[i]
}
func ParseMax_file_size_type(v string) (any, error) {
    result := MAX_FILE_SIZE_MAX_FILE_SIZE_TYPE
    switch v {
        case "max_file_size":
            result = MAX_FILE_SIZE_MAX_FILE_SIZE_TYPE
        default:
            return 0, errors.New("Unknown Max_file_size_type value: " + v)
    }
    return &result, nil
}
func SerializeMax_file_size_type(values []Max_file_size_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Max_file_size_type) isMultiValue() bool {
    return false
}

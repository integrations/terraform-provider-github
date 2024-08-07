package models
type RepositoryRuleMember2_type int

const (
    MAX_FILE_PATH_LENGTH_REPOSITORYRULEMEMBER2_TYPE RepositoryRuleMember2_type = iota
)

func (i RepositoryRuleMember2_type) String() string {
    return []string{"max_file_path_length"}[i]
}
func ParseRepositoryRuleMember2_type(v string) (any, error) {
    result := MAX_FILE_PATH_LENGTH_REPOSITORYRULEMEMBER2_TYPE
    switch v {
        case "max_file_path_length":
            result = MAX_FILE_PATH_LENGTH_REPOSITORYRULEMEMBER2_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMember2_type(values []RepositoryRuleMember2_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMember2_type) isMultiValue() bool {
    return false
}

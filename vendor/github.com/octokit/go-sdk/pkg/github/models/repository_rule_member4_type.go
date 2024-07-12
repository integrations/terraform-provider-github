package models
type RepositoryRuleMember4_type int

const (
    MAX_FILE_SIZE_REPOSITORYRULEMEMBER4_TYPE RepositoryRuleMember4_type = iota
)

func (i RepositoryRuleMember4_type) String() string {
    return []string{"max_file_size"}[i]
}
func ParseRepositoryRuleMember4_type(v string) (any, error) {
    result := MAX_FILE_SIZE_REPOSITORYRULEMEMBER4_TYPE
    switch v {
        case "max_file_size":
            result = MAX_FILE_SIZE_REPOSITORYRULEMEMBER4_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMember4_type(values []RepositoryRuleMember4_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMember4_type) isMultiValue() bool {
    return false
}

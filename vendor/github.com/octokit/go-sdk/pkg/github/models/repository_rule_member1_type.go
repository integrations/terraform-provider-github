package models
type RepositoryRuleMember1_type int

const (
    FILE_PATH_RESTRICTION_REPOSITORYRULEMEMBER1_TYPE RepositoryRuleMember1_type = iota
)

func (i RepositoryRuleMember1_type) String() string {
    return []string{"file_path_restriction"}[i]
}
func ParseRepositoryRuleMember1_type(v string) (any, error) {
    result := FILE_PATH_RESTRICTION_REPOSITORYRULEMEMBER1_TYPE
    switch v {
        case "file_path_restriction":
            result = FILE_PATH_RESTRICTION_REPOSITORYRULEMEMBER1_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMember1_type(values []RepositoryRuleMember1_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMember1_type) isMultiValue() bool {
    return false
}

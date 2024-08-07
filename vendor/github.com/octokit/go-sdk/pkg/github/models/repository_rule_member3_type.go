package models
type RepositoryRuleMember3_type int

const (
    FILE_EXTENSION_RESTRICTION_REPOSITORYRULEMEMBER3_TYPE RepositoryRuleMember3_type = iota
)

func (i RepositoryRuleMember3_type) String() string {
    return []string{"file_extension_restriction"}[i]
}
func ParseRepositoryRuleMember3_type(v string) (any, error) {
    result := FILE_EXTENSION_RESTRICTION_REPOSITORYRULEMEMBER3_TYPE
    switch v {
        case "file_extension_restriction":
            result = FILE_EXTENSION_RESTRICTION_REPOSITORYRULEMEMBER3_TYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeRepositoryRuleMember3_type(values []RepositoryRuleMember3_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleMember3_type) isMultiValue() bool {
    return false
}

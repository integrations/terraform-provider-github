package models
import (
    "errors"
)
type RepositoryRuleCodeScanning_type int

const (
    CODE_SCANNING_REPOSITORYRULECODESCANNING_TYPE RepositoryRuleCodeScanning_type = iota
)

func (i RepositoryRuleCodeScanning_type) String() string {
    return []string{"code_scanning"}[i]
}
func ParseRepositoryRuleCodeScanning_type(v string) (any, error) {
    result := CODE_SCANNING_REPOSITORYRULECODESCANNING_TYPE
    switch v {
        case "code_scanning":
            result = CODE_SCANNING_REPOSITORYRULECODESCANNING_TYPE
        default:
            return 0, errors.New("Unknown RepositoryRuleCodeScanning_type value: " + v)
    }
    return &result, nil
}
func SerializeRepositoryRuleCodeScanning_type(values []RepositoryRuleCodeScanning_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RepositoryRuleCodeScanning_type) isMultiValue() bool {
    return false
}

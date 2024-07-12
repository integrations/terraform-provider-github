package models
import (
    "errors"
)
// The type of label. Read-only labels are applied automatically when the runner is configured.
type RunnerLabel_type int

const (
    READONLY_RUNNERLABEL_TYPE RunnerLabel_type = iota
    CUSTOM_RUNNERLABEL_TYPE
)

func (i RunnerLabel_type) String() string {
    return []string{"read-only", "custom"}[i]
}
func ParseRunnerLabel_type(v string) (any, error) {
    result := READONLY_RUNNERLABEL_TYPE
    switch v {
        case "read-only":
            result = READONLY_RUNNERLABEL_TYPE
        case "custom":
            result = CUSTOM_RUNNERLABEL_TYPE
        default:
            return 0, errors.New("Unknown RunnerLabel_type value: " + v)
    }
    return &result, nil
}
func SerializeRunnerLabel_type(values []RunnerLabel_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i RunnerLabel_type) isMultiValue() bool {
    return false
}

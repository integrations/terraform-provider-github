package models
import (
    "errors"
)
type ContentSubmodule_type int

const (
    SUBMODULE_CONTENTSUBMODULE_TYPE ContentSubmodule_type = iota
)

func (i ContentSubmodule_type) String() string {
    return []string{"submodule"}[i]
}
func ParseContentSubmodule_type(v string) (any, error) {
    result := SUBMODULE_CONTENTSUBMODULE_TYPE
    switch v {
        case "submodule":
            result = SUBMODULE_CONTENTSUBMODULE_TYPE
        default:
            return 0, errors.New("Unknown ContentSubmodule_type value: " + v)
    }
    return &result, nil
}
func SerializeContentSubmodule_type(values []ContentSubmodule_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ContentSubmodule_type) isMultiValue() bool {
    return false
}

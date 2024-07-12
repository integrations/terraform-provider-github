package models
import (
    "errors"
)
// The process in which the Page will be built.
type Page_build_type int

const (
    LEGACY_PAGE_BUILD_TYPE Page_build_type = iota
    WORKFLOW_PAGE_BUILD_TYPE
)

func (i Page_build_type) String() string {
    return []string{"legacy", "workflow"}[i]
}
func ParsePage_build_type(v string) (any, error) {
    result := LEGACY_PAGE_BUILD_TYPE
    switch v {
        case "legacy":
            result = LEGACY_PAGE_BUILD_TYPE
        case "workflow":
            result = WORKFLOW_PAGE_BUILD_TYPE
        default:
            return 0, errors.New("Unknown Page_build_type value: " + v)
    }
    return &result, nil
}
func SerializePage_build_type(values []Page_build_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Page_build_type) isMultiValue() bool {
    return false
}

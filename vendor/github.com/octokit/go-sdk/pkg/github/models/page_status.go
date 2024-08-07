package models
import (
    "errors"
)
// The status of the most recent build of the Page.
type Page_status int

const (
    BUILT_PAGE_STATUS Page_status = iota
    BUILDING_PAGE_STATUS
    ERRORED_PAGE_STATUS
)

func (i Page_status) String() string {
    return []string{"built", "building", "errored"}[i]
}
func ParsePage_status(v string) (any, error) {
    result := BUILT_PAGE_STATUS
    switch v {
        case "built":
            result = BUILT_PAGE_STATUS
        case "building":
            result = BUILDING_PAGE_STATUS
        case "errored":
            result = ERRORED_PAGE_STATUS
        default:
            return 0, errors.New("Unknown Page_status value: " + v)
    }
    return &result, nil
}
func SerializePage_status(values []Page_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Page_status) isMultiValue() bool {
    return false
}

package models
import (
    "errors"
)
// The type of advisory.
type GlobalAdvisory_type int

const (
    REVIEWED_GLOBALADVISORY_TYPE GlobalAdvisory_type = iota
    UNREVIEWED_GLOBALADVISORY_TYPE
    MALWARE_GLOBALADVISORY_TYPE
)

func (i GlobalAdvisory_type) String() string {
    return []string{"reviewed", "unreviewed", "malware"}[i]
}
func ParseGlobalAdvisory_type(v string) (any, error) {
    result := REVIEWED_GLOBALADVISORY_TYPE
    switch v {
        case "reviewed":
            result = REVIEWED_GLOBALADVISORY_TYPE
        case "unreviewed":
            result = UNREVIEWED_GLOBALADVISORY_TYPE
        case "malware":
            result = MALWARE_GLOBALADVISORY_TYPE
        default:
            return 0, errors.New("Unknown GlobalAdvisory_type value: " + v)
    }
    return &result, nil
}
func SerializeGlobalAdvisory_type(values []GlobalAdvisory_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GlobalAdvisory_type) isMultiValue() bool {
    return false
}

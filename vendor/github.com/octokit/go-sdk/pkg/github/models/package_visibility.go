package models
import (
    "errors"
)
type Package_visibility int

const (
    PRIVATE_PACKAGE_VISIBILITY Package_visibility = iota
    PUBLIC_PACKAGE_VISIBILITY
)

func (i Package_visibility) String() string {
    return []string{"private", "public"}[i]
}
func ParsePackage_visibility(v string) (any, error) {
    result := PRIVATE_PACKAGE_VISIBILITY
    switch v {
        case "private":
            result = PRIVATE_PACKAGE_VISIBILITY
        case "public":
            result = PUBLIC_PACKAGE_VISIBILITY
        default:
            return 0, errors.New("Unknown Package_visibility value: " + v)
    }
    return &result, nil
}
func SerializePackage_visibility(values []Package_visibility) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Package_visibility) isMultiValue() bool {
    return false
}

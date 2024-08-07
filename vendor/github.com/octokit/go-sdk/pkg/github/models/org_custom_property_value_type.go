package models
import (
    "errors"
)
// The type of the value for the property
type OrgCustomProperty_value_type int

const (
    STRING_ORGCUSTOMPROPERTY_VALUE_TYPE OrgCustomProperty_value_type = iota
    SINGLE_SELECT_ORGCUSTOMPROPERTY_VALUE_TYPE
    MULTI_SELECT_ORGCUSTOMPROPERTY_VALUE_TYPE
    TRUE_FALSE_ORGCUSTOMPROPERTY_VALUE_TYPE
)

func (i OrgCustomProperty_value_type) String() string {
    return []string{"string", "single_select", "multi_select", "true_false"}[i]
}
func ParseOrgCustomProperty_value_type(v string) (any, error) {
    result := STRING_ORGCUSTOMPROPERTY_VALUE_TYPE
    switch v {
        case "string":
            result = STRING_ORGCUSTOMPROPERTY_VALUE_TYPE
        case "single_select":
            result = SINGLE_SELECT_ORGCUSTOMPROPERTY_VALUE_TYPE
        case "multi_select":
            result = MULTI_SELECT_ORGCUSTOMPROPERTY_VALUE_TYPE
        case "true_false":
            result = TRUE_FALSE_ORGCUSTOMPROPERTY_VALUE_TYPE
        default:
            return 0, errors.New("Unknown OrgCustomProperty_value_type value: " + v)
    }
    return &result, nil
}
func SerializeOrgCustomProperty_value_type(values []OrgCustomProperty_value_type) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OrgCustomProperty_value_type) isMultiValue() bool {
    return false
}

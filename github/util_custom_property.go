package github

import (
	"fmt"
	"strconv"

	"github.com/google/go-github/v89/github"
)

// flattenOrganizationRepositoryCustomPropertyDefaultValue normalises the
// polymorphic default_value returned by the API into a list of strings. The
// wire type depends on value_type: multi_select is an array, true_false is a
// stringified bool and the rest are plain strings.
func flattenOrganizationRepositoryCustomPropertyDefaultValue(cp *github.CustomProperty) ([]string, error) {
	if cp.DefaultValue == nil {
		return nil, nil
	}

	switch cp.ValueType {
	case github.PropertyValueTypeMultiSelect:
		if v, ok := cp.DefaultValueStrings(); ok {
			return v, nil
		}
	case github.PropertyValueTypeTrueFalse:
		if v, ok := cp.DefaultValueBool(); ok {
			return []string{strconv.FormatBool(v)}, nil
		}
	default:
		if v, ok := cp.DefaultValueString(); ok {
			return []string{v}, nil
		}
	}

	return nil, fmt.Errorf("default_value %#v could not be parsed for value_type %q", cp.DefaultValue, cp.ValueType)
}

// parseRepositoryCustomPropertyValueToStringSlice normalises the polymorphic
// value of a custom property set on a repository into a list of strings.
func parseRepositoryCustomPropertyValueToStringSlice(prop *github.CustomPropertyValue) ([]string, error) {
	switch value := prop.Value.(type) {
	case string:
		return []string{value}, nil
	case []string:
		return value, nil
	default:
		return nil, fmt.Errorf("custom property value couldn't be parsed as a string or a list of strings: %s", value)
	}
}

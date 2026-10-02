package github

import (
	"fmt"
	"strconv"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const organizationCustomPropertyReadErrorFormat = "error reading organization custom property %q: %v"

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

// allowedValuesForOrganizationRepositoryCustomProperty returns the allowed
// values only for the select types; the API may echo them for other types.
func allowedValuesForOrganizationRepositoryCustomProperty(cp *github.CustomProperty) []string {
	switch cp.ValueType {
	case github.PropertyValueTypeSingleSelect, github.PropertyValueTypeMultiSelect:
		return cp.AllowedValues
	default:
		return nil
	}
}

func organizationRepositoryCustomPropertyToMap(cp *github.CustomProperty) (map[string]any, error) {
	defaultValue, err := flattenOrganizationRepositoryCustomPropertyDefaultValue(cp)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"property_name":      cp.GetPropertyName(),
		"value_type":         string(cp.ValueType),
		"required":           cp.GetRequired(),
		"default_value":      defaultValue,
		"description":        cp.GetDescription(),
		"allowed_values":     allowedValuesForOrganizationRepositoryCustomProperty(cp),
		"values_editable_by": cp.GetValuesEditableBy(),
	}, nil
}

// setOrganizationRepositoryCustomPropertyState writes an API custom property
// definition into the resource or data source state.
func setOrganizationRepositoryCustomPropertyState(d *schema.ResourceData, cp *github.CustomProperty) diag.Diagnostics {
	values, err := organizationRepositoryCustomPropertyToMap(cp)
	if err != nil {
		return diag.Errorf(organizationCustomPropertyReadErrorFormat, cp.GetPropertyName(), err)
	}

	d.SetId(cp.GetPropertyName())
	for key, value := range values {
		if err := d.Set(key, value); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
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

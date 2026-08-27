package github

import (
	"testing"

	"github.com/google/go-github/v89/github"
)

func TestCustomPropertyDefaultValueString(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		prop *github.CustomProperty
		want string
	}{
		"nil property":            {prop: nil, want: ""},
		"nil default":             {prop: &github.CustomProperty{ValueType: github.PropertyValueTypeString}, want: ""},
		"string":                  {prop: &github.CustomProperty{ValueType: github.PropertyValueTypeString, DefaultValue: "dev"}, want: "dev"},
		"single_select":           {prop: &github.CustomProperty{ValueType: github.PropertyValueTypeSingleSelect, DefaultValue: "gold"}, want: "gold"},
		"url":                     {prop: &github.CustomProperty{ValueType: github.PropertyValueTypeURL, DefaultValue: "https://example.com"}, want: "https://example.com"},
		"true_false false":        {prop: &github.CustomProperty{ValueType: github.PropertyValueTypeTrueFalse, DefaultValue: "false"}, want: "false"},
		"true_false true":         {prop: &github.CustomProperty{ValueType: github.PropertyValueTypeTrueFalse, DefaultValue: "true"}, want: "true"},
		"multi_select":            {prop: &github.CustomProperty{ValueType: github.PropertyValueTypeMultiSelect, DefaultValue: []any{"a", "b"}}, want: "a,b"},
		"unexpected shape string": {prop: &github.CustomProperty{ValueType: github.PropertyValueTypeTrueFalse, DefaultValue: "not-a-bool"}, want: "not-a-bool"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := customPropertyDefaultValueString(tc.prop); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

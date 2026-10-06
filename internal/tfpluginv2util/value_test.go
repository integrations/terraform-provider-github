package tfpluginv2util

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGetValue(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		schema   map[string]*schema.Schema
		data     map[string]any
		key      string
		expected any
		ok       bool
		typeOK   bool
	}{
		{
			name:     "key_exists_and_type_matches",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}},
			data:     map[string]any{"foo": "bar"},
			key:      "foo",
			expected: "bar",
			ok:       true,
			typeOK:   true,
		},
		{
			name:     "key_exists_but_type_does_not_match",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeInt}},
			data:     map[string]any{"foo": 42},
			key:      "foo",
			expected: "",
			ok:       true,
			typeOK:   false,
		},
		{
			name:     "key_does_not_exist",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}},
			data:     map[string]any{},
			key:      "foo",
			expected: "",
			ok:       false,
			typeOK:   false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, tt.schema, tt.data)

			v, ok, typeOK := GetValue[string](d, tt.key)
			if v != tt.expected || ok != tt.ok || typeOK != tt.typeOK {
				t.Errorf("GetValue() = (%v, %v, %v), want (%v, %v, %v)", v, ok, typeOK, tt.expected, tt.ok, tt.typeOK)
			}
		})
	}
}

func TestGetOk(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		schema   map[string]*schema.Schema
		data     map[string]any
		key      string
		expected any
		ok       bool
	}{
		{
			name:     "key_exists_and_type_matches",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}},
			data:     map[string]any{"foo": "bar"},
			key:      "foo",
			expected: "bar",
			ok:       true,
		},
		{
			name:     "key_exists_but_type_does_not_match",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeInt}},
			data:     map[string]any{"foo": 42},
			key:      "foo",
			expected: "",
			ok:       false,
		},
		{
			name:     "key_does_not_exist",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}},
			data:     map[string]any{},
			key:      "foo",
			expected: "",
			ok:       false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, tt.schema, tt.data)

			v, ok := GetOk[string](d, tt.key)
			if v != tt.expected || ok != tt.ok {
				t.Errorf("GetOk() = (%v, %v), want (%v, %v)", v, ok, tt.expected, tt.ok)
			}
		})
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		schema   map[string]*schema.Schema
		data     map[string]any
		key      string
		expected any
	}{
		{
			name:     "key_exists_and_type_matches",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}},
			data:     map[string]any{"foo": "bar"},
			key:      "foo",
			expected: "bar",
		},
		{
			name:     "key_exists_but_type_does_not_match",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeInt}},
			data:     map[string]any{"foo": 42},
			key:      "foo",
			expected: "",
		},
		{
			name:     "key_does_not_exist",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}},
			data:     map[string]any{},
			key:      "foo",
			expected: "",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, tt.schema, tt.data)

			v := Get[string](d, tt.key)
			if v != tt.expected {
				t.Errorf("Get() = %v, want %v", v, tt.expected)
			}
		})
	}
}

func TestGetKeysOk(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		schema   map[string]*schema.Schema
		data     map[string]any
		keys     []string
		expected any
		ok       bool
	}{
		{
			name:     "first_key_exists_and_type_matches",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}, "bar": {Type: schema.TypeString}},
			data:     map[string]any{"foo": "baz"},
			keys:     []string{"foo", "bar"},
			expected: "baz",
			ok:       true,
		},
		{
			name:     "second_key_exists_and_type_matches",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}, "bar": {Type: schema.TypeString}},
			data:     map[string]any{"bar": "baz"},
			keys:     []string{"foo", "bar"},
			expected: "baz",
			ok:       true,
		},
		{
			name:     "key_exists_but_type_does_not_match",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeInt}, "bar": {Type: schema.TypeString}},
			data:     map[string]any{"foo": 42},
			keys:     []string{"foo", "bar"},
			expected: "",
			ok:       false,
		},
		{
			name:     "no_keys_exist",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeString}, "bar": {Type: schema.TypeString}},
			data:     map[string]any{},
			keys:     []string{"foo", "bar"},
			expected: "",
			ok:       false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, tt.schema, tt.data)

			v, ok := GetKeysOk[string](d, tt.keys...)
			if v != tt.expected || ok != tt.ok {
				t.Errorf("GetKeysOk() = (%v, %v), want (%v, %v)", v, ok, tt.expected, tt.ok)
			}
		})
	}
}

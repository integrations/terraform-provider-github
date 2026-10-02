package tfpluginv2util

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestGetSetValue(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		schema   map[string]*schema.Schema
		data     map[string]any
		key      string
		expected []string
		ok       bool
		typeOK   bool
	}{
		{
			name:     "key_exists_and_type_matches",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}}},
			data:     map[string]any{"foo": []any{"bar", "baz"}},
			key:      "foo",
			expected: []string{"bar", "baz"},
			ok:       true,
			typeOK:   true,
		},
		{
			name:     "key_exists_but_type_does_not_match",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeInt}}},
			data:     map[string]any{"foo": []any{1, 2}},
			key:      "foo",
			expected: nil,
			ok:       true,
			typeOK:   false,
		},
		{
			name:     "key_does_not_exist",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}}},
			data:     map[string]any{},
			key:      "foo",
			expected: nil,
			ok:       false,
			typeOK:   false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, tt.schema, tt.data)

			v, ok, typeOK := GetSetValue[string](d, tt.key)
			if !cmp.Equal(v, tt.expected) || ok != tt.ok || typeOK != tt.typeOK {
				t.Errorf("GetSetValue() = (%v, %v, %v), want (%v, %v, %v)", v, ok, typeOK, tt.expected, tt.ok, tt.typeOK)
			}
		})
	}
}

func TestGetSetOk(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name     string
		schema   map[string]*schema.Schema
		data     map[string]any
		key      string
		expected []string
		ok       bool
	}{
		{
			name:     "key_exists_and_type_matches",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}}},
			data:     map[string]any{"foo": []any{"bar", "baz"}},
			key:      "foo",
			expected: []string{"bar", "baz"},
			ok:       true,
		},
		{
			name:     "key_exists_but_type_does_not_match",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeInt}}},
			data:     map[string]any{"foo": []any{1, 2}},
			key:      "foo",
			expected: nil,
			ok:       false,
		},
		{
			name:     "key_does_not_exist",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}}},
			data:     map[string]any{},
			key:      "foo",
			expected: nil,
			ok:       false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, tt.schema, tt.data)

			v, ok := GetSetOk[string](d, tt.key)
			if !cmp.Equal(v, tt.expected) || ok != tt.ok {
				t.Errorf("GetSetOk() = (%v, %v), want (%v, %v)", v, ok, tt.expected, tt.ok)
			}
		})
	}
}

func TestGetSet(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name        string
		schema      map[string]*schema.Schema
		data        map[string]any
		key         string
		zeroAsEmpty bool
		expected    []string
	}{
		{
			name:        "key_exists_and_type_matches",
			schema:      map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}}},
			data:        map[string]any{"foo": []any{"bar", "baz"}},
			key:         "foo",
			zeroAsEmpty: false,
			expected:    []string{"bar", "baz"},
		},
		{
			name:        "key_exists_but_type_does_not_match",
			schema:      map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeInt}}},
			data:        map[string]any{"foo": []any{1, 2}},
			key:         "foo",
			zeroAsEmpty: false,
			expected:    nil,
		},
		{
			name:     "key_does_not_exist",
			schema:   map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}}},
			data:     map[string]any{},
			key:      "foo",
			expected: nil,
		},
		{
			name:        "key_does_not_exist_want_empty",
			schema:      map[string]*schema.Schema{"foo": {Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}}},
			data:        map[string]any{},
			key:         "foo",
			zeroAsEmpty: true,
			expected:    []string{},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := schema.TestResourceDataRaw(t, tt.schema, tt.data)

			v := GetSet[string](d, tt.key, tt.zeroAsEmpty)
			if !cmp.Equal(v, tt.expected) {
				t.Errorf("GetSet() = %v, want %v", v, tt.expected)
			}
		})
	}
}

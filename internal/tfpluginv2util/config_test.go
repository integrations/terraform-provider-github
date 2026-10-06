package tfpluginv2util

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestIsSet(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name   string
		schema map[string]*schema.Schema
		raw    cty.Value
		path   string
		want   bool
	}{
		{
			name:   "attribute_set",
			schema: map[string]*schema.Schema{"foo": {Type: schema.TypeString, Optional: true}},
			raw:    cty.ObjectVal(map[string]cty.Value{"foo": cty.StringVal("bar")}),
			path:   "foo",
			want:   true,
		},
		{
			name:   "attribute_not_set",
			schema: map[string]*schema.Schema{"foo": {Type: schema.TypeString, Optional: true}},
			raw:    cty.ObjectVal(map[string]cty.Value{}),
			path:   "foo",
			want:   false,
		},
		{
			name:   "attribute_not_set_with_default",
			schema: map[string]*schema.Schema{"foo": {Type: schema.TypeString, Optional: true, Default: "bar"}},
			raw:    cty.ObjectVal(map[string]cty.Value{}),
			path:   "foo",
			want:   false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			d := testResourceData{
				ResourceData: schema.TestResourceDataRaw(t, tt.schema, map[string]any{}),
				rawConfig:    tt.raw,
			}

			got := IsSet(d, tt.path)

			if got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

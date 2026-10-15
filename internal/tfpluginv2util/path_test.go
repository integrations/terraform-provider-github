package tfpluginv2util

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
)

func TestPath(t *testing.T) {
	t.Parallel()

	for _, tt := range []struct {
		name  string
		input string
		want  cty.Path
	}{
		{
			name:  "empty_string",
			input: "",
			want:  nil,
		},
		{
			name:  "single_attribute",
			input: "foo",
			want:  cty.Path{cty.GetAttrStep{Name: "foo"}},
		},
		{
			name:  "attribute_index_attribute",
			input: "foo.0.bar",
			want:  cty.Path{cty.GetAttrStep{Name: "foo"}, cty.IndexStep{Key: cty.NumberIntVal(0)}, cty.GetAttrStep{Name: "bar"}},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := Path(tt.input)

			if !got.Equals(tt.want) {
				t.Errorf("Path(%q) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

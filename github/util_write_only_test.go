package github

import (
	"errors"
	"strings"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type rawConfigReaderStub struct {
	config cty.Value
	diags  diag.Diagnostics
}

func (r rawConfigReaderStub) GetRawConfigAt(path cty.Path) (cty.Value, diag.Diagnostics) {
	if r.diags != nil {
		return cty.DynamicVal, r.diags
	}
	value, err := path.Apply(r.config)
	if err != nil {
		return cty.DynamicVal, diag.FromErr(errors.New("invalid raw configuration path"))
	}
	return value, nil
}

func TestReadRawWriteOnlyString(t *testing.T) {
	t.Parallel()

	const canary = "raw-write-only-canary"
	configuration := func(secret cty.Value) cty.Value {
		return cty.ObjectVal(map[string]cty.Value{
			"configuration": cty.ListVal([]cty.Value{cty.ObjectVal(map[string]cty.Value{
				"secret_wo": secret,
			})}),
		})
	}

	tests := map[string]struct {
		reader    rawConfigReaderStub
		path      cty.Path
		want      string
		wantError bool
	}{
		"known nested string": {
			reader: rawConfigReaderStub{config: configuration(cty.StringVal(canary))},
			path:   webhookSecretWriteOnlyPath,
			want:   canary,
		},
		"null": {
			reader:    rawConfigReaderStub{config: configuration(cty.NullVal(cty.String))},
			path:      webhookSecretWriteOnlyPath,
			wantError: true,
		},
		"unknown": {
			reader:    rawConfigReaderStub{config: configuration(cty.UnknownVal(cty.String))},
			path:      webhookSecretWriteOnlyPath,
			wantError: true,
		},
		"wrong type": {
			reader:    rawConfigReaderStub{config: configuration(cty.NumberIntVal(1))},
			path:      webhookSecretWriteOnlyPath,
			wantError: true,
		},
		"absent": {
			reader: rawConfigReaderStub{config: cty.ObjectVal(map[string]cty.Value{
				"configuration": cty.ListVal([]cty.Value{cty.EmptyObjectVal}),
			})},
			path:      webhookSecretWriteOnlyPath,
			wantError: true,
		},
		"malformed path": {
			reader:    rawConfigReaderStub{config: configuration(cty.StringVal(canary))},
			path:      cty.GetAttrPath("missing"),
			wantError: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, diags := readRawWriteOnlyString(test.reader, test.path)
			if got != test.want {
				t.Fatalf("value mismatch; got length %d, want length %d", len(got), len(test.want))
			}
			if diags.HasError() != test.wantError {
				t.Fatalf("diagnostic error = %t, want %t", diags.HasError(), test.wantError)
			}
			if strings.Contains(diagsString(diags), canary) {
				t.Fatal("diagnostics contained the write-only value")
			}
		})
	}
}

func diagsString(diags diag.Diagnostics) string {
	var result strings.Builder
	for _, diagnostic := range diags {
		result.WriteString(diagnostic.Summary)
		result.WriteString(diagnostic.Detail)
	}
	return result.String()
}

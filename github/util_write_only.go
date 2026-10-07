package github

import (
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

var webhookSecretWriteOnlyPath = cty.GetAttrPath("configuration").IndexInt(0).GetAttr("secret_wo")

// rawConfigAtReader is deliberately small so the raw-value handling can be
// tested without placing write-only values in ResourceData's ordinary model.
type rawConfigAtReader interface {
	GetRawConfigAt(cty.Path) (cty.Value, diag.Diagnostics)
}

func readRawWriteOnlyString(d rawConfigAtReader, path cty.Path) (string, diag.Diagnostics) {
	value, getDiags := d.GetRawConfigAt(path)
	if getDiags.HasError() {
		return "", writeOnlyValueDiagnostic(path, "could not be read from the raw configuration")
	}
	if !value.IsKnown() {
		return "", writeOnlyValueDiagnostic(path, "must be known when the webhook operation is applied")
	}
	if value.IsNull() {
		return "", writeOnlyValueDiagnostic(path, "must not be null when the webhook operation is applied")
	}
	if !value.Type().Equals(cty.String) {
		return "", writeOnlyValueDiagnostic(path, "must be a string")
	}

	return value.AsString(), nil
}

func writeOnlyValueDiagnostic(path cty.Path, problem string) diag.Diagnostics {
	return diag.Diagnostics{{
		Severity:      diag.Error,
		Summary:       "Invalid write-only webhook secret",
		Detail:        fmt.Sprintf("The write-only attribute configuration.0.secret_wo %s.", problem),
		AttributePath: path,
	}}
}

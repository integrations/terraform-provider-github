package github

import (
	"fmt"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func validateProjectV2Date(value any, path cty.Path) diag.Diagnostics {
	date, ok := value.(string)
	if !ok {
		return diag.Diagnostics{{Severity: diag.Error, Summary: "Invalid Projects V2 date", Detail: fmt.Sprintf("%s must be a string, got %T", path, value), AttributePath: path}}
	}
	if _, err := time.Parse(time.DateOnly, date); err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: "Invalid Projects V2 date", Detail: fmt.Sprintf("%s must use YYYY-MM-DD format: %v", path, err), AttributePath: path}}
	}
	return nil
}

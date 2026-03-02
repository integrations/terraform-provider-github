package checks

import (
	"golang.org/x/tools/go/analysis"

	L001 "github.com/integrations/terraform-provider-github/tools/tfproviderlint/checks/L001"
)

var AllChecks = []*analysis.Analyzer{
	L001.Analyzer,
}

package checks

import (
	"golang.org/x/tools/go/analysis"

	"github.com/integrations/terraform-provider-github/tools/tfproviderlint/checks/L001"
	"github.com/integrations/terraform-provider-github/tools/tfproviderlint/checks/L002"
)

var AllChecks = []*analysis.Analyzer{
	L001.Analyzer,
	L002.Analyzer,
}

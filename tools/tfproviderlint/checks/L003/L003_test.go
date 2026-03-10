package L003_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/integrations/terraform-provider-github/tools/tfproviderlint/checks/L003"
)

func TestL003(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, L003.Analyzer, "testdata/src/a")
}

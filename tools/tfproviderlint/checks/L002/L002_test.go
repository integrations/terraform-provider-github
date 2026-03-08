package L002_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/integrations/terraform-provider-github/tools/tfproviderlint/checks/L002"
)

func TestL002(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, L002.Analyzer, "testdata/src/a")
}

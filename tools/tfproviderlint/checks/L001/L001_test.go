package L001_test

import (
	"testing"

	"github.com/integrations/terraform-provider-github/tools/tfproviderlint/checks/L001"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestL001(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, L001.Analyzer, "testdata/src/a")
}

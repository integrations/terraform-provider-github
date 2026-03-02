// Package L001 defines an Analyzer that checks for
// Schema with ValidateFunc configured
package L001

import (
	"log"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
	"golang.org/x/tools/go/analysis"
)

const Doc = `check for Schema with ValidateFunc configured

The L001 analyzer reports cases of schemas which configures ValidateFunc instead of ValidateDiagFunc
, which will fail provider schema validation.

This is because ValidateFunc is deprecated.`

const analyzerName = "L001"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		schemainfo.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (any, error) {
	log.Println("$$$ L001: running")
	defer log.Println("$$$ L001: done")
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	schemaInfos := pass.ResultOf[schemainfo.Analyzer].([]*schema.SchemaInfo)
	for _, schemaInfo := range schemaInfos {
		if ignorer.ShouldIgnore(analyzerName, schemaInfo.AstCompositeLit) {
			continue
		}

		if !schemaInfo.DeclaresField("ValidateFunc") {
			continue
		}

		log.Printf("$$$ L001: reporting: %d", schemaInfo.AstCompositeLit.Pos())
		pass.Reportf(schemaInfo.AstCompositeLit.Pos(), "%s: schema should not configure ValidateFunc, replace it with ValidateDiagFunc", analyzerName)
	}
	log.Println("$$$ L001: returning nil")
	return nil, nil
}

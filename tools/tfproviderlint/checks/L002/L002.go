package L002

import (
	"go/ast"
	"go/token"

	tfplsh "github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourceinfo"
	"github.com/bflad/tfproviderlint/passes/helper/schema/schemainfo"
	"golang.org/x/tools/go/analysis"

	"github.com/integrations/terraform-provider-github/tools/tfproviderlint/helper/schema"
)

const Doc = `check for resources which configure Create, Update, Delete, or Importer.State instead of CreateContext, UpdateContext, DeleteContext, or Importer.StateContext

The L002 analyzer reports cases of resources which configure Create, Update, Delete, or Importer.State instead of CreateContext, UpdateContext, DeleteContext, or Importer.StateContext`

const analyzerName = "L002"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		resourceinfo.Analyzer,
		schemainfo.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (any, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	resources := pass.ResultOf[resourceinfo.Analyzer].([]*tfplsh.ResourceInfo)
	for _, resource := range resources {

		fields := []string{tfplsh.ResourceFieldCreate, tfplsh.ResourceFieldRead, tfplsh.ResourceFieldUpdate, tfplsh.ResourceFieldDelete}

		for _, field := range fields {
			if resource.DeclaresField(field) {
				if ignorer.ShouldIgnore(analyzerName, resource.Fields[field].Key) {
					continue
				}
				pass.Reportf(resource.Fields[field].Pos(), "%s: resource should not configure %s, replace it with %s", analyzerName, field, field+"Context")
			}
		}

		if resource.DeclaresField(tfplsh.ResourceFieldImporter) {
			var resourceImporterSchema *tfplsh.SchemaInfo
			switch value := resource.Fields[tfplsh.ResourceFieldImporter].Value.(type) {
			case *ast.UnaryExpr:
				if value.Op != token.AND || !schema.IsTypeResourceImporter(pass.TypesInfo.TypeOf(value.X)) {
					continue
				}
				resourceImporterSchema = tfplsh.NewSchemaInfo(value.X.(*ast.CompositeLit), pass.TypesInfo)
			}
			if resourceImporterSchema != nil && resourceImporterSchema.DeclaresField("State") {
				if ignorer.ShouldIgnore(analyzerName, resourceImporterSchema.Fields["State"].Key) {
					continue
				}
				pass.Reportf(resourceImporterSchema.Fields["State"].Pos(), "%s: resource should not configure %s, replace it with %s", analyzerName, "Importer.State", "Importer.StateContext")
			}
		}

	}

	return nil, nil
}

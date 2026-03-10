package L003

import (
	"go/ast"
	"go/token"
	"go/types"

	tfplsh "github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourceinfo"
	"golang.org/x/tools/go/analysis"

	"github.com/integrations/terraform-provider-github/tools/tfproviderlint/helper/schema"
)

const Doc = `check for direct calls to resource CRUD functions

The L003 analyzer reports cases where functions assigned to resource CRUD fields
(Create, Read, Update, Delete, Importer, etc.) are called directly from other code.
These functions are entry points for the Terraform SDK and must not be called directly.`

const analyzerName = "L003"

var Analyzer = &analysis.Analyzer{
	Name: analyzerName,
	Doc:  Doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
		resourceinfo.Analyzer,
	},
	Run: run,
}

var crudFields = []string{
	tfplsh.ResourceFieldCreate,
	tfplsh.ResourceFieldCreateContext,
	tfplsh.ResourceFieldCreateWithoutTimeout,
	tfplsh.ResourceFieldRead,
	tfplsh.ResourceFieldReadContext,
	tfplsh.ResourceFieldReadWithoutTimeout,
	tfplsh.ResourceFieldUpdate,
	tfplsh.ResourceFieldUpdateContext,
	tfplsh.ResourceFieldUpdateWithoutTimeout,
	tfplsh.ResourceFieldDelete,
	tfplsh.ResourceFieldDeleteContext,
	tfplsh.ResourceFieldDeleteWithoutTimeout,
	tfplsh.ResourceFieldExists,
}

func buildCRUDFuncSet(pass *analysis.Pass, resources []*tfplsh.ResourceInfo) map[types.Object]string {
	crudFuncs := make(map[types.Object]string)

	for _, resource := range resources {
		for _, field := range crudFields {
			if !resource.DeclaresField(field) {
				continue
			}
			addFunc(pass, resource.Fields[field].Value, field, crudFuncs)
		}
		addImporterFuncs(pass, resource, crudFuncs)
	}

	return crudFuncs
}

func addFunc(pass *analysis.Pass, value ast.Expr, fieldName string, crudFuncs map[types.Object]string) {
	ident, ok := value.(*ast.Ident)
	if !ok {
		return
	}

	obj, ok := pass.TypesInfo.Uses[ident]
	if !ok {
		return
	}

	crudFuncs[obj] = fieldName
}

func addImporterFuncs(pass *analysis.Pass, resource *tfplsh.ResourceInfo, crudFuncs map[types.Object]string) {
	if !resource.DeclaresField(tfplsh.ResourceFieldImporter) {
		return
	}

	value, ok := resource.Fields[tfplsh.ResourceFieldImporter].Value.(*ast.UnaryExpr)
	if !ok || value.Op != token.AND || !schema.IsTypeResourceImporter(pass.TypesInfo.TypeOf(value.X)) {
		return
	}

	importerSchema := tfplsh.NewSchemaInfo(value.X.(*ast.CompositeLit), pass.TypesInfo)
	if importerSchema == nil {
		return
	}

	for _, field := range []string{"State", "StateContext"} {
		if importerSchema.DeclaresField(field) {
			addFunc(pass, importerSchema.Fields[field].Value, "Importer."+field, crudFuncs)
		}
	}
}

func run(pass *analysis.Pass) (any, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	resources := pass.ResultOf[resourceinfo.Analyzer].([]*tfplsh.ResourceInfo)

	crudFuncs := buildCRUDFuncSet(pass, resources)
	if len(crudFuncs) == 0 {
		return nil, nil
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if ignorer.ShouldIgnore(analyzerName, callExpr) {
				return true
			}

			ident, ok := callExpr.Fun.(*ast.Ident)
			if !ok {
				return true
			}

			obj, ok := pass.TypesInfo.Uses[ident]
			if !ok {
				return true
			}

			fieldName, isCRUD := crudFuncs[obj]
			if !isCRUD {
				return true
			}

			pass.Reportf(callExpr.Pos(), "%s: function %s is a resource CRUD function (%s) and must not be called directly", analyzerName, ident.Name, fieldName)

			return true
		})
	}

	return nil, nil
}

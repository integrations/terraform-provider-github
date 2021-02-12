package analyzer

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

// Analyzer that checks if all struct's fields are initialized
var Analyzer = &analysis.Analyzer{
	Name:     "exhaustivestruct",
	Doc:      "Checks if all struct's fields are initialized",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		var name string

		compositeLit := node.(*ast.CompositeLit)

		i, ok := compositeLit.Type.(*ast.Ident)

		if ok {
			name = i.Name
		} else {
			s, ok := compositeLit.Type.(*ast.SelectorExpr)

			if !ok {
				return
			}

			name = s.Sel.Name
		}

		if compositeLit.Type == nil {
			return
		}

		t := pass.TypesInfo.TypeOf(compositeLit.Type)

		if t == nil {
			return
		}

		str, ok := t.Underlying().(*types.Struct)

		if !ok {
			return
		}

		if len(compositeLit.Elts) == 0 {
			return
		}

		missing := []string{}

		for i := 0; i < str.NumFields(); i++ {
			fieldName := str.Field(i).Name()
			exists := false

			if !str.Field(i).Exported() {
				continue
			}

			for _, e := range compositeLit.Elts {
				if k, ok := e.(*ast.KeyValueExpr); ok {
					if i, ok := k.Key.(*ast.Ident); ok {
						if i.Name == fieldName {
							exists = true
							break
						}
					}
				}
			}

			if !exists {
				missing = append(missing, fieldName)
			}
		}

		if len(missing) == 1 {
			pass.Reportf(node.Pos(), "%s is missing in %s", missing[0], name)
		} else if len(missing) > 1 {
			pass.Reportf(node.Pos(), "%s are missing in %s", strings.Join(missing, ", "), name)
		}
	})

	return nil, nil
}

package canonicalheader

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"net/http"
	"strconv"
	"unsafe"

	"github.com/go-toolsmith/astcast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/types/typeutil"
)

const (
	pkgPath = "net/http"
	name    = "Header"
)

var Analyzer = &analysis.Analyzer{
	Name:     "canonicalheader",
	Doc:      "canonicalheader checks whether net/http.Header uses canonical header",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (any, error) {
	spctor, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, fmt.Errorf("want %T, got %T", spctor, pass.ResultOf[inspect.Analyzer])
	}

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	var outerErr error

	spctor.Preorder(nodeFilter, func(n ast.Node) {
		if outerErr != nil {
			return
		}

		callExp, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}

		fn, ok := typeutil.Callee(pass.TypesInfo, callExp).(*types.Func)
		if !ok {
			return
		}

		signature, ok := fn.Type().(*types.Signature)
		if !ok {
			return
		}

		recv := signature.Recv()
		if recv == nil {
			return
		}

		object, ok := recv.Type().(*types.Named)
		if !ok {
			return
		}

		if !isHTTPHeader(object) {
			return
		}

		if !isValidMethod(astcast.ToSelectorExpr(callExp.Fun).Sel.Name) {
			return
		}

		arg, ok := callExp.Args[0].(*ast.BasicLit)
		if !ok {
			return
		}

		if arg.Kind != token.STRING {
			return
		}

		if len(arg.Value) < 2 {
			return
		}

		quote := arg.Value[0]
		headerKeyOriginal, err := strconv.Unquote(arg.Value)
		if err != nil {
			outerErr = err
			return
		}

		headerKeyCanonical := http.CanonicalHeaderKey(headerKeyOriginal)
		if headerKeyOriginal == headerKeyCanonical {
			return
		}

		newText := make([]byte, 0, len(headerKeyCanonical)+2)
		newText = append(newText, quote)
		newText = append(newText, unsafe.Slice(unsafe.StringData(headerKeyCanonical), len(headerKeyCanonical))...)
		newText = append(newText, quote)

		pass.Report(
			analysis.Diagnostic{
				Pos:     arg.Pos(),
				End:     arg.End(),
				Message: fmt.Sprintf("non-canonical header %q, instead use: %q", headerKeyOriginal, headerKeyCanonical),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: fmt.Sprintf("should replace %q with %q", headerKeyOriginal, headerKeyCanonical),
						TextEdits: []analysis.TextEdit{
							{
								Pos:     arg.Pos(),
								End:     arg.End(),
								NewText: newText,
							},
						},
					},
				},
			},
		)
	})

	return nil, outerErr
}

func isHTTPHeader(named *types.Named) bool {
	return named.Obj() != nil &&
		named.Obj().Pkg() != nil &&
		named.Obj().Pkg().Path() == pkgPath &&
		named.Obj().Name() == name
}

func isValidMethod(name string) bool {
	switch name {
	case "Get", "Set", "Add", "Del", "Values":
		return true
	default:
		return false
	}
}

// nouse provides a linter for forbidding the use of specific identifiers
package forbidigo

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"
	"regexp"
)

type Issue interface {
	Details() string
	Position() token.Position
	String() string
}

type UsedIssue struct {
	identifier string
	pattern    string
	position   token.Position
}

func (a UsedIssue) Details() string {
	return fmt.Sprintf("use of `%s` forbidden by pattern `%s`", a.identifier, a.pattern)
}

func (a UsedIssue) Position() token.Position {
	return a.position
}

func (a UsedIssue) String() string { return toString(a) }

func toString(i Issue) string {
	return fmt.Sprintf("%s at %s", i.Details(), i.Position())
}

type Linter struct {
	patterns []*regexp.Regexp
}

func DefaultPatterns() []string {
	return []string{`^fmt\.Print(|f|ln)$`}
}

func NewLinter(patterns []string) (*Linter, error) {
	if len(patterns) == 0 {
		patterns = DefaultPatterns()
	}
	compiledPatterns := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		re, err := regexp.Compile(p)
		if err != nil {
			return nil, fmt.Errorf("unable to compile pattern `%s`: %s", p, err)
		}
		compiledPatterns = append(compiledPatterns, re)
	}
	return &Linter{
		patterns: compiledPatterns,
	}, nil
}

type visitor struct {
	linter   *Linter
	comments []*ast.CommentGroup

	fset   *token.FileSet
	issues []Issue
}

func (l *Linter) Run(fset *token.FileSet, nodes ...ast.Node) ([]Issue, error) {
	var issues []Issue // nolint:prealloc // we don't know how many there will be
	for _, node := range nodes {
		var comments []*ast.CommentGroup
		if file, ok := node.(*ast.File); ok {
			comments = file.Comments
		}
		visitor := visitor{
			linter:   l,
			fset:     fset,
			comments: comments,
		}
		ast.Walk(&visitor, node)
		issues = append(issues, visitor.issues...)
	}
	return issues, nil
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.SelectorExpr:
	case *ast.Ident:
	default:
		return v
	}
	for _, p := range v.linter.patterns {
		if p.MatchString(v.textFor(node)) && !v.permit(node) {
			v.issues = append(v.issues, UsedIssue{
				identifier: v.textFor(node),
				pattern:    p.String(),
				position:   v.fset.Position(node.Pos()),
			})
		}
	}
	return nil
}

func (v *visitor) textFor(node ast.Node) string {
	buf := new(bytes.Buffer)
	if err := printer.Fprint(buf, v.fset, node); err != nil {
		log.Fatalf("ERROR: unable to print node at %s: %s", v.fset.Position(node.Pos()), err)
	}
	return buf.String()
}

func (v *visitor) permit(node ast.Node) bool {
	nodePos := v.fset.Position(node.Pos())
	var nolint = regexp.MustCompile(fmt.Sprintf(`^permit:%s\b`, regexp.QuoteMeta(v.textFor(node))))
	for _, c := range v.comments {
		commentPos := v.fset.Position(c.Pos())
		if commentPos.Line == nodePos.Line && nolint.MatchString(c.Text()) {
			return true
		}
	}
	return false
}

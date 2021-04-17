package gogrep

import (
	"go/ast"
	"go/token"
)

type StmtList = stmtList
type ExprList = exprList

func IsEmptyList(n ast.Node) bool {
	if list, ok := n.(nodeList); ok {
		return list.len() == 0
	}
	return false
}

// Parse creates a gogrep pattern out of a given string expression.
func Parse(fset *token.FileSet, expr string, strict bool) (*Pattern, error) {
	m := matcher{
		fset:    fset,
		strict:  strict,
		capture: make([]CapturedNode, 0, 8),
	}
	node, err := m.parseExpr(expr)
	if err != nil {
		return nil, err
	}
	return &Pattern{m: &m, Expr: node}, nil
}

// Pattern is a compiled gogrep pattern.
type Pattern struct {
	Expr ast.Node
	m    *matcher
}

// MatchData describes a successful pattern match.
type MatchData struct {
	Node    ast.Node
	Capture []CapturedNode
}

type CapturedNode struct {
	Name string
	Node ast.Node
}

func (data MatchData) CapturedByName(name string) (ast.Node, bool) {
	return findNamed(data.Capture, name)
}

// Clone creates a pattern copy.
func (p *Pattern) Clone() *Pattern {
	clone := *p
	clone.m = &matcher{}
	*clone.m = *p.m
	clone.m.capture = make([]CapturedNode, 0, 8)
	return &clone
}

// MatchNode calls cb if n matches a pattern.
func (p *Pattern) MatchNode(n ast.Node, cb func(MatchData)) {
	p.m.capture = p.m.capture[:0]
	if p.m.node(p.Expr, n) {
		cb(MatchData{
			Capture: p.m.capture,
			Node:    n,
		})
	}
}

func (p *Pattern) MatchStmtList(stmts []ast.Stmt, cb func(MatchData)) {
	p.matchNodeList(p.Expr.(stmtList), stmtList(stmts), cb)
}

func (p *Pattern) MatchExprList(exprs []ast.Expr, cb func(MatchData)) {
	p.matchNodeList(p.Expr.(exprList), exprList(exprs), cb)
}

func (p *Pattern) matchNodeList(pattern, list nodeList, cb func(MatchData)) {
	listLen := list.len()
	from := 0
	for {
		p.m.capture = p.m.capture[:0]
		matched, offset := p.m.nodes(pattern, list.slice(from, listLen), true)
		if matched == nil {
			break
		}
		cb(MatchData{
			Capture: p.m.capture,
			Node:    matched,
		})
		from += offset - 1
		if from >= listLen {
			break
		}
	}
}

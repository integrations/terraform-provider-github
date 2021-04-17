package ruleguard

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/quasilyte/go-ruleguard/internal/mvdan.cc/gogrep"
	"github.com/quasilyte/go-ruleguard/ruleguard/goutil"
)

type rulesRunner struct {
	state *engineState

	ctx   *RunContext
	rules *goRuleSet

	importer *goImporter

	filename string
	src      []byte

	filterParams filterParams
}

func newRulesRunner(ctx *RunContext, state *engineState, rules *goRuleSet) *rulesRunner {
	importer := newGoImporter(state, goImporterConfig{
		fset:         ctx.Fset,
		debugImports: ctx.DebugImports,
		debugPrint:   ctx.DebugPrint,
	})
	rr := &rulesRunner{
		ctx:      ctx,
		importer: importer,
		rules:    rules,
		filterParams: filterParams{
			env:      state.env.GetEvalEnv(),
			importer: importer,
			ctx:      ctx,
		},
	}
	rr.filterParams.nodeText = rr.nodeText
	return rr
}

func (rr *rulesRunner) nodeText(n ast.Node) []byte {
	if gogrep.IsEmptyList(n) {
		return nil
	}

	from := rr.ctx.Fset.Position(n.Pos()).Offset
	to := rr.ctx.Fset.Position(n.End()).Offset
	src := rr.fileBytes()
	if (from >= 0 && from < len(src)) && (to >= 0 && to < len(src)) {
		return src[from:to]
	}
	// Fallback to the printer.
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, rr.ctx.Fset, n); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func (rr *rulesRunner) fileBytes() []byte {
	if rr.src != nil {
		return rr.src
	}

	// TODO(quasilyte): re-use src slice?
	src, err := ioutil.ReadFile(rr.filename)
	if err != nil || src == nil {
		// Assign a zero-length slice so rr.src
		// is never nil during the second fileBytes call.
		rr.src = make([]byte, 0)
	} else {
		rr.src = src
	}
	return rr.src
}

func (rr *rulesRunner) run(f *ast.File) error {
	// TODO(quasilyte): run local rules as well.

	rr.filename = rr.ctx.Fset.Position(f.Pos()).Filename
	rr.filterParams.filename = rr.filename
	rr.collectImports(f)

	if rr.rules.universal.categorizedNum != 0 {
		ast.Inspect(f, func(n ast.Node) bool {
			if n == nil {
				return false
			}
			rr.runRules(n)
			return true
		})
	}

	return nil
}

func (rr *rulesRunner) runRules(n ast.Node) {
	switch n := n.(type) {
	case *ast.BlockStmt:
		rr.runStmtListRules(n.List)
	case *ast.CaseClause:
		rr.runStmtListRules(n.Body)
	case *ast.CommClause:
		rr.runStmtListRules(n.Body)

	case *ast.CallExpr:
		rr.runExprListRules(n.Args)
	case *ast.CompositeLit:
		rr.runExprListRules(n.Elts)
	case *ast.ReturnStmt:
		rr.runExprListRules(n.Results)
	}

	cat := categorizeNode(n)
	for _, rule := range rr.rules.universal.rulesByCategory[cat] {
		matched := false
		rule.pat.MatchNode(n, func(m gogrep.MatchData) {
			matched = rr.handleMatch(rule, m)
		})
		if matched {
			break
		}
	}
}

func (rr *rulesRunner) runExprListRules(list []ast.Expr) {
	for _, rule := range rr.rules.universal.rulesByCategory[nodeExprList] {
		matched := false
		rule.pat.MatchExprList(list, func(m gogrep.MatchData) {
			matched = rr.handleMatch(rule, m)
		})
		if matched {
			break
		}
	}
}

func (rr *rulesRunner) runStmtListRules(list []ast.Stmt) {
	for _, rule := range rr.rules.universal.rulesByCategory[nodeStmtList] {
		matched := false
		rule.pat.MatchStmtList(list, func(m gogrep.MatchData) {
			matched = rr.handleMatch(rule, m)
		})
		if matched {
			break
		}
	}
}

func (rr *rulesRunner) reject(rule goRule, reason string, m gogrep.MatchData) {
	if rule.group != rr.ctx.Debug {
		return // This rule is not being debugged
	}

	pos := rr.ctx.Fset.Position(m.Node.Pos())
	rr.ctx.DebugPrint(fmt.Sprintf("%s:%d: [%s:%d] rejected by %s",
		pos.Filename, pos.Line, filepath.Base(rule.filename), rule.line, reason))

	values := make([]gogrep.CapturedNode, len(m.Capture))
	copy(values, m.Capture)
	sort.Slice(values, func(i, j int) bool {
		return values[i].Name < values[j].Name
	})

	for _, v := range values {
		name := v.Name
		node := v.Node
		var expr ast.Expr
		switch node := node.(type) {
		case ast.Expr:
			expr = node
		case *ast.ExprStmt:
			expr = node.X
		default:
			continue
		}

		typ := rr.ctx.Types.TypeOf(expr)
		typeString := "<unknown>"
		if typ != nil {
			typeString = typ.String()
		}
		s := strings.ReplaceAll(goutil.SprintNode(rr.ctx.Fset, expr), "\n", `\n`)
		rr.ctx.DebugPrint(fmt.Sprintf("  $%s %s: %s", name, typeString, s))
	}
}

func (rr *rulesRunner) handleMatch(rule goRule, m gogrep.MatchData) bool {
	if rule.filter.fn != nil {
		rr.filterParams.match = m
		filterResult := rule.filter.fn(&rr.filterParams)
		if !filterResult.Matched() {
			rr.reject(rule, filterResult.RejectReason(), m)
			return false
		}
	}

	message := rr.renderMessage(rule.msg, m, true)
	node := m.Node
	if rule.location != "" {
		node, _ = m.CapturedByName(rule.location)
	}
	var suggestion *Suggestion
	if rule.suggestion != "" {
		suggestion = &Suggestion{
			Replacement: []byte(rr.renderMessage(rule.suggestion, m, false)),
			From:        node.Pos(),
			To:          node.End(),
		}
	}
	info := GoRuleInfo{
		Group:    rule.group,
		Filename: rule.filename,
		Line:     rule.line,
	}
	rr.ctx.Report(info, node, message, suggestion)
	return true
}

func (rr *rulesRunner) collectImports(f *ast.File) {
	rr.filterParams.imports = make(map[string]struct{}, len(f.Imports))
	for _, spec := range f.Imports {
		s, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			continue
		}
		rr.filterParams.imports[s] = struct{}{}
	}
}

func (rr *rulesRunner) renderMessage(msg string, m gogrep.MatchData, truncate bool) string {
	var buf strings.Builder
	if strings.Contains(msg, "$$") {
		buf.Write(rr.nodeText(m.Node))
		msg = strings.ReplaceAll(msg, "$$", buf.String())
	}
	if len(m.Capture) == 0 {
		return msg
	}

	capture := make([]gogrep.CapturedNode, len(m.Capture))
	copy(capture, m.Capture)
	sort.Slice(capture, func(i, j int) bool {
		return len(capture[i].Name) > len(capture[j].Name)
	})

	for _, c := range capture {
		n := c.Node
		key := "$" + c.Name
		if !strings.Contains(msg, key) {
			continue
		}
		buf.Reset()
		buf.Write(rr.nodeText(n))
		// Don't interpolate strings that are too long.
		var replacement string
		if truncate && buf.Len() > 60 {
			replacement = key
		} else {
			replacement = buf.String()
		}
		msg = strings.ReplaceAll(msg, key, replacement)
	}
	return msg
}

package godot

import (
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"
	"strings"
)

// specialReplacer is a replacer for some types of special lines in comments,
// which shouldn't be checked. For example, if comment ends with a block of
// code it should not necessarily have a period at the end.
const specialReplacer = "<godotSpecialReplacer>"

// getComments extracts comments from a file.
func getComments(file *ast.File, fset *token.FileSet, scope Scope) ([]comment, error) {
	if len(file.Comments) == 0 {
		return nil, nil
	}

	// Read original file. This is necessary for making a replacements for
	// inline comments. I couldn't find a better way to get original line
	// with code and comment without reading the file. Function `Format`
	// from "go/format" won't help here if the original file is not gofmt-ed.
	lines, err := readFile(file, fset)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}

	// Check consistency to avoid checking index in each function
	lastComment := file.Comments[len(file.Comments)-1]
	if p := fset.Position(lastComment.End()); len(lines) < p.Line {
		return nil, fmt.Errorf("inconsistence between file and AST: %s", p.Filename)
	}

	var comments []comment
	decl := getDeclarationComments(file, fset, lines)
	switch scope {
	case AllScope:
		// All comments
		comments = getAllComments(file, fset, lines)
	case TopLevelScope:
		// All top level comments and comments from the inside
		// of top level blocks
		comments = append(
			getBlockComments(file, fset, lines),
			getTopLevelComments(file, fset, lines)...,
		)
	default:
		// Top level declaration comments and comments from the inside
		// of top level blocks
		comments = append(getBlockComments(file, fset, lines), decl...)
	}

	// Set `decl` flag
	setDecl(comments, decl)

	return comments, nil
}

// getBlockComments gets comments from the inside of top level
// blocks: var (...), const (...).
func getBlockComments(file *ast.File, fset *token.FileSet, lines []string) []comment {
	var comments []comment
	for _, decl := range file.Decls {
		d, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		// No parenthesis == no block
		if d.Lparen == 0 {
			continue
		}
		for _, c := range file.Comments {
			// Skip comments outside this block
			if d.Lparen > c.Pos() || c.Pos() > d.Rparen {
				continue
			}
			// Skip comments that are not top-level for this block
			// (the block itself is top level, so comments inside this block
			// would be on column 2)
			// nolint: gomnd
			if fset.Position(c.Pos()).Column != 2 {
				continue
			}
			firstLine := fset.Position(c.Pos()).Line
			lastLine := fset.Position(c.End()).Line
			comments = append(comments, comment{
				ast:   c,
				lines: lines[firstLine-1 : lastLine],
			})
		}
	}
	return comments
}

// getTopLevelComments gets all top level comments.
func getTopLevelComments(file *ast.File, fset *token.FileSet, lines []string) []comment {
	var comments []comment // nolint: prealloc
	for _, c := range file.Comments {
		if fset.Position(c.Pos()).Column != 1 {
			continue
		}
		firstLine := fset.Position(c.Pos()).Line
		lastLine := fset.Position(c.End()).Line
		comments = append(comments, comment{
			ast:   c,
			lines: lines[firstLine-1 : lastLine],
		})
	}
	return comments
}

// getDeclarationComments gets top level declaration comments.
func getDeclarationComments(file *ast.File, fset *token.FileSet, lines []string) []comment {
	var comments []comment // nolint: prealloc
	for _, decl := range file.Decls {
		var cg *ast.CommentGroup
		switch d := decl.(type) {
		case *ast.GenDecl:
			cg = d.Doc
		case *ast.FuncDecl:
			cg = d.Doc
		}

		if cg == nil {
			continue
		}

		firstLine := fset.Position(cg.Pos()).Line
		lastLine := fset.Position(cg.End()).Line
		comments = append(comments, comment{
			ast:   cg,
			lines: lines[firstLine-1 : lastLine],
		})
	}
	return comments
}

// getAllComments gets every single comment from the file.
func getAllComments(file *ast.File, fset *token.FileSet, lines []string) []comment {
	var comments []comment //nolint: prealloc
	for _, c := range file.Comments {
		firstLine := fset.Position(c.Pos()).Line
		lastLine := fset.Position(c.End()).Line
		comments = append(comments, comment{
			ast:   c,
			lines: lines[firstLine-1 : lastLine],
		})
	}
	return comments
}

// getText extracts text from comment. If comment is a special block
// (e.g., CGO code), a block of empty lines is returned. If comment contains
// special lines (e.g., tags or indented code examples), they are replaced
// with `specialReplacer` to skip checks for it.
// The result can be multiline.
func getText(comment *ast.CommentGroup) (s string) {
	if len(comment.List) == 1 &&
		strings.HasPrefix(comment.List[0].Text, "/*") &&
		isSpecialBlock(comment.List[0].Text) {
		return ""
	}

	for _, c := range comment.List {
		text := c.Text
		isBlock := false
		if strings.HasPrefix(c.Text, "/*") {
			isBlock = true
			text = strings.TrimPrefix(text, "/*")
			text = strings.TrimSuffix(text, "*/")
		}
		for _, line := range strings.Split(text, "\n") {
			if isSpecialLine(line) {
				s += specialReplacer + "\n"
				continue
			}
			if !isBlock {
				line = strings.TrimPrefix(line, "//")
			}
			s += line + "\n"
		}
	}
	if len(s) == 0 {
		return ""
	}
	return s[:len(s)-1] // trim last "\n"
}

// readFile reads file and returns it's lines as strings.
func readFile(file *ast.File, fset *token.FileSet) ([]string, error) {
	fname := fset.File(file.Package)
	f, err := ioutil.ReadFile(fname.Name())
	if err != nil {
		return nil, err
	}
	return strings.Split(string(f), "\n"), nil
}

// setDecl sets `decl` flag to comments which are declaration comments.
func setDecl(comments, decl []comment) {
	for _, d := range decl {
		for i, c := range comments {
			if d.ast.Pos() == c.ast.Pos() {
				comments[i].decl = true
				break
			}
		}
	}
}

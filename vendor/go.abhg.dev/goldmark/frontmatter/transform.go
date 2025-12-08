package frontmatter

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// MetaTransformer is an AST transformer
// that converts the front matter of a document into a map[string]interface{}
// and stores it on the document as metadata.
//
// Access the metadata by calling the [ast.Document.Meta] method.
type MetaTransformer struct{}

var _ parser.ASTTransformer = (*MetaTransformer)(nil)

// Transform fetches the front matter from the parser context
// and sets it as the metadata of the document.
//
// This is a no-op if the front matter is not present in the context
// or it cannot be decoded.
func (t *MetaTransformer) Transform(node *ast.Document, _ text.Reader, pc parser.Context) {
	fm := Get(pc)
	if fm == nil {
		return
	}

	var data map[string]any
	if err := fm.Decode(&data); err != nil {
		return
	}

	node.SetMeta(data)
}

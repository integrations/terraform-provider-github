package frontmatter

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

// Mode specifies the mode in which the Extender operates.
// By default, the extender extracts the front matter from the document,
// but does not render it or do anything else with it.
//
// Change the mode by setting the Mode field of the Extender object.
type Mode int

//go:generate stringer -type Mode

const (
	// SetMetadata instructs the extender to convert the front matter
	// into a map[string]interface{} and set it as the metadata
	// of the document.
	//
	// This may be accessed by calling the Document.Meta() method.
	SetMetadata Mode = 1 << iota
)

// Extender adds support for front matter to a Goldmark Markdown parser.
//
// Use it by installing it into the [goldmark.Markdown] object upon creation.
// For example:
//
//	goldmark.New(
//		// ...
//		goldmark.WithExtensions(
//			// ...
//			&frontmatter.Extender{},
//		),
//	)
type Extender struct {
	// Formats lists the front matter formats
	// that are supported by the extender.
	//
	// If empty, DefaultFormats is used.
	Formats []Format

	// Mode specifies the mode in which the extender operates.
	// See documentation of the Mode type for more information.
	Mode Mode
}

var _ goldmark.Extender = (*Extender)(nil)

// Extend extends the provided Goldmark Markdown.
func (e *Extender) Extend(md goldmark.Markdown) {
	md.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(&Parser{
				Formats: e.Formats,
			}, 0),
		),
	)

	if e.Mode&SetMetadata != 0 {
		md.Parser().AddOptions(
			parser.WithASTTransformers(
				util.Prioritized(&MetaTransformer{}, 0),
			),
		)
	}
}

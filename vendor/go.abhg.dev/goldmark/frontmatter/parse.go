package frontmatter

import (
	"bytes"
	"sync"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// Parser parses front matter from a Markdown document.
type Parser struct {
	// Formats specifies the front matter formats
	// supported by the parser.
	//
	// If Formats is empty, DefaultFormats is used.
	Formats []Format

	once         sync.Once
	triggers     []byte          // list of open delimiters
	formatByOpen map[byte]Format // open delim => format
}

var _ parser.BlockParser = (*Parser)(nil)

func (p *Parser) init() {
	p.once.Do(func() {
		formats := p.Formats
		if len(formats) == 0 {
			formats = DefaultFormats
		}

		p.formatByOpen = make(map[byte]Format, len(formats))
		for _, format := range formats {
			p.triggers = append(p.triggers, format.Delim)
			p.formatByOpen[format.Delim] = format
		}
	})
}

// Trigger returns bytes that can trigger this parser.
//
// This implements [parser.BlockParser].
func (p *Parser) Trigger() []byte {
	p.init()

	return p.triggers
}

// Open begins parsing a frontmatter block,
// returning nil if a frontmatter block is not found.
//
// This implements [parser.BlockParser].
func (p *Parser) Open(_ ast.Node, reader text.Reader, _ parser.Context) (ast.Node, parser.State) {
	p.init()

	// Front matter must be at the beginning of the document.
	if lineno, _ := reader.Position(); lineno > 1 {
		return nil, parser.NoChildren
	}

	line, seg := reader.PeekLine()
	delim, delimCount := lineDelim(line)
	if delim == 0 {
		return nil, parser.NoChildren
	}

	format, ok := p.formatByOpen[delim]
	if !ok {
		return nil, parser.NoChildren
	}

	return &frontmatterNode{
		Format:     format,
		DelimCount: delimCount,
		Segment: text.Segment{
			Start: seg.Stop,
			Stop:  seg.Stop,
		},
	}, parser.NoChildren
}

// Continue continues parsing the following lines of a frontmatter block,
// transitioning to Close when the block is finished.
//
// This implements [parser.BlockParser].
func (p *Parser) Continue(node ast.Node, reader text.Reader, _ parser.Context) parser.State {
	p.init()

	n := node.(*frontmatterNode)
	line, seg := reader.PeekLine()

	if delim, count := lineDelim(line); delim != 0 {
		if delim == n.Format.Delim && count == n.DelimCount {
			reader.Advance(seg.Len())
			return parser.Close
		}
	}
	n.Segment.Stop = seg.Stop
	return parser.Continue | parser.NoChildren
}

// Close cleans up after parsing a frontmatter block.
//
// This implements [parser.BlockParser].
func (p *Parser) Close(node ast.Node, reader text.Reader, pc parser.Context) {
	p.init()

	n := node.(*frontmatterNode)
	raw := n.Segment.Value(reader.Source())

	(&Data{
		format: n.Format,
		raw:    raw,
	}).set(pc)

	parent := node.Parent()
	parent.RemoveChild(parent, node)
}

// CanInterruptParagraph reports that a frontmatter block cannot interrupt a paragraph.
//
// This implements [parser.BlockParser].
func (*Parser) CanInterruptParagraph() bool {
	return false
}

// CanAcceptIndentedLine reports that a frontmatter block cannot be indented.
//
// This implements [parser.BlockParser].
func (*Parser) CanAcceptIndentedLine() bool {
	return false
}

var _kind = ast.NewNodeKind("frontmatter")

// Hidden node to store information about the parse state
// before the front matter is placed in the parser context.
type frontmatterNode struct {
	ast.BaseBlock

	// Format holds the front matter format we matched.
	Format Format

	// Number of times the delimiter was repeated
	// in the opening line.
	DelimCount int

	// Segment holds the text range over which the front matter spans.
	Segment text.Segment
}

var _ ast.Node = (*frontmatterNode)(nil)

func (n *frontmatterNode) Dump(source []byte, level int) {
	ast.DumpHelper(n, source, level, map[string]string{
		"Format": n.Format.Name,
		"Data":   string(n.Segment.Value(source)),
	}, nil)
}

func (n *frontmatterNode) Kind() ast.NodeKind {
	return _kind
}

var (
	_cr = []byte("\r")
	_lf = []byte("\n")
)

// lineDelim interprets the line as a frontmatter delimiting line.
// If the line is a delimiting line, it returns the delimiter
// and the number of times it was repeated.
// Otherwise, it returns 0.
func lineDelim(line []byte) (delim byte, count int) {
	// CR and LF stripped separately
	// to handle both, CRLF and just LF.
	line = bytes.TrimSuffix(line, _lf)
	line = bytes.TrimSuffix(line, _cr)
	if len(line) < 3 {
		return 0, 0
	}

	delim = line[0]
	for _, c := range line[1:] {
		if c != delim {
			return 0, 0
		}
	}
	return delim, len(line)
}

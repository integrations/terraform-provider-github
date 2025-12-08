package frontmatter

import "github.com/yuin/goldmark/parser"

// _dataKey is the ContextKey under which the frontmatter data is stored
// in the [parser.Context].
var _dataKey = parser.NewContextKey()

// Data holds the front matter data.
// Use [Get] to retrieve the data from the [parser.Context].
type Data struct {
	raw    []byte
	format Format
}

// Get retrieves the front matter data from the [parser.Context].
// If the data is not present, it returns nil.
func Get(ctx parser.Context) *Data {
	d, _ := ctx.Get(_dataKey).(*Data)
	return d
}

// Decode decodes the front matter data into the provided value.
// The value must be a pointer to a struct or a map.
//
//	data := frontmatter.Get(ctx)
//	if data == nil {
//		return errors.New("no front matter")
//	}
//
//	var metadata struct {
//		Title string
//		Tags []string
//	}
//	if err := data.Decode(&metadata); err != nil {
//		return err
//	}
func (d *Data) Decode(dst any) error {
	return d.format.Unmarshal(d.raw, dst)
}

// set stores front matter data in the [parser.Context].
func (d *Data) set(ctx parser.Context) {
	ctx.Set(_dataKey, d)
}

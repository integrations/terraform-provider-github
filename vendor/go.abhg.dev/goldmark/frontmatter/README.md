# goldmark-frontmatter

[![Go Reference](https://pkg.go.dev/badge/go.abhg.dev/goldmark/frontmatter.svg)](https://pkg.go.dev/go.abhg.dev/goldmark/frontmatter)
[![CI](https://github.com/abhinav/goldmark-frontmatter/actions/workflows/ci.yml/badge.svg)](https://github.com/abhinav/goldmark-frontmatter/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/abhinav/goldmark-frontmatter/branch/main/graph/badge.svg?token=Q47RX5AA3O)](https://codecov.io/gh/abhinav/goldmark-frontmatter)

goldmark-frontmatter is an extension for the [goldmark] Markdown parser
that adds support for parsing YAML and TOML front matter from Markdown documents.

  [goldmark]: http://github.com/yuin/goldmark

## Features

- Parses YAML and TOML front matter out of the box
- Allows defining your own front matter formats
- Exposes front matter in via a types-safe API

#### Comparison to goldmark-meta

[yuin/goldmark-meta](https://github.com/yuin/goldmark-meta)
is another extension for goldmark that provides support for frontmatter.
Here's a quick comparison of the two extensions:

| Feature                  | goldmark-frontmatter | goldmark-meta |
| ------------------------ | -------------------- | ------------- |
| YAML support (`---`)     | ✅                   | ✅            |
| TOML support (`+++`)     | ✅                   | ❌            |
| Custom formats           | ✅                   | ❌            |
| Decode `map[string]any`  | ✅                   | ✅            |
| Decode type-safe structs | ✅                   | ❌            |
| Render as table          | ❌                   | ✅            |

### Demo

A web-based demonstration of the extension is available at
<https://abhinav.github.io/goldmark-frontmatter/demo/>.

## Installation

```bash
go get go.abhg.dev/goldmark/frontmatter@latest
```

## Usage

To use goldmark-frontmatter, import the `frontmatter` package.

```go
import "go.abhg.dev/goldmark/frontmatter"
```

Then include the `frontmatter.Extender` in the list of extensions
when constructing your [`goldmark.Markdown`].

  [`goldmark.Markdown`]: https://pkg.go.dev/github.com/yuin/goldmark#Markdown

```go
goldmark.New(
  goldmark.WithExtensions(
    // ...
    &frontmatter.Extender{},
  ),
).Convert(src, out)
```

By default, this won't have any effect except stripping the front matter
from the document.
See [Accessing front matter](#accessing-front-matter) on how to read it.

### Syntax

Front matter starts with three or more instances of a delimiter,
and must be the first line of the document.

The supported delimiters are:

- YAML: `-`

    For example:

    ```markdown
    ---
    title: goldmark-frontmatter
    tags: [markdown, goldmark]
    description: |
      Adds support for parsing YAML front matter.
    ---

    # Heading 1
    ```

- TOML: `+`

    For example:

    ```markdown
    +++
    title = "goldmark-frontmatter"
    tags = ["markdown", "goldmark"]
    description = """\
      Adds support for parsing YAML front matter.\
      """
    +++

    # Heading 1
    ```

The front matter block ends with the same number of instances of the delimiter.
So if the opening line used 10 occurrences, so must the closing.

    ---------------------------
    title: goldmark-frontmatter
    tags: [markdown, goldmark]
    ---------------------------

### Accessing front matter

You can use one of two ways to access front matter
parsed by goldmark-frontmatter:

* [Decode it into a struct with `frontmatter.Data`](#decode-a-struct)
* [Read from document metadata](#read-from-document-metadata)

#### Decode a struct

To decode front matter into a struct,
you must pass in a `parser.Context`
when you call `Markdown.Convert` or `Parser.Parse`.

```go
md := goldmark.New(
  goldmark.WithExtensions(&frontmatter.Extender{}),
  // ...
)

ctx := parser.NewContext()
md.Convert(src, out, parser.WithContext(ctx))
```

Following that, use `frontmatter.Get` to access a `frontmatter.Data` object.
Use `Data.Decode` to unmarshal your front matter into a data structure.

```go
d := frontmatter.Get(ctx)

var meta struct {
  Title string   `yaml:"title"`
  Tags  []string `yaml:"tags"`
  Desc  string   `yaml:"description"`
}
if err := d.Decode(&meta); err != nil {
  // ...
}
```

You're not limited to structs here.
You can also decode into `map[string]any` to access all fields.


```go
var meta map[string]any
if err := fm.Decode(&meta); err != nil {
  // ...
}
```

However, if you need that, it's easier to
[read it from the document metadata](#read-from-document-metadata).


#### Read from document metadata

You can install the extension with `frontmatter.SetMetadata` mode:

```go
md := goldmark.New(
  goldmark.WithExtensions(&frontmatter.Extender{
    Mode: frontmatter.SetMetadata,
  }),
  // ...
)
```

In this mode, the extension will decode the front matter
into a `map[string]any`,
and set it on the [Document](https://pkg.go.dev/github.com/yuin/goldmark/ast#Document).
You can access it with the [Document.Meta method](https://pkg.go.dev/github.com/yuin/goldmark/ast#Document.Meta).

```go
root := md.Parser().Parse(text.NewReader(src))
doc := root.OwnerDocument()
meta := doc.Meta()
```

## License

This software is made available under the BSD3 license.

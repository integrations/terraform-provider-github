package tfpluginv2util

import (
	"strconv"
	"strings"

	"github.com/hashicorp/go-cty/cty"
)

// Path converts a dot-separated string into a cty.Path, interpreting numeric segments as index steps and others as attribute steps.
func Path(pathStr string) cty.Path {
	if pathStr == "" {
		return nil
	}

	parts := strings.Split(pathStr, ".")
	path := make(cty.Path, len(parts))

	for i, part := range parts {
		if idx, err := strconv.Atoi(part); err == nil {
			path[i] = cty.IndexStep{Key: cty.NumberIntVal(int64(idx))}
		} else {
			path[i] = cty.GetAttrStep{Name: part}
		}
	}

	return path
}

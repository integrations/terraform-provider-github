package schema

import (
	"go/types"

	schemahelper "github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
)

const (
	TypeNameResourceImporter = `ResourceImporter`
)

// IsTypeResourceImporter returns if the type is ResourceImporter from the schema package.
func IsTypeResourceImporter(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return schemahelper.IsNamedType(t, TypeNameResourceImporter)
	case *types.Pointer:
		return IsTypeResourceImporter(t.Elem())
	default:
		return false
	}
}

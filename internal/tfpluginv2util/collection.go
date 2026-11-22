package tfpluginv2util

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// GetSetValue returns a typed slice for the specified key in a [schema.ResourceData] with a boolean representing if the key was set and a bool representing if the type was incorrect.
func GetSetValue[T any](d DataGetter, key string) ([]T, bool, bool) {
	valSet, ok, typeOk := GetValue[*schema.Set](d, key)
	if !ok || !typeOk {
		return nil, ok, typeOk
	}

	res := make([]T, valSet.Len())
	for i, v := range valSet.List() {
		vt, ok := v.(T)
		if !ok {
			return nil, true, false
		}

		res[i] = vt
	}

	return res, true, true
}

// GetSetOk returns a typed slice for the specified key in a [schema.ResourceData] with a boolean representing if the key was set. If an incorrect type is provided this will always return nil and false.
func GetSetOk[T any](d *schema.ResourceData, key string) ([]T, bool) {
	v, ok, typeOK := GetSetValue[T](d, key)
	if !ok || !typeOK {
		return v, false
	}

	return v, true
}

// GetSet returns a typed slice for the specified key in a [schema.ResourceData], if it's not set or the type is invalid an empty slice will be returned. If zeroAsEmpty is true then a zero value will be treated as an empty slice.
func GetSet[T any](d *schema.ResourceData, key string, zeroAsEmpty bool) []T {
	v, ok, typeOK := GetSetValue[T](d, key)
	if (ok && typeOK) || !zeroAsEmpty {
		return v
	}

	return make([]T, 0)
}

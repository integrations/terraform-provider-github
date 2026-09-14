package tfschemautil

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// GetValue returns a typed value for the specified key in a [schema.ResourceData] with a boolean representing if the key was set and a bool representing if the type was incorrect.
func GetValue[T any](d *schema.ResourceData, key string) (T, bool, bool) {
	var zero T

	v, ok := d.GetOk(key)
	if !ok {
		return zero, false, false
	}

	vt, ok := v.(T)
	if !ok {
		return zero, true, false
	}

	return vt, true, true
}

// GetOk returns a typed value for the specified key in a [schema.ResourceData] with a boolean representing if the key was set. If an incorrect type is provided this will always return the zero value and false.
func GetOk[T any](d *schema.ResourceData, key string) (T, bool) {
	v, ok, typeOK := GetValue[T](d, key)
	if !ok || !typeOK {
		return v, false
	}

	return v, true
}

// Get returns a typed value for the specified key in a [schema.ResourceData], if it's not set or the type is invalid the zero value will be returned.
func Get[T any](d *schema.ResourceData, key string) T {
	v, _, _ := GetValue[T](d, key)
	return v
}

// GetKeysOk is a helper function that checks multiple keys in the ResourceData and returns the first one that is set and a boolean indicating if any were set.
func GetKeysOk[T any](d *schema.ResourceData, keys ...string) (T, bool) {
	for _, key := range keys {
		v, ok, typeOK := GetValue[T](d, key)
		if ok && typeOK {
			return v, true
		}
	}

	var zero T
	return zero, false
}

// GetSetValue returns a typed slice for the specified key in a [schema.ResourceData] with a boolean representing if the key was set and a bool representing if the type was incorrect.
func GetSetValue[T any](d *schema.ResourceData, key string) ([]T, bool, bool) {
	val, ok := d.GetOk(key)
	if !ok {
		return nil, false, false
	}

	valSet, ok := val.(*schema.Set)
	if !ok {
		return nil, true, false
	}

	res := make([]T, valSet.Len())
	for i, v := range valSet.List() {
		vv, ok := v.(T)
		if !ok {
			return nil, true, false
		}

		res[i] = vv
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

package tfpluginv2util

// GetValue returns a typed value for the specified key in a [DataGetter] with a boolean representing if the key was set and a bool representing if the type was incorrect.
func GetValue[T any](d DataGetter, key string) (T, bool, bool) {
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

// GetOk returns a typed value for the specified key in a [DataGetter] with a boolean representing if the key was set. If an incorrect type is provided this will always return the zero value and false.
func GetOk[T any](d DataGetter, key string) (T, bool) {
	v, ok, typeOK := GetValue[T](d, key)
	if !ok || !typeOK {
		return v, false
	}

	return v, true
}

// Get returns a typed value for the specified key in a [DataGetter], if it's not set or the type is invalid the zero value will be returned.
func Get[T any](d DataGetter, key string) T {
	v, _, _ := GetValue[T](d, key)
	return v
}

// GetKeysOk is a helper function that checks multiple keys in a [DataGetter] and returns the first one that is set and a boolean indicating if any were set.
func GetKeysOk[T any](d DataGetter, keys ...string) (T, bool) {
	for _, key := range keys {
		v, ok, typeOK := GetValue[T](d, key)
		if ok && typeOK {
			return v, true
		}
	}

	var zero T
	return zero, false
}

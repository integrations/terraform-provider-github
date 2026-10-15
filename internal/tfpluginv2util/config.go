package tfpluginv2util

// IsSet checks if the given path in the resource data has been set.
func IsSet(d ConfigDataGetter, pathStr string) bool {
	if d == nil || pathStr == "" {
		return false
	}

	rawConfig := d.GetRawConfig()
	if rawConfig.IsNull() {
		return false
	}

	path := Path(pathStr)

	val, diag := d.GetRawConfigAt(path)
	if diag.HasError() {
		return false
	}

	return !val.IsNull()
}

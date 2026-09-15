package tfschemautil

type DataGetter interface {
	// GetOk returns the value for the specified key in a [schema.ResourceData] with a boolean representing if the key was set.
	GetOk(key string) (any, bool)
}

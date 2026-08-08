package github

import "fmt"

const (
	projectV2OwnerOrganization = "organization"
	projectV2OwnerUser         = "user"
)

type projectV2ValueGetter interface {
	Get(string) any
}

func projectV2Get[T any](data projectV2ValueGetter, key string) T {
	return projectV2As[T](data.Get(key), key)
}

func projectV2As[T any](value any, key string) T {
	typed, ok := value.(T)
	if !ok {
		panic(fmt.Sprintf("Projects V2 schema value %q has unexpected type %T", key, value))
	}
	return typed
}

func projectV2MapGet[T any](values map[string]any, key string) T {
	return projectV2As[T](values[key], key)
}

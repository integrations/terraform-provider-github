package tfpluginv2util

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// DataGetter is an interface that abstracts access to key-value data.
type DataGetter interface {
	// GetOk returns the value for the specified key with a boolean representing if the key was set.
	GetOk(key string) (any, bool)

	// Get returns the value for the specified key.
	Get(key string) any
}

// ConfigDataGetter is an interface that extends DataGetter to provide access to the raw configuration and specific paths within it.
type ConfigDataGetter interface {
	DataGetter

	GetRawConfig() cty.Value

	GetRawConfigAt(valPath cty.Path) (cty.Value, diag.Diagnostics)
}

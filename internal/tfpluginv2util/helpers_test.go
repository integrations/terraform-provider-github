package tfpluginv2util

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type testResourceData struct {
	*schema.ResourceData
	rawConfig cty.Value
}

func (d testResourceData) GetRawConfig() cty.Value {
	return d.rawConfig
}

func (d testResourceData) GetRawConfigAt(path cty.Path) (cty.Value, diag.Diagnostics) {
	value, err := path.Apply(d.rawConfig)
	if err != nil {
		return cty.DynamicVal, diag.Diagnostics{{
			Severity:      diag.Error,
			Summary:       "Invalid config path",
			Detail:        err.Error(),
			AttributePath: path,
		}}
	}

	return value, nil
}

// Code generated by "mapstructure-to-hcl2 -type OutputConfig"; DO NOT EDIT.
package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatOutputConfig is an auto-generated flat version of OutputConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatOutputConfig struct {
	OutputDir *string `mapstructure:"output_directory" required:"false" cty:"output_directory"`
}

// FlatMapstructure returns a new FlatOutputConfig.
// FlatOutputConfig is an auto-generated flat version of OutputConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*OutputConfig) FlatMapstructure() *FlatOutputConfig { return new(FlatOutputConfig) }

// HCL2Spec returns the hcl spec of a OutputConfig.
// This spec is used by HCL to read the fields of OutputConfig.
// The decoded values from this spec will then be applied to a FlatOutputConfig.
func (*OutputConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"output_directory": &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
	}
	return s
}

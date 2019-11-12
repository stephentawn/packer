package packer

import "github.com/hashicorp/hcl/v2/hcldec"

type HCL2Speccer interface {
	// ConfigSpec should return the hcl object spec used to configure the
	// builder. It will be used to tell the HCL parsing library how to
	// validate/configure a configuration.
	ConfigSpec() hcldec.ObjectSpec

	// FlatConfig returns a struct config config struct that is flattenned
	// based on mapstructure tags. The FlatConfig will then be loaded with that
	// was parsed using ConfigSpec and given back for setup.
	FlatConfig() interface{}
}

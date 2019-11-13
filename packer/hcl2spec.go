package packer

import "github.com/hashicorp/hcl/v2/hcldec"

type HCL2Speccer interface {
	// ConfigSpec should return the hcl object spec used to configure the
	// builder. It will be used to tell the HCL parsing library how to
	// validate/configure a configuration.
	ConfigSpec() hcldec.ObjectSpec
}

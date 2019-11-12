package hcl2template

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

type Decodable interface {
	HCL2Spec() map[string]hcldec.Spec
}

func decodeHCL2Spec(block *hcl.Block, ctx *hcl.EvalContext, dec Decodable) (cty.Value, hcl.Diagnostics) {
	spec := dec.HCL2Spec()
	return hcldec.Decode(block.Body, hcldec.ObjectSpec(spec), ctx)
}

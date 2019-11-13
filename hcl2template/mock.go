//go:generate mapstructure-to-hcl2 -type MockConfig,NestedMockConfig

package hcl2template

import (
	"context"
	"time"

	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
)

type NestedMockConfig struct {
	String          string            `mapstructure:"string"`
	Int             int               `mapstructure:"int"`
	Int64           int64             `mapstructure:"int64"`
	Bool            bool              `mapstructure:"bool"`
	Trilean         config.Trilean    `mapstructure:"trilean"`
	Duration        time.Duration     `mapstructure:"duration"`
	MapStringString map[string]string `mapstructure:"map_string_string"`
	SliceString     []string          `mapstructure:"slice_string"`
}

type MockConfig struct {
	NestedMockConfig `mapstructure:",squash"`
	Nested           NestedMockConfig   `mapstructure:"nested"`
	NestedSlice      []NestedMockConfig `mapstructure:"nested_slice"`
}

//////
// MockBuilder
//////

type MockBuilder struct {
	Config *MockConfig
}

var _ packer.Builder = new(MockBuilder)

func (b *MockBuilder) ConfigSpec() hcldec.ObjectSpec { return b.Config.HCL2Spec() }

func (b *MockBuilder) Prepare(raws ...interface{}) ([]string, error) {
	err := config.Decode(b.Config, &config.DecodeOpts{
		Interpolate: true,
	}, raws...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (b *MockBuilder) Run(ctx context.Context, ui packer.Ui, hook packer.Hook) (packer.Artifact, error) {
	return nil, nil
}

//////
// MockProvisioner
//////

type MockProvisioner struct {
	config *MockConfig
}

var _ packer.Provisioner = new(MockProvisioner)

func (b *MockProvisioner) ConfigSpec() hcldec.ObjectSpec { return b.config.HCL2Spec() }

func (b *MockProvisioner) Prepare(raws ...interface{}) error {
	return nil
}

func (b *MockProvisioner) Provision(ctx context.Context, ui packer.Ui, comm packer.Communicator) error {
	return nil
}

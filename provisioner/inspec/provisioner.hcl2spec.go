// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package inspec

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName      *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType    *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug          *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce          *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError        *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars       map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars  []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Command              *string           `cty:"command"`
	SubCommand           *string           `cty:"sub_command"`
	ExtraArguments       []string          `mapstructure:"extra_arguments" cty:"extra_arguments"`
	InspecEnvVars        []string          `mapstructure:"inspec_env_vars" cty:"inspec_env_vars"`
	Profile              *string           `mapstructure:"profile" cty:"profile"`
	AttributesDirectory  *string           `mapstructure:"attributes_directory" cty:"attributes_directory"`
	AttributesFiles      []string          `mapstructure:"attributes" cty:"attributes"`
	Backend              *string           `mapstructure:"backend" cty:"backend"`
	User                 *string           `mapstructure:"user" cty:"user"`
	Host                 *string           `mapstructure:"host" cty:"host"`
	LocalPort            *int              `mapstructure:"local_port" cty:"local_port"`
	SSHHostKeyFile       *string           `mapstructure:"ssh_host_key_file" cty:"ssh_host_key_file"`
	SSHAuthorizedKeyFile *string           `mapstructure:"ssh_authorized_key_file" cty:"ssh_authorized_key_file"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() *FlatConfig { return new(FlatConfig) }

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*Config) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"command":                    &hcldec.AttrSpec{Name: "command", Type: cty.String, Required: false},
		"sub_command":                &hcldec.AttrSpec{Name: "sub_command", Type: cty.String, Required: false},
		"extra_arguments":            &hcldec.AttrSpec{Name: "extra_arguments", Type: cty.List(cty.String), Required: false},
		"inspec_env_vars":            &hcldec.AttrSpec{Name: "inspec_env_vars", Type: cty.List(cty.String), Required: false},
		"profile":                    &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"attributes_directory":       &hcldec.AttrSpec{Name: "attributes_directory", Type: cty.String, Required: false},
		"attributes":                 &hcldec.AttrSpec{Name: "attributes", Type: cty.List(cty.String), Required: false},
		"backend":                    &hcldec.AttrSpec{Name: "backend", Type: cty.String, Required: false},
		"user":                       &hcldec.AttrSpec{Name: "user", Type: cty.String, Required: false},
		"host":                       &hcldec.AttrSpec{Name: "host", Type: cty.String, Required: false},
		"local_port":                 &hcldec.AttrSpec{Name: "local_port", Type: cty.Number, Required: false},
		"ssh_host_key_file":          &hcldec.AttrSpec{Name: "ssh_host_key_file", Type: cty.String, Required: false},
		"ssh_authorized_key_file":    &hcldec.AttrSpec{Name: "ssh_authorized_key_file", Type: cty.String, Required: false},
	}
	return s
}

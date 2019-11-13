// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package powershell

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName        *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType      *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug            *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce            *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError          *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars         map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars    []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	Binary                 *bool             `cty:"binary"`
	ExecuteCommand         *string           `mapstructure:"execute_command" cty:"execute_command"`
	Inline                 []string          `cty:"inline"`
	RemotePath             *string           `mapstructure:"remote_path" cty:"remote_path"`
	Script                 *string           `cty:"script"`
	Scripts                []string          `cty:"scripts"`
	ValidExitCodes         []int             `mapstructure:"valid_exit_codes" cty:"valid_exit_codes"`
	Vars                   []string          `mapstructure:"environment_vars" cty:"environment_vars"`
	EnvVarFormat           *string           `mapstructure:"env_var_format" cty:"env_var_format"`
	RemoteEnvVarPath       *string           `mapstructure:"remote_env_var_path" cty:"remote_env_var_path"`
	ElevatedExecuteCommand *string           `mapstructure:"elevated_execute_command" cty:"elevated_execute_command"`
	StartRetryTimeout      *string           `mapstructure:"start_retry_timeout" cty:"start_retry_timeout"`
	ElevatedEnvVarFormat   *string           `mapstructure:"elevated_env_var_format" cty:"elevated_env_var_format"`
	ElevatedUser           *string           `mapstructure:"elevated_user" cty:"elevated_user"`
	ElevatedPassword       *string           `mapstructure:"elevated_password" cty:"elevated_password"`
	ExecutionPolicy        *string           `mapstructure:"execution_policy" cty:"execution_policy"`
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
		"binary":                     &hcldec.AttrSpec{Name: "binary", Type: cty.Bool, Required: false},
		"execute_command":            &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"inline":                     &hcldec.AttrSpec{Name: "inline", Type: cty.List(cty.String), Required: false},
		"remote_path":                &hcldec.AttrSpec{Name: "remote_path", Type: cty.String, Required: false},
		"script":                     &hcldec.AttrSpec{Name: "script", Type: cty.String, Required: false},
		"scripts":                    &hcldec.AttrSpec{Name: "scripts", Type: cty.List(cty.String), Required: false},
		"valid_exit_codes":           &hcldec.AttrSpec{Name: "valid_exit_codes", Type: cty.List(cty.Number), Required: false},
		"environment_vars":           &hcldec.AttrSpec{Name: "environment_vars", Type: cty.List(cty.String), Required: false},
		"env_var_format":             &hcldec.AttrSpec{Name: "env_var_format", Type: cty.String, Required: false},
		"remote_env_var_path":        &hcldec.AttrSpec{Name: "remote_env_var_path", Type: cty.String, Required: false},
		"elevated_execute_command":   &hcldec.AttrSpec{Name: "elevated_execute_command", Type: cty.String, Required: false},
		"start_retry_timeout":        &hcldec.AttrSpec{Name: "start_retry_timeout", Type: cty.String, Required: false},
		"elevated_env_var_format":    &hcldec.AttrSpec{Name: "elevated_env_var_format", Type: cty.String, Required: false},
		"elevated_user":              &hcldec.AttrSpec{Name: "elevated_user", Type: cty.String, Required: false},
		"elevated_password":          &hcldec.AttrSpec{Name: "elevated_password", Type: cty.String, Required: false},
		"execution_policy":           &hcldec.AttrSpec{Name: "execution_policy", Type: cty.String, Required: false},
	}
	return s
}

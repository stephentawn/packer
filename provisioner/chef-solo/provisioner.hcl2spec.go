// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package chefsolo

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName            *string                `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType          *string                `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug                *bool                  `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce                *bool                  `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError              *string                `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars             map[string]string      `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars        []string               `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	ChefEnvironment            *string                `mapstructure:"chef_environment" cty:"chef_environment"`
	ChefLicense                *string                `mapstructure:"chef_license" cty:"chef_license"`
	ConfigTemplate             *string                `mapstructure:"config_template" cty:"config_template"`
	CookbookPaths              []string               `mapstructure:"cookbook_paths" cty:"cookbook_paths"`
	RolesPath                  *string                `mapstructure:"roles_path" cty:"roles_path"`
	DataBagsPath               *string                `mapstructure:"data_bags_path" cty:"data_bags_path"`
	EncryptedDataBagSecretPath *string                `mapstructure:"encrypted_data_bag_secret_path" cty:"encrypted_data_bag_secret_path"`
	EnvironmentsPath           *string                `mapstructure:"environments_path" cty:"environments_path"`
	ExecuteCommand             *string                `mapstructure:"execute_command" cty:"execute_command"`
	InstallCommand             *string                `mapstructure:"install_command" cty:"install_command"`
	RemoteCookbookPaths        []string               `mapstructure:"remote_cookbook_paths" cty:"remote_cookbook_paths"`
	Json                       map[string]interface{} `cty:"json"`
	PreventSudo                *bool                  `mapstructure:"prevent_sudo" cty:"prevent_sudo"`
	RunList                    []string               `mapstructure:"run_list" cty:"run_list"`
	SkipInstall                *bool                  `mapstructure:"skip_install" cty:"skip_install"`
	StagingDir                 *string                `mapstructure:"staging_directory" cty:"staging_directory"`
	GuestOSType                *string                `mapstructure:"guest_os_type" cty:"guest_os_type"`
	Version                    *string                `mapstructure:"version" cty:"version"`
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
		"packer_build_name":              &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":            &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                   &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                   &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":                &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":          &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":     &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"chef_environment":               &hcldec.AttrSpec{Name: "chef_environment", Type: cty.String, Required: false},
		"chef_license":                   &hcldec.AttrSpec{Name: "chef_license", Type: cty.String, Required: false},
		"config_template":                &hcldec.AttrSpec{Name: "config_template", Type: cty.String, Required: false},
		"cookbook_paths":                 &hcldec.AttrSpec{Name: "cookbook_paths", Type: cty.List(cty.String), Required: false},
		"roles_path":                     &hcldec.AttrSpec{Name: "roles_path", Type: cty.String, Required: false},
		"data_bags_path":                 &hcldec.AttrSpec{Name: "data_bags_path", Type: cty.String, Required: false},
		"encrypted_data_bag_secret_path": &hcldec.AttrSpec{Name: "encrypted_data_bag_secret_path", Type: cty.String, Required: false},
		"environments_path":              &hcldec.AttrSpec{Name: "environments_path", Type: cty.String, Required: false},
		"execute_command":                &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"install_command":                &hcldec.AttrSpec{Name: "install_command", Type: cty.String, Required: false},
		"remote_cookbook_paths":          &hcldec.AttrSpec{Name: "remote_cookbook_paths", Type: cty.List(cty.String), Required: false},
		"json":                           &hcldec.BlockAttrsSpec{TypeName: "json", ElementType: cty.String, Required: false},
		"prevent_sudo":                   &hcldec.AttrSpec{Name: "prevent_sudo", Type: cty.Bool, Required: false},
		"run_list":                       &hcldec.AttrSpec{Name: "run_list", Type: cty.List(cty.String), Required: false},
		"skip_install":                   &hcldec.AttrSpec{Name: "skip_install", Type: cty.Bool, Required: false},
		"staging_directory":              &hcldec.AttrSpec{Name: "staging_directory", Type: cty.String, Required: false},
		"guest_os_type":                  &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"version":                        &hcldec.AttrSpec{Name: "version", Type: cty.String, Required: false},
	}
	return s
}

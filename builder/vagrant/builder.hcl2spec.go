// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package vagrant

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	HTTPDir                   *string           `mapstructure:"http_directory" cty:"http_directory"`
	HTTPPortMin               *int              `mapstructure:"http_port_min" cty:"http_port_min"`
	HTTPPortMax               *int              `mapstructure:"http_port_max" cty:"http_port_max"`
	ISOChecksum               *string           `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum"`
	ISOChecksumURL            *string           `mapstructure:"iso_checksum_url" cty:"iso_checksum_url"`
	ISOChecksumType           *string           `mapstructure:"iso_checksum_type" cty:"iso_checksum_type"`
	RawSingleISOUrl           *string           `mapstructure:"iso_url" required:"true" cty:"iso_url"`
	ISOUrls                   []string          `mapstructure:"iso_urls" cty:"iso_urls"`
	TargetPath                *string           `mapstructure:"iso_target_path" cty:"iso_target_path"`
	TargetExtension           *string           `mapstructure:"iso_target_extension" cty:"iso_target_extension"`
	FloppyFiles               []string          `mapstructure:"floppy_files" cty:"floppy_files"`
	FloppyDirectories         []string          `mapstructure:"floppy_dirs" cty:"floppy_dirs"`
	FloppyLabel               *string           `mapstructure:"floppy_label" cty:"floppy_label"`
	BootGroupInterval         *string           `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval"`
	BootWait                  *string           `mapstructure:"boot_wait" cty:"boot_wait"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" cty:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" cty:"temporary_key_pair_name"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" cty:"ssh_private_key_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" cty:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `cty:"ssh_public_key"`
	SSHPrivateKey             []byte            `cty:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm"`
	OutputDir                 *string           `mapstructure:"output_dir" required:"false" cty:"output_dir"`
	SourceBox                 *string           `mapstructure:"source_path" required:"true" cty:"source_path"`
	GlobalID                  *string           `mapstructure:"global_id" required:"true" cty:"global_id"`
	Checksum                  *string           `mapstructure:"checksum" required:"false" cty:"checksum"`
	ChecksumType              *string           `mapstructure:"checksum_type" required:"false" cty:"checksum_type"`
	BoxName                   *string           `mapstructure:"box_name" required:"false" cty:"box_name"`
	InsertKey                 *bool             `mapstructure:"insert_key" required:"false" cty:"insert_key"`
	Provider                  *string           `mapstructure:"provider" required:"false" cty:"provider"`
	VagrantfileTpl            *string           `mapstructure:"vagrantfile_template" cty:"vagrantfile_template"`
	TeardownMethod            *string           `mapstructure:"teardown_method" required:"false" cty:"teardown_method"`
	BoxVersion                *string           `mapstructure:"box_version" required:"false" cty:"box_version"`
	Template                  *string           `mapstructure:"template" required:"false" cty:"template"`
	SyncedFolder              *string           `mapstructure:"synced_folder" cty:"synced_folder"`
	SkipAdd                   *bool             `mapstructure:"skip_add" required:"false" cty:"skip_add"`
	AddCACert                 *string           `mapstructure:"add_cacert" required:"false" cty:"add_cacert"`
	AddCAPath                 *string           `mapstructure:"add_capath" required:"false" cty:"add_capath"`
	AddCert                   *string           `mapstructure:"add_cert" required:"false" cty:"add_cert"`
	AddClean                  *bool             `mapstructure:"add_clean" required:"false" cty:"add_clean"`
	AddForce                  *bool             `mapstructure:"add_force" required:"false" cty:"add_force"`
	AddInsecure               *bool             `mapstructure:"add_insecure" required:"false" cty:"add_insecure"`
	SkipPackage               *bool             `mapstructure:"skip_package" required:"false" cty:"skip_package"`
	OutputVagrantfile         *string           `mapstructure:"output_vagrantfile" cty:"output_vagrantfile"`
	PackageInclude            []string          `mapstructure:"package_include" cty:"package_include"`
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
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":               &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"iso_checksum":                 &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_checksum_url":             &hcldec.AttrSpec{Name: "iso_checksum_url", Type: cty.String, Required: false},
		"iso_checksum_type":            &hcldec.AttrSpec{Name: "iso_checksum_type", Type: cty.String, Required: false},
		"iso_url":                      &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_urls":                     &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_target_path":              &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"iso_target_extension":         &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"floppy_files":                 &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"floppy_dirs":                  &hcldec.AttrSpec{Name: "floppy_dirs", Type: cty.List(cty.String), Required: false},
		"floppy_label":                 &hcldec.AttrSpec{Name: "floppy_label", Type: cty.String, Required: false},
		"boot_keygroup_interval":       &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"boot_wait":                    &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"boot_command":                 &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"output_dir":                   &hcldec.AttrSpec{Name: "output_dir", Type: cty.String, Required: false},
		"source_path":                  &hcldec.AttrSpec{Name: "source_path", Type: cty.String, Required: false},
		"global_id":                    &hcldec.AttrSpec{Name: "global_id", Type: cty.String, Required: false},
		"checksum":                     &hcldec.AttrSpec{Name: "checksum", Type: cty.String, Required: false},
		"checksum_type":                &hcldec.AttrSpec{Name: "checksum_type", Type: cty.String, Required: false},
		"box_name":                     &hcldec.AttrSpec{Name: "box_name", Type: cty.String, Required: false},
		"insert_key":                   &hcldec.AttrSpec{Name: "insert_key", Type: cty.Bool, Required: false},
		"provider":                     &hcldec.AttrSpec{Name: "provider", Type: cty.String, Required: false},
		"vagrantfile_template":         &hcldec.AttrSpec{Name: "vagrantfile_template", Type: cty.String, Required: false},
		"teardown_method":              &hcldec.AttrSpec{Name: "teardown_method", Type: cty.String, Required: false},
		"box_version":                  &hcldec.AttrSpec{Name: "box_version", Type: cty.String, Required: false},
		"template":                     &hcldec.AttrSpec{Name: "template", Type: cty.String, Required: false},
		"synced_folder":                &hcldec.AttrSpec{Name: "synced_folder", Type: cty.String, Required: false},
		"skip_add":                     &hcldec.AttrSpec{Name: "skip_add", Type: cty.Bool, Required: false},
		"add_cacert":                   &hcldec.AttrSpec{Name: "add_cacert", Type: cty.String, Required: false},
		"add_capath":                   &hcldec.AttrSpec{Name: "add_capath", Type: cty.String, Required: false},
		"add_cert":                     &hcldec.AttrSpec{Name: "add_cert", Type: cty.String, Required: false},
		"add_clean":                    &hcldec.AttrSpec{Name: "add_clean", Type: cty.Bool, Required: false},
		"add_force":                    &hcldec.AttrSpec{Name: "add_force", Type: cty.Bool, Required: false},
		"add_insecure":                 &hcldec.AttrSpec{Name: "add_insecure", Type: cty.Bool, Required: false},
		"skip_package":                 &hcldec.AttrSpec{Name: "skip_package", Type: cty.Bool, Required: false},
		"output_vagrantfile":           &hcldec.AttrSpec{Name: "output_vagrantfile", Type: cty.String, Required: false},
		"package_include":              &hcldec.AttrSpec{Name: "package_include", Type: cty.List(cty.String), Required: false},
	}
	return s
}

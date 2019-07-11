// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package googlecomputeexport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	AccountFile         *string           `mapstructure:"account_file" cty:"account_file"`
	DiskSizeGb          *int64            `mapstructure:"disk_size" cty:"disk_size"`
	DiskType            *string           `mapstructure:"disk_type" cty:"disk_type"`
	MachineType         *string           `mapstructure:"machine_type" cty:"machine_type"`
	Network             *string           `mapstructure:"network" cty:"network"`
	Paths               []string          `mapstructure:"paths" cty:"paths"`
	Subnetwork          *string           `mapstructure:"subnetwork" cty:"subnetwork"`
	VaultGCPOauthEngine *string           `mapstructure:"vault_gcp_oauth_engine" cty:"vault_gcp_oauth_engine"`
	Zone                *string           `mapstructure:"zone" cty:"zone"`
	ServiceAccountEmail *string           `mapstructure:"service_account_email" cty:"service_account_email"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{} { return new(FlatConfig) }

// HCL2Spec returns the hcldec.Spec of a FlatConfig.
// This spec is used by HCL to read the fields of FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"account_file":               &hcldec.AttrSpec{Name: "account_file", Type: cty.String, Required: false},
		"disk_size":                  &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"disk_type":                  &hcldec.AttrSpec{Name: "disk_type", Type: cty.String, Required: false},
		"machine_type":               &hcldec.AttrSpec{Name: "machine_type", Type: cty.String, Required: false},
		"network":                    &hcldec.AttrSpec{Name: "network", Type: cty.String, Required: false},
		"paths":                      &hcldec.AttrSpec{Name: "paths", Type: cty.List(cty.String), Required: false},
		"subnetwork":                 &hcldec.AttrSpec{Name: "subnetwork", Type: cty.String, Required: false},
		"vault_gcp_oauth_engine":     &hcldec.AttrSpec{Name: "vault_gcp_oauth_engine", Type: cty.String, Required: false},
		"zone":                       &hcldec.AttrSpec{Name: "zone", Type: cty.String, Required: false},
		"service_account_email":      &hcldec.AttrSpec{Name: "service_account_email", Type: cty.String, Required: false},
	}
	return s
}
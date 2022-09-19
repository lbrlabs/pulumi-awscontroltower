package awscontroltower

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	controlTowerShim "github.com/idealo/terraform-provider-controltower/shim"
	"github.com/lbrlabs/pulumi-awscontroltower/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	controlTowerPkg = "awscontroltower"
	// modules:
	controlTowerMod = "index" // the xyz module
)

var namespaceMap = map[string]string{
	controlTowerPkg: "AwsControlTower",
}

func makeMember(moduleTitle string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(moduleTitle)
	namespaceMap[moduleName] = moduleTitle
	fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + fn
	return tokens.ModuleMember(controlTowerPkg + ":" + token + ":" + mem)
}

func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

func makeResource(mod string, res string) tokens.Type {
	return makeType(mod, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(controlTowerShim.NewProvider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "awscontroltower",
		Description:       "A Pulumi package for creating and managing control tower accounts.",
		Keywords:          []string{"pulumi", "aws", "controltower", "lbrlabs"},
		License:           "Apache-2.0",
		Homepage:          "https://leebriggs.co.uk/projects#pulumi-awscontroltower",
		Repository:        "https://github.com/lbrlabs/pulumi-awscontroltower",
		PluginDownloadURL: "github://api.github.com/lbrlabs",
		GitHubOrg:         "idealo", // not in the terraform-providers repo
		Publisher:         "lbrlabs",
		DisplayName:       "AWS Control Tower",
		Config:            map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"controltower_aws_account": {Tok: makeResource(controlTowerMod, "controlTowerAwsAccount")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			PackageName: "@lbrlabs/pulumi-awscontroltower",
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PackageName: "lbrlabs_pulumi_aws_controltower",
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/lbrlabs/pulumi-%[1]s/sdk/", controlTowerPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				controlTowerPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RootNamespace: "Lbrlabs.PulumiPackage",
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}

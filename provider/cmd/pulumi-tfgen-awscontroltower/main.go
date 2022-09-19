package main

import (
	awscontroltower "github.com/lbrlabs/pulumi-awscontroltower/provider"
	"github.com/lbrlabs/pulumi-awscontroltower/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
)

func main() {
	// Modify the path to point to the new provider
	tfgen.Main("awscontroltower", version.Version, awscontroltower.Provider())
}

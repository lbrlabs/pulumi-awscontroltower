package main

import (
	awscontroltower "github.com/lbrlabs/pulumi-awscontroltower/provider"
	"github.com/lbrlabs/pulumi-awscontroltower/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

func main() {
	// Modify the path to point to the new provider
	tfbridge.Main("awscontroltower", version.Version, awscontroltower.Provider(), pulumiSchema)
}

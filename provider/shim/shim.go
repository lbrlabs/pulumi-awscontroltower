package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/lbrlabs/pulumi-awscontroltower/provider/pkg/version"
	"github.com/idealo/terraform-provider-controltower/internal/provider"
)

func NewProvider() *schema.Provider {
	prov := provider.New(version.Version)()
	return prov
}
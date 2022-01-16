package shim

import (
	"github.com/go-kit/kit/metrics/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/idealo/terraform-provider-controltower/internal/provider"
)

func NewProvider() *schema.Provider {
	return provider.New()
}

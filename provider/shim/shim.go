package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/idealo/terraform-provider-controltower/internal/provider"
)

func NewProvider() *schema.Provider {
	return provider.New("dev")()
}

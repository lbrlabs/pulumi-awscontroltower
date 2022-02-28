module github.com/idealo/terraform-provider-controltower/shim

go 1.17

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.9.0
	github.com/idealo/terraform-provider-controltower v1.1.0
)

require (
	github.com/go-kit/kit v0.12.0
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/terraform-plugin-go v0.5.0 // indirect
)

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019194827-62530c6537a4

module github.com/jaxxstorm/pulumi-awscontroltower/provider

go 1.16

require (
	github.com/idealio/terraform-provider-controltower/shim v0.0.0
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.16.0
	github.com/pulumi/pulumi/sdk/v3 v3.22.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
	github.com/idealio/terraform-provider-controltower/shim => ./shim
)

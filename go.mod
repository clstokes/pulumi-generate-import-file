module github.com/clstokes/pulumi-generate-import-file

go 1.16

require (
	github.com/pulumi/pulumi-aws/provider/v4 v4.0.0-20210920231643-9f1e8dcb314c
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.7.0
	github.com/pulumi/pulumi-vault/provider/v4 v4.0.0-20210903135450-8b58a54bc872
)

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
	github.com/hashicorp/terraform-plugin-test => github.com/hashicorp/terraform-plugin-test v1.3.0
	github.com/pulumi/pulumi-terraform-bridge/v3 => github.com/pulumi/pulumi-terraform-bridge/v3 v3.7.0
	github.com/terraform-providers/terraform-provider-aws => github.com/pulumi/terraform-provider-aws v1.38.1-0.20210919140801-9f75bddc51fd
)

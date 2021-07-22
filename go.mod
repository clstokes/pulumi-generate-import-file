module github.com/clstokes/pulumi-generate-import-file

go 1.16

require (
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.4.0
	github.com/pulumi/pulumi-vault/provider/v4 v4.0.0-20210712133926-d34f23ec5472
// github.com/pulumi/pulumi-vault v4.0.0-20210712133926-d34f23ec5472
)

replace github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
replace github.com/hashicorp/terraform-plugin-test => github.com/hashicorp/terraform-plugin-test v1.3.0

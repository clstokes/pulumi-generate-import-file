ensure::
	go mod download

build::
	@# We need this ldflag to pass a version to the provider or the provider will panic
	@go run -ldflags '-X github.com/pulumi/pulumi-vault/provider/v4/pkg/version.Version=v4.0.0 -X github.com/pulumi/pulumi-aws/provider/v4/pkg/version.Version=v4.0.0' .

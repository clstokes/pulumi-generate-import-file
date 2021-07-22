# pulumi-generate-import-file

[Pulumi](https://www.pulumi.com/docs/get-started/install/) supports a few ways for [importing existing cloud resources](https://www.pulumi.com/docs/guides/adopting/import/) into your Pulumi stacks. This program will allow you to take an existing Terraform state file and generate a file that the [pulumi import](./pulumi-generate-import-file examples/terraform-vault/terraform.tfstate) CLI command can use to import those resources and generate the corresponding Go, TypeScript, etc. code for the imported resources.

## Usage

1. Build the tool with `make build`
1. Execute the tool with `./pulumi-generate-import-file ~/path/to/terraform.tfstate`
1. Copy the output to a `resources.json` or similarly named file
    - you can also pipe the output directly to a file instead of copying/pasting it
    - e.g. `./pulumi-generate-import-file ~/path/to/terraform.tfstate > resources.json`
1. From your Pulumi project run `pulumi import --file ~/path/to/resources.json`

## Important Notes

- Only v4 terraform state files are supported
- The names of the resource imports will be different than the original Terraform resource name. The `pulumi-generate-import-file` binary will add a prefix to the resource name with the resource type and add a suffix of the resource index. This is primarily due to https://github.com/pulumi/pulumi/issues/6032. [Resource aliases](https://www.pulumi.com/docs/intro/concepts/resources/#aliases) can be used to refactor these names after import.

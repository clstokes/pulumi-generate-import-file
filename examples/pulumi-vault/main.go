package main

import (
	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault"
	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault/generic"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := generic.NewSecret(ctx, "vault_generic_secret_example0", &generic.SecretArgs{
			DataJson:    pulumi.ToSecret("{\"foo\":\"bar\",\"pizza\":\"cheese\"}").(pulumi.StringOutput),
			DisableRead: pulumi.Bool(false),
			Path:        pulumi.String("secret/foo-0"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = generic.NewSecret(ctx, "vault_generic_secret_example1", &generic.SecretArgs{
			DataJson:    pulumi.ToSecret("{\"foo\":\"bar\",\"pizza\":\"cheese\"}").(pulumi.StringOutput),
			DisableRead: pulumi.Bool(false),
			Path:        pulumi.String("secret/foo-1"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = vault.NewPolicy(ctx, "vault_policy_example0", &vault.PolicyArgs{
			Name:   pulumi.String("dev-team-0"),
			Policy: pulumi.String("path \"secret/my_app\" {\n  capabilities = [\"update\"]\n}\n"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		_, err = vault.NewPolicy(ctx, "vault_policy_example1", &vault.PolicyArgs{
			Name:   pulumi.String("dev-team-1"),
			Policy: pulumi.String("path \"secret/my_app\" {\n  capabilities = [\"update\"]\n}\n"),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}
		return nil
	})
}

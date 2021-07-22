import * as pulumi from "@pulumi/pulumi";
import * as vault from "@pulumi/vault";

const vault_generic_secret_example0 = new vault.generic.Secret("vault_generic_secret_example0", {
    dataJson: pulumi.secret("{\"foo\":\"bar\",\"pizza\":\"cheese\"}"),
    disableRead: false,
    path: "secret/foo-0",
}, {
    protect: true,
});
const vault_generic_secret_example1 = new vault.generic.Secret("vault_generic_secret_example1", {
    dataJson: pulumi.secret("{\"foo\":\"bar\",\"pizza\":\"cheese\"}"),
    disableRead: false,
    path: "secret/foo-1",
}, {
    protect: true,
});
const vault_policy_example0 = new vault.Policy("vault_policy_example0", {
    name: "dev-team-0",
    policy: `path "secret/my_app" {
  capabilities = ["update"]
}
`,
}, {
    protect: true,
});
const vault_policy_example1 = new vault.Policy("vault_policy_example1", {
    name: "dev-team-1",
    policy: `path "secret/my_app" {
  capabilities = ["update"]
}
`,
}, {
    protect: true,
});

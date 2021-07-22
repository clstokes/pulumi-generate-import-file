resource "vault_policy" "example" {
  count = 2
  name = "dev-team-${count.index}"

  policy = <<EOT
path "secret/my_app" {
  capabilities = ["update"]
}
EOT
}

resource "vault_generic_secret" "example" {
  count = 2
  path = "secret/foo-${count.index}"

  data_json = <<EOT
{
  "foo":   "bar",
  "pizza": "cheese"
}
EOT
}

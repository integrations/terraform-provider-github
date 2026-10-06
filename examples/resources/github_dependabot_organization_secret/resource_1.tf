# Un-encrypted Secret Example

resource "github_dependabot_organization_secret" "example" {
  secret_name = "EXAMPLE_SECRET_NAME"
  value       = "example-value"
  visibility  = "all"
}

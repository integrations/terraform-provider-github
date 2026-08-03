# Un-encrypted Secret Example

resource "github_dependabot_secret" "example" {
  repository  = "example-repo"
  secret_name = "EXAMPLE_SECRET_NAME"
  value       = "example-value"
}

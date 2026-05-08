# Retrieve information about multiple GitHub users.
data "github_users" "example" {
  usernames = ["example1", "example2", "example3"]
}

output "valid_users" {
  value = data.github_users.example.logins
}

output "invalid_users" {
  value = data.github_users.example.unknown_logins
}

# Retrieve information about a GitHub user.
data "github_user" "example" {
  username = "example"
}

# Retrieve information about the currently authenticated user.
data "github_user" "current" {
  username = ""
}

output "current_github_login" {
  value = data.github_user.current.login
}


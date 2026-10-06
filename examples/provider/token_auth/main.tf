provider "github" {
  owner = "octocat"

  auth_mode = "token" # or `GITHUB_AUTH_MODE=token`

  token = var.token # or `GITHUB_TOKEN`
}

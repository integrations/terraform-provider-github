resource "github_actions_organization_secret" "example" {
  secret_name     = "mysecret"
  plaintext_value = "foo"
  visibility      = "selected"
}

resource "github_repository" "example" {
  name       = "myrepo"
  visibility = "public"
}

resource "github_actions_organization_secret_repository" "example" {
  secret_name   = github_actions_organization_secret.example.name
  repository_id = github_repository.example.repo_id
}

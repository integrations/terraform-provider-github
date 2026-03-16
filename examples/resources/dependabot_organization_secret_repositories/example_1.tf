resource "github_dependabot_organization_secret" "example" {
	secret_name     = "mysecret"
	plaintext_value = "foo"
	visibility      = "selected"
}

resource "github_repository" "example" {
	name       = "myrepo"
	visibility = "public"
}

resource "github_dependabot_organization_secret_repositories" "example" {
  secret_name             = github_dependabot_organization_secret.example.name
  selected_repository_ids = [ github_repository.example.repo_id ]
}

resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_organization_immutable_releases" "example" {
  enforced_repositories   = "selected"
  selected_repository_ids = [github_repository.example.repo_id]
}

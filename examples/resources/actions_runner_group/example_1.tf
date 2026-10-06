resource "github_repository" "example" {
  name = "my-repository"
}

resource "github_actions_runner_group" "example" {
  name                    = github_repository.example.name
  visibility              = "selected"
  selected_repository_ids = [github_repository.example.repo_id]
}

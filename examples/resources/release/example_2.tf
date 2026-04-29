resource "github_repository" "example" {
  name      = "repo"
  auto_init = true
}

resource "github_branch" "example" {
  repository    = github_repository.example.name
  branch        = "branch_name"
  source_branch = github_repository.example.default_branch
}

resource "github_release" "example" {
  repository       = github_repository.example.name
  tag_name         = "v1.0.0"
  target_commitish = github_branch.example.branch
  draft            = false
  prerelease       = false
}

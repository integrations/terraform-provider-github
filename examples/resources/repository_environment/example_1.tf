data "github_user" "current" {
  username = ""
}

resource "github_repository" "example" {
  name        = "A Repository Project"
  description = "My awesome codebase"
}

resource "github_repository_environment" "example" {
  environment         = "example"
  repository          = github_repository.example.name
  prevent_self_review = true
  reviewers {
    users = [data.github_user.current.id]
  }
  deployment_branch_policy {
    protected_branches     = true
    custom_branch_policies = false
  }
}

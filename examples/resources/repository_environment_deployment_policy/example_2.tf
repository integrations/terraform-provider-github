
data "github_user" "current" {
  username = ""
}

resource "github_repository" "test" {
  name = "tf-acc-test-%s"
}

resource "github_repository_environment" "test" {
  repository  = github_repository.test.name
  environment = "environment/test"
  wait_timer  = 10000
  reviewers {
    users = [data.github_user.current.id]
  }
  deployment_branch_policy {
    protected_branches     = false
    custom_branch_policies = true
  }
}

resource "github_repository_environment_deployment_policy" "test" {
  repository  = github_repository.test.name
  environment = github_repository_environment.test.environment
  tag_pattern = "v*"
}

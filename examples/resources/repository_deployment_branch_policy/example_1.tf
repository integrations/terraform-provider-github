resource "github_repository_environment" "env" {
  repository  = "my_repo"
  environment = "my_env"
  deployment_branch_policy {
    protected_branches     = false
    custom_branch_policies = true
  }
}

resource "github_repository_deployment_branch_policy" "foo" {
  depends_on = [github_repository_environment.env]

  repository       = "my_repo"
  environment_name = "my_env"
  name             = "foo"
}

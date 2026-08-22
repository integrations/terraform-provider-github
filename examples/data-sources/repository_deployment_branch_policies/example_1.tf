data "github_repository_deployment_branch_policies" "example" {
  repository       = "example-repository"
  environment_name = "env_name"
}

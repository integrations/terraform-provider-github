resource "github_repository" "terraformed" {
  name        = "terraformed"
  description = "A repository created by Terraform"
}

resource "github_repository_environment" "terraformed" {
  repository = github_repository.terraformed.name
  environment = "terraformed"  
  deployment_branch_policy {
    protected_branches          = false
    custom_branch_policies = true
  }
}

resource "github_repository_environment_deployment_policy" "terraformed" {
  repository = github_repository.terraformed.name
  environment = github_repository_environment.terraformed.environment
  branch_pattern = "terraformed/*"
}
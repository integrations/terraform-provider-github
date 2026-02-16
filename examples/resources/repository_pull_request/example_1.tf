resource "github_repository_pull_request" "example" {
  base_repository = "example-repository"
  base_ref        = "main"
  head_ref        = "feature-branch"
  title           = "My newest feature"
  body            = "This will change everything"
}

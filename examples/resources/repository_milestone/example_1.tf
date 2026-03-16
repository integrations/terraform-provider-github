# Create a milestone for a repository
resource "github_repository_milestone" "example" {
  owner      = "example-owner"
  repository = "example-repository"
  title      = "v1.1.0"
}

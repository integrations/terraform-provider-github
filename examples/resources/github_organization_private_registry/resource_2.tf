# Username and password registry, restricted to selected repositories.
resource "github_repository" "example" {
  name = "example"
}

resource "github_organization_private_registry" "example" {
  registry_type           = "maven_repository"
  url                     = "https://maven.example.com/releases"
  auth_type               = "username_password"
  username                = "deploy"
  value                   = var.maven_password
  visibility              = "selected"
  selected_repository_ids = [github_repository.example.repo_id]
}

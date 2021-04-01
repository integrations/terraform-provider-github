resource "github_repository" "app_installation_example" {
  name        = "appinstallationexample"
  description = "A repository to install an application in."
}

resource "github_app_installation_repository" "test"{
    repository      = github_repository.app_installation_example.name
    installation_id = var.installation_id
}

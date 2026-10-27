resource "github_repository" "example" {
  name = "example"
}

resource "github_repository_custom_properties" "example" {
  repository = github_repository.example.name

  property {
    name  = "environment"
    value = ["production"]
  }

  property {
    name  = "team"
    value = ["platform"]
  }
}

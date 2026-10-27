resource "github_repository_custom_properties" "example" {
  repository = "my-repo"

  property {
    name  = "languages"
    value = ["go", "typescript", "python"]
  }

  property {
    name  = "environment"
    value = ["staging"]
  }
}

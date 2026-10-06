# NOTE: This assumes there already is a custom property defined on the org level called `my-cool-string` of type `string`

resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"
}

resource "github_repository_custom_property" "example" {
  repository     = github_repository.example.name
  property_name  = "my-cool-string"
  property_type  = "string"
  property_value = ["test"]
}

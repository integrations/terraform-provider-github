# NOTE: This assumes there already is a custom property defined on the org level called `my-cool-property` of type `string`

resource "github_repository" "example" {
  name        = "example"
  description = "My awesome codebase"
}

resource "github_repository_custom_property" "string" {
  repository     = github_repository.example.name
  property_name  = "my-cool-property"
  property_type  = "string"
  property_value = ["test"]
}

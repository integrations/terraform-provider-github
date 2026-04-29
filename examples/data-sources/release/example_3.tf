data "github_release" "example" {
  repository  = "example-repository"
  owner       = "example-owner"
  retrieve_by = "tag"
  release_tag = "v1.0.0"
}

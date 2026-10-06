data "github_release" "example" {
  repository  = "example-repository"
  owner       = "example-owner"
  retrieve_by = "latest"
}

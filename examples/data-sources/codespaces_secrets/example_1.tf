data "github_codespaces_secrets" "example" {
  name = "example_repository"
}

data "github_codespaces_secrets" "example_2" {
  full_name = "org/example_repository"
}

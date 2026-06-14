resource "github_enterprise_app_installation_repositories" "example" {
  enterprise_slug       = "my-enterprise"
  organization          = "my-org"
  installation_id       = "12345678"
  selected_repositories = ["repo-a", "repo-b"]
}

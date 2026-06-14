resource "github_enterprise_app_installation" "example" {
  enterprise_slug      = "my-enterprise"
  organization         = "my-org"
  client_id            = "Iv23liABCDEFGH012345"
  repository_selection = "selected"
  repositories         = ["repo-a", "repo-b"]
}

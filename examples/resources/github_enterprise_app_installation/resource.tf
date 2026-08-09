resource "github_enterprise_app_installation" "all_repos" {
  enterprise_slug      = "my-enterprise"
  organization         = "my-org"
  client_id            = "Iv1.abc123def456"
  repository_selection = "all"
}

resource "github_enterprise_app_installation" "selected_repos" {
  enterprise_slug       = "my-enterprise"
  organization          = "my-org"
  client_id             = "Iv1.789ghi012jkl"
  repository_selection  = "selected"
  selected_repositories = ["my-repo-1", "my-repo-2"]
}

data "github_repository_pull_requests" "example" {
  base_repository = "example-repository"
  base_ref        = "main"
  sort_by         = "updated"
  sort_direction  = "desc"
  state           = "open"
}

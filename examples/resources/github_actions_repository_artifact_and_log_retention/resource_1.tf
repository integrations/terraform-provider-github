resource "github_actions_repository_artifact_and_log_retention" "example" {
  repository = "example-repo"
  days       = 7
}

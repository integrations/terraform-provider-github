resource "github_team_project" "platform" {
  project_id = github_project.planning.id
  team_slug  = "platform"
}

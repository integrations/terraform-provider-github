resource "github_project_repository" "planning" {
  project_id = github_project.planning.id
  repository = "planning"
}

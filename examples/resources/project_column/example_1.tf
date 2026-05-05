resource "github_organization_project" "project" {
  name = "A Organization Project"
  body = "This is an organization project."
}

resource "github_project_column" "column" {
  project_id = github_organization_project.project.id
  name       = "a column"
}

resource "github_organization_project" "project" {
  name = "An Organization Project"
  body = "This is an organization project."
}

resource "github_project_column" "column" {
  project_id = github_organization_project.project.id
  name       = "Backlog"
}

resource "github_project_card" "card" {
  column_id = github_project_column.column.column_id
  note      = "## Unaccepted 👇"
}

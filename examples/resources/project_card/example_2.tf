resource "github_repository" "test" {
  name = "myrepo"
  has_projects = true
  has_issues   = true
}

resource "github_issue" "test" {
  repository       = github_repository.test.id
  title            = "Test issue title"
  body             = "Test issue body"
}

resource "github_repository_project" "test" {
  name            = "test"
  repository      = github_repository.test.name
  body            = "this is a test project"
}

resource "github_project_column" "test" {
  project_id = github_repository_project.test.id
  name       = "Backlog"
}

resource "github_project_card" "test" {
  column_id    = github_project_column.test.column_id
  content_id   = github_issue.test.issue_id
  content_type = "Issue"
}

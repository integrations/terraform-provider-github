# Create an issue with milestone and project assignment
resource "github_repository" "test" {
  name       = "tf-acc-test-%s"
  auto_init  = true
  has_issues = true
}

resource "github_repository_milestone" "test" {
  owner       = split("/", "${github_repository.test.full_name}")[0]
  repository  = github_repository.test.name
  title       = "v1.0.0"
  description = "General Availability"
  due_date    = "2022-11-22"
  state       = "open"
}

resource "github_issue" "test" {
  repository       = github_repository.test.name
  title            = "My issue"
  body             = "My issue body"
  labels           = ["bug", "documentation"]
  assignees        = ["bob-github"]
  milestone_number = github_repository_milestone.test.number
}

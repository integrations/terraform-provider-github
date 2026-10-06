# Create a simple issue
resource "github_repository" "test" {
  name       = "tf-acc-test-%s"
  auto_init  = true
  has_issues = true
}

resource "github_issue" "test" {
  repository = github_repository.test.name
  title      = "My issue title"
  body       = "The body of my issue"
}

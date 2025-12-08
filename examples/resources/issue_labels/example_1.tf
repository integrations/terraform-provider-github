# Create a new, red colored label
resource "github_issue_labels" "test_repo" {
  repository = "test-repo"

  label {
    name  = "Urgent"
    color = "FF0000"
  }

  label {
    name  = "Critical"
    color = "FF0000"
  }
}

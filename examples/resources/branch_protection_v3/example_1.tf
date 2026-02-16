# Protect the main branch of the foo repository. Only allow a specific user to merge to the branch.
resource "github_branch_protection_v3" "example" {
  repository = github_repository.example.name
  branch     = "main"

  restrictions {
    users = ["foo-user"]
  }
}

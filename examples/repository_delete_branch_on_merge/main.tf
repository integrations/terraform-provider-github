resource "github_repository" "delete_branch_on_merge" {
  name                   = "delete_branch_on_merge"
  description            = "A repository with delete-branch-on-merge configured"
  default_branch         = "main"
  private                = true
  delete_branch_on_merge = true
}

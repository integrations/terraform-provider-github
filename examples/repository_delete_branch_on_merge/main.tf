resource "github_repository" "delete_branch_on_merge" {
  name                   = "delete_branch_on_merge"
  description            = "A repository with delete-branch-on-merge configured"
  visibility             = "private"
  delete_branch_on_merge = true
}
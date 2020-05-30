data "github_release" "by_tag" {
  repository  = var.repository
  owner       = var.owner
  release_tag = var.release_tag
  retrieve_by = "tag"
}

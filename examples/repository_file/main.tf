resource "github_repository_file" "this" {
  repository = var.repository
  file       = var.file
  content    = var.content

  branch = var.branch

  commit_author  = var.commit_author
  commit_message = var.commit_message
  commit_email   = var.commit_email
}

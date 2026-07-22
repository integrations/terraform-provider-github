variable "issue_node_id" {
  type = string
}

resource "github_project_item" "issue" {
  project_id = github_project.planning.id
  content_id = var.issue_node_id
}

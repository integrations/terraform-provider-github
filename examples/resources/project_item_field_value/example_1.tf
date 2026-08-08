variable "status_option_id" {
  type = string
}

resource "github_project_item_field_value" "status" {
  project_id              = github_project.planning.id
  item_id                 = github_project_item.issue.id
  field_id                = github_project_field.status.id
  single_select_option_id = var.status_option_id
}

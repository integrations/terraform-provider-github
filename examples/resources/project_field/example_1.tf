resource "github_project_field" "status" {
  project_id = github_project.planning.id
  name       = "Status"
  data_type  = "SINGLE_SELECT"

  single_select_option {
    name        = "To Do"
    description = "Ready to start"
    color       = "GRAY"
  }

  single_select_option {
    name        = "In Progress"
    description = "Work has started"
    color       = "YELLOW"
  }
}

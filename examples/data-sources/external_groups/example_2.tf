data "github_external_groups" "example_external_groups_filtered" {
  display_name = "my-group"
}

locals {
  filtered_groups = data.github_external_groups.example_external_groups_filtered
}

output "groups" {
  value = local.filtered_groups
}

data "github_external_groups" "example_external_groups" {}

locals {
  local_groups = data.github_external_groups.example_external_groups
}

output "groups" {
  value = local.local_groups
}

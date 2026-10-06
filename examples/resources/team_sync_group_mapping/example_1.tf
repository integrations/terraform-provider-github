
data "github_organization_team_sync_groups" "example_groups" {}

resource "github_team_sync_group_mapping" "example_group_mapping" {
  team_slug = "example"

  dynamic "group" {
    for_each = [for g in data.github_organization_team_sync_groups.example_groups.groups : g if g.group_name == "some_team_group"]
    content {
      group_id          = group.value.group_id
      group_name        = group.value.group_name
      group_description = group.value.group_description
    }
  }
}

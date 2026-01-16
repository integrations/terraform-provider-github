---
layout: "github"
page_title: "GitHub: github_emu_group_mapping"
description: |-
  Manages mappings between external groups for enterprise managed users.
---

# github_emu_group_mapping

This resource manages mappings between external groups for enterprise managed users and GitHub teams. It wraps the [Teams#ExternalGroups API](https://docs.github.com/en/rest/reference/teams#external-groups). Note that this is a distinct resource from `github_team_sync_group_mapping`. `github_emu_group_mapping` is special to the Enterprise Managed User (EMU) external group feature, whereas `github_team_sync_group_mapping` is specific to Identity Provider Groups.

!> **Warning:**: This resources `Read` function has a fundamental bug. It doesn't verify that the group is actually linked to the team. Someone could modify the linked group outside of Terraform and the resource would not detect it.

## Example Usage

```hcl
resource "github_emu_group_mapping" "example_emu_group_mapping" {
  team_slug = "emu-test-team" # The GitHub team name to modify
  group_id = 28836 # The group ID of the external group to link
}
```

## Argument Reference

The following arguments are supported:

- `team_slug` - (Required) Slug of the GitHub team
- `group_id`  - (Required) Integer corresponding to the external group ID to be linked

## Import

GitHub EMU External Group Mappings can be imported using the external `group_id` and `team_slug` separated by a colon, e.g.

```sh
$ terraform import github_emu_group_mapping.example_emu_group_mapping 28836:emu-test-team
```

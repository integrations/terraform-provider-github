provider "github" {
}

terraform {
  required_providers {
    github = {
      source  = "integrations/github"
    }
  }
}

# resource "github_team" "emu-test-team" {
# 	name = "emu-test-team"
# }

resource "github_emu_group_mapping" "test" {

  team_slug = "emu-test-team-2"
  group = {
    group_id = "28836"
    group_name = "terraform-emu-test-group"
    group_description = "AAD group for testing EMUs"
  }
}

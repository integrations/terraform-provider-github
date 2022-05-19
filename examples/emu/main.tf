provider "github" {
}

terraform {
  required_providers {
    github = {
      source  = "integrations/github"
    }
  }
}

resource "github_emu_group_mapping" "example_emu_group_mapping" {
  team_slug = "emu-test-team"
  group_id = 28836
}

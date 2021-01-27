# NOTE: There are errors when using the following syntax. See links for details.
# /cc https://github.com/hashicorp/terraform/issues/23529
# /cc https://github.com/hashicorp/terraform/issues/4149
# /cc https://github.com/integrations/terraform-provider-github/issues/500
# 
# data "github_team" "writers" {
#   slug = "writer-team"
# }
# 
# resource "github_team_repository" "writers" {
#   for_each   = data.github_team.writers.id
#   # or for multiple teams something like:
#   # for_each   = { for obj in [data.github_team.writers] : obj.id => obj.id }
#   team_id    = each.value
#   repository = "repo"
#   permission = "push"
# }

data "github_team" "writers" {
  slug = "writers"
}

data "github_team" "editors" {
  slug = "editors"
}

locals {
  teams = [data.github_team.writers.id, data.github_team.editors.id]
}

resource "github_repository" "writers" {
  name = "writers"
  auto_init = true
}

resource "github_team_repository" "writers" {
  count = length(local.teams)
  team_id = local.teams[count.index]
  repository = github_repository.writers.name 
  permission = "push"
}

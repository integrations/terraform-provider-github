data "github_enterprise" "example" {
  slug = "example-co"
}

resource "github_enterprise_actions_runner_group" "example" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "example-runner-group"
  visibility      = "all"
}

resource "github_enterprise_actions_hosted_runner" "example" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "example-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_enterprise_actions_runner_group.example.id
}

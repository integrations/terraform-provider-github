data "github_enterprise" "example" {
  slug = "example-co"
}

resource "github_enterprise_actions_runner_group" "advanced" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "advanced-runner-group"
  visibility      = "selected"
}

resource "github_enterprise_actions_hosted_runner" "advanced" {
  enterprise_slug = data.github_enterprise.example.slug
  name            = "advanced-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size              = "8-core"
  runner_group_id   = github_enterprise_actions_runner_group.advanced.id
  maximum_runners   = 10
  public_ip_enabled = true
}

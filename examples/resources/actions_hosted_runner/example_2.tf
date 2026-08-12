resource "github_actions_runner_group" "advanced" {
  name       = "advanced-runner-group"
  visibility = "selected"
}

resource "github_actions_hosted_runner" "advanced" {
  name = "advanced-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size              = "8-core"
  runner_group_id   = github_actions_runner_group.advanced.id
  maximum_runners   = 10
  public_ip_enabled = true
}

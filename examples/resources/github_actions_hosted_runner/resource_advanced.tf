resource "github_actions_runner_group" "example" {
  name       = "example-runner-group"
  visibility = "all"
}

resource "github_actions_hosted_runner" "example" {
  name = "example-hosted-runner"

  image {
    id     = "2306"
    source = "github"
  }

  size              = "8-core"
  runner_group_id   = github_actions_runner_group.example.id
  maximum_runners   = 10
  public_ip_enabled = true
}

resource "github_actions_hosted_runner" "example" {
  name = "example-hosted-runner"
  
  image {
    id     = "2306"
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_actions_runner_group.example.id

  timeouts {
    delete = "15m"
  }
}

resource "github_actions_runner_group" "example" {
  name       = "example-runner-group"
  visibility = "all"
}

resource "github_actions_hosted_runner" "example" {
  name = "example-hosted-runner"
  
  image {
    id     = "ubuntu-latest"
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_actions_runner_group.example.id
}

# Advanced example with optional parameters
resource "github_actions_hosted_runner" "advanced" {
  name = "advanced-hosted-runner"
  
  image {
    id     = "ubuntu-latest"
    source = "github"
  }

  size             = "8-core"
  runner_group_id  = github_actions_runner_group.example.id
  maximum_runners  = 10
  enable_static_ip = true
}

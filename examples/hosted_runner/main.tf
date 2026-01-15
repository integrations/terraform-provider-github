resource "github_actions_runner_group" "example" {
  name       = "example-runner-group"
  visibility = "all"
}

# NOTE: You must first query available images using the GitHub API:
# GET /orgs/{org}/actions/hosted-runners/images/github-owned
# The image ID is numeric, not a string like "ubuntu-latest"
resource "github_actions_hosted_runner" "example" {
  name = "example-hosted-runner"
  
  image {
    id     = "2306"  # Ubuntu Latest (24.04) - query your org for available IDs
    source = "github"
  }

  size            = "4-core"
  runner_group_id = github_actions_runner_group.example.id
}

# Advanced example with optional parameters
resource "github_actions_hosted_runner" "advanced" {
  name = "advanced-hosted-runner"
  
  image {
    id     = "2306"  # Ubuntu Latest (24.04) - query your org for available IDs
    source = "github"
  }

  size             = "8-core"
  runner_group_id  = github_actions_runner_group.example.id
  maximum_runners  = 10
  enable_static_ip = true
}

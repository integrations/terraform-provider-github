resource "github_organization_network_configuration" "example" {
  name                 = "my-network-configuration"
  compute_service      = "actions"
  network_settings_ids = ["23456789ABCDEF1"]
}

resource "github_actions_runner_group" "example" {
  name                     = "my-runner-group"
  visibility               = "all"
  network_configuration_id = github_organization_network_configuration.example.id
}

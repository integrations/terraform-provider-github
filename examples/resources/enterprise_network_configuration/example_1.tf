resource "github_enterprise_network_configuration" "example" {
  enterprise_slug      = "my-enterprise"
  name                 = "my-network-configuration"
  compute_service      = "actions"
  network_settings_ids = ["23456789ABCDEF1"]
}

resource "github_enterprise_actions_runner_group" "example" {
  enterprise_slug          = "my-enterprise"
  name                     = "my-runner-group"
  visibility               = "all"
  network_configuration_id = github_enterprise_network_configuration.example.id
}

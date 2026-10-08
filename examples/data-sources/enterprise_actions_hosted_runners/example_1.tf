data "github_enterprise_actions_hosted_runners" "example" {
  enterprise_slug = "example-co"
}

output "runner_names" {
  value = [for runner in data.github_enterprise_actions_hosted_runners.example.runners : runner.name]
}

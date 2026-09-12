data "github_repository_rulesets" "example" {
  repository = "example-repo"
}

output "all_rulesets" {
  value = data.github_repository_rulesets.example.rulesets
}

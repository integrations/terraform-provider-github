data "github_actions_environment_secrets" "example" {
  name        = "exampleRepo"
  environment = "exampleEnvironment"
}

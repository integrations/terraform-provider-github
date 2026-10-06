data "github_repositories" "example" {
  query           = "org:hashicorp language:Go"
  include_repo_id = true
}

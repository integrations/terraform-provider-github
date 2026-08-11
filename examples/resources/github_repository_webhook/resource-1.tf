resource "github_repository" "example_repo" {
  name         = "example-repo"
  description  = "Terraform acceptance tests"
  homepage_url = "http://example.com/"

  visibility = "public"
}

resource "github_repository_webhook" "example_webhook" {
  repository = github_repository.example_repo.name

  configuration {
    url          = "https://google.de/"
    content_type = "form"
    insecure_ssl = false
  }

  active = false

  events = ["issues"]
}

variable "webhook_secret" {
  type      = string
  sensitive = true
  ephemeral = true
}

resource "github_repository" "repo" {
  name         = "foo"
  description  = "Terraform acceptance tests"
  homepage_url = "http://example.com/"

  visibility = "public"
}

resource "github_repository_webhook" "foo" {
  repository = github_repository.repo.name

  configuration {
    url               = "https://google.de/"
    content_type      = "form"
    insecure_ssl      = false
    secret_wo         = var.webhook_secret
    secret_wo_version = 1
  }

  active = false

  events = ["issues"]
}

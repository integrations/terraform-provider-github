variable "webhook_secret" {
  type      = string
  sensitive = true
  ephemeral = true
}

resource "github_organization_webhook" "foo" {
  name = "web"

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

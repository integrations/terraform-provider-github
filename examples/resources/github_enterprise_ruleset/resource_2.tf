# Block large files and secrets-prone extensions on push, for two organizations only.

resource "github_enterprise_ruleset" "push_restrictions" {
  enterprise_slug = "my-enterprise"
  name            = "push-restrictions"
  target          = "push"
  enforcement     = "active"

  conditions {
    organization_name {
      include = ["platform", "payments"]
      exclude = []
    }

    repository_name {
      include = ["~ALL"]
      exclude = []
    }
  }

  rules {
    max_file_size {
      max_file_size = 50
    }

    file_extension_restriction {
      restricted_file_extensions = ["*.pem", "*.p12"]
    }
  }
}

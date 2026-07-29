# The `repository` target is only available for enterprise rulesets. It governs the
# repositories themselves rather than their contents.

resource "github_enterprise_ruleset" "repository_lifecycle" {
  enterprise_slug = "my-enterprise"
  name            = "repository-lifecycle"
  target          = "repository"
  enforcement     = "active"

  conditions {
    organization_name {
      include = ["~ALL"]
      exclude = []
    }

    repository_name {
      include = ["~ALL"]
      exclude = []
    }
  }

  rules {
    repository_delete   = true
    repository_transfer = true

    repository_visibility {
      internal = true
      private  = true
    }
  }
}

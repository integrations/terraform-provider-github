# Target organizations and repositories by custom property instead of by name.
# Note that organization properties have no `source` — only repository properties do.

resource "github_enterprise_ruleset" "production_by_property" {
  enterprise_slug = "my-enterprise"
  name            = "production-by-property"
  target          = "branch"
  enforcement     = "active"

  conditions {
    organization_property {
      include {
        name            = "environment"
        property_values = ["production"]
      }

      exclude {
        name            = "environment"
        property_values = ["sandbox"]
      }
    }

    repository_property {
      include {
        name            = "language"
        property_values = ["Go"]
        source          = "custom"
      }
    }

    ref_name {
      include = ["~DEFAULT_BRANCH"]
      exclude = []
    }
  }

  rules {
    deletion            = true
    required_signatures = true
  }
}

# Example: Tag target ruleset for protecting tags
# This ruleset applies to tags across the enterprise

resource "github_enterprise_ruleset" "tag_protection" {
  enterprise_slug = "your-enterprise"
  name            = "tag-protection-ruleset"
  target          = "tag"
  enforcement     = "active"

  # Allow organization admins to bypass tag rules
  bypass_actors {
    actor_id    = 1
    actor_type  = "OrganizationAdmin"
    bypass_mode = "always"
  }

  # Conditions define which organizations, repositories, and refs this ruleset applies to
  conditions {
    # Target all organizations
    organization_name {
      include = ["~ALL"]
    }

    # Target all repositories
    repository_name {
      include = ["~ALL"]
    }

    # Target specific tag patterns (required for tag target)
    ref_name {
      include = ["v*", "release/*"]
      exclude = ["*-beta", "*-alpha"]
    }
  }

  # Rules that apply to matching tags
  rules {
    # Prevent tag creation without bypass permission
    creation = true

    # Prevent tag updates (tags should be immutable)
    update = true

    # Prevent tag deletion without bypass permission
    deletion = true

    # Require signed commits for tags
    required_signatures = true

    # Tag name pattern (only for tag target)
    tag_name_pattern {
      name     = "Semantic Version Tags"
      operator = "regex"
      pattern  = "^v[0-9]+\\.[0-9]+\\.[0-9]+(-[a-zA-Z0-9.]+)?$"
      negate   = false
    }

    # Commit message pattern for tagged commits
    commit_message_pattern {
      name     = "Release Commit Message"
      operator = "starts_with"
      pattern  = "Release:"
      negate   = false
    }

    # Require specific commit author email pattern
    commit_author_email_pattern {
      name     = "Release Manager Email"
      operator = "contains"
      pattern  = "release@your-company.com"
      negate   = false
    }
  }
}

# Example: Less restrictive tag ruleset for development tags
resource "github_enterprise_ruleset" "dev_tag_protection" {
  enterprise_slug = "your-enterprise"
  name            = "dev-tag-protection-ruleset"
  target          = "tag"
  enforcement     = "active"

  conditions {
    organization_name {
      include = ["~ALL"]
    }

    repository_name {
      include = ["~ALL"]
    }

    # Only apply to development/snapshot tags
    ref_name {
      include = ["*-SNAPSHOT", "*-dev"]
    }
  }

  rules {
    # Allow tag creation
    creation = false

    # Allow tag updates for development tags
    update = false

    # Prevent tag deletion
    deletion = true

    # Tag name pattern for development tags
    tag_name_pattern {
      name     = "Development Tag Pattern"
      operator = "regex"
      pattern  = "^v[0-9]+\\.[0-9]+\\.[0-9]+-[a-zA-Z0-9.]+(SNAPSHOT|dev)$"
      negate   = false
    }
  }
}

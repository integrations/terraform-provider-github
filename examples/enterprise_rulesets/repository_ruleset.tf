
# Example: Repository target ruleset for repository management
# This ruleset controls repository creation, deletion, and naming
resource "github_enterprise_ruleset" "repository_management" {
  enterprise_slug = "your-enterprise"
  name            = "repository-management-ruleset"
  target          = "repository"
  enforcement     = "active"

  # Allow organization admins to bypass repository rules
  bypass_actors {
    actor_id    = 1
    actor_type  = "OrganizationAdmin"
    bypass_mode = "always"
  }

  # Conditions define which organizations and repositories this ruleset applies to
  # Note: ref_name is NOT used for repository target
  conditions {
    # Target all organizations
    organization_name {
      include = ["~ALL"]
    }

    # Target all repositories
    repository_name {
      include = ["~ALL"]
    }
  }

  # Repository-specific rules (only valid for repository target)
  rules {
    # Prevent repository creation without bypass permission
    repository_creation = true

    # Prevent repository deletion without bypass permission
    repository_deletion = true

    # Prevent repository transfer without bypass permission
    repository_transfer = true

    # Enforce repository naming conventions
    repository_name {
      pattern = "^[a-z][a-z0-9-]*$" # lowercase letters, numbers, and hyphens only
      negate  = false
    }

    # Control repository visibility changes
    repository_visibility {
      internal = true  # Allow internal visibility
      private  = true  # Allow private visibility
      # Note: public visibility is implicitly allowed if not restricted
    }
  }
}

# Example: Stricter repository ruleset for production organizations
resource "github_enterprise_ruleset" "production_repository_rules" {
  enterprise_slug = "your-enterprise"
  name            = "production-repository-rules"
  target          = "repository"
  enforcement     = "active"

  bypass_actors {
    actor_id    = 1
    actor_type  = "OrganizationAdmin"
    bypass_mode = "always"
  }

  conditions {
    # Only apply to production organizations
    organization_name {
      include = ["*-production", "*-prod"]
    }

    repository_name {
      include = ["~ALL"]
    }
  }

  rules {
    # Block repository creation, deletion, and transfer
    repository_creation = true
    repository_deletion = true
    repository_transfer = true

    # Strict naming: must start with org prefix and follow kebab-case
    repository_name {
      pattern = "^prod-[a-z][a-z0-9-]*$"
      negate  = false
    }

    # Only allow private repositories in production
    repository_visibility {
      internal = false
      private  = true
    }
  }
}

# Example: Repository ruleset with organization ID targeting
resource "github_enterprise_ruleset" "org_id_repository_rules" {
  enterprise_slug = "your-enterprise"
  name            = "org-id-repository-rules"
  target          = "repository"
  enforcement     = "evaluate" # Test mode - doesn't block, just reports

  conditions {
    # Use organization_id instead of organization_name
    # This is useful when you know the specific org IDs
    organization_id = [123456, 789012]

    # Use repository_id for specific repositories
    repository_id = [111111, 222222]
  }

  rules {
    repository_creation = true

    repository_name {
      pattern = "^[a-z0-9-]+$"
      negate  = false
    }

    repository_visibility {
      internal = true
      private  = true
    }
  }
}

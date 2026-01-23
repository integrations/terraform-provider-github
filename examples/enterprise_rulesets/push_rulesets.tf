# Example: Push target ruleset for file and content restrictions
# This ruleset applies to all pushes across the enterprise

resource "github_enterprise_ruleset" "push_restrictions" {
  enterprise_slug = "your-enterprise"
  name            = "push-restrictions-ruleset"
  target          = "push"
  enforcement     = "active"

  # Allow deploy keys and organization admins to bypass
  bypass_actors {
    actor_type  = "DeployKey"
    bypass_mode = "always"
  }

  bypass_actors {
    actor_id    = 1
    actor_type  = "OrganizationAdmin"
    bypass_mode = "always"
  }

  # Conditions define which organizations and repositories this ruleset applies to
  # Note: ref_name is NOT used for push target
  conditions {
    # Target all organizations
    organization_name {
      include = ["~ALL"]
    }

    # Target all repositories
    repository_name {
      include = ["~ALL"]
      exclude = ["sandbox-*"]
    }
  }

  # Rules that apply to all pushes
  rules {
    # Restrict specific file paths from being pushed
    file_path_restriction {
      restricted_file_paths = [
        "secrets.txt",
        "*.key",
        "*.pem",
        ".env",
        "credentials/*"
      ]
    }

    # Limit maximum file size to prevent large files
    max_file_size {
      max_file_size = 100 # Max 100 MB
    }

    # Limit maximum file path length
    max_file_path_length {
      max_file_path_length = 255
    }

    # Restrict specific file extensions
    file_extension_restriction {
      restricted_file_extensions = [
        "*.exe",
        "*.dll",
        "*.so",
        "*.dylib",
        "*.zip",
        "*.tar.gz"
      ]
    }

    # Commit message pattern
    commit_message_pattern {
      name     = "Valid Commit Message"
      operator = "regex"
      pattern  = "^(feat|fix|docs|style|refactor|test|chore)(\\(.+\\))?: .+"
      negate   = false
    }

    # Commit author email pattern
    commit_author_email_pattern {
      name     = "Corporate Email"
      operator = "ends_with"
      pattern  = "@your-company.com"
      negate   = false
    }

    # Committer email pattern
    committer_email_pattern {
      name     = "Corporate Email"
      operator = "ends_with"
      pattern  = "@your-company.com"
      negate   = false
    }
  }
}

# Example: Security-focused push ruleset
resource "github_enterprise_ruleset" "security_push_restrictions" {
  enterprise_slug = "your-enterprise"
  name            = "security-push-restrictions"
  target          = "push"
  enforcement     = "active"

  conditions {
    organization_name {
      include = ["~ALL"]
    }

    repository_name {
      include = ["*-prod", "*-production"]
    }
  }

  rules {
    # Block common secret file patterns
    file_path_restriction {
      restricted_file_paths = [
        "*.pem",
        "*.key",
        "*.cert",
        "*.p12",
        "*.pfx",
        ".env",
        ".env.*",
        "secrets.yml",
        "credentials.json"
      ]
    }

    # Strict file size limits for production
    max_file_size {
      max_file_size = 50 # Max 50 MB
    }

    # Block executable and archive files
    file_extension_restriction {
      restricted_file_extensions = [
        "*.exe",
        "*.dll",
        "*.so",
        "*.dylib",
        "*.bin",
        "*.dmg"
      ]
    }

    # Require signed commits
    required_signatures = true
  }
}

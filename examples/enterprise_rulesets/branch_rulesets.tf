# Example: Branch target ruleset with comprehensive branch protection rules
# This ruleset applies to branches across the enterprise

resource "github_enterprise_ruleset" "branch_protection" {
  enterprise_slug = "your-enterprise"
  name            = "branch-protection-ruleset"
  target          = "branch"
  enforcement     = "active"

  # Optional: Allow certain users/teams to bypass the ruleset
  bypass_actors {
    actor_id    = 1
    actor_type  = "OrganizationAdmin"
    bypass_mode = "always"
  }

  bypass_actors {
    actor_type  = "DeployKey"
    bypass_mode = "always"
  }

  # Conditions define which organizations, repositories, and refs this ruleset applies to
  conditions {
    # Target all organizations in the enterprise
    organization_name {
      include = ["~ALL"]
      exclude = []
    }

    # Target all repositories
    repository_name {
      include = ["~ALL"]
      exclude = ["test-*"] # Exclude test repositories
    }

    # Target all branches (required for branch target)
    ref_name {
      include = ["~DEFAULT_BRANCH", "main", "master", "release/*"]
      exclude = ["experimental/*"]
    }
  }

  # Rules that apply to matching branches
  rules {
    # Prevent branch creation without bypass permission
    creation = true

    # Prevent branch updates without bypass permission
    update = false

    # Prevent branch deletion without bypass permission
    deletion = true

    # Require linear history (no merge commits)
    required_linear_history = true

    # Require signed commits
    required_signatures = true

    # Prevent force pushes
    non_fast_forward = true

    # Pull request requirements
    pull_request {
      dismiss_stale_reviews_on_push     = true
      require_code_owner_review         = true
      require_last_push_approval        = true
      required_approving_review_count   = 2
      required_review_thread_resolution = true
      allowed_merge_methods             = ["squash", "merge"]
    }

    # Status check requirements
    required_status_checks {
      strict_required_status_checks_policy = true
      do_not_enforce_on_create             = false

      required_check {
        context        = "ci/build"
        integration_id = 0
      }

      required_check {
        context        = "ci/test"
        integration_id = 0
      }
    }

    # Commit message pattern requirements
    commit_message_pattern {
      name     = "Conventional Commits"
      operator = "regex"
      pattern  = "^(feat|fix|docs|style|refactor|test|chore)(\\(.+\\))?: .{1,50}"
      negate   = false
    }

    # Commit author email pattern
    commit_author_email_pattern {
      name     = "Corporate Email Only"
      operator = "regex"
      pattern  = "@your-company\\.com$"
      negate   = false
    }

    # Committer email pattern
    committer_email_pattern {
      name     = "Corporate Email Only"
      operator = "regex"
      pattern  = "@your-company\\.com$"
      negate   = false
    }

    # Branch name pattern (only for branch target)
    branch_name_pattern {
      name     = "Valid Branch Names"
      operator = "regex"
      pattern  = "^(main|master|develop|feature/|bugfix/|hotfix/|release/)"
      negate   = false
    }

    # Code scanning requirements
    required_code_scanning {
      required_code_scanning_tool {
        tool                       = "CodeQL"
        alerts_threshold           = "errors"
        security_alerts_threshold  = "high_or_higher"
      }
    }

    # Copilot code review (if enabled)
    copilot_code_review {
      review_on_push           = true
      review_draft_pull_requests = false
    }
  }
}

resource "github_enterprise_ruleset" "branch_by_property" {
  enterprise_slug = "your-enterprise"
  name            = "production-repos-branch-protection"
  target          = "branch"
  enforcement     = "active"

  conditions {
    organization_name {
      include = ["~ALL"]
      exclude = []
    }

    # Target repositories based on custom properties
    repository_property {
      include {
        name            = "environment"
        property_values = ["production", "staging"]
        source          = "custom"
      }

      exclude {
        name            = "lifecycle"
        property_values = ["deprecated", "archived"]
      }
    }

    ref_name {
      include = ["~DEFAULT_BRANCH", "refs/heads/release/*"]
      exclude = []
    }
  }

  rules {
    deletion         = true
    non_fast_forward = true
  }
}
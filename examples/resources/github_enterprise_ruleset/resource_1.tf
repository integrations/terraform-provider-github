# Protect the default branch of every repository in the enterprise.

resource "github_enterprise_ruleset" "protect_default_branch" {
  enterprise_slug = "my-enterprise"
  name            = "protect-default-branch"
  target          = "branch"
  enforcement     = "active"

  bypass_actors {
    actor_type  = "EnterpriseOwner"
    bypass_mode = "always"
  }

  conditions {
    organization_name {
      include = ["~ALL"]
      exclude = []
    }

    repository_name {
      include = ["~ALL"]
      exclude = []
    }

    ref_name {
      include = ["~DEFAULT_BRANCH"]
      exclude = []
    }
  }

  rules {
    deletion         = true
    non_fast_forward = true

    pull_request {
      required_approving_review_count = 1
      require_code_owner_review       = true
    }
  }
}

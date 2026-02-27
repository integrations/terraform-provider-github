resource "github_organization_ruleset" "example" {
  name        = "example"
  target      = "branch"
  enforcement = "active"

  conditions {
    ref_name {
      include = ["~ALL"]
      exclude = []
    }
  }

  bypass_actors {
    actor_id    = 13473
    actor_type  = "Integration"
    bypass_mode = "always"
  }

  rules {
    creation                = true
    update                  = true
    deletion                = true
    required_linear_history = true
    required_signatures     = true

    branch_name_pattern {
      name     = "example"
      negate   = false
      operator = "starts_with"
      pattern  = "ex"
    }

    required_workflows {
      do_not_enforce_on_create = true
      required_workflow {
        repository_id = 1234
        path          = ".github/workflows/ci.yml"
        ref           = "main"
      }
    }

    required_code_scanning {
      required_code_scanning_tool {
        alerts_threshold          = "errors"
        security_alerts_threshold = "high_or_higher"
        tool                      = "CodeQL"
      }
    }
  }
}

# Example with push ruleset
# Note: Push targets must NOT have ref_name in conditions, only repository_name or repository_id
resource "github_organization_ruleset" "example_push" {
  name        = "example_push"
  target      = "push"
  enforcement = "active"

  conditions {
    repository_name {
      include = ["~ALL"]
      exclude = []
    }
  }

  rules {
    # Push targets only support these rules:
    # file_path_restriction, max_file_size, max_file_path_length, file_extension_restriction
    file_path_restriction {
      restricted_file_paths = [".github/workflows/*", "*.env"]
    }

    max_file_size {
      max_file_size = 100  # 100 MB
    }

    max_file_path_length {
      max_file_path_length = 255
    }

    file_extension_restriction {
      restricted_file_extensions = ["*.exe", "*.dll", "*.so"]
    }
  }
}

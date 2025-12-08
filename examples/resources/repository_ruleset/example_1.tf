resource "github_repository" "example" {
  name        = "example"
  description = "Example repository"
}

resource "github_repository_ruleset" "example" {
  name        = "example"
  repository  = github_repository.example.name
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

    required_deployments {
      required_deployment_environments = ["test"]
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
resource "github_repository_ruleset" "example_push" {
  name        = "example_push"
  repository  = github_repository.example.name
  target      = "push"
  enforcement = "active"

  rules {
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

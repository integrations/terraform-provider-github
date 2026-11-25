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
      max_file_size = 100 # 100 MB
    }

    max_file_path_length {
      max_file_path_length = 255
    }

    file_extension_restriction {
      restricted_file_extensions = ["*.exe", "*.dll", "*.so"]
    }
  }
}

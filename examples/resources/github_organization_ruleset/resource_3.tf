# Example with repository ruleset
# Note: Repository targets must NOT have ref_name in conditions, only repository_name or repository_id
resource "github_organization_ruleset" "example_repository" {
  name        = "example_repository"
  target      = "repository"
  enforcement = "active"

  conditions {
    repository_name {
      include = ["~ALL"]
      exclude = []
    }
  }

  rules {
    # Repository targets only support these rules:
    # repository_create, repository_delete, repository_name, repository_transfer, repository_visibility
    repository_create   = true
    repository_delete   = true
    repository_transfer = true

    repository_name {
      pattern = "^team-"
      negate  = false
    }

    repository_visibility {
      internal = true
      private  = true
    }
  }
}

resource "github_repository" "terraformed" {
  name        = "terraformed"
  description = "A repository created by terraform"
  visibility  = "public"

  security_and_analysis {
    # Cannot set advanced_security for public repositories as it is always on by default.
    # advanced_security {
    #   status = "enabled"
    # }
    secret_scanning {
      status = "enabled"
    }
    secret_scanning_push_protection {
      status = "enabled"
    }
  }
}


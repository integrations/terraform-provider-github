# Configure the GitHub Provider
provider "github" {
  version = "~> 5.0"
}

# Add a user to the organization
resource "github_membership" "membership_for_user_x" {
  # ...
}

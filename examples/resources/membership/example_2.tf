# Invite someone to the organization by email, e.g. when they
# do not have a GitHub account yet
resource "github_membership" "membership_for_some_email" {
  email = "someuser@example.com"
  role  = "member"
}

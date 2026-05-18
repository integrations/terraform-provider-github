# Manage organization membership by stable GitHub user ID.
# Recommended over `username` for production: if the user renames their
# account, the membership stays in sync without drift.
resource "github_membership" "membership_by_user_id" {
  user_id = 1
  role    = "member"
}

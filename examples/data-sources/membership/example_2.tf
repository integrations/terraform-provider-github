# Look up a membership by the stable GitHub user ID.
# The numeric ID does not change when the user renames their account.
data "github_membership" "by_user_id" {
  user_id = 1
}

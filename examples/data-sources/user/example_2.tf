# Retrieve information about a GitHub user by their stable numeric ID.
# Useful when the user may rename themselves: the lookup keeps working.
data "github_user" "by_id" {
  user_id = 1
}

# Downgrade a member to an outside collaborator when the resource is destroyed
resource "github_membership" "outside_collaborator_on_destroy" {
  username = "SomeUser"
  role     = "member"

  downgrade_on_destroy = true
  downgrade_to         = "outside_collaborator"
}

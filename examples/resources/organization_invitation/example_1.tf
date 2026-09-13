# Invite by email address
resource "github_organization_invitation" "by_email" {
  email = "newmember@example.com"
  role  = "direct_member"
}

# Invite by GitHub user ID
data "github_user" "invitee" {
  username = "someuser"
}

resource "github_organization_invitation" "by_id" {
  invitee_id = data.github_user.invitee.id
  role       = "direct_member"
}

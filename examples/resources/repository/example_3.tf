resource "github_repository" "example" {
  name       = "example"
  visibility = "private"

  custom_properties = {
    securityTier = "tier1"
    owningTeam   = "platform"
  }
}

resource "github_repository" "public" {
  name        = "public-visibility"
  description = "A public-visible repository created by Terraform"
  visibility  = "public"
}

resource "github_repository" "private" {
  name        = "private-visibility"
  description = "A private-visible repository created by Terraform"
  visibility  = "private"
}

# NOTE: Expect an error when testing with a non-org account
# > Error: PATCH https://api.github.com/repos/:org/:repo
# > 422 Only organizations associated with an enterprise can set visibility to internal []
resource "github_repository" "internal" {
  name        = "internal-visibility"
  description = "A internal-visible repository created by Terraform"
  visibility  = "internal"
}

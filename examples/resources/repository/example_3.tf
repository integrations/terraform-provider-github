resource "github_repository" "forked_repo" {
  name         = "forked-repository"
  description  = "This is a fork of another repository"
  fork         = true
  source_owner = "some-org"
  source_repo  = "original-repository"
}

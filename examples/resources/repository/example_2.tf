resource "github_repository" "example" {
  name        = "example"
  description = "My awesome web page"

  private = false

  pages {
    source {
      branch = "master"
      path   = "/docs"
    }
  }
}

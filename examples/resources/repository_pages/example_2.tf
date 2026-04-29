resource "github_repository" "example" {
  name       = "my-repo"
  visibility = "public"
  auto_init  = true

  lifecycle {
    ignore_changes = [
      pages,
    ]
  }
}

resource "github_repository_pages" "example" {
  repository = github_repository.example.name
  build_type = "workflow"
}

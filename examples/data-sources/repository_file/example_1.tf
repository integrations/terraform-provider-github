data "github_repository_file" "foo" {
  repository = github_repository.foo.name
  branch     = "main"
  file       = ".gitignore"
}


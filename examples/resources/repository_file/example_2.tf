
resource "github_repository" "bar" {
  name      = "example2"
  auto_init = true
}

resource "github_branch" "bar" {
  branch     = "does/not/exist"
  repository = github_repository.bar.name

}

resource "github_repository_file" "bar" {
  repository          = github_repository.bar.name
  branch              = "does/not/exist"
  file                = ".gitignore"
  content             = "**/*.tfstate"
  commit_message      = "Managed by Terraform"
  commit_author       = "Terraform User"
  commit_email        = "terraform@example.com"
  overwrite_on_create = true
}

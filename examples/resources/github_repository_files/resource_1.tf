resource "github_repository" "foo" {
  name      = "example"
  auto_init = true
}

resource "github_repository_files" "foo" {
  repository     = github_repository.foo.name
  branch         = "main"
  commit_message = "Managed by Terraform"
  commit_author  = "Terraform User"
  commit_email   = "terraform@example.com"

  file {
    path    = ".gitignore"
    content = "**/*.tfstate"
  }
  file {
    path    = "CODEOWNERS"
    content = "* @octocat\n"
  }
  file {
    path    = "config/app.yaml"
    content = "feature_flag: true\n"
  }
}

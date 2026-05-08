resource "github_repository" "repo" {
  name        = "my-repo"
  description = "GitHub repo managed by Terraform"

  private = false
}

resource "github_repository_autolink_reference" "autolink" {
  repository = github_repository.repo.name

  key_prefix = "TICKET-"

  target_url_template = "https://example.com/TICKET?query=<num>"
}

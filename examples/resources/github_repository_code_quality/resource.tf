# Analyze Python code on self-hosted runners labeled "code-quality".
# Omit runner_type and runner_label to use GitHub-hosted runners.

resource "github_repository" "example" {
  name       = "my-repo"
  visibility = "private"
}

resource "github_repository_code_quality" "example" {
  repository   = github_repository.example.name
  languages    = ["python"]
  runner_type  = "labeled"
  runner_label = "code-quality"
}

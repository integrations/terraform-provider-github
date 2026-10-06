data "github_repository" "test" {
  name = "test"
}

resource "github_repository_topics" "test" {
  repository = github_repository.test.name
  topics     = ["topic-1", "topic-2"]
}

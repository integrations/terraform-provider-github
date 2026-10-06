data "github_repository" "this" {
  name = "example"
}

data "github_branch" "this" {
  branch     = data.github_repository.this.default_branch
  repository = data.github_repository.this.name
}

data "github_tree" "this" {
  recursive  = false
  repository = data.github_repository.this.name
  tree_sha   = data.github_branch.this.sha
}

output "entries" {
  value = data.github_tree.this.entries
}


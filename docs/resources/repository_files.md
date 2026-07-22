---
page_title: "github_repository_files (Resource) - GitHub"
description: |-
  Manages a set of files within a GitHub repository, writing all changes in a single commit per apply
---

# github_repository_files (Resource)

Manages a set of files within a GitHub repository, writing all changes in a single commit per apply via the GitHub Git Data API (blob → tree → commit → ref update).

Use this resource instead of multiple `github_repository_file` resources when you need atomic multi-file commits and want to avoid `409` conflicts caused by parallel per-file writes to the same branch.

~> **Adoption semantics.** Paths listed in `file {}` blocks become the source of truth: existing files at those paths are overwritten on first apply, and removing a `file {}` block deletes that path in the next commit. Files in the repository that are *not* listed in any `file {}` block are never touched — even on destroy. The base tree is preserved and only the listed paths are modified.

~> **Concurrency.** When the branch advances during a commit (e.g., another writer pushed concurrently), the resource transparently re-reads HEAD, rebases the change onto the new base, and retries the ref update — bounded to a handful of attempts with exponential backoff.

## Example Usage

### Basic

```terraform
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
```

### Dynamic file generation

```terraform
locals {
  tenant_namespaces = {
    test01 = { name = "test01" }
    test02 = { name = "test02" }
    test03 = { name = "test03" }
  }
}

resource "github_repository_files" "tenants" {
  repository     = "example"
  branch         = "main"
  commit_message = "chore: sync tenants"
  commit_author  = "Terraform"
  commit_email   = "tf@example.com"

  dynamic "file" {
    for_each = local.tenant_namespaces
    content {
      path    = "tenants/${file.key}.yaml"
      content = yamlencode(file.value)
    }
  }
}
```

## Argument Reference

The following arguments are supported:

- `repository` - (Required) The repository to commit files to. Renaming the repository in GitHub is detected via `repository_id` and treated as a rename, not a recreate.

- `file` - (Required, block, one or more) The set of files this resource manages. Each block contributes one entry to the commit's tree.

  - `path` - (Required) The path of the file in the repository, relative to the repo root.

  - `content` - (Required) The file's content. Stored verbatim; encoded as base64 when sent to the GitHub API so any byte sequence is supported.

- `branch` - (Optional) The branch to commit to. Defaults to the repository's default branch. The branch must already exist; use the `github_branch` resource to create branches.

- `commit_message` - (Optional) The commit message. Auto-generated from the change scope (added / updated / removed counts) if empty.

- `commit_author` - (Optional) Committer author name. **NOTE:** GitHub App users may omit author and email so GitHub can verify commits as the GitHub App.

- `commit_email` - (Optional) Committer email address. **NOTE:** GitHub App users may omit author and email so GitHub can verify commits as the GitHub App.

## Attributes Reference

The following additional attributes are exported:

- `repository_id` - The repository's numeric ID. Used internally to distinguish a repository rename from a recreate.

- `ref` - The fully-qualified ref (`refs/heads/<branch>`) that this resource commits to.

- `commit_sha` - The SHA of the most recent commit on the managed branch.

- `tree_sha` - The tree SHA of that commit.

- `file.*.sha` - The blob SHA of each managed file's current content on the branch.

## Differences from `github_repository_file`

This resource is intentionally smaller than the per-file `github_repository_file`:

- `overwrite_on_create` is **not** supported. Paths listed in `file {}` blocks are always overwritten on first apply — the batch resource is a desired-state declaration.
- `autocreate_branch`, `autocreate_branch_source_branch`, and `autocreate_branch_source_sha` are **not** supported. These are already deprecated on `github_repository_file`. Use the `github_branch` resource to create branches, or `auto_init = true` on `github_repository` for the initial commit.

## Migrating from `github_repository_file`

Cross-resource migration cannot be done by the provider; it's a state-surgery + import workflow you run once per branch.

1. Remove each per-file resource from state (this does **not** delete the file from GitHub):

   ```sh
   terraform state list \
     | grep '^github_repository_file\.tenant\[' \
     | xargs -n1 terraform state rm
   ```

2. Replace the per-file resources in your configuration with a single `github_repository_files` block (use `dynamic "file"` if your file set is generated).

3. Import the new resource. The ID is `<repository>` (uses the default branch) or `<repository>:<branch>`.

   ```sh
   terraform import github_repository_files.tenants my-repo:main
   ```

4. Run `terraform plan`. If the `content` in your new configuration matches what's already on the branch for each path, the plan should be a no-op. Otherwise the first apply will produce a single commit reconciling the differences.

## Import

Repository files can be imported using the format `<repository>` (which uses the default branch) or `<repository>:<branch>`. After import, run `terraform plan` to see the diff between the imported (empty) file set and your configuration.

```sh
terraform import github_repository_files.tenants example:main
```

---
layout: "github"
page_title: "GitHub: github_actions_hosted_runner_custom_image_versions"
description: |-
  Get a list of versions for a custom image for GitHub-hosted runners
---

# github_actions_hosted_runner_custom_image_versions

Use this data source to retrieve all versions of a specific custom image for GitHub-hosted runners in your organization.

## Example Usage

```hcl
data "github_actions_hosted_runner_custom_image_versions" "my_image_versions" {
  image_id = 123
}

output "versions" {
  value = data.github_actions_hosted_runner_custom_image_versions.my_image_versions.versions
}
```

## Argument Reference

* `image_id` - (Required) The custom image definition ID.

## Attributes Reference

* `versions` - A list of image versions. Each version has the following attributes:
  * `version` - Version string (e.g., `1.0.0`).
  * `size_gb` - Size of the image version in GB.
  * `state` - State of the version (e.g., `Ready`).
  * `created_on` - Timestamp when the version was created.

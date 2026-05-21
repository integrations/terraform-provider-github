---
layout: "github"
page_title: "GitHub: github_actions_hosted_runner_custom_image"
description: |-
  Get a single custom image definition for GitHub-hosted runners
---

# github_actions_hosted_runner_custom_image

Use this data source to retrieve details of a specific custom image for GitHub-hosted runners in your organization.

## Example Usage

```hcl
data "github_actions_hosted_runner_custom_image" "my_image" {
  image_id = 123
}

output "image_name" {
  value = data.github_actions_hosted_runner_custom_image.my_image.name
}

output "image_state" {
  value = data.github_actions_hosted_runner_custom_image.my_image.state
}
```

## Argument Reference

* `image_id` - (Required) The custom image definition ID.

## Attributes Reference

* `platform` - Platform of the image (e.g., `linux-x64`).
* `name` - Name of the custom image.
* `source` - Source of the image.
* `versions_count` - Number of versions of this image.
* `total_versions_size` - Total size of all versions in GB.
* `latest_version` - Latest version string.
* `state` - State of the image (e.g., `Ready`).

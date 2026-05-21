---
layout: "github"
page_title: "GitHub: github_actions_hosted_runner_custom_images"
description: |-
  Get a list of custom images for GitHub-hosted runners in an organization
---

# github_actions_hosted_runner_custom_images

Use this data source to retrieve a list of custom images available for GitHub-hosted runners in your organization.

## Example Usage

```hcl
data "github_actions_hosted_runner_custom_images" "all" {
}

output "custom_images" {
  value = data.github_actions_hosted_runner_custom_images.all.images
}
```

## Attributes Reference

* `images` - A list of custom images. Each image has the following attributes:
  * `id` - The custom image definition ID.
  * `platform` - Platform of the image (e.g., `linux-x64`).
  * `name` - Name of the custom image.
  * `source` - Source of the image.
  * `versions_count` - Number of versions of this image.
  * `total_versions_size` - Total size of all versions in GB.
  * `latest_version` - Latest version string.
  * `state` - State of the image (e.g., `Ready`).

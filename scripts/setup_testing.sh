#!/bin/bash

# Ask for the resource name
read -p "Enter the resource name to use in the test(i.e. issue_label for resource_github_issue_label): " resourceName

# Define the directory and file path
DIR="examples/testing/${resourceName}"
FILE="${DIR}/main.tf"

# Create the directory, if it doesn't already exist
mkdir -p "$DIR"

# Create (or overwrite) the file and add the contents with the provided resource name
cat <<EOF >"$FILE"
// Empty configuration
provider "github" {
  //owner = "octokit"
  //token = "ABC123"
}

// Terraform global config for providers
terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }

  // Example resource placeholder using provided resource name
  // Documentation: https://registry.terraform.io/providers/integrations/github/latest/docs/resources/${resourceName}
  resource "${resourceName}" "test" {

  }
}
EOF

echo "File ${FILE} has been created with the specified contents."
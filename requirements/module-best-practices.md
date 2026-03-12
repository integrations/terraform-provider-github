In this comprehensive blog, I have shared detailed Terraform module development best practices.

Whether you're creating modules for your team or contributing to open-source projects, these tips will help you write clean, reusable, and easy-to-maintain infrastructure code.

By the end of this guide, you will have learned:

How to build reusable and well-structured Terraform modules.
Best practices for naming, inputs, outputs, and file organization.
How to test, document, and version your modules effectively.
How to use CI/CD to automate checks and keep your modules up to date.
How to prepare modules for open-source use with proper GitHub setup.
The article series is divided into two parts.

In Part 1, we'll look at best practices for developing open-source Terraform modules and talk about the commonly used Terraform module repository template pattern that many organizations use.
In Part 2, we'll dive into a well-designed, pre-configured template that implements all best practices from Part 1, meets the requirements of most Terraform projects, and can be bootstrapped in 5 minutes.
Terraform Modules
Terraform modules are a key paradigm for managing infrastructure as code. They reduce complexity, increase reusability, and ensure consistency across projects.

By encapsulating infrastructure configurations into modular components, teams can standardize resource setups, reduce errors, and share best practices across projects.

However, the value of these modules is dependent on their internal structure. A well-structured module guarantees that infrastructure stays understandable and manageable as complexity grows.

It allows development teams to make changes over time without introducing additional risks or obstacles.

In contrast, badly constructed modules can cause confusion, make debugging more difficult, and require more maintenance due to hidden bugs and other issues.

⚠️
Please keep in mind that HashiCorp changed the Terraform licensing, leading the community to create a fork called OpenTofu, which is licensed under the MPL-2.0 license. All tips, tricks, and best practices shared in this article are fully compatible with OpenTofu.
Let’s get started!

Terraform Module Development Best Practices
In this section, we will look at best practices for Terraform module development that may help for a wide range of users.

I recommend you adopt only those practices that are relevant to your project or organization. Although they are suggestions, implementing as many of these into practice as you can will improve the quality and maintainability of your Terraform code.

Terraform-related configuration
In this section we will look at best practices realated to terraform configurations.

#1: Use clear input and output definitions
Name: To demonstrate the purpose of variables and outputs, give them descriptive names.
Type: Always declare variable types to ensure consistency. Avoid using any type or missing type.
Default value: Set default values whenever possible. Having a default value is always better to leaving it empty.
Validations: To verify inputs, use Terraform's built-in validation.
Sensitive: Mark variables or outputs handling secrets as sensitive = true to safeguard sensitive data.
Write documentation: Document variables and outputs, including their purpose, types, and default values.
💡
Use tools like terraform-docs to generate this automatically.
Think about backward compatibility: Adding a variable with a default value is backwards-compatible. Removing a variable is backwards-incompatible; therefore, avoid it wherever possible.
Minimize unnecessary variables: Before creating a variable, consider: "Will this value ever need to change?" If the response is no, use the local value.
#2: Keep Logical Structure
Use a dedicated repository for each Terraform module. Submodules can be included in the module, but they must be directly related to the main module and not implement unrelated functionality.

💡
Use the standard module structure for Terraform modules
Organize module files logically:

main.tf for core resources.
variables.tf for input variables.
outputs.tf for outputs.versions.tf for the minimal terraform version, and provider configurations (if required).
README.md for the repository's basic documentation of the module and all of its nested folders.
To improve readability, split logic into separate files.

For instance, you may split the main.tf file into something like this if your module is in charge of setting up an Application Load Balancer (ALB), EC2 instances, and S3 buckets:

network.tf: Contains the configuration for the Application Load Balancer (ALB) and related networking resources.
vms.tf or ec2.tf: Manages the configuration for EC2 instances.
bucket.tf or s3.tf: Includes the S3 bucket configuration.
Use the following directories:

modules/ for nested or reusable sub-modules.
examples/ for sample configurations and usage examples.
docs/ for additional documentation, such as detailed design or architecture diagrams.
tests/ for infrastructure tests
scripts/ for custom scripts that Terraform may call during execution
helpers/ for organizing helper scripts required for automation or maintenance but not immediately invoked by Terraform.
files/ for static files that Terraform references but does not run. It can be startup scripts, configuration files
wrappers/ for implementing Terragrunt wrapper pattern.
#3: Use nested submodules for complex logic (modules/ directory)
When you dealing with complex Terraform configurations, use nested submodules to break down logic into smaller, reusable units. Follow these best practices:

Place submodules in the modules/ directory using the naming pattern modules/${module_name}.

modules/
├── network/
├── compute/
└── storage/
Nested modules should be considered as private unless explicitly stated in the module’s documentation.

Users should avoid using nested modules directly unless the developer has specified that they are safe and designed for external usage. Consider it, and remember to update the sub-module information in the documentation during the development.
Always refer to the provided documentation for guidance on module usage.
Terraform does not track refactored resources. When you move resources from a top-level module to submodule or just rename them, Terraform may consider them new resources and try to recreate them. This can lead to significant issues for users, such as resource outages or data loss. 

To mitigate this behavior, use moved in scope of refactoring.

moved {
  from = aws_instance.old_name
  to   = module.new_submodule.aws_instance.new_name
}
If refactoring is unavoidable, let users know about it as a significant change in the CHANGELOG and give them detailed guidance on how to manage the changeover.

#4: Try to avoid custom scripts (scripts/ directory)
General rule: Custom scripts should be considered a last resort. Try to avoid them if it is possible. Resources created by such scripts are not tracked or managed by Terraform state, which can lead to inconsistencies.

Exceptions to the rule: Custom scripts may be necessary when Terraform does not support the desired functionality.

Before starting to use scripts, consider these alternatives:

Provider-defined functions: Explore the existing capabilities of your Terraform provider
Try to use alternative providers: For example, in the case of AWS you can use awscc (AWS Cloud Control API) which is generated from CloudFormation resource definitions and has different resource scopes.
3. Use something like TerraCurl for unsupported resources.

You can check more information about it in the article Writing Terraform for unsupported resources.

If custom scripts are indeed necessary and if none of the above options work, follow these two key steps:

Document scripts clearly: Provide a detailed explanation for why the custom script exists. Include a deprecation plan, if possible, and write clear documentation for the script and its usage.
2. Use Provisioners for Execution: Use Terraform’s provisioners (e.g., local-exec) to call custom scripts when needed.

#5: Automate toil with helper scripts (helpers/ directory)
For any repetitive or manual tasks that need automation, create simple scripts and store them in the helpers/ directory. These scripts should abide by the following principles:

Clear Documentation
Provide a brief overview of the script's purpose and why it exists.
Include usage instructions and any other relevant context.
2. Usability Standards

Implement argument validation to handle incorrect or missing inputs gracefully (in case of using bash, you may check “How to Use Bash Getopts With Examples" article).
Include a --help option that explains the script’s functionality, arguments, and examples of usage.
#6: Store static files in a separate directory (files/ directory)
Templates, assets, and other items that Terraform references but does not run directly are known as static files. To properly organize these files, adhere to the following rules:

Separate lengthy HereDocs
Move lengthy HereDoc content into external files for better readability.
Use the file() function to reference these files in your Terraform code.
2. Use .tftpl or .tpl file extensions for templates

When working with Terraform's templatefile() function, use .tftpl or .tpl extensions for template files. While Terraform doesn't enforce this naming standard (you can use any that you like), this extension will help your editor understand the content and might offer a better editing experience as a result.
3. Place Templates in a templates/ subdirectory

Organize all template files within a templates/ subdirectory inside the files/ directory.
A typical example of content within the files/ directory is an AWS Lambda function, commonly placed in files/<your-lambda-function>/.

In this case, the files/<your-lambda-function>/ directory can be considered the root folder, and if the Lambda function is written in Python, this directory may contain an internal structure such as Python packages, dependencies, build scripts or other related files.

#7: Implement Terragrunt Wrapper
Many people use Terragrunt to follow DRY principles. Terragrunt can handle a wide range of tasks across your stack, but module development requires some preparation as well.

In some cases, it is not feasible to use Terraform's native for_each feature. Hence, the Terragrunt wrapper, also known as the single-module wrapper pattern, comes into play. This technique helps in situations like:

You need to deploy multiple copies of a module with different configurations.
Your environment's limitations or complexities prevent you from using native for_each Terraform function.
To apply this approach, establish a new directory wrappers/ with at least three files:

main.tf
outputs.tf
variables.tf
Let’s check them in more detail.

main.tf file contains:

Terraform source
for_each loop over an items variable
All inputs that the module has in the following format: `my_input_variable  = try(each.value.my_input_variable, var.defaults.my_input_variable, {})`
For example:

module "wrapper" {
  source = "../"

  for_each = var.items

  my_input_variable = try(each.value.my_input_variable, 
  var.defaults.my_input_variable, {})
}
variables.tf file contains:

defaults variable, that is, a map of default values. These default values will be used as a backup for each input variable if it is empty or missing.
items variable, which maps items to create a wrapper from
For example:

variable "defaults" {
  description = "A map of default values. These default values will be used as a backup for each input variable if it is empty or missing."
  type        = any
  default     = {}
}

variable "items" {
  description = "Maps of items to create a wrapper from. Values are passed through to the module."
  type        = any
  default     = {}
}
outputs.tf file contains:

wrapper output with a map of outputs from a wrapper.
For example:

output "wrapper" {
  description = "Map of outputs from a wrapper."
  value          = module.wrapper
}
#8: Avoid providers or backend configurations
Shared modules must not configure providers or backends. These settings must be defined in root modules in order to maintain flexibility and reusability.

It is also a good practice to use the required providers block to define only the bare minimum of necessary provider versions. If you need to let the user know about state management in your module, include this information in the README.md file.

versions.tf Example:

terraform {
  required_version = ">= 1.0"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.0.0"
    }
  }
}

As shown in the example, it also contains a minimal Terraform version.

Establishing it is a good practice if you know your code needs a specific version or newer to work properly. If you are unclear about the minimum version required for your module, set >=1.0 as the default definition.

Test the Terraform code
Testing your Terraform code is just as important as testing other types of programming code. A lot of options exist in the Terraform ecosystem to assist you in ensuring the dependability and correctness of your modules.

Let's look at some of the best practices for module testing.

Static analysis
Typically, this involves analyzing the syntax, structure, and deployment configuration using linters, static code analysis tools, or Terraform dry-runs.

These tools are used before the deployment. This topic will be discussed in further detail in the "Code Linting and Quality" section.

Module integration testing
Testing modules in isolation is a useful way to confirm that they work as expected. Module integration testing entails deploying the module to a test environment and ensuring that it generates the expected resources. The following testing frameworks can help you write and manage tests:

Terraform-tests: a built-in Terraform feature that allows you to write tests with HCL.
Terraform-check: a built-in Terraform feature that may help you to validate infrastructure outside the usual resource lifecycle.
Terratest: a Go library that provides patterns and helper functions for testing infrastructure. To gain hands-on experience, start with the article "Automating AWS Infrastructure Testing With Terratest".
terraform-compliance: a lightweight, security and compliancefocused BDD test framework
Google tools
Google's blueprint testing framework
InSpec
End-to-end testing
Integration testing is important to verify that multiple modules work together properly.

In the scope of this method, all modules that exist in your environment are deployed together. Following that, tests are checking that everything works properly using API calls or any other methods.

Although this method is long and expensive, it offers the highest level of verification and can ensure that changes do not cause problems with your production systems.

Some highlights
Don't use old testing tools (like Kitchen-Terraform) and remember to keep the version of tools updated.
Only the first two approaches, static analysis and module integration testing, are relevant in the context of module development. This is because we can't predict how or where the user will use the module. As a result, end-to-end testing is ineffective in this scenario.
All of the tools I outlined for integration testing are also applicable to end-to-end testing.
To learn more about Terraform testing, check the following links;
Testing HashiCorp Terraform
Implement end-to-end Terratest testing on Terraform projects
Write a documentation
A Terraform module's success depends on clear and effective documentation.

A well-structured README.md makes usage easier, reduces complexity, and improves the overall user experience. Conversely, badly organized documentation might have a negative impact on the module's popularity and usefulness.

To create high-quality module documentation, adopt the following principles:

Provide comprehensive examples: Include detailed sample configurations that explain module usage. It's best to give as many examples as possible. Use the examples/ folder for this.
Automate documentation generation: Use terraform-docs to automatically generate and maintain up-to-date documentation
Generate input/output documentation automatically.
Add documentation for:
Detailed examples in the examples/ directory.
Submodules: They should be properly documented, with explanations of their purpose, usage, and specific configurations.
Create a well-structured README.md .Essentially, include the following sections:
An overview of the module.
Examples of usage.
Input and output definitions.
Version compatibility and known limits.
Any badges like “Terraform support” or “OpenTofu support”
Git configuration
Git has various configuration settings that improve project management. The following are essential .git configuration files and their purposes.

.gitignore
This file specifies which files and directories Git should ignore, so they are not tracked or committed to the repository. The Terraform codebase's gitignore file has to include the following exclusions:

Terraform
state files
lock files 
init directory
extra tfvars files
Any build artifacts generated by non-Terraform code, such as Python bytecode files (.pyc) or anything else
Log files
Example:

**/.terraform/*
*.tfstate
*.tfstate.*
crash.log
crash.*.log
*.tfvars
*.tfvars.json

override.tf
override.tf.json
*_override.tf
*_override.tf.json

.terraform.lock.hcl
.terraform.tfstate.lock.info

.terraformrc
terraform.rc

*.pyc
.gitattributes
This file keeps uniform file formatting and prevents line ending issues, especially in projects with contributors from different operating systems. Here is an example of essential configuration:

# Normalize Terraform files
*.tf text

# Enforce LF (Unix-style) line endings
*.sh eol=lf
*.ts text eol=lf
*.json text eol=lf
*.md text eol=lf
*.tf text eol=lf
*.txt text eol=lf

language-configuration.json linguist-language=jsonc
.vscode/**.json linguist-language=jsonc
.gitconfig
This file defines project-specific Git settings for enforcing standards and managing author information. By default, it's good to include the following configuration:

[core] 
autocrlf = input 
longpaths = true
autocrlf: Manages line-ending normalization across operating systems (e.g., Windows, Linux, MacOS). When set to input, Git only changes Windows-style line endings (CRLF) to Unix-style line endings (LF) only when a file is committed

longpaths: Allows Git to support file paths longer than 260 characters, the default maximum path length on Windows. This option prevents errors when working with deeply nested directories or repositories that have long file names.

Use a trunk-based development
Trunk-based development is the most frequently recommended Git branching strategy, according to DevOps Research and Assessment (DORA), and Google DevOps capabilities. With this method, you can always maintain the main branch clean of unnecessary files or artifacts and ready for deployment.

Trunk-based development for Terraform module repositories can be adopted in the following way:

main branch
The main branch contains the most recent code and serves as the main development branch.

It should always remain clean, protected, and ready for use.

Feature and bug-fix branches
Development takes place in the feature and bug-fix branches. These branches are chopped off from the main branch.

Naming convention:

Feature branches: feature/$feature_name or feat/$feature_name
Bug-fix branches: fix/$bugfix_name
Release Branches
A release branch is used to prepare code for the upcoming release. When the release branch is stable and confirmed, it should be merged with the main branch.

Naming convention: release/$version_number (e.g., release/v1.2.0).

Pull Requests
Once a feature or bug fix has been completed, use a pull request to merge it back into the main branch.

To reduce merge conflicts, rebase branches before merging.

GitHub configuration
GitHub provides a range of powerful built-in features to enhance repository management and collaboration.

These features can be activated by adding specific files to your repository, typically within the .github folder. Using these features can help attract contributors, build interest in your repository, and improve the development experience.


tofuutils/tenv: A repository that uses almost every GitHub feature.
PULL_REQUEST_TEMPLATE.md
Adding this file to your repository enables project contributors to see a predefined structure in the pull request body automatically.

It’s considered best practice to include key instructions or questions to help streamline the review process.

Example:

## Description
Provide a summary of the PR changes.

## Checklist
- [ ] Code is formatted (`terraform fmt`).
- [ ] Documentation is updated.
- [ ] Tests are added or updated.

ISSUE_TEMPLATE directory
This directory contains a collection of .md files that are designed to standardize the submission process for issues.

These structured templates provide clear instructions to contributors when they report problems, seek enhancements, or deal with other types of issues.


A dialog window with pre-defined issue templates
In addition to templates, you may change issue submission behavior by including a .github/ISSUE_TEMPLATE/config.yml file in your repository. This configuration file allows you to specify the default behaviors and settings for issue templates. Example:

blank_issues_enabled: true
In this file blank_issues_enabled means:

true: Allows users to submit blank issues without requiring a template.
false: Blocks blank issues entirely, ensuring that all issues are submitted using only predefined templates.
CONTRIBUTING.md
The CONTRIBUTING.md file contains explicit guidelines for contributing to the project. It may include:

Set up the development environment: Follow step-by-step instructions to clone the repository, install dependencies, and run the project locally.
Submitting pull request: Clearly outline the workflow for creating pull requests, including branch naming rules, explicit commit messages, and code standards.
Code review expectations: Define what contributors should expect from the review process and how feedback will be provided.
Issues and feature requests: Explain how contributors can submit bug reports, suggest new features, and participate in discussions.
LICENSE
The LICENSE file specifies the terms under which your project may be used, modified, and distributed. It ensures that users and contributors understand their rights and responsibilities.

It's better to choose a license that aligns with your project's goals:

MIT: A permissive license that allows virtually unrestricted use, modification, and distribution.
Apache 2.0: Provides similar freedoms to MIT but adds protections such as patent rights.
GPL: A copyleft license that mandates derivative works to be open source under the same license.
If you're confused about which license to choose, visit the Choose a License. It provides you with simple instructions and examples to help you choose the best license for your purposes.

Also, as a default license, you can use the MIT license, which is ideal for a wide range of projects, including Terraform modules.

CODE_OF_CONDUCT.md
The CODE_OF_CONDUCT.md file allows you to create community standards, indicate a friendly and inclusive project, and specify anticipated behavior.

Typically, this file may:

Define acceptable and inappropriate user behavior.
Specify how community members can report violations.
Outline the procedures for resolving disputes and ensuring enforcement.
Emphasize creating a welcoming space for everyone, regardless of background or experience.
FUNDING.yml
The FUNDING.md file allows the repository to display funding options. It might provide links to platforms where contributors can make donations to support the project. Examples of such platforms are:

Patreon
Open Collective
GitHub Sponsors
Other appropriate platforms based on your requirements.
Don't forget to explain how the funds will be used, whether for development, community projects, or operational purposes.


A banner that GitHub displays if FUNDING.md is enabled.
CODEOWNERS
The CODEOWNERS file assigns ownership to specified files or directories, which automates pull request reviews. Example:

# Assign maintainers for specific parts of the project.
*.tf             @team-infra
*.md             @alexander-sharov

For further details, check the official GitHub documentation.

MAINTAINERS.md
The MAINTAINERS.md file contains a list of project maintainers, their duties, and contact information. Example:

# MAINTAINERS

## Maintainers List

| Name          | Role                | GitHub Handle      | Contact             |
|---------------|---------------------|--------------------|---------------------|
| Alice Johnson | Lead Developer      | @alice-johnson     | alice@example.com   |
| Bob Smith     | Documentation Lead  | @bob-smith         | bob@example.com     |

## Responsibilities
- **Lead Developer**: Manages core development and architecture.
- **Documentation Lead**: Ensures project documentation is correct and up-to-date.
SECURITY.md
The SECURITY.md file defines the security policy and vulnerability-handling procedure for your project. Example:

# SECURITY POLICY

## Reporting a Vulnerability

If you discover a security vulnerability, please report it by emailing **security@example.com**. Provide as much detail as possible to help us address the issue promptly.

## What to Expect
- We will acknowledge receipt of your report within 72 hours.
- A timeline for the investigation and resolution will be communicated.
SUPPORT.md
The SUPPORT.md file instructs users on how to properly report problems or obtain assistance.

For instance:

# SUPPORT

## How to Get Help

- For general inquiries or usage questions, please open an issue in the repository.
- If you encounter a bug, file a detailed issue with the following information:
  - Steps to reproduce the problem
  - Expected behavior
  - Actual behavior
  - Logs or screenshots, if applicable

## Contact Us

For urgent or private matters, contact **support@example.com**. We aim to respond within 72 hours.
Continuous integration configuration
Continuous Integration (CI) plays an essential role for open-source repositories since it ensures the reliability, quality, and maintainability of the codebase. When many people make changes, CI pipelines help to ensure consistency and prevent bugs from being merged into the main branch.

Although there is no single "ideal" CI pipeline, certain key components should always be covered, including: 

Linting
Formatting 
Code Validation
Compliance Checks
PR verification automation
All of the mentioned checks can be automated with CI tools like Jenkins, GitLab CI, and GitHub Actions.

Every time code is pushed or a pull request (PR) is sent, these checks ought to be executed as CI jobs. In addition, since many people are contributing through pull requests, PR validation and main branch validation must be your top priorities.

To maximize CI pipeline efficiency, implement as many features as possible from the "Code Linting and Quality" section.

Beyond code validation, CI automation should also cover:

Documentation updates: Ensure that documentation stays in sync with code changes.
Dependency Management: Handle package updates efficiently.
Versioning & Releases: Automate version control, release management, and change tracking.
Test automation: CI pipeline should include automated checks for all types of tests. These processes align with best practices stated in the "Test Terraform Code" section. 
Now let's take a closer look at such automation.

Setup Versioning
Each module should have a version to manage its usage effectively. For your releases, I advise using semantic versioning (e.g., v1.0.0):

MAJOR version (1.x.x): Introduces breaking changes.
MINOR version (x.1.x): Adds functionality in a backward-compatible manner.
PATCH version (x.x.1): Fixes bugs or includes improvements without introducing new functionality or breaking backward compatibility.
Ideally, version management should be automated using your CI tool, such as Jenkins, GitLab CI/CD, GitHub Actions, or other similar platforms. In the bulk of my projects, I utilize the GitHub action kvendingoldo/semver-action and other similar platforms, which do the following automatically:

Set the semantic version as a Git tag for each commit.
Create release branches
Create GitHub releases
This configuration is an excellent starting point, but feel free to experiment with different tools that best suit your workflow. Check such links as:

Semantic release
Semantic release action
💡
Note: It's important to use the GitHub Releases features with public modules. If you decide to use a custom versioning process, make sure that you integrate GitHub Releases into your workflow so that your users have clear visibility and accessibility.

GitHub release produced by kvendingoldo/semver-action.
Automate CHANGELOG updates
A CHANGELOG.md is a file that provides a curated, chronologically ordered list of notable changes for each version of a project. While it can be optional for personal projects, it is mandatory for public projects.

This file helps users easily track significant changes between releases or versions.

In the context of Terraform module development, it is important to follow these best practices:

Maintain a CHANGELOG.md file to document changes across versions.

For example:


# Changelog

## [1.0.0] - 2025-01-01
### Added
- Initial release of the Terraform module.

## [0.1.0] - 2024-12-01
### Added
- Added basic AWS VPC code.

### Fixed
- Resolved some typos.
Include changelog updates in pull requests (PRs): Add a note that the PR should include text describing any changes that will be included in the changelog to ensure proper documentation of updates.

Automate changelog updates to streamline the process and maintain consistency. You can use the following approaches:

Release Drafter tool
Conventional Changelog tool
Go-changelog from Hashicorp (You can check Cloudflare as an example)
💡
For additional information on building and maintaining a changelog, visit keep a changelog
Setup Dependency Management
Currently, there are two standard products available for automatic dependency updates that can be used for public Terraform modules.

Let's look at both of them.

1. Dependabot
Dependabot is a GitHub-native tool that automatically checks and updates dependencies. You can configure it to manage Terraform modules, GitHub Actions, Python PIP packages, and any other ecosystem.

Here is a sample configuration for managing Terraform dependencies that checks changes daily and, if found, creates a pull request with a new version:

# .github/dependabot.yml
version: 2
updates:
  - package-ecosystem: "terraform"
    directory: "/"
    schedule:
      interval: "daily"
Key features of Dependabot:

Automatically opens pull requests for dependency updates.
Supports a wide range of ecosystems, including Terraform, npm, Python, GitHub Actions, and more.
Includes built-in GitHub security alarms for dependencies' vulnerabilities.
GitHub native
2. Renovate
Renovate is yet another very customisable tool for managing dependencies. It has more flexibility than Dependabot, making it ideal for projects with complicated dependency management requirements.

Here is an renovate.json sample that does the same function as the DepedaBot configuration:

{
  "extends": ["config:base"],
  "terraform": {
    "enabled": true
  },
  "schedule": ["daily"]
   "automerge": false
}
Key features of Renovate:

Manages several repositories and ecosystems at once.
Provides granular control over update rules, including grouping updates or setting specific policies.
Automatically updates configuration files, such as Terraform’s versions.tf or .tf.lock.hcl.
Integrates with self-hosted and cloud platforms
Work better with non-GitHub platforms like GitLab CI.
Configure jobs to reduce toil
Maintaining a public project frequently requires you to oversee multiple tasks at once. Any project management task that is performed at least twice ought to be automated to optimize your workflow and time management.

Here are two useful automations that you may use while developing Terraform modules to improve project efficiency.

A job that simply closes existing issues and PRs without a new activity.

name: 'Close stale issues and PRs'
on:
 schedule:
   - cron: '0 0 * * *'

jobs:
 stale:
   runs-on: ubuntu-22.04
   steps:
     - uses: actions/stale@v9
       with:
         repo-token: ${{ secrets.GITHUB_TOKEN }}
         days-before-stale: 30
         stale-issue-label: stale
         stale-pr-label: stale
             stale-issue-message: |
           This issue has been automatically marked as stale because it has been open 30 days
           with no activity. Remove stale label or comment or this issue will be closed in 10 days
         stale-pr-message: |
           This PR has been automatically marked as stale because it has been open 30 days
           with no activity. Remove stale label or comment or this PR will be closed in 10 days
         # Not stale if have this labels or part of milestone
         exempt-issue-labels: bug,wip,on-hold
         exempt-pr-labels: bug,wip,on-hold
         exempt-all-milestones: true
         # Close issue operations
         # Label will be automatically removed if the issues are no longer closed nor locked.
         days-before-close: 10
         delete-branch: true
         close-issue-message: This issue was automatically closed because of stale in 10 days
         close-pr-message: This PR was automatically closed because of stale in 10 days
A job that validate PR title

name: 'Validate PR title'

on:
 pull_request_target:
   types:
     - opened
     - edited
     - synchronize

jobs:
 main:
   name: Validate PR title
   runs-on: ubuntu-22.04
   steps:
     - uses: amannn/action-semantic-pull-request@v5.5.3
      env:
         GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
       with:
         # Configure which types are allowed.
         # Default: https://github.com/commitizen/conventional-commit-types
         types: |
           fix
           feat
           docs
           ci
           chore
         # Configure that a scope must always be provided.
         requireScope: false
         # Configure additional validation for the subject based on a regex.
         # This example ensures the subject starts with an uppercase character.
         subjectPattern: ^[A-Z].+$
         # If `subjectPattern` is configured, you can use this property to override
         # the default error message that is shown when the pattern doesn't match.
         # The variables `subject` and `title` can be used within the message.
         subjectPatternError: |
           The subject "{subject}" found in the pull request title "{title}"
           didn't match the configured pattern. Please ensure that the subject
           starts with an uppercase character.
         # For work-in-progress PRs you can typically use draft pull requests
         # from Github. However, private repositories on the free plan don't have
         # this option and therefore this action allows you to opt-in to using the
         # special "[WIP]" prefix to indicate this state. This will avoid the
         # validation of the PR title and the pull request checks remain pending.
         # Note that a second check will be reported if this is enabled.
  wip: true
         # When using "Squash and merge" on a PR with only one commit, GitHub
         # will suggest using that commit message instead of the PR title for the
         # merge commit, and it's easy to commit this by mistake. Enable this option
         # to also validate the commit message for one commit PRs.
         validateSingleCommit: false
GitHub actions usage
Existing GitHub Actions are important building blocks for creating GitHub pipelines.

👨‍💻
My key advice: avoid writing custom actions wherever possible. Instead, explore the GitHub Marketplace to find pre-existing actions that fit your requirements. If a relevant action lacks a specific feature, consider contributing to its development rather than rebuilding the wheel.
Creating custom actions is time-consuming and rarely justified for small to medium-sized organizations. Leveraging existing solutions not only saves effort but also ensures better maintenance and community support.

Code Linting and Quality
Linting is an essential practice for maintaining high-quality Terraform code.

It helps ensure consistency, compliance with best practices, and early detection of potential issues. Well-linted code is easier to review, debug, and manage; therefore, it's an important stage in the development lifecycle.

Let us have a look at linting best practices.

Make sure no sensitive data has been committed to a repository. Use the tools listed below for that:
Gitleaks or Git-Secret CI jobs
Talisman pre-commit checks
Use Infracost to get an approximate breakdown of the cloud infrastructure costs.
Use linters like tflint or terraform validate to catch syntax errors and enforce Terraform best practices.
Format all code using terraform fmt and enforce this via pre-commit hooks.
To detect compliance and security issues, use the security linters listed below:
Checkov: An open-source static code analysis tool developed by Bridgecrew. It supports multiple IaC frameworks, including Terraform, CloudFormation, Kubernetes, Helm charts, and Dockerfiles. Checkov performs comprehensive scans to detect misconfigurations and policy violations, helping maintain adherence to security best practices.
Terrascan: Developed by Tenable, Terrascan is designed to detect compliance and security violations across various IaC frameworks, including Terraform, CloudFormation, ARM Templates, Kubernetes, Helm, Kustomize, and Dockerfiles. It uses the Open Policy Agent (OPA) for policy as code, allowing for customizable and extensive policy enforcement.
Trivy: Aqua Security's Trivy is a versatile security scanner that can detect vulnerabilities in container images, file systems, and Git repositories. It has Terraform support, which allows you to detect security vulnerabilities in your IaC configurations. Notably, tfsec, previously recommended for Terraform security scanning, has been incorporated into Trivy, combining their capabilities (tfsec#1994).
KICS (Keeping Infrastructure as Code Secure): Checkmarx developed KICS, which identifies security vulnerabilities, compliance issues, and infrastructure misconfigurations early in the development cycle. It supports multiple types of IaC frameworks, allowing you to secure your infrastructure before deployment.
To learn more about static code analysis tools, see "IaC Security Analysis: Checkov vs. tfsec vs. Terrascan – A Comparative Evaluation" and "Comparing Checkov vs. tfsec vs. Terrascan" articles. 
Use Open Policy Agent (OPA). It is a policy-as-code engine that enables flexible and automated enforcement of compliance, security, and operational rules within Terraform IAC configurations.

OPA uses the Rego language to establish custom policies, validate Terraform files, and plan outputs for misconfigurations.

It works smoothly with CI/CD pipelines to guarantee policy compliance before deployment and includes pre-built policy libraries for typical use cases. OPA contributes to the maintenance of secure and consistent infrastructure standards by identifying concerns early on.
If you use GitLab, then use GitLab IaC Scanning. Infrastructure as code scanning is a built-in feature of the platform, available even on the free plan. It includes a preset CI/CD template that performs helpful static analysis tests on the IaC files in your project. 
Use GitHub Super-Linter to maintain consistent code quality across multiple file types, such as YAML, JSON, Markdown, and Terraform.

If Super-Linter is not an acceptable choice (e.g., you can have self-hosted Jenkins), try using yamllint or jsonlint to effectively validate and format YAML and JSON files. Example of a super-linter GitHub workflow:

name: Super-Linter
on: [push, pull_request]
jobs:
  lint:
    runs-on: ubuntu-22.04
    steps:
      - uses: actions/checkout@v3
      - uses: github/super-linter@v4
        env:
          VALIDATE_TERRAFORM: true
          VALIDATE_YAML: true
          VALIDATE_JSON: true
          VALIDATE_MARKDOWN: true
Developer-environment configuration
Now let's look at some of the key Developer-environment configurations.

#1: Configure editor
Developers use a wide range of IDEs and text editors, from lightweight alternatives like Vim to powerful, full-featured IDEs like IntelliJ IDEA or VSCode.

It is highly recommended to add a .editorconfig file in your repository to ensure uniform code formatting and standards across different development environments.

The .editorconfig file is supported by most modern IDEs and text editors.

By setting shared code conventions, teams can ensure that developers, regardless of the tools they use, follow the same formatting guidelines. This minimizes issues caused by inconsistent code formatting, streamlines collaboration, and promotes a cleaner codebase.

Example:

# EditorConfig is awesome: https://editorconfig.org

# top-most EditorConfig file
root = true

[*]
charset = utf-8
end_of_line = lf
indent_size = 2
indent_style = space
insert_final_newline = true
max_line_length = 80
trim_trailing_whitespace = true

[*.{tf,tfvars}]
indent_size = 2
indent_style = space

[*.md]
max_line_length = 0
trim_trailing_whitespace = false

[Makefile]
tab_width = 2
indent_style = tab
[COMMIT_EDITMSG]
max_line_length = 0
Aside from editorconfig, the second good option is to keep the configuration for VSCode, which will be used by the majority of developers in 2025.

Example of .vscode/settings.json:

{
  "editor.formatOnSave": true,
  "terraform.format": {
    "enable": true
  },
  "editor.tabSize": 2
}
Either is not a recommendation, but it is a good idea to specify all plugins that may be relevant during development.

Example of .vscode/extensions.json:

{
  "recommendations": [
    "hashicorp.terraform",
    "redhat.vscode-yaml",
    "streetsidesoftware.code-spell-checker"
  ]
}
#2: Use a pre-commit framework
The pre-commit framework is a powerful tool for automating code quality checks and enforcing standards before changes are committed to your repository. It provides immediate feedback to developers, helping them catch and address issues in their personal workspace before committing the code.

These hooks check the state of the code before committing and can stop the process if tests fail, ensuring that only high-quality code is pushed to the repository.

To use it, you can create .pre-commit-config.yaml in the root of your repository, and run it manually before the commit via pre-commit run --all-files. 

The following example contains multiple best practices, which we have already examined in the article:

# .pre-commit-config.yaml

repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.5.0
    hooks:
      - id: check-yaml
      - id: check-json
      - id: trailing-whitespace
        args: ["--markdown-linebreak-ext=md"]
      - id: check-added-large-files
      - id: check-executables-have-shebangs
      - id: check-shebang-scripts-are-executable
      - id: check-merge-conflict
      - id: check-vcs-permalinks
      - id: detect-private-key
      - id: detect-aws-credentials
        args: ["--allow-missing-credentials"]
      - id: end-of-file-fixer
      - id: no-commit-to-branch
      - id: pretty-format-json
        args:
          - --autofix
  - repo: https://github.com/zricethezav/gitleaks
    rev: v8.18.2
    hooks:
      - id: gitleaks
  - repo: https://github.com/bridgecrewio/checkov.git
    rev: 3.2.43
    hooks:
     - id: checkov
        args:
          - --config-file
          - .checkov.yaml
  - repo: https://github.com/antonbabenko/pre-commit-terraform
    rev: v1.88.3
    hooks:
      - id: terraform_fmt
      - id: terraform_validate
        args:
        - --hook-config=--retry-once-with-cleanup=true
      - id: terraform_docs
        args:
          - --args=--config=.terraform-docs.yaml
      - id: terraform_tflint
        args:
          - --args=--config=__GIT_WORKING_DIR__/.tflint.hcl
You may use the AWS Terraform template repository as a starting point for developing your pre-commit configuration.

It offers a large number of standard checks that you can apply to your project. Furthermore, pay attention to these two repositories that provide common pre-commit hooks for Terraform/OpenTofu:

Pre-commit hooks for Terraform
Pre-commit hooks for OpenTofu
Terraform-module template pattern
Previously in this post, we explored several guidelines relevant to most Terraform modules.

However, managing tens or hundreds of modules distributed across multiple repositories makes it practically impossible to keep all configurations, CI processes, and linting rules up to date.  Using a copy-pasting method is neither efficient nor sustainable.

To make a maintenance task easier, use a GitHub/GitLab repository or GitHub fork methodology for your Terraform modules. This repository can contain all of the essential CI pipelines, linters, tools, and code configurations, which are implemented as a common template for your Terraform modules. 

In my projects, I use both of these approaches under the moniker "Terraform Module Template Repository". This repository allows me to constantly keep my configurations up to date and deliver updates rapidly.

Furthermore, users can quickly build a new module by clicking via "Use this template button" within the template repository.

Before we go any further, I'd want to discuss the differences between a fork and a template repository, which will be used as a template for future repositories.

If you use the GitHub template, it may be difficult to transmit updates from the parent repository to the child repository because they are technically independent.

In the case of forks, the repositories are linked, but I do not suggest backmerging, because you may spend a lot of time at merge conflicts.

In the scope of both methods, I advocate implementing a custom CI job in the template repository that will distribute updates to children via easy Git and bash commands manually or by CRON. We'll look at an example of such a job in Part 2 of this article series.

Returning to the "Terraform Module Template Repository" approach, it's important to note that applying it for personal projects or smaller-scale infrastructure may be overkill due to the amount of required work (believe me!). 

The upfront costs of creating and maintaining a reusable, parameterized template may exceed the benefits. If you only have a few repositories, it can be much easier to copy-paste configurations, without implementing an additional automation.

Finally, if you expect to manage dozens or hundreds of modules, you should create a Terraform Module Template Repository from the ground up at the beginning stage of project development.

For smaller-scale cases, you can look for open-source alternatives (one of which will be discussed in Part 2 of this series), or just copy and paste your preferred configurations between Terraform modules.

Conclusion
Thank you for reading! I hope you found this guide helpful. Stay tuned for Part 2, where I will cover:

A ready-to-use Terraform module template that implements the majority of best practices outlined in this article.
An example of CI jobs that deliver updates from the template repository to their children repositories.
Practical steps to adopt a ready-to-use module template for your organization in 5 minutes
More code samples, as well as tips and advice based on projects that have already used the template repository and other best practices outlined in this article.

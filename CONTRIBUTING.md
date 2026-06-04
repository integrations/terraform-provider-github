# Contributing

Hi there! We're thrilled that you'd like to contribute to this project. Your help is essential to keep it great.

Contributions to this project are [released](https://help.github.com/articles/github-terms-of-service/#6-contributions-under-repository-license) to the public under the [project's open source license](LICENSE).

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.

Before submitting an issue or a pull request, please search the repository for existing content. Issues and PRs with multiple comments, reactions, and experience reports increase the likelihood of a merged change.

## Table of Contents <!-- omit from toc -->

- [AI Use Policy and Guidelines](#ai-use-policy-and-guidelines)
  - [Using AI as a Coding Assistant](#using-ai-as-a-coding-assistant)
  - [Using AI for Communication](#using-ai-for-communication)
- [Submitting a pull request](#submitting-a-pull-request)
- [Documentation Update Checklist](#documentation-update-checklist)
- [Quick End-To-End Example](#quick-end-to-end-example)
  - [Local Development Setup](#local-development-setup)
  - [Local Development Iteration](#local-development-iteration)
  - [Debugging the terraform provider](#debugging-the-terraform-provider)
- [Manual Testing](#manual-testing)
  - [Using a local version of the provider](#using-a-local-version-of-the-provider)
  - [Cleaning Up Test Resources](#cleaning-up-test-resources)
  - [GitHub Organization](#github-organization)
- [Environment Variable Reference](#environment-variable-reference)
  - [Example _.vscode/settings.json_ file](#example-vscodesettingsjson-file)

## AI Use Policy and Guidelines

Our goal in this project is to develop a stable and well-maintained provider library. This requires careful attention to detail in every change we integrate. Maintainer time and attention is very limited, so it's important that changes you ask us to review represent your _best_ work.

You are encouraged to use tools that help you write good code, including AI tools. However, you always need to understand and explain the changes you're proposing to make, whether or not you used an LLM as part of your process to produce them. The answer to "Why did you make change X?" should never be "I'm not sure. The AI did it."

**Do not submit an AI-generated PR you haven't personally understood and tested**, as this wastes maintainers' time. PRs that appear to violate this guideline will be closed without review. If you do submit a largely AI-generated PR, clearly mark it as such in the description. Note that maintainers may still close it without further review if it does not seem worthwhile.

### Using AI as a Coding Assistant

1. Don't skip **becoming familiar with the part of the codebase** you're working on. This will let you write better prompts and validate their output if you use an LLM. Code assistants can be a useful search engine/discovery tool in this process, but don't trust claims they make about how Terraform, Terraform providers or the GitHub API works. LLMs are often wrong, even about details that are clearly answered in [Terraform Provider SDKv2 documentation](https://developer.hashicorp.com/terraform/plugin/sdkv2), [Terraform documentation](https://developer.hashicorp.com/terraform), or [GitHub API documentation](https://docs.github.com/en/rest).
2. Split up your changes into **coherent commits**, even if an LLM generates them all in one go. This makes it easier for maintainers to review and understand your changes, and also helps you keep track of your own work.
3. Don't simply ask an LLM to add **code comments**, as it will likely produce a bunch of text that unnecessarily explains what's already clear from the code. If using an LLM to generate comments, be really specific in your request, demand succinctness, and carefully edit the result.

### Using AI for Communication

Contributors are expected to communicate with intention, to avoid wasting maintainer time with long, sloppy writing. We strongly prefer clear and concise communication about points that actually require discussion over long AI-generated comments.

When you use an LLM to write a message for you, it remains **your responsibility** to read through the whole thing and make sure that it makes sense to you and represents your ideas concisely. A good rule of thumb is that if you can't make yourself carefully read some LLM output that you generated, nobody else wants to read it either.

Here are some concrete guidelines for using LLMs as part of your communication workflows:

1. When writing a pull request description, **do not include anything that's obvious** from looking at your changes directly (e.g., files changed, functions updated, etc.). Instead, focus on the _why_ behind your changes. Don't ask an LLM to generate a PR description on your behalf based on your code changes, as it will simply regurgitate the information that's already there.
2. Similarly, when responding to a pull request comment, **explain _your_ reasoning**. Don't prompt an LLM to re-describe what can already be seen from the code.
3. Verify that **everything you write is accurate**, whether or not an LLM generated any part of it. The maintainers will be unable to review your contributions if you misrepresent your work (e.g., wrongly describing your code changes, their effect, or your testing process).
4. Complete all parts of the **PR description template**, including the checklists. Don't simply overwrite the template with LLM output.
5. **Clarity and succinctness** are much more important than perfect grammar, so you shouldn't feel obliged to pass your writing through an LLM. If you do ask an LLM to clean up your writing style, be sure it does _not_ make it longer in the process. Demand succinctness in your prompt.
6. Quoting an LLM answer is usually less helpful than linking to **relevant primary sources**, like source code, or reference materials. If you do need to quote an LLM answer in a discussion, clearly distinguish LLM output from your own thoughts.

## Submitting a pull request

1. Fork and clone the repository.
2. Create a new branch: `git switch -c my-branch-name`.
3. Make your change, add tests, and make sure the tests still pass.
4. Make sure the [documentation has been updated](#documentation-update-checklist)
5. Ensure formatting and linting are passing. (`make fmt` and `make lint` can be used to check this locally.)
6. Push to your fork and submit a pull request.
7. Pat yourself on the back and wait for your pull request to be reviewed and merged.

Here are a few things you can do that will increase the likelihood of your pull request being accepted:

- Allow your pull request to receive edits by maintainers.
- Discuss your changes with the community in an issue.
- Write tests.
- Keep your change as focused as possible. If there are multiple changes you would like to make that are not dependent upon each other, please submit them as separate pull requests.
- Write a [good commit message](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).

## Documentation Update Checklist

When your change impacts a resource, data source, or provider behavior, make sure documentation changes include all of the following:

1. Ensure `Description` fields in Resource/Data Source schemas are up to date.
2. Update the relevant template files when needed (for example, `templates/data-sources/users.md.tmpl`).
3. Update `examples/**` where necessary.
   - Always add at least one example for any new resource or data source.
   - Ensure changed fields are reflected in examples.
   - Ensure complex fields (such as List, Set, and Map fields) have examples.
   - Ensure all examples are valid Terraform configuration.
4. Regenerate docs with `make generatedocs`.
5. Review generated output and confirm it matches expected behavior and schema.

## Quick End-To-End Example

This section describes a typical sequence performed when developing locally. Full details of available tooling are available in the section on [Manual Testing](#manual-testing).

### Local Development Setup

Once you have the repository cloned, there's a couple of additional steps you'll need to take. Since most of the testing is acceptance or integration testing, we need to manipulate real GitHub resources in order to run it. Useful setup steps are listed below:

- If you haven't already, [create a GitHub organization you can use for testing](#github-organization).
  - Optional: you may find it beneficial to create a test user as well in order to avoid potential rate-limiting issues on your main account.
  - Your organization _must_ have a repository called `terraform-template-module`. The [terraformtesting/terraform-template-module](https://github.com/terraformtesting/terraform-template-module) repo is a good, re-usable example.
    - You _must_ make sure that the "Template Repository" item in Settings is checked for this repo.
- If you haven't already, generate a Personal Access Token (PAT) for authenticating your test runs.
- Export the necessary configuration for authenticating your provider with GitHub

  ```sh
  export GH_TEST_AUTH_MODE="organization"
  export GITHUB_OWNER="<name of an organization>"
  export GITHUB_USERNAME="<username of the user who created the token>"
  export GITHUB_TOKEN="<token of a user with an organization account>"
  export GITHUB_LEGACY_CLIENT="false"
  ```

- Build the project with `make build`
- Try an example test run from the default (`main`) branch, like `TF_LOG=DEBUG make testacc T=TestAccGithubRepositories`. All those tests should pass.

### Local Development Iteration

1. Write a test describing what you will fix. See [`github_ip_ranges`](./github/data_source_github_ip_ranges_test.go) for an example using the preferred `ConfigStateChecks` pattern, and [ARCHITECTURE.md](ARCHITECTURE.md#test-structure) for full guidance.
2. Run your test and observe it fail. Enabling debug output allows for observing the underlying requests and responses made during the acceptance test run.

```sh
TF_LOG=DEBUG make testacc T=TestAccGithubIssueLabel
```

1. Align the resource's implementation to your test case and observe it pass:

```sh
make testacc T=TestAccGithubIssueLabel
```

Note that some resources still use a previous format that is incompatible with automated test runs, which depend on using the `skipUnlessMode` helper. When encountering these resources, tests should be rewritten to the latest format.

Also note that there is no build / `terraform init` / `terraform plan` sequence here. It is uncommon to run into a bug or feature that requires iteration without using tests. When these cases arise, the `examples/` directory is used to approach the problem, which is detailed in the next section.

### Debugging the terraform provider

Println debugging can easily be used to obtain information about how code changes perform. If the `TF_LOG=DEBUG` level is set, debug messages will be printed. Use `tflog.Debug(ctx, "your message here", map[string]any{...})` for new code. Some existing code still uses `log.Printf("[DEBUG] ...")` — see [ARCHITECTURE.md](ARCHITECTURE.md#logging) for the migration pattern.

If a full debugger is desired, VSCode may be used. In order to do so,

1. Create a launch.json file with this configuration:

```json
{
 "name": "Attach to Process",
 "type": "go",
 "request": "attach",
 "mode": "local",
 "processId": 0,
}
```

Setting a `processId` of 0 allows a dropdown to select the process of the provider.

1. Add a sleep call (e.g. `time.Sleep(10 * time.Second)`) in `providerConfigure` (in `github/provider.go`) before the immediate `return` call. This will allow time to connect the debugger while the provider is initializing, before any critical logic happens.

2. Build the terraform provider with debug flags enabled and copy it to the appropriate bin folder with a command like `go build -gcflags="all=-N -l" -o ~/go/bin/`.

3. Create or edit a `dev.tfrc` that points toward the newly-built binary, and export the `TF_CLI_CONFIG_FILE` variable to point to it. Further instructions on this process may be found in the [Building the provider](#using-a-local-version-of-the-provider) section.

4. Run a terraform command (e.g. `terraform apply`). While the provider pauses on initialization, go to VSCode and click "Attach to Process". In the search box that appears, type `terraform-provi` and select the terraform provider process.

5. The debugger is now connected! During a typical terraform command, the plugin will be invoked multiple times. If the debugger disconnects and the plugin is invoked again later in the run, the developer will have to re-attach each time as the process ID changes.

## Manual Testing

> **Note:** Automated test coverage is incomplete ([#1414](https://github.com/integrations/terraform-provider-github/issues/1414)). Manual testing on each PR is essential until this is resolved.

### Using a local version of the provider

Build the provider with debug flags for attaching a debugger:

```sh
go build -gcflags="all=-N -l" -o ~/go/bin/
```

This enables verifying your locally built provider using examples available in the `examples/` directory.
Note that you will first need to configure your shell to map our provider to the local build:

```sh
export TF_CLI_CONFIG_FILE=path/to/project/examples/dev.tfrc
```

An example file is available in our `examples` directory and resembles:

```hcl
provider_installation {
  dev_overrides {
    "integrations/github" = "~/go/bin/"
  }

  direct {}
}
```

See <https://www.terraform.io/docs/cli/config/config-file.html> for more details.

When running examples, you should spot the following warning to confirm you are using a local build:

```console
Warning: Provider development overrides are in effect

The following provider development overrides are set in the CLI configuration:
 - integrations/github in /Users/jcudit/go/bin
```

See the [Environment Variable Reference](#environment-variable-reference) below for the full list of configuration options.

There are also a small number of unit tests in the provider. Due to the nature of the provider, such tests are currently only recommended for exercising functionality completely internal to the provider. These may be executed by running `make test`.

### Cleaning Up Test Resources

Acceptance tests create real GitHub resources prefixed with `tf-acc-test-`. If tests fail or are interrupted, these resources may be left behind. Run the sweeper to clean them up:

```sh
make sweep
```

This removes leaked test repositories and teams matching the `tf-acc-test-` prefix.

### GitHub Organization

If you do not have an organization already that you are comfortable running tests against, you will need to [create one](https://help.github.com/en/articles/creating-a-new-organization-from-scratch). The free "Team for Open Source" org type is fine for these tests. The name of the organization must then be exported in your environment as `GITHUB_OWNER`.

Make sure that your organization has a `terraform-template-module` repository ([terraformtesting/terraform-template-module](https://github.com/terraformtesting/terraform-template-module) is an example you can clone) and that its "Template repository" item in Settings is checked.

If you are interested in using and/or testing GitHub's [Team synchronization](https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github) feature, please contact a maintainer as special arrangements can be made for your convenience.

## Environment Variable Reference

Commonly required environment variables are listed below:

```sh
# Enable debug logging
export TF_LOG=DEBUG

# Enables acceptance tests
export TF_ACC="1"

# Configure the URL override for GHES.
export GITHUB_BASE_URL=

# Use the modern client for testing.
export GITHUB_LEGACY_CLIENT="false"

# Configure acceptance testing mode; one of anonymous, individual, organization, team or enterprise. If not set will default to anonymous
export GH_TEST_AUTH_MODE=

# Configure authentication for testing
export GITHUB_OWNER=
export GITHUB_USERNAME=
export GITHUB_TOKEN=

# Configure user level values
export GH_TEST_USER_REPOSITORY=

# Configure values for the organization under test
export GH_TEST_ORG_USER=
export GH_TEST_ORG_SECRET_NAME=
export GH_TEST_ORG_REPOSITORY=
export GH_TEST_ORG_TEMPLATE_REPOSITORY=
export GH_TEST_ORG_APP_INSTALLATION_ID=

# Configure external (non-org) users
export GH_TEST_EXTERNAL_USER=
export GH_TEST_EXTERNAL_USER_TOKEN=
export GH_TEST_EXTERNAL_USER2=

# Configure values for the enterprise under test
export GH_TEST_ENTERPRISE_EMU_GROUP_ID=
export GITHUB_ENTERPRISE_SLUG=

# Configure test options
export GH_TEST_ADVANCED_SECURITY=

# Configure if the enterprise is an EMU enterprise
export GH_TEST_ENTERPRISE_IS_EMU=
```

### Example _.vscode/settings.json_ file

To run acceptance tests the `TF_ACC` environment variable must be set. Below is an example `settings.json` file for VSCode that sets this variable and the other necessary environment variables when running tests from the editor.

```json
{
  "go.testEnvVars": {
    "TF_ACC": "1",
    "GITHUB_TOKEN": "<TOKEN>",
    "GITHUB_BASE_URL": "https://api.github.com/",
    "GITHUB_ENTERPRISE_SLUG": "",
    "GITHUB_OWNER": "<ORGANIZATION>",
    "GITHUB_USERNAME": "<USERNAME>",
    "GITHUB_LEGACY_CLIENT": "false",
    "GH_TEST_AUTH_MODE": "organization",
    "GH_TEST_USER_REPOSITORY": "",
    "GH_TEST_ORG_USER": "",
    "GH_TEST_ORG_SECRET_NAME": "",
    "GH_TEST_ORG_REPOSITORY": "",
    "GH_TEST_ORG_TEMPLATE_REPOSITORY": "",
    "GH_TEST_ORG_APP_INSTALLATION_ID": "",
    "GH_TEST_EXTERNAL_USER": "",
    "GH_TEST_EXTERNAL_USER_TOKEN": "",
    "GH_TEST_EXTERNAL_USER2": "",
    "GH_TEST_ADVANCED_SECURITY": "false",
  },
  "go.testTimeout": "3600s",
  "go.testFlags": [
    "-v",
    "-count=1",
  ]
}
```

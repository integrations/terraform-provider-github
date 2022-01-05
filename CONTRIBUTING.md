# Contributing

Hi there! We're thrilled that you'd like to contribute to this project. Your help is essential for keeping it great.

Contributions to this project are [released](https://help.github.com/articles/github-terms-of-service/#6-contributions-under-repository-license) to the public under the [project's open source license](LICENSE).

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.

## Submitting a pull request

0. Fork and clone the repository
0. Create a new branch: `git checkout -b my-branch-name`
0. Make your change, add tests, and make sure the tests still pass
0. Push to your fork and submit a pull request
0. Pat your self on the back and wait for your pull request to be reviewed and merged.

Here are a few things you can do that will increase the likelihood of your pull request being accepted:

- Discuss your changes with the community in an issue.
- Allow your pull request to receive edits by maintainers.
- Write tests.
- Keep your change as focused as possible. If there are multiple changes you would like to make that are not dependent upon each other, consider submitting them as separate pull requests.
- Write a [good commit message](http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html).

## Quick End-To-End Example

This section describes a typical sequence performed when developing locally. Full details of available tooling are available in the next section on [Automated And Manual Testing](#automated-and-manual-testing).

### Local Development Setup

Once you have the repository cloned, there's a couple of additional steps you'll need to take. Since most of the testing is acceptance or integration testing, we need to manipulate GitHub resources in order to run it. Useful setup steps are listed below:

- If you haven't already, [create a GitHub organization you can use for testing](#github-organization).
  - Optional: some may find it beneficial to create a test user as well in order to avoid potential rate-limiting issues on your main account.
  - Your organization _must_ have a repository called `terraform-module-template`. The [terraformtesting/terraform-template-module](https://github.com/terraformtesting/terraform-template-module) repo is a good, re-usable example.
    - You _must_ make sure that the "Template Repository" item in Settings is checked for this repo.
- If you haven't already, [generate a Personal Access Token (PAT) for authenticating your test runs](#github-personal-access-token).
- Export the necessary configuration for authenticating your provider with GitHub
  ```sh
  export GITHUB_TOKEN=<token of a user with an organization account>
  export GITHUB_ORGANIZATION=<name of an organization>
  ```
- Build the project with `make build`
- Try an example test run from the default (`master`) branch, like `TF_LOG=DEBUG TF_ACC=1 go test -v   ./... -run ^TestAccGithubRepositories`. All those tests should pass.

### Local Development Iteration

1. Write a test describing what you will fix. See [`github_label`](./github/resource_github_issue_label_test.go) for an example format.
1. Run your test and observe it fail. Enabling debug output allows for observing the underlying requests and responses made as well as viewing state (search `STATE:`) generated during the acceptance test run.
```sh
TF_LOG=DEBUG TF_ACC=1  go test -v   ./... -run ^TestAccGithubIssueLabel
```
1. Align the resource's implementation to your test case and observe it pass:
```sh
TF_ACC=1  go test -v   ./... -run ^TestAccGithubIssueLabel
```

Note that some resources still use a previous format that is incompatible with automated test runs, which depend on using the `skipUnlessMode` helper. When encountering these resources, tests are rewritten to the latest format.

Also note that there is no build / `terraform init` / `terraform plan` sequence here.  It is uncommon to run into a bug or feature that requires iteration without using tests. When these cases arise, the `examples/` directory is used to approach the problem, which is detailed in the next section.

### Debugging the terraform provider

Println debugging can easily be used to obtain information about how code changes perform. If the `TF_LOG=DEBUG` level is set, calls to `log.Printf("[DEBUG] your message here")` will be printed in the program's output.

If a full debugger is desired, VSCode may be used. In order to do so,

1. create a launch.json file with this configuration:
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

2. Add a sleep call (e.g. `time.Sleep(15 * time.Second)`) in the [func providerConfigure(p *schema.Provider](https://github.com/integrations/terraform-provider-github/blob/main/github/provider.go#L176) before the immediate `return` call. This will allow time to connect the debugger while the provider is initializing, before any critical logic happens.

2. Build the terraform provider with debug flags enabled and copy it to a bin folder with a command like `go build -gcflags="all=-N -l" -o ~/go/bin`.

3. Create or edit a `dev.tfrc` that points toward the newly-built binary, and export the `TF_CLI_CONFIG_FILE` variable to point to it. Further instructions on this process may be found in the [Building the provider](#building-the-provider) section.

4. Run a terraform command (e.g. `terraform apply`). While the provider pauses on initialization, go to VSCode and click "Attach to Process". In the search box that appears, type `terraform-provi` and select the terraform provider process.

5. The debugger is now connected! During a typical terraform command, the plugin may be invoked multiple times. If the debugger disconnects and the plugin is invoked again later in the run, the developer will have to re-attach each time as the process ID changes.


## Automated And Manual Testing

### Overview

When raising a pull request against this project, automated tests will be launched to run a subset of our test suite.

Full acceptance testing is run [daily][acc-daily]. In line with Terraform Provider testing best practices, these tests exercise against a live, public GitHub deployment (referred to as `dotcom`). Tests may also run against an Enterprise GitHub deployment (referred to as `ghes`), which is sometimes available during parts of a month. If your change requires testing against a specific version of GitHub, please let a maintainer know and this may be arranged.

Partial acceptance testing can be run manually by creating a branch prefixed with `test/`.  Simple detection of changes and related test files is performed and a subset of acceptance tests are run against commits to these branches. This is a useful workflow for reviewing PRs submitted by the community, but local testing is preferred for contributors while iterating towards publishing a PR.

### Building The Provider

Clone the provider
```sh
$ git clone git@github.com:integrations/terraform-provider-github.git
```

Enter the provider directory and build the provider while specifying an output directory:

```sh
$ go build -o ~/go/bin/
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

See https://www.terraform.io/docs/cli/config/config-file.html for more details.

When running examples, you should spot the following warning to confirm you are using a local build:

```console
Warning: Provider development overrides are in effect

The following provider development overrides are set in the CLI configuration:
 - integrations/github in /Users/jcudit/go/bin
```

### Developing The Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.13+ is *required*).

You may also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`. Recent Go releases may have removed the need for this step however.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-github
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
# run all tests through `make`
$ make testacc
# run all tests directly
$ go test -v   ./...
# run specific test
$ go test -v   ./... -run TestAccProviderConfigure
```

Commonly required environment variables are listed below:

```sh
# enable debug logging
export TF_LOG=DEBUG

# enable testing of organization scenarios instead of individual or anonymous
export GITHUB_ORGANIZATION=

# enable testing of individual scenarios instead of organizaiton or anonymous
export GITHUB_OWNER=

# enable testing of enterprise appliances
export GITHUB_BASE_URL=

# leverage helper accounts for tests requiring them
# examples include:
# - https://github.com/github-terraform-test-user
# - https://github.com/terraformtesting
export GITHUB_TEST_OWNER=
export GITHUB_TEST_ORGANIZATION=
export GITHUB_TEST_USER_TOKEN=
```

See [this project](https://github.com/terraformtesting/acceptance-tests) for more information on how tests are run automatically.

### GitHub Personal Access Token

You will need to create a [personal access token](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line) for
testing. It will need to have the following scopes selected:
* repo
* admin:org
* admin:public_key
* admin:repo_hook
* admin:org_hook
* user
* delete_repo
* admin:gpg_key

Once the token has been created, it must be exported in your environment as `GITHUB_TOKEN`.

### GitHub Organization

If you do not have an organization already that you are comfortable running tests against, you will need to [create one](https://help.github.com/en/articles/creating-a-new-organization-from-scratch). The free "Team for Open Source" org type is fine for these tests. The name of the organization must then be exported in your environment as `GITHUB_ORGANIZATION`.

Make sure that your organization has a `terraform-module-template` repository ([terraformtesting/terraform-template-module](https://github.com/terraformtesting/terraform-template-module) is an example you can clone) and that its "Template repository" item in Settings is checked.

If you are interested in using and/or testing Github's [Team synchronization](https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github) feature, please contact a maintainer as special arrangements can be made for your convenience.

## Resources

- [How to Contribute to Open Source](https://opensource.guide/how-to-contribute/)
- [Using Pull Requests](https://help.github.com/articles/about-pull-requests/)
- [GitHub Help](https://help.github.com)


[acc-daily]: https://github.com/integrations/terraform-provider-github/actions?query=workflow%3A%22Acceptance+Tests+%28All%29%22

## Contributing

Hi there! We're thrilled that you'd like to contribute to this project. Your help is essential for keeping it great.

Contributions to this project are [released](https://help.github.com/articles/github-terms-of-service/#6-contributions-under-repository-license) to the public under the [project's open source license](LICENSE.md).

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

Enter the provider directory and build the provider while specifying the output directory:

```sh
$ go build -o ~/.terraform.d/plugins/terraform-provider-github
```

This enables verifying your locally built provider using examples available in the `examples/` directory. Just ensure you are not specifying a provider version so that `terraform init` falls back to using the build found under `~/.terraform.d/plugins/terraform-provider-github`.  If your example directory already has a `.terraform` directory from a previous run, remove that directory and `terraform init` again to consume the local build.

### Developing The Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). 

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

If you are interested in using and/or testing Github's [Team synchronization](https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github) feature, please contact a maintainer as special arrangements can be made for your convenience.

## Resources

- [How to Contribute to Open Source](https://opensource.guide/how-to-contribute/)
- [Using Pull Requests](https://help.github.com/articles/about-pull-requests/)
- [GitHub Help](https://help.github.com)


[acc-daily]: https://github.com/integrations/terraform-provider-github/actions?query=workflow%3A%22Acceptance+Tests+%28All%29%22
Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.svg)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-github`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone https://github.com/terraform-providers/terraform-provider-github.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-github
$ make build
# or if you're on a mac:
$ gnumake build
```

Using the provider
----------------------

Detailed documentation for the GitHub provider can be found [here](https://www.terraform.io/docs/providers/github/index.html).

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

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
$ make testacc
```

Acceptance test prerequisites
-----------------------------
In order to successfully run the full suite of acceptance tests, you will need to have the following:

export `https://api.github.com/` as the environment variable `GITHUB_BASE_URL`.

### GitHub personal access token
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

### GitHub organization
If you do not have an organization already that you are comfortable running tests against, you will need to [create one](https://help.github.com/en/articles/creating-a-new-organization-from-scratch). The free "Team for Open Source" org type is fine for these tests. The name of the
organization must then be exported in your environment as `GITHUB_OWNER`. If you are interested in using and/or testing Github's [Team synchronization](https://help.github.com/en/github/setting-up-and-managing-organizations-and-teams/synchronizing-teams-between-your-identity-provider-and-github) feature, you will need to have an organization that uses Github Enterprise Cloud in addition to the requirements defined in the Github docs and set the environment variable `ENTERPRISE_ACCOUNT` to `true`. 

### Test repositories
In the organization you are using above, create the following test repositories:

* `test-repo`
  * The description should be `Test description, used in GitHub Terraform provider acceptance test.`
  * The website url should be `http://www.example.com`
  * Create two topics within the repo named `test-topic` and `second-test-topic`
  * In the repo settings, make sure all features and merge button options are enabled.
  * Create a `test-branch` branch
* `test-repo-template`
  * Configure the repository to be a [Template repository](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/creating-a-template-repository)
  * Create a release on the repository with `tag = v1.0`

Export an environment variable corresponding to `GITHUB_TEMPLATE_REPOSITORY=test-repo-template`.

### GitHub users
Export your github username (the one you used to create the personal access token above) as `GITHUB_TEST_USER`. You will need to export a
different github username as `GITHUB_TEST_COLLABORATOR`. Please note that these usernames cannot be the same as each other, and both of them
must be real github usernames. The collaborator user does not need to be added as a collaborator to your test repo or organization, but as
the acceptance tests do real things (and will trigger some notifications for this user), you should probably make sure the person you specify
knows that you're doing this just to be nice. You can also export `GITHUB_TEST_COLLABORATOR_TOKEN` in order to test the invitation acceptance.

Additionally the user exported as `GITHUB_TEST_USER` should have a public email address configured in their profile; this should be exported
as `GITHUB_TEST_USER_EMAIL` and the Github name exported as `GITHUB_TEST_USER_NAME` (this could be different to your GitHub login).

Finally, export the ID of the release created in the template repository as `GITHUB_TEMPLATE_REPOSITORY_RELEASE_ID`
